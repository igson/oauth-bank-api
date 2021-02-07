package routers

import (
	"net/http"

	"github.com/igson/banking/src/controller"
	"github.com/igson/banking/src/datasources/banking"
	"github.com/igson/banking/src/domain/repository"
	"github.com/igson/banking/src/domain/service"
)

var (
	accountRepo       = repository.NewAccountRepository(banking.GetDbClient())
	accountService    = service.NewAccountService(accountRepo)
	accountController = controller.NewAccountController(accountService)
)

var rotasAccounts = []Rota{
	{
		URI:                "/accounts/:account_id",
		Metodo:             http.MethodGet,
		Funcao:             accountController.BuscarContaPorID,
		RequerAutenticacao: true,
		Name:               "BuscarContaDeCliente",
	},
	{
		URI:                "/accounts",
		Metodo:             http.MethodPost,
		Funcao:             accountController.RegistrarTransacao,
		RequerAutenticacao: true,
		Name:               "ListarContas",
	},
	{
		URI:                "/customers/:customer_id/account/:account_id",
		Metodo:             http.MethodPost,
		Funcao:             accountController.RegistrarTransacao,
		RequerAutenticacao: true,
		Name:               "RealizarTransacao",
	},
}
