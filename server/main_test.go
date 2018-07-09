package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestMockServerSavesPayload(t *testing.T) {
	f, err := os.Open("fixture.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	jsonData, _ := (ioutil.ReadAll(f))

	req, _ := http.NewRequest(http.MethodPost, "https://api.segment.io/v1/t", bytes.NewBuffer(jsonData))
	segmentMockHandler(&responseWriter{}, req)
	var j interface{}
	if err := json.Unmarshal(jsonData, &j); err != nil {
		t.Errorf("Could not unmarshall json fixture (%s)", err)
	}

	if !reflect.DeepEqual(actual.Body, string(jsonData)) {
		t.Errorf("Expected (%s), received (%s)", jsonData, actual.Body)
	}

}

type responseWriter struct{}

func (responseWriter) WriteHeader(int)           {}
func (responseWriter) Write([]byte) (int, error) { return 0, nil }
func (responseWriter) Header() http.Header       { return map[string][]string{} }
