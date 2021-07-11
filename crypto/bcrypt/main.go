package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "SecretPass0639!"
	hash, err := generateHash(password)
	if err != nil {
		log.Fatal("Error occured: ", err)
	}
	fmt.Println("PASSWORD: ", password)
	fmt.Println("HASH: ", hash)

	fmt.Println("------------------------")

	//To get error modify the hash variable
	match, err := checkHash(password, hash)
	if err != nil {
		fmt.Println(password, "is NOT matching with", hash)
		fmt.Println("MATCH:", match)
		log.Fatal(err)
	}
	fmt.Println(password, "is matching with", hash)
	fmt.Println("MATCH:", match)
}

func generateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
