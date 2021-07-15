package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 5
	arr[1] = 10
	arr[2] = 17
	arr[3] = 20
	arr[4] = 25

	var total int
	for _, value := range arr {
		total += value
	}

	fmt.Println(float64(total) / float64(len(arr)))
}
