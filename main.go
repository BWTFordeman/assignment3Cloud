package main

import (
	"discordBot/goBot/bot"
	"discordBot/goBot/config"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	log.Println("http.ListenAndServe", http.ListenAndServe(":"+os.Getenv("PORT"), nil), nil)
}
