package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type SomeClient struct {
	*http.Client
}

func main() {
	c := SomeClient{
		&http.Client{Timeout: time.Microsecond * 100},
	}

	res, err := c.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		// get `url.Error` struct pointer from `err` interface
		urlErr := err.(*url.Error)

		// check if error occured due to timeout
		if urlErr.Timeout() {
			fmt.Println("Error occured due to a timeout.")
		}
		log.Fatal("Error:", err)
	}

	fmt.Println("Success:", res.StatusCode)
}
