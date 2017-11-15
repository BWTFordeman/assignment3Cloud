package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

//outgoingPost is the structure of json values posting to dialogflow
type outgoingPost struct {
	Contexts  map[int]string `json:"contexts"`
	Lang      string         `json:"lang"`
	Query     string         `json:"query"`
	SessionID string         `json:"sessionId"`
	TimeZone  string         `json:"timezone"`
}

//IncomingPost gets the json values from dialogflow
type IncomingPost struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Lang      string `json:"lang"`
	Result    struct {
	} `json:"result"`
	Status struct {
	} `json:"status"`
	SessionID string `json:"sessionId"`
}

//FromDialogFlow recieved base and target currency from dialogFlow
type FromDialogFlow struct {
	Result struct {
		Parameters struct {
			BaseCurrency   string `json:"baseCurrency"`
			TargetCurrency string `json:"targetCurrency"`
		} `json:"parameters"`
	} `json:"result"`
}

type CurrencyRequest struct {
	DisplayText string `json:"displayText"`
	Speech      string `json:"speech"`
}

func main() {
	http.HandleFunc("/dialogflow", testhandler)

	log.Println("http.ListenAndServe", http.ListenAndServe(":"+os.Getenv("PORT"), nil), nil)
}

func testhandler(w http.ResponseWriter, r *http.Request) {

	var l FromDialogFlow

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, "Error decoding post request for average", http.StatusBadRequest)
		return
	}
	//Send to assignment2

	//send back to dialogflow
	var dialogResponse CurrencyRequest
	str := ""
	str += "The rate between ..."

	dialogResponse.DisplayText = str
	dialogResponse.Speech = str

	http.Header.Add(w.Header(), "content-type", "application/json")
	err = json.NewEncoder(w).Encode(dialogResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
	}

}
