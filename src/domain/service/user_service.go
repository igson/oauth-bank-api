package service

import (
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

type userService struct {
	userRepository interfaces.IUserRepository
}

//NewUserService acesso ao repositório
func NewUserService(userRepository interfaces.IUserRepository) interfaces.IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Login(usuario string, senha string) *errors.RestErroAPI {

	erro := u.userRepository.Login(usuario, senha)

	if erro != nil {
		return erro
	}

	return nil
}
