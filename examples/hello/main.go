package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 10; i++ {
		switch {
		case i == 5:
			break outer
		default:
			fmt.Println(i)
		}
	}
}
