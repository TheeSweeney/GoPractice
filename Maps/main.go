package main

import "fmt"

func main() {
	colors := map[string]string{
		//we are declaring a map where all keys are type sting, and all values are type string
		"red":   "#ff0000",
		"green": "#4bf745",
	} //as with structs, maps must end with comma

	fmt.Println(colors)
}
