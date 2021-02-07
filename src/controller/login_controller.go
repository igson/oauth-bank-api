package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/autenticacao"
	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	service interfaces.IUserService
}

//NewLoginController construtor pra injeção das dependências
func NewLoginController(userService interfaces.IUserService) LoginController {
	return &loginController{
		service: userService,
	}
}

func (c *loginController) Login(ctx *gin.Context) {

	login := models.Login{}

	if err := ctx.ShouldBindJSON(&login); err != nil {
		erroMessage := errors.NewBadRequestError("invalid json error body")
		ctx.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	if loginErr := c.service.Login(login.Username, login.Password); loginErr != nil {
		ctx.JSON(loginErr.StatusCode, loginErr)
		return
	}

	id, erro := strconv.Atoi(login.Username)

	if erro != nil {
		restErr := errors.NewBadRequestError("Username ID deve ser número")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	fmt.Println("ID informado:", id)

	token, tokenErr := autenticacao.CriarToken(uint64(id))

	if tokenErr != nil {
		ctx.JSON(tokenErr.StatusCode, tokenErr)
		return
	}

	ctx.JSON(http.StatusOK, token)

}
