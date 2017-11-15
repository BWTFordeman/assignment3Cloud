package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

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
	//Get incoming post Request with json values from dialogflow
	if r.Method == "POST" {
		/*	decoder := json.NewDecoder(r.Body)
			var f IncomingPost

			err := decoder.Decode(&f)
			if err != nil {
				http.Error(w, "Error decoding post request for average", http.StatusBadRequest)
			} else {
		*/
		//Edit the values in order to send to assignment2 application

		//Send response back to dialogflow with correct values in json format.
		url := "https://api.dialogflow.com/v1/query?v=20150910"
		var body = []byte(`{"title":"Do you think I am dumb?"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "Application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		//}
	} else {
		http.Error(w, "Invalid request type", http.StatusMethodNotAllowed)
	}
}
