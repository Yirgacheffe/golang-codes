package main

import (
	"log"
	"net/http"
	"os"
)

func Setup() *http.Client {

	c := http.Client{}
	t := Decorate(&http.Transport{}, Logger(log.New(os.Stdout, "", 0)), BasicAuth("username", "password"))

	c.Transport = t

	return &c

}
