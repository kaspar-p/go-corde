package main

import (
	"github.com/kaspar-p/go-corde/src/discord"

	"github.com/spf13/viper"
)

func GetConfig() discord.Config {
	viper.SetConfigFile(".env.yaml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	appId := viper.GetString("APP_ID")
	botToken := viper.GetString("BOT_TOKEN")
	testChannel := viper.GetString("TEST_CHANNEL")
	testingBot := viper.GetString("TESTING_BOT")

	return discord.Config{
		AppId:       appId,
		BotToken:    botToken,
		TestChannel: testChannel,
		TestingBot:  testingBot,
	}
}

func main() {
	config := GetConfig()
	tester := discord.CreateTester(config)

	tester.ExpectSending(".wyd <@" + config.TestingBot + ">").ToGetResponse("nothing much, wbu ;)")
}
