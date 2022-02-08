package gourd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type DiscordSendChain struct {
	session       *discordgo.Session
	channelId     string
	returnChannel chan string
}

type ConditionFunc func(message *discordgo.MessageCreate) bool

func (sendChain *DiscordSendChain) checkCondition(desc string, conditionFunc ConditionFunc, passedString string) *DiscordSendChain {
	sendChain.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if conditionFunc(m) {
			sendChain.returnChannel <- fmt.Sprintf("Error in assert\n\tExpected actual to %s: %s\n\tActual message: %s\n", desc, passedString, m.Content)
		} else {
			sendChain.returnChannel <- ""
		}
	})

	matched := <-sendChain.returnChannel
	if matched != "" {
		panic(errors.New(matched))
	}

	return sendChain
}

func (sendChain *DiscordSendChain) ToContain(substring string) SendChain {
	return sendChain.checkCondition(
		"contain",
		func(message *discordgo.MessageCreate) bool {
			return !strings.Contains(message.Content, substring)
		},
		substring,
	)
}

func (sendChain *DiscordSendChain) ToReturn(expected string) SendChain {
	return sendChain.checkCondition(
		"equal",
		func(message *discordgo.MessageCreate) bool {
			return message.Content != expected
		},
		expected,
	)
}
