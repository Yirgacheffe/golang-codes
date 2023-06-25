package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var errPosArgSpecified = errors.New("positional arguments specified")

type config struct {
	nbrOfTimes int
	printUsage bool
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
A greeter application which prints the name you entered <integer> number of times.`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintln(w, usageString)
}

func validateArgs(c config) error {
	if !(c.nbrOfTimes > 0) {
		return errors.New("must specify a number greater then 0")
	}

	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprint(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errPosArgSpecified
	}

	return name, nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}

	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		var usageString = `
		A greeter application which prints the name you entered a specified number of times.
		
		Usage of %s: <options> [name]`
		fmt.Fprintf(w, usageString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Option: ")
		fs.PrintDefaults()
	}

	fs.IntVar(&c.nbrOfTimes, "n", 0, "Number of times of greet")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}

	return c, nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}

	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Hello %s\n", name)
	for i := 0; i < c.nbrOfTimes; i++ {
		fmt.Fprintln(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		if errors.Is(err, errPosArgSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
