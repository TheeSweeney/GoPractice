package main

import "fmt"

type person struct {
	// firstName string
	// lastName  string
	// since both parts of the struct are string, it can be written as
	firstName, lastName string
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// if we do not specify, fields will be assigned in the order they exist in the struct
	// not great in case we want to reorder the struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// ^^^ this way we can add fields in any order, and can impartially complete a struct - if a struct has three fields, we can assign values to the firs and third
	// we can also create structs like so
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)

}
