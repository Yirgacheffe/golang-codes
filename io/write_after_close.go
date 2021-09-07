package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func write(f io.Writer, s string) {
	n, err := io.WriteString(f, s)
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		fmt.Printf("Successful! %d bytes has been written.\n", n)
	}
}

func main() {
	f, err := os.OpenFile("info.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error:", err)
	}
	write(f, "Hello")

	fmt.Printf("File name: %v\n", f.Name())
	fmt.Printf("File description: %v\n", f.Fd())

	f.Close()
	write(f, "Hello Again!")
}
