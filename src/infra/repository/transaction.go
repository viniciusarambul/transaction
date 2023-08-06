package repository

import (
	"github.com/viniciusarambul/transaction/src/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) entity.TransactionRepository {
	return &TransactionRepository{DB}
}

func (t *TransactionRepository) Create(transaction *entity.Transaction) error {

	err := t.DB.Create(&transaction)

	return err.Error
}

func (t *TransactionRepository) SumTotalBalance(accountId int) (entity.Transaction, error) {
	return entity.Transaction{}, nil

}
