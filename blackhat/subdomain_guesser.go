package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/miekg/dns"
)

type result struct {
	IPAddress string
	Hostname  string
}

type empty struct{}

func lookupA(fqdn, dnsAddr string) ([]string, error) {
	var m dns.Msg
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)

	var ips []string
	in, err := dns.Exchange(&m, dnsAddr)
	if err != nil {
		return ips, err
	}

	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
			return ips, nil
		}
	}

	return ips, nil
}

func lookupCNAME(fqdn, dnsAddr string) ([]string, error) {
	var m dns.Msg
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)

	var fqdns []string
	in, err := dns.Exchange(&m, dnsAddr)
	if err != nil {
		return fqdns, err
	}

	if len(in.Answer) < 1 {
		return fqdns, errors.New("no answer")
	}

	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}

	return fqdns, nil
}

func lookup(fqdn, dnsAddr string) []result {
	var results []result
	var cfqdn = fqdn

	for {
		cnames, err := lookupCNAME(cfqdn, dnsAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue
		}

		ips, err := lookupA(cfqdn, dnsAddr)
		if err != nil {
			break
		}

		for _, ip := range ips {
			results = append(results, result{IPAddress: ip, Hostname: fqdn})
		}
		break
	}

	return results
}

func worker(tracker chan empty, fqdns chan string, gather chan []result, serverAddr string) {
	for fqdn := range fqdns {
		results := lookup(fqdn, serverAddr)
		if len(results) > 0 {
			gather <- results
		}
	}

	tracker <- empty{}
}

func main() {
	var (
		flDomain      = flag.String("domain", "", "The domain to perform guessing.")
		flWordlist    = flag.String("wordlist", "", "THe worklist to guessing.")
		flWorkerCount = flag.Int("c", 50, "The amount of workers to use.")
		flDNSAddr     = flag.String("dns", "8.8.8.8:53", "The DNS server to use.")
	)

	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		log.Println("-domain and -wordlist are required")
		os.Exit(1)
	}

	var results []result

	fqdns := make(chan string, *flWorkerCount)
	gather := make(chan []result)

	tracker := make(chan empty)

	fh, err := os.Open(*flWordlist)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	for i := 0; i < *flWorkerCount; i++ {
		go worker(tracker, fqdns, gather, *flDNSAddr)
	}

	for scanner.Scan() {
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), *flDomain) // i.e. calendar.googleapis.com
	}

	go func() {
		for r := range gather {
			results = append(results, r...)
		}
		var e empty
		tracker <- e
	}()

	close(fqdns)
	for i := 0; i < *flWorkerCount; i++ {
		<-tracker
	}

	close(gather)
	<-tracker

	w := tabwriter.NewWriter(os.Stdout, 0, 8, ' ', ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Hostname, r.IPAddress)
	}

	w.Flush()
}
