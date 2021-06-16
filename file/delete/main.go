// Enes Can Güneş - kodlib.com -2021
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Remove("file.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists.")
		}
		log.Fatal(err)
	}
	fmt.Println("File deleted.")

}
