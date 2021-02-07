package dto

import (
	"strings"

	"github.com/igson/banking/src/errors"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

//Validate validação de campos do account
func (ar NewAccountRequest) Validate() *errors.RestErroAPI {
	if ar.Amount < 5000 {
		return errors.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(ar.AccountType) != "saving" && strings.ToLower(ar.AccountType) != "checking" {
		return errors.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
