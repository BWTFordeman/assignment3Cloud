package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	var data FromDialogFlow
	data.Result.Parameters.BaseCurrency = "EUR"
	data.Result.Parameters.TargetCurrency = "USD"
	data.Result.Parameters.Number = "7"
	m, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/dialogflow", ioutil.NopCloser(strings.NewReader(string(m))))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req) // Test run number 1  (each has different values sent in)

	// Test run number 2
	data.Result.Parameters.Number = ""
	handler.ServeHTTP(rr, req)

	//Test run number 3

	data.Result.Parameters.Number = ""
	m, err = json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/dialogflow", ioutil.NopCloser(strings.NewReader(string(m))))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Test run number 4
	data.Result.Parameters.Average = "average"
	m, err = json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/dialogflow", ioutil.NopCloser(strings.NewReader(string(m))))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
}
