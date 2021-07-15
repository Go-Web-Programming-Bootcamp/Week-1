package main

import "fmt"

func main() {
	var x = map[string]int{
		"Steven": 1,
		"Hans":   2,
	}
	fmt.Println(x)
	delete(x, "Steven")
	fmt.Println(x)
}
