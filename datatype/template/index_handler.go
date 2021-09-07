package main

import (
	"fmt"
	"io"
	"net/http"
)

var helloIndexHTML = "<DOCTYPE-TYPE html><html><head><title>Hello World</title></head><body>Hello, World!</body></html>"

func helloHTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, helloIndexHTML)
}

func main() {
	fmt.Println("---------------------------------------------")

	http.HandleFunc("/hello", helloHTTPHandler)
	http.ListenAndServe(":9000", nil)
	http.Handle("/assets/", http.FileServer(http.Dir("assets")))
}
