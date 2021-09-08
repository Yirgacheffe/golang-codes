package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func encryptAES(key []byte, plainText string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(plainText))
	c.Encrypt(out, []byte(plainText))

	return hex.EncodeToString(out)
}

func decryptAES(key []byte, cipherText string) {
	ct, _ := hex.DecodeString(cipherText)
	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ct))
	c.Decrypt(pt, ct)

	fmt.Println("Decrypted:", string(pt[:]))
}

func main() {

	plaintext := []byte("My name is Astaxie")

	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	key := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key = os.Args[2]
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	tt := encryptAES([]byte(key), string(plaintext))
	fmt.Println(plaintext)
	fmt.Println(len(key))

	decryptAES([]byte(key), tt)

	fmt.Println("-------CFB Encrypter --------------------")
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)

	fmt.Printf("%s=>%x\n", plaintext, cipherText)

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, cipherText)

	fmt.Printf("%x=>%s\n", cipherText, plaintextCopy)
}
