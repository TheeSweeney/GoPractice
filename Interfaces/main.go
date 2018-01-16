package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

//for examples sake, imagine the two getGreeting funcs contain very custom logic
func (englishBot) getGreeting() string {
	//^^^ because we are not using the receiver in the function, we do not need to include a reference to it (for example func (eb englishBot) getGreeting()...)
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
