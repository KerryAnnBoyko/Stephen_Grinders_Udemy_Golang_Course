package main

import "fmt"

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

func (englishBot) getGreeting() string {
	// VERY SPECIFIC LOGIC FOR AN ENGLISH BOT -- contrived example
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/* What we'd need to do without interfaces */
// func printGreetingEB(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreetingSB(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
