package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := []string{"zero", "one", "two", "three"}
	randomIndex := rand.Intn(len(arr))

	fmt.Println("Array: ", arr)
	fmt.Println("Selected Index: ", randomIndex)
	fmt.Println("Selected Element: ", arr[randomIndex])
}
