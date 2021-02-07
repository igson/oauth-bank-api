package autenticacao

import (
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

type Token struct {
	Username string `json:"usernameId"`
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

func claimsForUser(login models.Login) jwt.MapClaims {
	return jwt.MapClaims{
		"usernameId": login.Username,
		"role":       login.Role,
		"exp":        time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
