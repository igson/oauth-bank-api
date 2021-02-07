package dto

type AccountResponse struct {
	Id          int64   `json:"account_id"`
	CustomerId  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}
