package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

//FromDialogFlow recieved base and target currency from dialogFlow
type FromDialogFlow struct {
	Result struct {
		Parameters struct {
			BaseCurrency   string `json:"baseCurrency"`
			TargetCurrency string `json:"targetCurrency"`
			Average        string `json:"average"`
			Number         string `json:"number"`
		} `json:"parameters"`
	} `json:"result"`
}

//CurrencyRequest holds the structure of sending data to dialogFlow
type CurrencyRequest struct {
	DisplayText string `json:"displayText"`
	Speech      string `json:"speech"`
}

//LatestRequest sends a latest request to an assignment2 application
type LatestRequest struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}

func main() {
	http.HandleFunc("/dialogflow", testhandler)
	log.Println("http.ListenAndServe", http.ListenAndServe(":"+os.Getenv("PORT"), nil), nil)
}

func testhandler(w http.ResponseWriter, r *http.Request) {

	//Get base and target currency from dialogflow
	var l FromDialogFlow

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, "Error decoding post request for average", http.StatusBadRequest)
		return
	}

	//Post request to other application with currency values:
	message := LatestRequest{}
	message.BaseCurrency = l.Result.Parameters.BaseCurrency
	message.TargetCurrency = l.Result.Parameters.TargetCurrency
	toSend, err := json.Marshal(message)

	var URL string
	str := ""

	if l.Result.Parameters.Average == "average" {
		URL = "https://evil-barrow-41137.herokuapp.com/assignment2/average/"
		str += "The average value between "
	} else {
		URL = "https://evil-barrow-41137.herokuapp.com/assignment2/latest/"
		if l.Result.Parameters.Number == "" {
			str += "The rate between "
		}
	}

	resp, err := http.Post(URL, "application/json", bytes.NewReader(toSend))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	var current float64
	err = json.NewDecoder(resp.Body).Decode(&current)

	if l.Result.Parameters.Number != "" {
		number, err2 := strconv.ParseFloat(l.Result.Parameters.Number, 64)
		if err2 != nil {
			status := http.StatusBadRequest
			http.Error(w, http.StatusText(status), 400)
		}
		current *= number
		str += l.Result.Parameters.Number
	}

	//Make result as string
	var dialogResponse CurrencyRequest
	str += l.Result.Parameters.BaseCurrency
	str += " to "
	str += l.Result.Parameters.TargetCurrency
	str += " is "
	str += strconv.FormatFloat(float64(current), 'f', -1, 32)
	str += "."

	//Send back result to user:
	dialogResponse.DisplayText = str
	dialogResponse.Speech = str

	http.Header.Add(w.Header(), "content-type", "application/json")
	err = json.NewEncoder(w).Encode(dialogResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
