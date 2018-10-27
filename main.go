package main

import "fmt"

// new custom type named bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for english greeting
	return "hi there"
}

func (sb spanishBot) getGreeting() string {
	// very custom logic for spanish greeting
	return "hasta la vista"
}