package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostRequest(t *testing.T) {
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

	postRequest("https://evil-barrow-41137.herokuapp.com/assignment2/", rr, req)
}

func TestPostRequest2(t *testing.T) {
	var data FromDialogFlow
	data.Result.Parameters.BaseCurrency = "EUR"
	data.Result.Parameters.TargetCurrency = "USD"
	data.Result.Parameters.Number = "notANumber"
	data.Result.Parameters.Average = "randomTrash"
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

	postRequest("https://evil-barrow-41137.herokuapp.com/assignment2/", rr, req)
}

func TestPostRequest3(t *testing.T) {
	var data FromDialogFlow
	data.Result.Parameters.BaseCurrency = "EUR"
	data.Result.Parameters.TargetCurrency = ""
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

	postRequest("https://evil-barrow-41137.herokuapp.com/assignment2/", rr, req)
}

func TestPostRequest4(t *testing.T) {
	var data FromDialogFlow
	data.Result.Parameters.BaseCurrency = "EUR"
	data.Result.Parameters.TargetCurrency = "NOK"
	data.Result.Parameters.Average = "average"
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

	postRequest("https://evil-barrow-41137.herokuapp.com/assignment2/", rr, req)
}

func TestPostRequest5(t *testing.T) {
	var data CurrencyRequest
	data.DisplayText = "bla"
	data.Speech = "blo"
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

	postRequest("https://evil-barrow-41137.herokuapp.com/assignment2/", rr, req)
}

func TestPostRequest6(t *testing.T) {
	var data FromDialogFlow
	data.Result.Parameters.BaseCurrency = "EUR"
	data.Result.Parameters.TargetCurrency = "NOK"
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

	postRequest("https://evil-barrow-41137.herok", rr, req)
}
