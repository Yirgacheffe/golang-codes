package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	handler http.Handler
}

func (lm *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lm.handler.ServeHTTP(w, r)
	log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
}

type PoliteServer struct {
	//
}

func (ps *PoliteServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! Thanks for visiting!\n")
}

// Wrapped by http.Handler
func anotherLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {

	ps := &PoliteServer{}
	to := http.TimeoutHandler(ps, 2*time.Second, "time out")
	lm := &LoggingMiddleware{
		handler: to,
	}

	log.Fatal(http.ListenAndServe(":9090", lm))
}
