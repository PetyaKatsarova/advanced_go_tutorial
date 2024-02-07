// learning from: https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_
// building microservices with go Nic Jackson 1/21

//1. run the file, 2. in another terminal, on windows only in commmand prompt or any linux: curl -v localhost:9090
//3. in the first terminal, where the file was run will see the mesage from / which is HOLA WORLD, u can do curl -v localhost:9090/goodbye
// to send data: curl -v -d 'this is my data'localhost:9090

package main

import (
	"log"
	"net/http"
)

// type Handler interface { ServHTTP(RresponseWriter, *Request) }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		l := log.New(os.Stdout, "product-api", log.LstdFlags)
		
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye WORLD")
	})

	http.ListenAndServe(":9090", nil) // nil uses default serveMux is http handler
}
