package main

import "fmt"

func increasingNumberGenerator() func() int {
	x := 0
	increase := func() int {
		x++
		return x
	}
	return increase
}

func main() {
	generator := increasingNumberGenerator()
	fmt.Println(generator())
}
