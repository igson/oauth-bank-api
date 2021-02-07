package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/autenticacao"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

type CustomerController interface {
	GetAllCustomers(ctx *gin.Context)
	GetCustomer(ctx *gin.Context)
}

type customerController struct {
	service interfaces.ICustomerService
}

//NewCustomerController construtor pra injeção das dependências
func NewCustomerController(customerService interfaces.ICustomerService) CustomerController {
	return &customerController{
		service: customerService,
	}
}

func (c *customerController) GetAllCustomers(ctx *gin.Context) {

	status := ctx.Param("status")

	status = "active"
	/* if status == "" {
		restErr := errors.NewBadRequestError("Status inválido")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}
	*/
	customers, err := c.service.GetAllCustomer(status)

	if err != nil {
		ctx.JSON(err.StatusCode, err)
	}

	uID, erro := autenticacao.ExtrairUsuarioID(ctx.Request)

	if erro != nil {
		ctx.JSON(err.StatusCode, err)
	}

	fmt.Println("UID", uID)

	ctx.JSON(http.StatusOK, customers)

}

func (c *customerController) GetCustomer(ctx *gin.Context) {
	fmt.Printf("Acesso o Medoto %s", "GetCustomer")

	id, error := strconv.ParseInt(ctx.Param("customer_id"), 10, 64)

	if error != nil {
		restErr := errors.NewBadRequestError("Customer ID deve ser número")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	userTokenID, erro := autenticacao.ExtrairUsuarioID(ctx.Request)

	if erro != nil {
		ctx.JSON(erro.StatusCode, erro.Message)
		return
	}

	if userTokenID != id {
		restErr := errors.NewBadRequestError("Você não tem permissão pra essa operação")
		ctx.JSON(restErr.StatusCode, restErr.Message)
		return
	}

	customer, err := c.service.GetCustomer(id)

	if err != nil {
		ctx.JSON(err.StatusCode, err)
	}

	ctx.JSON(http.StatusOK, customer)

}
