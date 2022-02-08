package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Config struct {
	AppId    string
	BotToken string
}

type Tester interface {
	SetChannel(channelId string)
	SendTrigger(message string)
	AssertNextCommandIs(expected string) (bool, string)
}

func CreateConnection(config Config) Tester {
	session, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Printf("Error encountered while creating a bot with token %s. Err: %v\n", config.BotToken, err)

		panic(errors.Wrap(err, "Error encountered while creating a bot with token: "+config.BotToken))
	}

	var t Tester = &DiscordTester{session: session}

	return t
}
