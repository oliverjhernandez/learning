package airportrobot

import "fmt"

// This exercise does not have tests for each individual task.
// Write your code here.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

// German
type German struct{}

func (gg German) LanguageName() string {
	return "German"
}

func (gg German) Greet(visitorName string) string {
	return fmt.Sprintf("Hallo %s!", visitorName)
}

// Italian
type Italian struct{}

func (gg Italian) LanguageName() string {
	return "Italian"
}

func (gg Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

// Portuguese
type Portuguese struct{}

func (gg Portuguese) LanguageName() string {
	return "Portuguese"
}

func (gg Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(visitorName))
}
