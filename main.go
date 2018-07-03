package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func main() {

	var expectedJSON, actualJSON []byte
	var err error
	http.HandleFunc("/application-payload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if expectedJSON, err = ioutil.ReadAll(r.Body); err != nil {
				log.Fatal(err)
			}
		}
	})

	http.HandleFunc("/expected-payload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if actualJSON, err = ioutil.ReadAll(r.Body); err != nil {
				log.Fatal(err)
			}
			eq := reflect.DeepEqual(expectedJSON, actualJSON)
			fmt.Println(eq)

		}

	})

	http.ListenAndServe(":6000", nil)
}
