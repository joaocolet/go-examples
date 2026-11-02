package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{base: 10, height: 5}
	s := square{side: 4}

	fmt.Println("Triangle area:", printArea(t))
	fmt.Println("Square area:", printArea(s))
}

func printArea(s shape) float64 {
	return s.getArea()
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}
