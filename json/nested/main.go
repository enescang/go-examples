//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"encoding/json"
	"log"
)

//we can separate each "struct"

//Please, check here if you curious about "tags"
//https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go
type Person struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Interests []struct {
		Language string
	} `json:"interests"`
	Address struct {
		HomeAddress string
		WorkAddress string
	} `json:"address"`
}

func main() {
	jsonToStruct()
	structToJson()
}

func jsonToStruct() {
	var myJson string = `
	{
		"name": "Enes",
		"email": "test@test.com",
		"interests": [
			{
				"language": "Go"
			},
			{
				"language": "NodeJs"
			}
		],
		"address": {
			"homeAddress": "At world",
			"workAddress": "At 01"
		}
	}	
	`

	var person1 Person
	err := json.Unmarshal([]byte(myJson), &person1)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("Person's name: " + person1.Name)
	log.Println("Person's email: " + person1.Email)
	log.Println("Person's interests: ", person1.Interests)
	log.Println("Person's address: ", person1.Address)
	log.Println(person1)
}

func structToJson() {
	var person1 Person
	person1.Name = "Enes"
	person1.Email = "test@test.com"
	person1.Interests = []struct{ Language string }{{"Go Programming Language"}, {"Nodejs"}}
	person1.Address.HomeAddress = "My home"
	person1.Address.WorkAddress = "My comp"
	//or
	//person1.Address = struct{HomeAddress string; WorkAddress string}{"My home","My comp"}
	output, err := json.Marshal(person1)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("person1 : ", string(output))
}
