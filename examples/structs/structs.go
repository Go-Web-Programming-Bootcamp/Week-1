package main

import "fmt"

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	circle := new(Circle)
	fmt.Println(circle.x, circle.y, circle.r)
}
