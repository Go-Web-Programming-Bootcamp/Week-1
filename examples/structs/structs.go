package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func (circle *Circle) calculateArea() float64 { // allows struct modification by using pointer
	return math.Pi * circle.r * circle.r
}

func main() {
	circle := Circle{1, 2, 3}
	fmt.Println(circle.x, circle.y, circle.r)
	fmt.Println(circle.calculateArea())
}
