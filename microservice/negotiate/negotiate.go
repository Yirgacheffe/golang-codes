package main

import (
	"net/http"

	"github.com/unrolled/render"
)

type Negotiator struct {
	ContentType string
	*render.Render
}

func GetNegotiator(r *http.Request) *Negotiator {
	contentType := r.Header.Get("Content-Type")
	return &Negotiator{
		ContentType: contentType, Render: render.New(),
	}
}
