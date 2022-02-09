package gourd

type Verb interface {
	ToNotContain(substring string) Verb
	ToNotReturn(substring string) Verb
	ToReturn(expected string) Verb
	ToContain(substring string) Verb
}

type Tester interface {
	ExpectSending(content string) Verb
}

type Config struct {
	AppId       string
	BotToken    string
	TestChannel string
	TestingBot  string
}
