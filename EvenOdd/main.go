package main

import "fmt"

func oddEven(n int) {
	if n%2 == 0 {
		fmt.Printf("%v is Even\n", n)
	} else {
		fmt.Printf("%v is Odd\n", n)
	}
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range numbers {
		oddEven(n)
	}
}
