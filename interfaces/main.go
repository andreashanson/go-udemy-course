package main

import "fmt"

type bot interface{
	getGreeting() string
}
type spanishBot struct{}
type englishBot struct{}

func main() {
	fmt.Println("Interfaces")
	sb := spanishBot{}
	eb := englishBot{}
	printGreeting(sb)
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
