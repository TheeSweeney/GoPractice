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

	fmt.Printf("%+v", jim)

}
