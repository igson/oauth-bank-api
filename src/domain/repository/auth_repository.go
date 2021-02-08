package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/igson/oauth-bank-api/src/domain/models"
	"github.com/igson/oauth-bank-api/src/errors"
	"github.com/igson/oauth-bank-api/src/interfaces"
)

const (
	queryLoginUser = "SELECT username, role FROM users WHERE username = ? and password = ?"
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

func (u *userRepository) Login(usuario string, senha string) (*models.Login, *errors.RestErroAPI) {

	var login models.Login
	fmt.Printf("Usuário: %s e Senha: %s", usuario, senha)
	err := u.client.Get(&login, queryLoginUser, usuario, senha)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err.Error())
			return nil, errors.NewNotFoundErro("Usuário ou seha inválido")
		} else {
			fmt.Println(err.Error())
		}
	}

	return &login, nil

}
