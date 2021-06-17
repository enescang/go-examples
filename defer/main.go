//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"fmt"
	"log"
	"os"
)

//Blog: https://blog.golang.org/defer-panic-and-recover

func main() {
	fileClose()
	LIFO()
}

//Defer is working based on "Last In First Out" logic
func LIFO() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}

	//with 	  defer: 9876543210
	//without defer: 0123456789
}

//Please, visit this Blog to more understandable example
//https://blog.golang.org/defer-panic-and-recover
func fileClose() {
	file, err := os.OpenFile("file.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
