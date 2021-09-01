package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
)

func Base64_Urlller() {
	data := []byte("Raw Data!")
	v := base64.URLEncoding.EncodeToString(data)
	fmt.Printf("Encoding: %s\n", v)

	d, err := base64.URLEncoding.DecodeString(v)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Decoding: %s\n", d)
}

func Base64_Encoder() {
	buf := bytes.Buffer{}

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer func() {
		if err := encoder.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err := encoder.Write([]byte("Other Data!")); err != nil {
		log.Println(err)
	}
	fmt.Println("StdEncode: ", buf.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	r, err := ioutil.ReadAll(decoder)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("StdEncode Decoder: ", string(r))
}

type Pos struct {
	X      int
	Y      int
	Object string
}

func Gob_EncodeTest() {
	buf := bytes.Buffer{}
	p := Pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}
	// gob.Register() // If p was an interface

	e := gob.NewEncoder(&buf)
	if err := e.Encode(&p); err != nil {
		log.Println(err)
	}
	fmt.Println("Gob Encoded Len: ", len(buf.Bytes()))

	p2 := Pos{}
	d := gob.NewDecoder(&buf)
	if err := d.Decode(&p2); err != nil {
		log.Println(err)
	}
	fmt.Println("Gob Decode Value: ", p2)
}

func main() {
	Base64_Urlller()
	Base64_Encoder()
	Gob_EncodeTest()
}
