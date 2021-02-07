package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/domain/dto"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

type AccountController interface {
	BuscarContaPorID(ctx *gin.Context)
	RegistrarTransacao(ctx *gin.Context)
}

type accountController struct {
	accountService interfaces.IAccountService
}

//NewAccountController construtor pra injeção das dependências
func NewAccountController(accountService interfaces.IAccountService) AccountController {
	return &accountController{
		accountService: accountService,
	}
}

func (c *accountController) BuscarContaPorID(ctx *gin.Context) {

	userID, erro := getID(ctx.Param("account_id"))

	if erro != nil {
		ctx.JSON(erro.StatusCode, erro)
		return
	}

	account, erroGerUser := c.accountService.BuscarContaPorID(userID)

	if erroGerUser != nil {
		ctx.JSON(erroGerUser.StatusCode, erroGerUser)
		return
	}

	ctx.JSON(http.StatusOK, account)

}

func (c *accountController) RegistrarTransacao(ctx *gin.Context) {

	accountId, error := strconv.ParseInt(ctx.Param("account_id"), 10, 64)

	if error != nil {
		restErr := errors.NewBadRequestError("Account ID deve ser número")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	customerId, error := strconv.ParseInt(ctx.Param("customer_id"), 10, 64)

	if error != nil {
		restErr := errors.NewBadRequestError("Custumer ID deve ser número")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	var transaction = dto.TransactionRequest{
		Id:         accountId,
		CustomerId: customerId,
	}

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		erroMessage := errors.NewBadRequestError("invalid json error body")
		ctx.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	fmt.Println(transaction)

	register, erroTransaction := c.accountService.RegistrarTransacao(transaction)

	if erroTransaction != nil {
		ctx.JSON(erroTransaction.StatusCode, erroTransaction)
		return
	}

	ctx.JSON(http.StatusOK, register)
}

func getID(userID string) (int64, *errors.RestErroAPI) {
	ID, erro := strconv.ParseInt(userID, 10, 64)
	if erro != nil {
		return 0, errors.NewBadRequestError("ID deve ser número")
	}
	return ID, nil
}
