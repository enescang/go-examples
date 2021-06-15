//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"encoding/json"
	"log"
)

func main() {
	jsonToStruct()
	structToJson()
}

func jsonToStruct() {
	myJson := `
	{
		"name": "github",
		"email": "test@test.com"
	}
	`
	var model myModel
	//or you can use map
	//var model map[string]interface{}

	//give the pointer of model with (&)
	//otherwise Unmarshal will return: json: Unmarshal(non-pointer main.myModel)
	err := json.Unmarshal([]byte(myJson), &model)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Println(model)
}

func structToJson() {
	var person myModel
	person.Name = "google"
	person.Email = "test1@test1.com"

	output, err := json.Marshal(person)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	//convert []byte(...) to string
	var jsonData string = string(output)
	log.Println(jsonData)
}

type myModel struct {
	Name  string
	Email string
}
