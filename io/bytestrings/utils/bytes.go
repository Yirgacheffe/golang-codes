package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// WorkWithBuffer will make use of the buffer created by the
// Buffer function
func WorkWithBuffer() error {

	rawString := "It's easy to encode unicode into a byte array"
	b := Buffer(rawString)

	// convert back to bytes easily
	b.Bytes()
	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	// we can also take our bytes and create a bytes reader
	// these readers implement io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner interfaces
	reader := bytes.NewReader([]byte(s))

	// we can also plug it into a scanner that allows
	// buffered reading and tokenzation
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	// this is no '\n' after scanner, make it looks nice
	fmt.Println()
	return nil

}

func AnotherByteWriteTest() {

	var buffer bytes.Buffer

	buffer.Write([]byte("this is byte buffer"))
	fmt.Fprintf(&buffer, " ,another strnig!\n")

	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)

	buffer.Reset()
	buffer.Write([]byte("Write go again!"))

	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())

	b := make([]byte, 3)

	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Read %s bytes: %d\n", b, n)
	}

}
