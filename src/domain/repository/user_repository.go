package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

const (
	queryLoginUser = "SELECT * FROM users WHERE username = ? and password = ?"
)

type userRepository struct {
	client *sqlx.DB
}

//NewUserRepository acesso ao repositório
func NewUserRepository(dbClient *sqlx.DB) interfaces.IUserRepository {
	return &userRepository{
		client: dbClient,
	}
}

func (u *userRepository) Login(usuario string, senha string) *errors.RestErroAPI {

	var l models.Login

	err := u.client.Get(&l, queryLoginUser, usuario, senha)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err.Error())
			return errors.NewNotFoundErro("Usuário ou seha inválido")
		} else {
			fmt.Println(err.Error())
			return errors.NewUnexpectedError("Unexpected database error")
		}
	}

	return nil

}
