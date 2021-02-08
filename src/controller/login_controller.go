package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/igson/oauth-bank-api/src/domain/dto"
	"github.com/igson/oauth-bank-api/src/errors"
	"github.com/igson/oauth-bank-api/src/interfaces"
)

type LoginController interface {
	Login(ctx *gin.Context)
	Verify(ctx *gin.Context)
}

type loginController struct {
	service interfaces.IAuthService
}

//NewLoginController construtor pra injeção das dependências
func NewLoginController(userService interfaces.IAuthService) LoginController {
	return &loginController{
		service: userService,
	}
}

func (c *loginController) Login(ctx *gin.Context) {

	request := dto.LoginRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		erroMessage := errors.NewBadRequestError("erro de formato JSON")
		ctx.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	tokenString, tokenErr := c.service.Login(request)

	if tokenErr != nil {
		ctx.JSON(tokenErr.StatusCode, tokenErr)
		return
	}

	ctx.JSON(http.StatusOK, tokenString)

}

func (c *loginController) NotImplementedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Handler not implemented...")
}

/*
  Sample URL string
 http://localhost:8181/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&customer_id=2000&account_id=95470
*/
func (c *loginController) Verify(ctx *gin.Context) {

	log.Println("Acesso -------> Login Verify")

	urlParams := make(map[string]string)

	// converting from Query to map type
	for k := range ctx.Request.URL.Query() {
		urlParams[k] = ctx.Request.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {

		isAuthorized, appError := c.service.Verify(urlParams)

		if appError != nil {
			ctx.JSON(http.StatusForbidden, notAuthorizedResponse())
		} else {

			if isAuthorized {
				ctx.JSON(http.StatusOK, authorizedResponse())
			} else {
				ctx.JSON(http.StatusForbidden, notAuthorizedResponse())
			}

		}

	} else {
		ctx.JSON(http.StatusForbidden, "Token ausente")
	}

}

func notAuthorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": false}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}
