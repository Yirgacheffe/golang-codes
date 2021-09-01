package main

import (
	"bytes"
	"fmt"
)

func main() {
	byteArr := []byte{'H', 'E', 'L', 'L', 'O'}

	str1 := string(byteArr[:])
	str2 := bytes.NewBuffer(byteArr).String()
	str3 := fmt.Sprintf("%s", byteArr)

	fmt.Printf("%s:%s:%s\n", str1, str2, str3)

	// byte[] <-> string
	sArr := []byte("testdfs")
	sStr := string([]byte{'t', 'e', 's', 't'})

	fmt.Println(sArr)
	fmt.Println(sStr)
}
