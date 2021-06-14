// Enes Can Güneş - kodlib.com -2021
package main

import "fmt"

func main() {
	//Declare a map (empty) v1
	languages := make(map[string]string)

	//Declare a map (empty) v2
	// languages := map[string]string{}

	//Declare a map (initial values) v3
	// languages := map[stirng]string{
	// 	"GO": "Golang",
	// 	"JS": "JavaScript",
	// }

	//add new element
	languages["GO"] = "Golang"
	languages["JS"] = "JavaScript"
	languages["C#"] = "Csharp"

	//get element
	golang := languages["GO"]
	valueJava, isExists := languages["JAVA"]

	fmt.Println("I am coding " + golang + " with love")
	fmt.Println("Is JAVA exists? :", isExists, valueJava)

	//delete element
	delete(languages, "C#")

	//length of map
	fmt.Println("There are", len(languages), "languages")

	fmt.Println(languages)

	fmt.Println("------------------------")

	//Mixed type map any-any
	mixedMap := map[interface{}]interface{}{
		"GO":    "Golang",
		2007:    "designed at",
		2009.11: "released at",
	}

	fmt.Println(mixedMap)

	fmt.Println("------------------------")

	// string-struct
	m := make(map[string]Model)

	//Can not assign
	//Please, chech this issue: https://github.com/golang/go/issues/3117
	//m["FIRST"].x = "Hello Go" //=> It is wrong

	//use like this
	temp := m["FIRST"]
	temp.x = "Hello Go"
	temp.y = 721
	m["FIRST"] = temp

	fmt.Println(m)
}

type Model struct {
	x string
	y int
}
