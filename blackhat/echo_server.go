package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		s, err := conn.Read(buf[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Fatalln("Unexpected error")
			break
		}

		log.Printf("Received %d bytes: %s\n", s, string(buf))
		log.Println("Writing")

		if _, err := conn.Write(buf[0:s]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func echo_buf(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	s, err := r.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}

	log.Printf("Read %d bytes: %s\n", len(s), s)
	log.Println("Writing")

	w := bufio.NewWriter(conn)
	if _, err := w.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}

	w.Flush()
}

func echo_cp(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {

	listener, err := net.Listen("tcp", ":20076")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}

}
