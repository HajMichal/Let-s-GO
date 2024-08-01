package main

import "fmt"

type bot interface {
	getGreeting() string
}

type enghlishBot struct {}
type spanishBot struct {}

func main () {
	eb := enghlishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (enghlishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}