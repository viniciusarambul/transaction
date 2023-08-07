package entity

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type (
	Account struct {
		ID        int
		Document  string
		LimitMax  decimal.Decimal
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AccountInput struct {
		Document  string          `json:"document"`
		LimitMax  decimal.Decimal `json:"limit_max"`
		CreateAt  time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
	}

	AccountOutput struct {
		ID       int             `json:"id"`
		Document string          `json:"document"`
		LimitMax decimal.Decimal `json:"limit_max"`
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

func RemoveLGPDFromResponse(document string) string {
	documentRegex := regexp.MustCompile(document)
	maskedDocument := documentRegex.ReplaceAllString(document, fmt.Sprintf("%s******-%s", document[0:3], document[len(document)-2:]))

	return maskedDocument
}
