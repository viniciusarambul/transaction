package entity

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type (
	Transaction struct {
		ID              int
		IdempotencyKey  string
		AccountId       int
		OperationTypeId int
		Amount          decimal.Decimal
		EventDate       time.Time
	}

	TransactionInput struct {
		AccountId       int             `json:"account_id"`
		OperationTypeId int             `json:"operation_type_id"`
		Amount          decimal.Decimal `json:"amount"`
		IdempotencyKey  string          `json:"idempotency_key"`
	}

	TransactionRepository interface {
		Create(transaction *Transaction) error
		SumTotalBalance(accountId int) (Transaction, error)
	}

	TransactionUseCase interface {
		Create(transactionInput *TransactionInput) (Transaction, error)
	}

	TransactionHandler interface {
		Create(context *gin.Context)
	}
)
