package main

import "fmt"

func isEven(n int) bool {
	isEven := true

	if n%2 == 0 {
		fmt.Printf("%v is Even\n", n)
	}
	if n%2 != 0 {
		fmt.Printf("%v is Odd\n", n)
		isEven = false
	}

	return isEven
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range numbers {
		isEven(n)
	}
}
