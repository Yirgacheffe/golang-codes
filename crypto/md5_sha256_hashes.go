package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

var (
	md5hash    = ""
	sha256hash = ""
)

func main() {

	fh, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {

		pass := scanner.Text()
		hash := fmt.Sprintf("%x", md5.Sum([]byte(pass)))

		if hash == md5hash {
			fmt.Printf("[+] Password found (MD5): %s\n", pass)
		}

		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
		if hash == sha256hash {
			fmt.Printf("[+] Password found (SHA-256): %s\n", pass)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

}
