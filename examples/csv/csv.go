package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, error := os.Open("state_table.csv")
	if error != nil {
		panic(error)
	}
	csvReader := csv.NewReader(file)
	content, error := csvReader.ReadAll()
	if error != nil {
		panic(error)
	}
	fmt.Println(content)
}
