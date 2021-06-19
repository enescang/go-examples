//Enes Can Güneş - kodlib.com - 2021
//Tutorial: Jake Wright, https://www.youtube.com/watch?v=LvgVSSpwND8
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		write("Hello Go")
		wg.Done()
	}()

	wg.Wait()
}

func write(str string) {
	for index := 0; index < 6; index++ {
		fmt.Println("Index: ", index, " string: ", str)
		time.Sleep(500 * time.Millisecond)
	}
}
