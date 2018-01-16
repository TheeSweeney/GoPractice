package main

import "fmt"

type bot interface {
	getGreeting() string
}

//This is basically saying:
//yo program, we have a new type called "bot" (need not be named that way, just done to match the englishBot or spanishBot)
//if you are a type in this program with a function called "getGreeting", you are now an honorary memnber of the type "bot"

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

//because englishBot and spanishBot both implement a "getGreeting" function, they are also type "bot"
//as such, they have access to the printGreeting func

//for examples sake, imagine the two getGreeting funcs contain very custom logic
func (englishBot) getGreeting() string {
	//^^^ because we are not using the receiver in the function, we do not need to include a reference to it (for example func (eb englishBot) getGreeting()...)
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
