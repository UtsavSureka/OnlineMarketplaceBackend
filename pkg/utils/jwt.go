package utils

import (
	"Ecomm/pkg/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWTSecretKey = "HelloWorld"

type Claims struct {
	Id       int
	Username string `json:"username"`
	IsAdmin  bool   `json:isAdmin`
	jwt.StandardClaims
}

func GenerateJwtToken(user models.User) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Id:       user.Id,
		Username: user.User_name,
		IsAdmin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Create encoded token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Generate token string using token secret key

	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", errors.New("error generating token string")
	}

	return tokenString, nil

}

func ValidateJwtToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
