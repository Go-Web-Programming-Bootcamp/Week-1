package main

import "fmt"

func main() {
	var arr = []int{2, 5, 10, 11}

	var total int
	for _, value := range arr {
		total += value
	}

	fmt.Println(float64(total) / float64(len(arr)))
}
