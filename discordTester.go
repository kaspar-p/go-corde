package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type DiscordSendChain struct {
	session   *discordgo.Session
	channelId string
}

func (sendChain *DiscordSendChain) ToGetResponse(expected string) {
	channel, err := sendChain.session.Channel(sendChain.channelId)
	if err != nil {
		log.Println("Error while getting channel with id:", sendChain.channelId, ". Error: ", err)
		panic(err)
	}

	if len(channel.Messages) == 0 {
		panic(fmt.Errorf("Error in assert.\n\tExpected: %s\n\tActual: <nothing>\n", expected))
	}

	lastMessage := channel.Messages[len(channel.Messages)-1]
	matches := lastMessage.Content == expected

	if !matches {
		panic(fmt.Errorf("Error in assert.\n\tExpected: %s\n\tActual: %s\n", expected, lastMessage.Content))
	}
}

type DiscordTester struct {
	session   *discordgo.Session
	channelId string
}

func (tester *DiscordTester) ExpectSending(message string) SendChain {
	_, err := tester.session.ChannelMessageSend(tester.channelId, message)
	if err != nil {
		log.Println("Error while sending message '"+message+"'. Error:", err)
		panic(err)
	}
	time.Sleep(5 * time.Second)

	return &DiscordSendChain{session: tester.session, channelId: tester.channelId}
}
