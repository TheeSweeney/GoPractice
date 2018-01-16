package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	//^^^type of map being iterated over
	for color, hex := range c {
		//color - key for this step through the loop
		//hex - value for this step through the loop
		fmt.Println("Hex code for", color, "is", hex)
	}
}
