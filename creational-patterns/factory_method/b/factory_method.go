package main

import "fmt"

type Shape interface {
	Area() int
}

type Circle struct {
	radius int
}

type Square struct {
	side int
}

type Rectangle struct {
	side int
}

func (c *Circle) Area() int {
	return c.radius * c.radius * 22 / 7
}

func (s *Square) Area() int {
	return s.side * s.side
}

func GetArea(s Shape) int {
	return s.Area()
}

func main() {
	s := &Square{5}
	fmt.Println(GetArea(s))

	c := &Circle{10}
	fmt.Println(GetArea(c))

	// will error, because type "Rectangle does not fulfill interface"
	// r := &Rectangle{5}
	// fmt.Println(GetArea(r))
}
