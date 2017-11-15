package bot

import (
	"fmt"
	"strings"

	"discordBot/goBot/config"

	"github.com/bwmarrin/discordgo"
)

//BotID holds the id of the bot for sending packages.
var BotID string
var goBat *discordgo.Session

//Start .starts our bot and does stuff.
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

//messageHandler works with messages sent though the discord chat.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Something Christian worked on before assignment2 for fun:
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		//Won't need to respond on what bot writes in chat.
		if m.Author.ID == BotID {
			return
		}

		if m.Content == "!gday" {
			_, _ = s.ChannelMessageSend(m.ChannelID, string("Good day "+m.Author.Username))
		}
		if m.Content == "!ILU" {
			_, _ = s.ChannelMessageSend(m.ChannelID, string("I love you "+m.Author.Username))
		}
		if m.Content == "!help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, string("Commands I listen to are:\n!gday\n !help\n !ILU\n !TTS 'message'."))
		}
		if strings.HasPrefix(m.Content, "!TTS ") {
			text := m.Content[5:]
			_, _ = s.ChannelMessageSendTTS(m.ChannelID, text)
		}

	}
}
