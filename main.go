package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func main() {

	segmentMockServer := mux.NewRouter()
	segmentMockServer.HandleFunc("/", segmentMockHandler)

	req := mux.NewRouter()
	req.HandleFunc("/assert", assertionHandler).Methods("GET")
	req.HandleFunc("/expected", expectationHandler).Methods("POST")

	go func() {
		log.Println(http.ListenAndServe(":8001", segmentMockServer))
	}()

	log.Println(http.ListenAndServe(":8002", req))

}

type request struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Body   string `json:"body"`
}

var expected, actual request
var err error

func assertionHandler(w http.ResponseWriter, r *http.Request) {
	eq := reflect.DeepEqual(expected, actual)
	fmt.Println(eq)
}

func expectationHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := dec.Decode(&expected); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func segmentMockHandler(w http.ResponseWriter, r *http.Request) {
	actual.Path = r.URL.Path
	actual.Method = r.Method
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	actual.Body = string(body)
}
