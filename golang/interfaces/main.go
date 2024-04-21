package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type (
	EnglishBot struct{}
	SpanishBot struct{}
)

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb EnglishBot) getGreeting() string {
	// very custom logic and super different than SpanishBot
	return "Hello there!"
}

func (sb SpanishBot) getGreeting() string {
	// very custom logic and super different than EnglishBot
	return "Hola!"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
