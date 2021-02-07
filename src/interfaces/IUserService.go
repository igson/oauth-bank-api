package interfaces

import (
	"github.com/igson/banking/src/errors"
)

type IUserService interface {
	Login(usuario string, senha string) *errors.RestErroAPI
}
