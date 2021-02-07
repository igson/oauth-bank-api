package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Logger informação das rotas invocadas
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("\n %s http://%s%s", c.Request.Method, c.Request.Host, c.Request.RequestURI)
		c.Next()
	}
}
