package interfaces

import (
	"github.com/igson/banking/src/domain/dto"
	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
)

//IAccountService interface pra implementação do servico
type IAccountService interface {
	CriarConta(string) (*models.Account, *errors.RestErroAPI)
	RegistrarTransacao(transaction dto.TransactionRequest) (*dto.TransactionResponse, *errors.RestErroAPI)
	BuscarContaPorID(accountId int64) (*dto.AccountResponse, *errors.RestErroAPI)
}
