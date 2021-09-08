package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide an integer")
		os.Exit(1)
	}

	nbr, _ := strconv.ParseInt(os.Args[1], 10, 64)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, nbr)
	if err != nil {
		fmt.Println("Little Endian:", err)
	}

	fmt.Printf("%d is %x in Little Endian\n", nbr, buf)

	buf.Reset()
	err = binary.Write(buf, binary.BigEndian, nbr)
	if err != nil {
		fmt.Println("Big Endian:", err)
	}

	fmt.Printf("And %x in Big Endian\n", buf)
}
