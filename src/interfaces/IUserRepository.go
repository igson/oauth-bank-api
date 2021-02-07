package interfaces

import (
	"github.com/igson/oauth-bank-api/src/domain/models"
	"github.com/igson/oauth-bank-api/src/errors"
)

//IUserRepository interface metodos
type IUserRepository interface {
	Login(usuario string, senha string) (*models.Login, *errors.RestErroAPI)
}
