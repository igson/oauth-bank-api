package interfaces

import (
	"github.com/igson/oauth-bank-api/src/domain/dto"
	"github.com/igson/oauth-bank-api/src/errors"
)

type IAuthService interface {
	Login(login dto.LoginRequest) (*string, *errors.RestErroAPI)
}
