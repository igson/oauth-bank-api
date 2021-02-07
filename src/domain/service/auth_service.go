package service

import (
	"github.com/igson/oauth-bank-api/src/domain/autenticacao"
	"github.com/igson/oauth-bank-api/src/domain/dto"
	"github.com/igson/oauth-bank-api/src/domain/models"
	"github.com/igson/oauth-bank-api/src/errors"
	"github.com/igson/oauth-bank-api/src/interfaces"
)

type authService struct {
	userRepository interfaces.IUserRepository
}

//NewAuthService acesso ao repositório
func NewAuthService(userRepository interfaces.IUserRepository) interfaces.IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (a *authService) Login(request dto.LoginRequest) (*string, *errors.RestErroAPI) {

	user, erro := a.userRepository.Login(request.Username, request.Password)

	if erro != nil {
		return nil, erro
	}

	login := models.Login{Username: user.Username, Role: user.Role}

	token, erro := autenticacao.GerarToken(login)

	if erro != nil {
		return nil, erro
	}

	return token, nil

}
