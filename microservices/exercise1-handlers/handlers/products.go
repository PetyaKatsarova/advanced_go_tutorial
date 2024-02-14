package handlers

// learning from "github.com/nicholasjackson/building-microservices-youtube/product-api"
/*
json: encode is faster and no memory allocation, marshal for bigger files: stores in buffer and allocates memory and single threads
*/
// curl localhost:9090 -XDELETE -v
// https://www.youtube.com/watch?v=UZbHLVsjpF0&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&index=4  //video 4

import (
	"advanced_go_tutorial/microservices/exercise1-handlers/data"
	"regexp"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`) // capture group () 1 or more
		g := reg.FindAllStringSubmatch(r.URI.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 1 { 
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString	:= g[0][1]
		id, err		:= strconv.Atoi(idString)
		if err != nil {
			return
		}
		p.l.Println("dot id", id)
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts() // lp: list products
	// d, err := json.Marshal(lp) // encodes recursevly in json slower than json.Encode
	err := lp.ToJSON(rw) // writing json
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// curl localhost:9090 -XPOST -v
// curl -v -X POST -H "Content-Type: application/json" -d "{\"id\": 1, \"name\": \"bla\", \"description\": \"bla bla describing\"}" http://localhost:9090

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body) // reader we are going to use is response body
	if err != nil {
		http.Error(rw, "Unable to unmarshal/decode json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}
