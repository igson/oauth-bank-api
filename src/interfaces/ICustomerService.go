package interfaces

import (
	"github.com/igson/banking/src/domain/dto"
	"github.com/igson/banking/src/errors"
)

type ICustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errors.RestErroAPI)
	GetCustomer(int64) (*dto.CustomerResponse, *errors.RestErroAPI)
}
