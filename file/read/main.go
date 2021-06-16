// Enes Can Güneş - kodlib.com -2021
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./file.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("File Data: ", string(data))
}
