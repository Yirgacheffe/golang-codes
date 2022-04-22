package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const RespTmpl = `
{ 
	"code": "E0000", 
	"describe": "Request succeed!", 
	"result": { 
		"invoiceSerialNum": "%s" 
	}
}`

func InvoiceHandler(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	serial := uuid.New()
	result := fmt.Sprintf(RespTmpl, serial)
	w.Write([]byte(result))

}

func main() {

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/invoices", InvoiceHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	log.Println("Nisemono will start, listening on 9090 ...")
	server.ListenAndServe()

}
