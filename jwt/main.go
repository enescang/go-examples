//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

func main() {
	secretToken, err := createToken("github")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token: ", secretToken)

	claims, err := verifyToken(secretToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Verified: ", claims.UserID)
}

func createToken(user string) (string, error) {
	claims := jwt.MapClaims{}
	claims["UserID"] = user

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte("SECRET_KEY"))
}

func verifyToken(token string) (*MyClaims, error) {
	check, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	claims, ok := check.Claims.(*MyClaims)
	if ok && check.Valid {
		return claims, nil
	}
	return nil, err
}

type MyClaims struct {
	jwt.StandardClaims
	UserID string `json:"UserID"`
}
