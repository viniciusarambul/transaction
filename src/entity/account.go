package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type (
	Account struct {
		ID       int
		Document string
		LimitMax decimal.Decimal
	}

	AccountInput struct {
		Document string `json:"document"`
	}

	AccountOutput struct {
		ID       int             `json:"id"`
		Document string          `json:"document"`
		LimitMax decimal.Decimal `json:"limit"`
	}

	AccountRepository interface {
		Find(id int) (Account, error)
		Create(account *Account) error
	}

	AccountPresenter interface {
		Output(account Account) AccountOutput
	}

	AccountUseCase interface {
		Create(accountInput AccountInput) (Account, error)
		Find(id int) (AccountOutput, error)
	}

	AccountHandler interface {
		Create(context *gin.Context)
	}
)
