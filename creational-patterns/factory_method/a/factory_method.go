package main

import "fmt"

// Product
type ShapeProduct interface {
	Area() float64
}

// Concrete Product
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 22 / 7 * c.Radius * c.Radius
}

type Square struct {
	Size float64
}

func (s *Square) Area() float64 {
	return s.Size * s.Size
}

// creator

func main() {
	fmt.Println("this is factory method")
}
