package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

func rsa_oaep_encrypt(secMsg string, key rsa.PublicKey) string {
	label := []byte("oaep encrypted")
	rng := rand.Reader
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secMsg), label)

	CheckError(err)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func rsa_oaep_decrypt(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("oaep encrypted")
	rng := rand.Reader
	plainText, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)

	CheckError(err)
	fmt.Println("Plaintext:", string(plainText))
	return string(plainText)
}

func main() {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)
	pubKey := privKey.PublicKey
	secMsg := "Very secret!"

	encryptedMsg := rsa_oaep_encrypt(secMsg, pubKey)
	fmt.Println("Cipher Text:", encryptedMsg)

	decryptedMsg := rsa_oaep_decrypt(encryptedMsg, *privKey)
	fmt.Println("Decrypted Msg:", decryptedMsg)
}
