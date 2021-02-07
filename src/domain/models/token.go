package models

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

type Claims struct {
	Username   string `json:"username"`
	Expiry     int64  `json:"exp"`
	Role       string `json:"role"`
	Authorized bool   `json:"authorized"`
}

//GenerateToken gera o token a partir dos dados do usuário
func GenerateToken() (*string, error) {

	var claims jwt.MapClaims

	//if o atributo account e customer não foram nulos criar role de usuário, caso seja, role de admin
	if l.Accounts.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}

	// 4º É invocado o método jwt.NewWithClaims pra criação das claims com dados do usuário
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 5º É retornado a numeração do token no formato String por base nas clams geradas
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))

	if err != nil {
		log.Println("Failed while signing token: " + err.Error())
		return nil, errors.New("cannot generate token")
	}
	return &signedTokenAsString, nil
}
