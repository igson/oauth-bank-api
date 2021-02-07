package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/middlewares"
)

/* var clienteController = controller.NewClienteController(service.NewClienteService())

func (r Rotas) registrarRotasCliente(rota *gin.RouterGroup) {
	rota.GET("/clientes/:cliente_id", clienteController.GetByID)
	rota.POST("/clientes", clienteController.Create)
}
*/

//Rota objeto de configuração das rotas
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(g *gin.Context)
	RequerAutenticacao bool
	Name               string
	Permissoes         []gin.HandlerFunc
	Grupo              string
}

//Configurar carregar lista de rotas
func Configurar(r *gin.Engine) *gin.Engine {
	//rotas := rotasAccounts
	var rotas []Rota
	rotas = append(rotas, rotasCustomer...)
	rotas = append(rotas, rotasLogin...)
	//func(httpMethod string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	//func (h *Handler) Me(c *gin.Context) {}
	// g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
	//func AuthUser(s model.TokenService) gin.HandlerFunc { middleware
	r.Use(middlewares.Logger())
	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.Handle(rota.Metodo, rota.URI, middlewares.Autenticar(), rota.Funcao)
			fmt.Println(fmt.Sprintf("Rota %s requer autenticação", rota.Name))
		} else {
			r.Handle(rota.Metodo, rota.URI, rota.Funcao)
			fmt.Println(fmt.Sprintf("Rota %s não requer autenticação", rota.Name))
		}

	}

	return r
}
