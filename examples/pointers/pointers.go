package main

import "fmt"

func zero(num *int) {
	*num = 0
}

func main() {
	num := 5
	zero(&num)
	fmt.Println(num)
}
