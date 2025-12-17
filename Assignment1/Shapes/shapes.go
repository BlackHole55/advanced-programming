package Shapes

import "math"

type Shapes interface {
	Area() float64
	Perimeter() float64
}

// Rectangle
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

// Square
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Triangle
type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	semiPerimeter := t.Perimeter() / 2

	return math.Sqrt(semiPerimeter * (semiPerimeter - t.SideA) * (semiPerimeter - t.SideB) * (semiPerimeter - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
