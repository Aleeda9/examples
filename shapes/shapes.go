package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
}

type Rectangle struct {
	a, b float64
}

func (r *Rectangle) perimeter() float64 {
	return 2*r.a + 2*r.b
}

type Circle struct {
	r float64
}

func (c *Circle) perimeter() float64 {
	return c.r * 2 * math.Pi
}

func shapePerimeter(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Println("perimeter", shape.perimeter())
	}
}

func main() {
	c := Circle{r: 5}
	r := Rectangle{a: 2, b: 3}

	shapePerimeter(&c, &r)
}
