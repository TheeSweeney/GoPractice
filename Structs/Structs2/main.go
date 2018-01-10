package main

import (
	"fmt"
)

type person struct {
	firstName, lastName string
}

func main() {

	jim := person{
		firstName: "butts",
		lastName:  "mcGee",
	}

	fmt.Println(jim)

}
