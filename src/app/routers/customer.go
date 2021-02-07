package routers

import (
	"net/http"

	"github.com/igson/banking/src/controller"
	"github.com/igson/banking/src/datasources/banking"
	"github.com/igson/banking/src/domain/repository"
	"github.com/igson/banking/src/domain/service"
)

var (
	customerRepo       = repository.NewCustomerRepository(banking.GetDbClient())
	customerService    = service.NewCustomerService(customerRepo)
	customerController = controller.NewCustomerController(customerService)
)

const (
	GRUPO_CLIENTES_URI = "clientes"
)

var rotasCustomer = []Rota{

	{
		URI:                "/customers",
		Metodo:             http.MethodGet,
		Funcao:             customerController.GetAllCustomers,
		RequerAutenticacao: true,
		Name:               "ListarClientes",
		Grupo:              GRUPO_CLIENTES_URI,
		//Permissoes:         []gin.HandlerFunc{middlewares.Logger(), middlewares.Autenticar(), customerController.GetAllCustomers},
	},
	{
		URI:                "/customers/:customer_id",
		Metodo:             http.MethodGet,
		Funcao:             customerController.GetCustomer,
		RequerAutenticacao: true,
		Name:               "BuscarCliente",
		Grupo:              GRUPO_CLIENTES_URI,
		//Permissoes:         []gin.HandlerFunc{middlewares.Logger(), middlewares.Autenticar(), customerController.GetAllCustomers},
	},
}
