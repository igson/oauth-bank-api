package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/igson/oauth-bank-api/src/domain/dto"
	"github.com/igson/oauth-bank-api/src/errors"
	"github.com/igson/oauth-bank-api/src/interfaces"
)

type LoginController interface {
	Login(ctx *gin.Context)
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
