package models

//Login armazenar dados de login
type Login struct {
	Username string `db:"username"`
	Role     string `db:"role"`
}
