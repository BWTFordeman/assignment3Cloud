package main

import (
	"bytes"
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

/*{
  "id": "f4f70dee-5ac4-4d4a-8c46-4e3b7d7f327f",
  "timestamp": "2017-11-15T12:36:02.275Z",
  "lang": "en",
  "result": {
    "source": "agent",
    "resolvedQuery": "hi",
    "action": "input.unknown",
    "actionIncomplete": false,
    "parameters": {},
    "contexts": [],
    "metadata": {
      "intentId": "459844df-995f-41c3-b301-47904e3f9b09",
      "webhookUsed": "false",
      "webhookForSlotFillingUsed": "false",
      "intentName": "Default Fallback Intent"
    },
    "fulfillment": {
      "speech": "One more time?",
      "messages": [
        {
          "type": 0,
          "speech": "Sorry, what was that?"
        }
      ]
    },
    "score": 1
  },
  "status": {
    "code": 200,
    "errorType": "success",
    "webhookTimedOut": false
  },
  "sessionId": "edd422d8-9619-4871-9e31-53270f020cec"
}*/

func main() {
	http.HandleFunc("/", testhandler)

	log.Println("http.ListenAndServe", http.ListenAndServe(":"+os.Getenv("PORT"), nil), nil)
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	//Getting message from slack:
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing", http.StatusTeapot)
	}
	form := r.Form

	form.Get("text")
	log.Println(form)

	//sending message to dialogflow:

	//sending message to other Application

	//Use correct body for sending message to slack:

	//Sending message to slack:
	webhook := "https://hooks.slack.com/services/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"
	var body = []byte(`{"text":"Whats up?"}`)
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error parsing", http.StatusTeapot)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error posting", http.StatusTeapot)
	}
	defer resp.Body.Close()

}
