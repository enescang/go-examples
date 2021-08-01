//Enes Can Güneş - kodlib.com - 2021
package main

import (
	validator "./validator"
)

type TestTag struct {
	Name string `my_validator:"string,min=1,max=4"`
}

func main() {
	validator.Validate(TestTag{Name: "database"})
}
