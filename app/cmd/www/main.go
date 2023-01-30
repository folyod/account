package main

import (
	"io"
	"net/http"
)

func main() {
	requestHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "f")
	}

	serv := http.Server{
		Addr:    ":8123",
		Handler: http.HandlerFunc(requestHandler),
	}
	_ = serv.ListenAndServe()
}
