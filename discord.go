package gourd

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func CreateTester(config Config) (tester Tester, disconnect func()) {
	var session *discordgo.Session

	preparedSession := make(chan *discordgo.Session)

	// Create the new bot session
	session, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Printf("Error encountered while creating a bot with token %s. Err: %v\n", config.BotToken, err)

		panic(errors.Wrap(err, "Error encountered while creating a bot with token: "+config.BotToken))
	}

	// Wait until the bot has connected
	session.AddHandler(func(s *discordgo.Session, ready *discordgo.Ready) {
		preparedSession <- s
	})
	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

	// Open the bot
	err = session.Open()
	if err != nil {
		log.Println("Error connecting to discord:", err)
		panic(err)
	}

	// Wait for async function to finish
	session = <-preparedSession

	return &discordTester{session: session, channelId: config.TestChannel}, func() {
		err := session.Close()
		if err != nil {
			panic(errors.Wrap(err, "Error closing discord session"))
		}
	}
}
