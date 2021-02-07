package banking

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

//GetDbClient importar a cliente de conexão
func GetDbClient() *sqlx.DB {
	dbUser := "tce"
	dbPasswd := "tce"
	dbAddr := "localhost"
	dbPort := "3306"
	dbName := "banking"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName))
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	fmt.Println("Conexção com o banco de dados pagbank estabelecida")
	return client
}
