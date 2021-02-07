package models

import "github.com/igson/banking/src/domain/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	Id              int64   `db:"transaction_id"`
	AccountId       int64   `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.Id,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
