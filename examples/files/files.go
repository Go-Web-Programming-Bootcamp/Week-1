package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, error := os.Open("./hello.txt")
	defer file.Close()
	if error != nil {
		log.Fatalln("There is an error in opening the intended file")
	}

	text, error := ioutil.ReadAll(file)
	if error != nil {
		log.Fatalln(("There is an error in reading the file"))
	}
	fmt.Println(string(text))
}
