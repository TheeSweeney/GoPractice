package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName, lastName string
	contactInfo
}

func main() {

	jim := person{
		firstName: "butts",
		lastName:  "mcGee",
		contactInfo: contactInfo{
			email:   "butts@gmail.com",
			zipCode: 11580,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}
