package dto

type CustomerResponse struct {
	Id          int64  `json:"customer_id" xml:"customer_id"`
	Name        string `json:"full_name" xml:"full_name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateofBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}
