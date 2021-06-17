// Enes Can Güneş - kodlib.com -2021
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("file.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists!")
		}
	}
	defer file.Close()

	myData := make([]byte, 0)
	myData = append(myData, 'g')
	myData = append(myData, 'i')
	myData = append(myData, 't')
	myData = append(myData, 'h')
	myData = append(myData, 'u')
	myData = append(myData, 'b')

	//Starting index is: 0(zero)
	totalBytes, err := file.Write(myData)
	if err != nil {
		log.Fatal("Failed to write data")
	}

	fmt.Println("Total Bytes Written: ", totalBytes)
}
