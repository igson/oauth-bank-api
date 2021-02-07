package dto

import (
	"github.com/igson/banking/src/errors"
)

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

type TransactionRequest struct {
	Id              int64   `json:"account_id"`
	Amount          float64 `json:"amount,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	CustomerId      int64   `json:"-,omitempty"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errors.RestErroAPI {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errors.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errors.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}
