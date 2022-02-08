package main

import (
	"github.com/spf13/viper"
)

func GetConfig() Config {
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

	return Config{
		AppId:       appId,
		BotToken:    botToken,
		TestChannel: testChannel,
		TestingBot:  testingBot,
	}
}

func main() {
	config := GetConfig()
	tester := CreateTester(config)

	tester.ExpectSending(".wyd <@" + config.TestingBot + ">").ToGetResponse("nothing much, wbu ;)")
}
