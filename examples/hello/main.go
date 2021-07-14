package main

import "fmt"

func main() {
	str := "Hello World!"
	for ind, char := range str {
		fmt.Println(ind, char) // Print the index and the character (in integer)
	}
}
