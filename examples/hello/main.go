package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
