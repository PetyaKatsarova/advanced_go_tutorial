package main

import (
	"net/http"
	"log"
	"io"
	"fmt"
)


type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("HOLA WORLD")
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Opps", http.StatusBadRequest) // doesnt terminate the flow: u need return
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("oops"))
			return
		}
		fmt.Printf("Data %s\n", data)
		fmt.Fprintf(rw, "Hello %s", data) //NB:!!!: responsewriter writes the data in the terminal where we run the code,
		// request data is on the server!!! the other terminal where we run the curl command
}
