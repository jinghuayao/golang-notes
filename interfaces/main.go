/*
Rules of Interfaces


Concrete type vs Interface type

concrete type: can create instance;
interface type: can not create instance directly

use a concrete type to realize an interface type,
then create instance using the concrete type that realizes the interface type
*/

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
	// custom logic
	return "Hi There"
}


func (spanishBot) getGreeting() string {
	// custom logic
	return "Hola!"
}


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}