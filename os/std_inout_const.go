package main

import (
	"fmt"
	"syscall"
)

type readfd int

func (r readfd) Read(buf []byte) (int, error) {
	return syscall.Read(int(r), buf)
}

type writefd int

func (w writefd) Write(buf []byte) (int, error) {
	return syscall.Write(int(w), buf)
}

const (
	Stdin  = readfd(0)
	Stdout = writefd(1)
	Stderr = writefd(2)
)

func main() {
	fmt.Fprintf(Stdout, "Hello world!")
}
