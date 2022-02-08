package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordTester struct {
	session       *discordgo.Session
	activeChannel string
}

func (tester *DiscordTester) getChannel() *discordgo.Channel {
	channel, err := tester.session.Channel(tester.activeChannel)
	if err != nil {
		log.Println("Error while getting channel with id:", tester.activeChannel, ". Error: ", err)
		panic(err)
	}

	return channel
}

func (tester *DiscordTester) SetChannel(channelId string) {
	tester.activeChannel = channelId
}

func (tester *DiscordTester) SendTrigger(message string) {
	_, err := tester.session.ChannelMessageSend(tester.activeChannel, message)
	if err != nil {
		log.Println("Error while sending message '"+message+"'. Error:", err)
		panic(err)
	}
}

func (tester *DiscordTester) AssertNextCommandIs(expected string) (bool, string) {
	channel := tester.getChannel()
	lastMessage := channel.Messages[len(channel.Messages)-1]
	matches := lastMessage.Content == expected

	if !matches {
		return false, fmt.Sprintf("Error in assert.\n\tExpected: %s\n\tActual: %s\n", expected, lastMessage.Content)
	} else {
		return true, ""
	}
}
