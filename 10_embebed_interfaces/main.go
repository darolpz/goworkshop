package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Material interface {
	Shape
	Object
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("Type: %T value: %v\n", m, m)
	fmt.Printf("Type: %T value: %v\n", s, s)
	fmt.Printf("Type: %T value: %v\n", o, o)

	fmt.Printf("Area: %v\n", m.Area())
	fmt.Printf("Area: %v\n", o.(Material).Area())
}
