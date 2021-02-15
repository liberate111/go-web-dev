package main

import "fmt"

type square struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (c circle) area() float64 {
	return PI * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("Area: %f\n", s.area())
}

const PI = 3.14159

func main() {
	sq1 := square{
		width:  5,
		height: 10,
	}

	c1 := circle{
		radius: 1,
	}

	info(sq1)
	info(c1)
}
