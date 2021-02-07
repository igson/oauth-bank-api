package dto

type TransactionResponse struct {
	TransactionId   int64   `json:"transaction_id"`
	AccountId       int64   `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
