package models

import "github.com/igson/banking/src/domain/dto"

type Account struct {
	Id          int64   `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

//AccountRepository interface de implementação pra acesso ao metodos
/* type AccountRepository interface {
	CriarConta(string) (*Account, *errors.RestErroAPI)
	RegistrarTransacao(transaction Transaction) (*Transaction, *errors.RestErroAPI)
	BuscarContaPorID(accountId string) (*Account, *errors.RestErroAPI)
}
*/

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

func (a Account) ToAccountResponseDTO() dto.AccountResponse {
	return dto.AccountResponse{
		Id:          a.Id,
		CustomerId:  a.CustomerId,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
}
