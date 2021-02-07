package app

import (
	"fmt"
	"os"

	"github.com/igson/banking/src/errors"
)

//StartApplication - dá início a aplicação
func StartApplication() {

	if erro := enviroumentsChecks(); erro != nil {
		panic(erro.Message)
	} else {
		rota := GerarRotas()
		rota.Run(":8080")
	}

}

func enviroumentsChecks() *errors.RestErroAPI {

	envProps := []string{
		//"SERVER_ADDRESS",
		//"SERVER_PORT",
		//"DB_USER",
		//"DB_PASSWD",
		//"DB_ADDR",
		//	"DB_PORT",
		//"DB_NAME",
		"user_db_username",
	}

	for _, k := range envProps {
		if os.Getenv(k) == "" {
			fmt.Printf(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
			return errors.NewInternalServerError(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}

	return nil

}
