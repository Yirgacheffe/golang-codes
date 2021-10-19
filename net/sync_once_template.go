package main

import (
	"html/template"
	"net/http"
	"sync"
)

func handleTemplate(files ...string) http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplerr = template.ParseFiles(files...)
		})

		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		// use the tpl
		tpl.Execute(w, "")
	}
}
