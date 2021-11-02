package main

import (
	"io"
	"log"
	"net"
)

// const site = "proxyserver.io:80"

const (
	target   = "httpbin.org:80"
	locproxy = "local.proxy.io"
)

func handler(src net.Conn) {
	dst, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatalln("Unable to connect to target host")
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
	// --------------------------------------------------------
}

func main() {
	listener, err := net.Listen("tcp", ":20083")
	if err != nil {
		log.Fatalln("Unable bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handler(conn)
	}
}
