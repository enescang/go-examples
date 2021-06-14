// Enes Can Güneş - kodlib.com -2021
package main

import "fmt"

func main() {

	str := []string{"12", "34", "6300"}

	for i := 0; i < len(str); i++ {
		fmt.Println("Index: ", i, " Value: ", str[i])
	}

	fmt.Println("------------------------")

	//It is similar to foreach loop like in javascript
	for index, value := range str {
		fmt.Println("Index: ", index, " Value: ", value)
	}

	fmt.Println("------------------------")

	//Simple while loop, like in javascript, php etc.
	//We can use condition in for loop
	var index int
	index = 0
	for index <= 10 {
		fmt.Println("for loop is running. Current index: ", index)
		index = index + 1
	}

	//Infinite loop
	// for {
	// 	fmt.Println("It will run forever. To stop: CTRL+C")
	// }
}
