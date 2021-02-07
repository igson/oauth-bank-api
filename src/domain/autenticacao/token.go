package autenticacao

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/igson/banking/src/errors"
)

const (
	TOKEN_DURATION     = time.Hour
	HMAC_SAMPLE_SECRET = "hmacSampleSecret"
)

//CriarToken gerar token
func CriarToken(usuarioId uint64) (*string, *errors.RestErroAPI) {
	fmt.Println("Criando Token", usuarioId)

	permissoes := jwt.MapClaims{
		"authorized": true,
		"usuarioId":  usuarioId,
		"exp":        time.Now().Add(TOKEN_DURATION).Unix(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	signedTokenAsString, err := newToken.SignedString([]byte(HMAC_SAMPLE_SECRET))

	if err != nil {
		log.Println("Failed while signing token: " + err.Error())
		return nil, errors.NewValidationError("Problema na geração do token")
	}

	fmt.Println("Number Token: ", signedTokenAsString)

	return &signedTokenAsString, nil

}

//ValidarToken validação do token de acesso
func ValidarToken(r *http.Request) *errors.RestErroAPI {

	tokenString := extrairToken(r)

	if tokenString == "" {
		return errors.NewValidationError("Sem chave de token válido")
	}

	token, erro := retornarChaveDeVerificacao(tokenString)

	if erro != nil {
		return errors.NewValidationError(erro.Message)
	}

	// retorn ok se assinatura do token for válida
	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return nil
	}

	return errors.NewValidationError("Token Inválido")
}

func extrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	//Bearer dividi a string em 2 partes pra saber se o token veio informado na requisição. Ex.: Bearer 102A0DI032K2DM29.0321k121k201
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""

}

//ExtrairUsuarioID extrair o ID do usuário
func ExtrairUsuarioID(r *http.Request) (int64, *errors.RestErroAPI) {

	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})

	if erro != nil {
		return 0, errors.NewValidationError("Erro no token!")
	}
	//Bearer dividi a string em 2 partes pra saber se o token veio informado na requisição. Ex.: Bearer 102A0DI032K2DM29.0321k121k201

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		usuarioId, erro := strconv.ParseInt(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)

		if erro != nil {
			return 0, errors.NewValidationError("Você não tem permissão pra essa operação.")
		}

		return usuarioId, nil

	}

	return 0, nil

}

// verifica assinatura do token enviado na requisição pra saber se está com assinatura correta
func retornarChaveDeVerificacao(tokenString string) (*jwt.Token, *errors.RestErroAPI) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})

	if err != nil {
		return nil, errors.NewValidationError("Método de assinatura inesperado!")
	}

	return token, nil

}
