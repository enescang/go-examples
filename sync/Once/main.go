//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"fmt"
	"sync"
)

var singleton sync.Once
var dbInstance string

func main() {
	//Let's call DatabaseConnection function multiple times to see differences
	for i := 0; i < 10; i++ {
		DatabaseConnection()
	}
	//sync.Once runs exactly one time
	//and if you need to create singleton pattern in your application
	//sync.Once is better option then traditional way.
	//Please check: https://pkg.go.dev/sync#Once
}

func DatabaseConnection() string {
	singleton.Do(func() {
		fmt.Println("Db Connection Started")
		dbInstance = "connect_your_database"
		fmt.Println("Db Connected")
	})
	fmt.Println("Other Stuff")
	return dbInstance
}
