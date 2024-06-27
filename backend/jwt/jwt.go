package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("kenshilabskey")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	fmt.Println("tokenString", tokenString)
	return tokenString, nil
}

func RefreshToken(username string) (string, error) {
	newAccessToken, err := GenerateJWT(username)
	if err != nil {
		fmt.Printf("Error generating new access token: %s", err.Error())
		return "", err
	}
	return newAccessToken, nil
}
