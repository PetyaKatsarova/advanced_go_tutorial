//https://www.youtube.com/watch?v=rfXSrgIGrKo

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

// start server in terminal: go run main.go
// in bash, i project folder: curl -X POST -H "Content-Type: application/json" -d @testdata/payload.json http://localhost:8080/
// output in bash: pull req id: 191568743
// or instead of the curl command, do it in postman, content type application/json, post, in body raw: put the json

func main() {
	http.HandleFunc("/", handle)
	fmt.Println("server starting on port 8080")
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

/*
The main reason for this nested structure is likely due to the JSON data that this Go code is expected to handle. If the JSON data representing a pull
 request looks something like this:
{
  "pull_request": {
    "ID": 12345
  }
}
Then the nested PullReq struct inside pullReq is used to directly map this JSON structure to the Go struct. The outer struct pullReq represents the
whole JSON object, and the inner PullReq struct corresponds to the pull_request JSON object inside it.
*/

type pullReq struct {
	PullReq struct{ ID int } `json:"pull_request"`
}

/*
prPool is a pool of pullReq objects (using sync.Pool).
to efficiently reuse pullReq instances, reducing the need to allocate new memory for each request.
*/
var prPool = sync.Pool{
	New: func() interface{} { return new(pullReq) },
}

func handle(w http.ResponseWriter, r *http.Request) {
	data := prPool.Get().(*pullReq)
	defer prPool.Put(data) // pullReq object is put back into the pool after the function finishes.
	data.PullReq.ID = 0

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		logrus.Errorf("could ot decode req: %v", err)
		http.Error(w, "could not decode req", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "pull req id: %d", data.PullReq.ID)
}
