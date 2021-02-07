package interfaces

import (
	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
)

//IAccountRepository interface pra implementação
type IAccountRepository interface {
	CriarConta(string) (*models.Account, *errors.RestErroAPI)
	RegistrarTransacao(transaction models.Transaction) (*models.Transaction, *errors.RestErroAPI)
	BuscarContaPorID(accountId int64) (*models.Account, *errors.RestErroAPI)
}
