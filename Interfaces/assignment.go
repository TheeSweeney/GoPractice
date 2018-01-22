/*
write a program that creates two custom struct types called 'triangle; and ;square;

the 'square' type should be a struct with a field called 'sideLength' of type float64

The 'triangle' type should be a struct with a field called 'height' of type float64 and a field of type 'base' of type float64

Both types should have function called 'getArea' that returns the calculated area of the square or triangle

Area of a triangle = 0.5*base*height
area of a square = sideLength * sideLength

Add a 'shape' interface that defines a function called 'printArea'. This function should calculate the area o the given hape and print it out to the terminal Design the interface so that the 'printArea' function can be called with either a triangle or a square.
*/

package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	t := triangle{
		base:   10,
		height: 10,
	}

	sq := square{
		sideLength: 10,
	}

	printArea(t)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Printf("the %T has an area of: %v \n", s, s.getArea())
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
