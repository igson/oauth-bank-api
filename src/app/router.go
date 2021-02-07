package app

import (
	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/app/routers"
	"github.com/igson/banking/src/controller"
	"github.com/igson/banking/src/datasources/banking"
	"github.com/igson/banking/src/domain/repository"
	"github.com/igson/banking/src/domain/service"
)

var (
	customerRepo       = repository.NewCustomerRepository(banking.GetDbClient())
	customerService    = service.NewCustomerService(customerRepo)
	customerController = controller.NewCustomerController(customerService)

	repo            = repository.NewUserRepository(banking.GetDbClient())
	uService        = service.NewUserService(repo)
	loginController = controller.NewLoginController(uService)
)

//GerarRotas iniciar rotas
func GerarRotas() *gin.Engine {
	rota := gin.Default()

	//r.Use(middlewares.Logger())
	//r.Use(middlewares.Autenticar(customerController.GetAllCustomers))
	//r.GET("/customers", middlewares.Autenticar(), customerController.GetAllCustomers)
	//r.GET("/customers/:customer_id", customerController.GetCustomer)
	//r.POST("/login", loginController.Login)
	return routers.Configurar(rota)
	//return r
}
