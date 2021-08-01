package validator

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type validator interface {
	Validate(interface{}) (bool, error)
}

func Validate(s interface{}) {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("my_validator")
		my_validator := getValidator(tag)
		ok, err := my_validator.Validate(v.Field(i).Interface())
		if !ok && err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(tag)
	}
}

func getValidator(tag string) validator {
	split := strings.Split(tag, ",")

	switch split[0] {
	case "string":
		sv := StringValidator{}
		fmt.Sscanf(strings.Join(split[1:], ","), "min=%d,max=%d", &sv.Min, &sv.Max)
		return sv
	}

	return Default{}
}

type Default struct{}

func (d Default) Validate(val interface{}) (bool, error) {
	return true, nil
}

type StringValidator struct {
	Min int
	Max int
}

func (sv StringValidator) Validate(val interface{}) (bool, error) {
	strLen := len(val.(string))
	if strLen < sv.Min {
		return false, fmt.Errorf("must be greater than %d", sv.Min)
	}
	if sv.Max > sv.Min && strLen >= sv.Max {
		return false, fmt.Errorf("must be less than %d", sv.Max)
	}

	return true, nil
}