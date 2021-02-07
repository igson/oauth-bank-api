package banking

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersDbUsername = "user_db_username"
	usersDbPassword = "user_db_password"
	usersDbHost     = "user_db_host"
	usersDbSchema   = "user_db_schema"
)

var (
	//Conexao conexão com o banco de dados
	Conexao  *sql.DB
	erro     error
	username = os.Getenv(usersDbUsername)
	password = os.Getenv(usersDbPassword)
	host     = os.Getenv(usersDbHost)
	schema   = os.Getenv(usersDbSchema)
)

func init() {

	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema))
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)

	Conexao, erro = sql.Open("mysql", dataSourceName)

	if erro != nil {
		panic(erro)
	}

	if erro = Conexao.Ping(); erro != nil {
		panic(erro)
	}

	log.Println("Conexão com banco de dados realizada com sucesso")

}
