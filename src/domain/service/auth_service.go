package service

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/igson/oauth-bank-api/src/domain/autenticacao"
	"github.com/igson/oauth-bank-api/src/domain/dto"
	"github.com/igson/oauth-bank-api/src/domain/models"
	"github.com/igson/oauth-bank-api/src/errors"
	"github.com/igson/oauth-bank-api/src/interfaces"
)

type authService struct {
	userRepository interfaces.IUserRepository
	permissoes     models.RolePermissions
}

//NewAuthService acesso ao repositório
func NewAuthService(userRepository interfaces.IUserRepository, permissoes models.RolePermissions) interfaces.IAuthService {
	return &authService{
		userRepository: userRepository,
		permissoes:     permissoes,
	}
}

func (a *authService) Login(request dto.LoginRequest) (*string, *errors.RestErroAPI) {

	user, erro := a.userRepository.Login(request.Username, request.Password)

	if erro != nil {
		return nil, erro
	}

	login := models.Login{Username: user.Username, Role: user.Role}

	fmt.Println(login)

	token, erro := autenticacao.GerarToken(login)

	if erro != nil {
		return nil, erro
	}

	return token, nil

}

func (a *authService) Verify(urlParams map[string]string) (bool, *errors.RestErroAPI) {
	// convert the string token to JWT struct
	fmt.Println("Token Parameter Value -----> ", urlParams["token"])
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		fmt.Println("Erro ao emitir o token.")
		errParse := errors.NewInternalServerError(err.Error())
		return false, errParse
	} else {

		if jwtToken.Valid {
			// type cast the token claims to jwt.MapClaims
			mapClaims := jwtToken.Claims.(jwt.MapClaims)
			// converting the token claims to Claims struct
			if claims, err := autenticacao.BuildClaimsFromJwtMapClaims(mapClaims); err != nil {
				fmt.Println("Erro ao emitir capturar o claims")
				errParseClaim := errors.NewInternalServerError(err.Error())
				return false, errParseClaim
			} else {
				/* if Role if user then check if the account_id and customer_id
				   coming in the URL belongs to the same token
				*/

				if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
					fmt.Println("Não Tem Autorização")
					return false, nil
				}
				// verify of the role is authorized to use the route
				isAuthorized := a.permissoes.IsAuthorizedFor(claims.Role, urlParams["routeName"])

				return isAuthorized, nil

			}

		} else {
			errToken := errors.NewValidationError("Token inválido")
			return false, errToken
		}

	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(autenticacao.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		log.Println("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
