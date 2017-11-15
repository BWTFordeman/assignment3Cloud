package main

import (
	"discordBot/goBot/bot"
	"discordBot/goBot/config"
	"fmt"
	"log"
	"net/http"
	"os"
)

//IncomingPost gets the json values from dialogflow
/*type IncomingPost struct {
  ID string `json:"id"`
}*/

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
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	log.Println("http.ListenAndServe", http.ListenAndServe(":"+os.Getenv("PORT"), nil), nil)
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Testing")
	//Get incoming post Request with json values from dialogflow

	//Edit the values in order to send to assignment2 application

	//Send response back to dialogflow with correct values in json format.
}
