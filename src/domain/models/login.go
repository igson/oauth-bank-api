package models

//Login armazenar dados de login
type Login struct {
	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"password"`
	Role       string `db:"role" json:"-"`
	CustomerId string `db:"customer_id" json:"-"`
	Created_on string `-`
}
