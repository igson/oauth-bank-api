package routers

import (
	"net/http"

	"github.com/igson/oauth-bank-api/src/controller"
	"github.com/igson/oauth-bank-api/src/datasources/banking"
	"github.com/igson/oauth-bank-api/src/domain/repository"
	"github.com/igson/oauth-bank-api/src/domain/service"
)

var (
	repo            = repository.NewUserRepository(banking.GetDbClient())
	authService     = service.NewAuthService(repo)
	loginController = controller.NewLoginController(authService)
)

var rotasLogin = []Rota{

	{
		URI:                "/oauth/login",
		Metodo:             http.MethodPost,
		Funcao:             loginController.Login,
		RequerAutenticacao: false,
		Name:               "Login",
	},
}
