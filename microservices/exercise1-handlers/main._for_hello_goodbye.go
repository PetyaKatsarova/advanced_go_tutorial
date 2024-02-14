// learning from:
// building microservices with go Nic Jackson 1/21

//1. run the file, 2. in another terminal, on windows only in commmand prompt or any linux: curl -v localhost:9090
//3. in the first terminal, where the file was run will see the mesage from / which is HOLA WORLD, u can do curl -v localhost:9090/goodbye
// to send data: curl -v -d 'this is my data' localhost:9090

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	// "github.com/PetyaKatsarova/advanced_go_tutorial"
)

// type Handler interface { ServHTTP(RresponseWriter, *Request) }

// rename to main only
func main_for_hello_goodbye() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloMsg := NewHello(l)
	byeMsg := NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", helloMsg)
	sm.Handle("/goodbye", byeMsg)

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("goodbye WORLD")
	// })
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	} // a type

	go func() {
		err := s.ListenAndServe() // not blocking in go func
		if err != nil {
			l.Fatal(err)
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
