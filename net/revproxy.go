package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	remote, err := url.Parse("http://cn.bing.com")
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)

			r.Host = remote.Host
			w.Header().Set("X-Proxy-By", "SoftSprit")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))

	if err = http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
