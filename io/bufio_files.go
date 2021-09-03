package main

import (
	"bufio"
	"fmt"
	"os"
)

func write() {
	f, e := os.OpenFile("buffertest.txt", os.O_WRONLY, 0666)
	CheckError(e)

	// create a sized buffer of 4 bytes and the default is 4096 bytes
	bw := bufio.NewWriterSize(f, 4)

	// write to buffer
	bw.Write([]byte("H"))
	bw.Write([]byte("e"))
	bw.Write([]byte("l"))
	bw.Write([]byte("l"))
	bw.Write([]byte("o"))
	bw.Write([]byte(" "))
	bw.Write([]byte("w"))
	bw.Write([]byte("o"))
	bw.Write([]byte("r"))
	bw.Write([]byte("l"))
	bw.Write([]byte("d"))

	// check how much is inside waiting to be written
	fmt.Println(bw.Buffered()) // 3

	// check available space left
	fmt.Println(bw.Available()) // 1

	// bw.Flush()
}

func read() {
	// open file for reading
	f, e := os.Open("buffertest.txt")
	CheckError(e)

	br := bufio.NewReader(f)

	// peek n bytes bbuf is a byte buffer of size 10
	bbuf := make([]byte, 10)
	bbuf, e = br.Peek(6)
	CheckError(e)

	// bbuf contents
	fmt.Println(string(bbuf)) // Hello

	// num read
	nr, e := br.Read(bbuf)
	CheckError(e)

	fmt.Println("Num bytes read", nr) // 6

	// read single byte
	singleByte, e := br.ReadByte()
	CheckError(e)

	fmt.Println("Single byte is", string(singleByte)) // w

	// reset buffer
	br.Reset(f)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	write()
	read()
}
