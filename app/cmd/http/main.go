package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	requestHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Fatal(io.WriteString(w, "f"))
	}

	serv := http.Server{
		Addr:    ":8123",
		Handler: http.HandlerFunc(requestHandler),
	}
	log.Fatal(serv.ListenAndServe())
}
