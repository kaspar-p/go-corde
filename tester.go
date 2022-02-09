package gourd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type discordTester struct {
	session   *discordgo.Session
	channelId string
}

func (tester *discordTester) ExpectSending(content string) Verb {
	asyncMessage := make(chan *discordgo.Message)

	// Add the handler that listens to every message
	tester.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		asyncMessage <- m.Message
	})

	// Send the trigger to the bot
	_, err := tester.session.ChannelMessageSend(tester.channelId, content)
	if err != nil {
		log.Println("Error while sending message '"+content+"'. Error:", err)
		panic(err)
	}

	// See and ignore the trigger for the bot
	<-asyncMessage

	// Capture the bot's response
	response := <-asyncMessage

	return &discordVerb{
		message: response,
	}
}
