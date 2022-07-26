package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	links map[string]string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	links = map[string]string{}

	http.HandleFunc("/", Home)
	http.HandleFunc("/links/add", addLink)
	http.HandleFunc("/links/get/", getLink)

	log.Fatal(http.ListenAndServe(":9090", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	log.Println("Get Home")

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	var response string
	for s, l := range links {
		response += fmt.Sprintf("Link: <a href=\"http://localhost:9090/links/get/%s\">http://localhost:9090/links/get/%s</a> \t\t Short Link: %s\n", s, s, l)
	}

	fmt.Fprintf(w, "<h2>Hello and Welcome to the Go Demo of URL Shortener!<h2><br>\n")
	fmt.Fprintf(w, response)
	return
}

func addLink(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Link")

	k, ok := r.URL.Query()["link"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to add link, no link provided")
		return
	}

	if !validLink(k[0]) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not create shortlink need absolute path link. Ex: /links/add?link=https://github.com/")
		return
	}

	log.Println(k)

	if _, ok = links[k[0]]; !ok {
		genKey := randStringKey(10)
		links[genKey] = k[0]

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusAccepted)

		href := fmt.Sprintf("<a href=\"http://localhost:9090/links/get/%s\">http://localhost:9090/links/get/%s</a>", genKey, genKey)
		fmt.Fprintf(w, "Short link added\n")
		fmt.Fprintf(w, href)
		return
	} else {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "This link has been added already")
		return
	}
}

func getLink(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println("Get link:", path)

	pathArgs := strings.Split(path, "/")
	shortKey := pathArgs[3]

	if len(shortKey) < 1 {
		w.WriteHeader(http.StatusNotFound)
		http.Redirect(w, r, "http://localhost:9090/", http.StatusTemporaryRedirect)
		return
	} else {
		url := links[shortKey]
		log.Printf("Redirect to: %s", url)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect) // ----- ------ ---
		return
	}

}

func validLink(link string) bool {
	r, err := regexp.Compile("^(http|https)://")
	if err != nil {
		return false
	}

	l := strings.TrimSpace(link)
	log.Printf("Checking for valid link: %s", l)
	return r.MatchString(l)
}

func randStringKey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

//  -- - - - - -- - - - - -- - - - - -- - - - - -- - - - - -- - - - - -- - - - - -- - - - - -- - - - - -- - - - -
