//Enes Can Güneş - kodlib.com - 2021
//Tutorial: Jake Wright, https://www.youtube.com/watch?v=LvgVSSpwND8
package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello Go")
	write("Hello JavaScript")
}

func write(str string) {
	var index int
	for {
		fmt.Println("Index: ", index, " string: ", str)
		time.Sleep(500 * time.Millisecond)
		index++
	}
}
