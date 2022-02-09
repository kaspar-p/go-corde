package gourd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type (
	discordVerb struct {
		message *discordgo.Message
	}
)

func (verb *discordVerb) ToContain(substring string) Verb {
	condition := !strings.Contains(verb.message.Content, substring)
	checkCondition(verb.message.Content, substring, "contain", condition)

	return verb
}

func (verb *discordVerb) ToReturn(expected string) Verb {
	condition := verb.message.Content != expected
	checkCondition(verb.message.Content, expected, "equal", condition)

	return verb
}

func (verb *discordVerb) ToNotContain(substring string) Verb {
	condition := strings.Contains(verb.message.Content, substring)
	checkCondition(verb.message.Content, substring, "not contain", condition)

	return verb
}

func (verb *discordVerb) ToNotReturn(expected string) Verb {
	condition := verb.message.Content == expected
	checkCondition(verb.message.Content, expected, "not equal", condition)

	return verb
}

func checkCondition(actual, expected, verb string, failureCondition bool) {
	if failureCondition {
		panic(fmt.Errorf("Error in assert\n\tExpected actual to %s: %s\n\tActual message: %s\n", verb, expected, actual))
	}
}
