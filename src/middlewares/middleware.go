package middlewares

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/autenticacao"
)

//Autenticar autenticação do token
func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Atenticando Usuário")

		if erro := autenticacao.ValidarToken(c.Request); erro != nil {
			fmt.Println("Token inválido...")
			c.JSON(erro.StatusCode, erro.Message)
			c.Abort()
			return
		}

		fmt.Println("Usuário Autenticado com sucesso.")

		c.Next()

	}

}

//Logger informação das rotas invocadas
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("\n %s http://%s%s", c.Request.Method, c.Request.Host, c.Request.RequestURI)
		c.Next()
	}
}
