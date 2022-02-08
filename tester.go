package gourd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordTester struct {
	session   *discordgo.Session
	channelId string
}

func (tester *DiscordTester) ExpectSending(content string) SendChain {
	c := make(chan string, 1)

	_, err := tester.session.ChannelMessageSend(tester.channelId, content)
	if err != nil {
		log.Println("Error while sending message '"+content+"'. Error:", err)
		panic(err)
	}

	return &DiscordSendChain{
		session:       tester.session,
		channelId:     tester.channelId,
		returnChannel: c,
	}
}
