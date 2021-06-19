//Enes Can Güneş - kodlib.com - 2021
//Tutorial: Jake Wright, https://www.youtube.com/watch?v=LvgVSSpwND8
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	go write("Hello Go", c)

	//Single message
	// msg := <-c
	// fmt.Println(msg)

	for message := range c {
		fmt.Println(message, "---")
	}

}

func write(str string, c chan string) {
	for index := 0; index < 6; index++ {
		var message string = "Index: " + strconv.Itoa(index) + " string: " + str
		c <- message
		time.Sleep(500 * time.Millisecond)
	}

	close(c)
}
