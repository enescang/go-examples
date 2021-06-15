//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	myJson := `
		[
			{
				"name": "github",
				"email": "test@tes.com"
			},
			{
				"name": "google",
				"email": "test@tes.com"
			}
		]
	`
	//If our json has multiple element use [] to get all of them
	var arrPerson []Person

	err := json.Unmarshal([]byte(myJson), &arrPerson)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println(arrPerson)
}

type Person struct {
	Name  string
	Email string
}
