package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// We need to rand.Seed
	// Please, check here for more information:
	// https://tour.golang.org/basics/1 and
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value

	rand.Seed(time.Now().UnixNano())

	//random int
	println("Random Int: ", rand.Int())

	//random int between 0-(n)
	println("Random Int Between 0-338: ", rand.Intn(338))

	//random int between (n1)-(n2) [min, max)
	min, max := 1, 2
	println("Random Int Between 100-200: ", rand.Intn(max-min)+min)

	//random float32 0-1
	fmt.Println("Random float64 0-1: ", rand.Float32())

	

}
