package repository

import (
	"github.com/viniciusarambul/transaction/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

type Balance struct {
	amount string
}

func NewTransactionRepository(DB *gorm.DB) entity.TransactionRepository {
	return &TransactionRepository{DB}
}

func (t *TransactionRepository) Create(transaction *entity.Transaction) error {

	err := t.DB.Create(&transaction)

	return err.Error
}

func (t *TransactionRepository) SumTotalBalance(accountId int) (entity.Transaction, error) {
	var transaction entity.Transaction

	query := t.DB.Select("sum(amount) as amount").Where("account_id = ?", accountId)

	err := query.Find(&transaction)
	return transaction, err.Error
}
