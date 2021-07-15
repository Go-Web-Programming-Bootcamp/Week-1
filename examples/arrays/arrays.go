package main

import "fmt"

func main() {
	var arr1 = []int64{1, 2, 3}
	fmt.Println(arr1)
	arr1 = append(arr1, 4, 5, 6, 7)
	fmt.Println(arr1)

	var arr2 = make([]int64, 3)
	arr2[0] = 8
	arr2[1] = 9
	arr1 = append(arr1, arr2[0:1]...)
	fmt.Println(arr1)
}
