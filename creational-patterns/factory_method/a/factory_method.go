package main

import "fmt"

const (
	Circle = 1
	Square = 2
)

// Product
type ShapeProduct interface {
	Area() int // factory method
}

// Concrete Product
type CircleConcreteBuilder struct {
	Radius int
}

func (c *CircleConcreteBuilder) Area() int {
	return 22 / 7 * c.Radius * c.Radius
}

type SquareConcreteBuilder struct {
	Size int
}

func (s *SquareConcreteBuilder) Area() int {
	return s.Size * s.Size
}

// factory
func ShapeFactory(m, num int) ShapeProduct {
	switch m {
	case Circle:
		return &CircleConcreteBuilder{num}
	case Square:
		return &SquareConcreteBuilder{Size: num}
	default:
		return nil
	}
}

func main() {
	circle := ShapeFactory(1, 5)
	fmt.Println(circle.Area())

	square := ShapeFactory(2, 5)
	fmt.Println(square.Area())
}
