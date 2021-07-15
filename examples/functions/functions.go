package main

import "fmt"

func main() {
	arr := []float64{1, 2, 3, 4, 5}
	average := func(arr []float64) float64 {
		var total float64
		for _, val := range arr {
			total += val
		}
		return total / float64(len(arr))
	}
	fmt.Println(average(arr))
}
