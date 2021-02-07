package routers

import (
	"net/http"

	"github.com/igson/banking/src/controller"
	"github.com/igson/banking/src/datasources/banking"
	"github.com/igson/banking/src/domain/repository"
	"github.com/igson/banking/src/domain/service"
)

var (
	repo            = repository.NewUserRepository(banking.GetDbClient())
	uService        = service.NewUserService(repo)
	loginController = controller.NewLoginController(uService)
)

var rotasLogin = []Rota{

	{
		URI:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             loginController.Login,
		RequerAutenticacao: false,
		Name:               "Login",
	},
}
