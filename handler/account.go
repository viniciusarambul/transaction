package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viniciusarambul/transaction/entity"
)

type AccountHandler struct {
	AccountUseCase entity.AccountUseCase
}

func NewAccountHandler(engine *gin.Engine, AccountUseCase entity.AccountUseCase) entity.AccountHandler {
	handler := &AccountHandler{AccountUseCase}
	engine.POST("/accounts", handler.Create)
	engine.GET("/accounts/:accountId", handler.Find)

	return handler
}

func (accountHandler AccountHandler) Create(context *gin.Context) {
	var accountInput entity.AccountInput
	err := context.ShouldBindJSON(&accountInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := accountHandler.AccountUseCase.Create(accountInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, account.ID)
}

func (accountHandler AccountHandler) Find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("accountId"))

	account, err := accountHandler.AccountUseCase.Find(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, account)
}
