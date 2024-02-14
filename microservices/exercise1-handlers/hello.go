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

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello request")
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error radind body", err)
			http.Error(rw, "Unable to read request body", http.StatusBadRequest) // doesnt terminate the flow: u need return
			return
		}
		fmt.Printf("From Hello with response writer:  %s\n", data)
		fmt.Fprintf(rw, "Hello %s", data) //NB:!!!: responsewriter writes the data in the terminal where we run the code,
		// request data is on the server!!! the other terminal where we run the curl command
}
