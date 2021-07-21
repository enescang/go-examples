package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello Web Assembly")
	c := make(chan interface{})
	js.Global().Set("add", js.FuncOf(add))

	<-c
}

func add(this js.Value, p []js.Value) interface{} {
	sum := p[0].Int() + p[1].Int()
	return js.ValueOf(sum)
}
