/*
create a program that reaeds the contents of a text file the prints its contents to the terminal

the file to open should be provided as an argument to the program when it is executed at the terminal. For example 'go run main.go myfile.txt' should open up the myfile.txt file

To read in the arguments provided to a program, you can reference the variable 'os.Args', which is a slice of type string

To open a file, check out the documentation for the 'Open' function in the 'os' package

What interfaces does the 'file' type implement?

if the 'file' type implements the 'reader' interface, you might be able to reuse that io.Copy function
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(f))
}
