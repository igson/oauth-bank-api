package interfaces

import (
	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
)

//ICustomerRepository acesso aos metodos de implementação
type ICustomerRepository interface {
	FindAll(status string) ([]models.Customer, *errors.RestErroAPI)
	ById(int64) (*models.Customer, *errors.RestErroAPI)
}
