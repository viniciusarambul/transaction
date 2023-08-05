package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	Account struct {
		ID        uuid.UUID
		AccountId int
		Document  string
		Limit     decimal.Decimal
	}
	AccountInput struct {
		AccountId int             `json:"account_id"`
		Document  string          `json:"document"`
		Limit     decimal.Decimal `json:"limit"`
	}
	AccountOutput struct {
		AccountId int             `json:"account_id"`
		Document  string          `json:"document"`
		Limit     decimal.Decimal `json:"limit"`
	}
)
