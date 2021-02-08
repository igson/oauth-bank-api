package autenticacao

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/igson/oauth-bank-api/src/domain/models"
	"github.com/igson/oauth-bank-api/src/errors"
)

const (
	TOKEN_DURATION     = time.Hour
	HMAC_SAMPLE_SECRET = "hmacSampleSecret"
)

type Claims struct {
	Username string `json:"username"`
	Expiry   int64  `json:"exp"`
	Role     string `json:"role"`
}

//GerarToken gera o token a partir dos dados do usuário
func GerarToken(login models.Login) (*string, *errors.RestErroAPI) {

	claims := claimsForUser(login)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))

	if err != nil {
		log.Println("Failed while signing token: " + err.Error())
		return nil, errors.NewValidationError("Erro inesperado na geração do token")
	}

	return &signedTokenAsString, nil

}

func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, error) {

	bytes, err := json.Marshal(mapClaims)

	if err != nil {
		return nil, err
	}

	var c Claims

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (c *Claims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	fmt.Println("User Name Claims ----> ", c.Username)
	fmt.Println("Router Name MAPS ----> ", urlParams["routeName"])
	fmt.Println("User Name MAPS ----> ", urlParams["username"])
	fmt.Println("User Type ----> ", c.Role)
	if c.Username != urlParams["username"] {
		return false
	}
	return true
}

func claimsForUser(login models.Login) jwt.MapClaims {
	return jwt.MapClaims{
		"username": login.Username,
		"role":     login.Role,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
