package main

import "fmt"

func fillWithOne(arr *[10]int) {
	for ind := range *arr {
		arr[ind] = 1
		// or (*arr)[ind] = 1
	}
}

func fillWithOneV2(arr [10]int) {
	for ind := range arr {
		arr[ind] = 1
	}
}

func main() {
	arr := [10]int{}
	fillWithOne(&arr)
	fmt.Println(arr)

	arr = [10]int{}
	fmt.Println(arr)
	fillWithOne(&arr)
	fmt.Println(arr)
}
