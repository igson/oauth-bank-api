package app

import (
	"github.com/gin-gonic/gin"

	"github.com/igson/oauth-bank-api/src/app/routers"
)

//GerarRotas iniciar rotas
func GerarRotas() *gin.Engine {
	rota := gin.Default()

	return routers.Configurar(rota)
}
