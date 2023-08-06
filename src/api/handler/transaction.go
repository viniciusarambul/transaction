package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusarambul/transaction/src/entity"
)

type TransactionHandler struct {
	TransactionUseCase entity.TransactionUseCase
}

func NewTransactionHandler(engine *gin.Engine, TransactionUseCase entity.TransactionUseCase) entity.TransactionHandler {
	handler := &TransactionHandler{TransactionUseCase: TransactionUseCase}
	engine.POST("/transactions", handler.Create)

	return handler
}

func (transactionHandler TransactionHandler) Create(context *gin.Context) {
	var transactionInput entity.TransactionInput
	err := context.ShouldBindJSON(&transactionInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := transactionHandler.TransactionUseCase.Create(transactionInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"idempotency_key": transaction.IdempotencyKey})
}
