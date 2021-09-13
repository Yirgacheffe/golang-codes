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

	fmt.Println(bytes.Compare([]byte("abc"), []byte("cba")))
	fmt.Println(bytes.Contains([]byte("abc"), []byte("a")))

	fmt.Println(bytes.ContainsRune([]byte("xyz"), 'z'))
	fmt.Println(bytes.Equal([]byte("xyz"), []byte("abc")))

	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))

	fmt.Println(bytes.Fields([]byte("   aaaa          bbb        ccc  ")))      // [[97 97 97 97] [98 98 98] [99 99 99]]
	fmt.Printf("%q", bytes.Fields([]byte("   aaaa          bbb        ccc  "))) // ["aaaa" "bbb" "ccc"]
	fmt.Println(bytes.Index([]byte("abcdefghi"), []byte("cd")))                 // 2

	s := [][]byte{[]byte("abc"), []byte("def"), []byte("ghi")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))

	fmt.Println(string(bytes.Repeat([]byte("abc"), 3)))
	fmt.Printf("%q\n", bytes.Split([]byte("abababab"), []byte("b")))
	fmt.Printf("%q\n", bytes.Trim([]byte("abcdefghi  "), " "))

}
