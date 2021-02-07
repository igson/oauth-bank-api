package models

//Login armazenar dados de login
type Login struct {
	Username string `db:"username" json:"username"`
	Role     string `db:"role" json:"role"`
}
