package main

import (
	"advanced_go_tutorial/microservices/exercise1-handlers/handlers"
	"log"
	"net/http"
	"os"
	"github.com/nicholasjackson/env"
	"time"
	"fmt"
	"context"
	"os/signal"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

// curl localhost:9090 | jq 
// for linux only jq: this is for the second new terminal after you run main.go
func main() {
	env.Parse()

	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	ph := handlers.NewProducts(l) // product handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	} // a type

	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe() // not blocking in go func
		if err != nil {	
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("Recieved terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc) // gracefully shutdown: server will not take requests but will serve the left over work
	// http.ListenAndServe(":9090", sm) // nil uses default serveMux is http handler
}