package gourd_test

import (
	"os"
	"testing"

	"github.com/kaspar-p/gourd"
	"github.com/spf13/viper"
)

func getConfig() gourd.Config {
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

	return gourd.Config{
		AppId:       appId,
		BotToken:    botToken,
		TestChannel: testChannel,
		TestingBot:  testingBot,
	}
}

var (
	tester     gourd.Tester
	testConfig gourd.Config
)

func TestMain(m *testing.M) {
	testConfig = getConfig()

	var disconnect func()

	tester, disconnect = gourd.CreateTester(testConfig)
	defer disconnect()

	code := m.Run()

	os.Exit(code)
}

// Note that this function requires busybee-dev to be running alongside it
func TestToReturnSplit(t *testing.T) {
	c := tester.ExpectSending(".wyd <@" + testConfig.TestingBot + ">")
	c.ToReturn("nothing much \\;)")
	c.ToNotReturn("something else? idk")
}

func TestToContainInline(t *testing.T) {
	tester.ExpectSending(".wyd <@" + testConfig.TestingBot + ">").ToContain("\\;)").ToNotContain("\\:)")
}
