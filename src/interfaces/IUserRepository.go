package interfaces

import (
	"github.com/igson/banking/src/errors"
)

//IUserRepository interface metodos
type IUserRepository interface {
	Login(usuario string, senha string) *errors.RestErroAPI
}
