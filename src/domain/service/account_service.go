package service

import (
	"fmt"
	"time"

	"github.com/igson/banking/src/domain/dto"
	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

const dbTSLayout = "2006-01-02 15:04:05"

//AccountService account service
type accountService struct {
	accountRepository interfaces.IAccountRepository
}

//NewAccountService acesso ao repositório
func NewAccountService(accountRepository interfaces.IAccountRepository) interfaces.IAccountService {
	return &accountService{
		accountRepository: accountRepository,
	}
}

func (s *accountService) CriarConta(account string) (*models.Account, *errors.RestErroAPI) {
	fmt.Println("Camada de serviço")
	s.accountRepository.CriarConta(account)
	return nil, nil
}

func (s *accountService) RegistrarTransacao(transaction dto.TransactionRequest) (*dto.TransactionResponse, *errors.RestErroAPI) {

	if erro := transaction.Validate(); erro != nil {
		return nil, erro
	}

	if transaction.IsTransactionTypeWithdrawal() {
		account, err := s.accountRepository.BuscarContaPorID(transaction.Id)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(transaction.Amount) {
			return nil, errors.NewValidationError("Insufficient balance in the account")
		}
	}

	t := models.Transaction{
		AccountId:       transaction.Id,
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	newTransaction, erro := s.accountRepository.RegistrarTransacao(t)

	if erro != nil {
		fmt.Println("Erro ----->", erro)
		return nil, erro
	}

	respTransactionDto := newTransaction.ToDto()

	return &respTransactionDto, nil

}

func (s *accountService) BuscarContaPorID(accountId int64) (*dto.AccountResponse, *errors.RestErroAPI) {

	if accountId == 0 {
		return nil, errors.NewBadRequestError("ID da conta não informado")
	}

	if account, erro := s.accountRepository.BuscarContaPorID(accountId); erro != nil {
		return nil, erro
	} else {
		response := account.ToAccountResponseDTO()
		return &response, nil
	}

}
