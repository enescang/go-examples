// Enes Can Güneş - kodlib.com -2021
package main

import (
	"fmt"
	"log"
	"os"
)

var (
	file *os.File
	err  error
)

func main() {
	//I just want to use different way of defining variables
	//We can use also like this:
	//file, err := os.Create("file.txt")
	file, err = os.Create("file.txt")
	if err != nil {
		log.Fatal("Failed to create file")
	}

	fmt.Println("File created:", file.Name())
}
