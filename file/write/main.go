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
			log.Fatal("File is not exists!")
		}
	}

	myData := make([]byte, 0)
	fmt.Println(string(myData), " len", len(myData))
	myData = append(myData, 'g')
	myData = append(myData, 'i')
	myData = append(myData, 't')
	myData = append(myData, 'h')
	myData = append(myData, 'u')
	myData = append(myData, 'b')
	fmt.Println(string(myData), " len", len(myData))

	//Starting index is: 0(zero)
	totalBytes, err := file.Write(myData)
	if err != nil {
		log.Fatal("Failed to write data")
	}

	fmt.Println("Total Bytes Written: ", totalBytes)
}
