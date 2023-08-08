package repository

import (
	"github.com/viniciusarambul/transaction/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) entity.AccountRepository {
	return &AccountRepository{DB}
}

func (a *AccountRepository) Find(id int) (entity.Account, error) {
	var account entity.Account

	err := a.DB.First(&account, id)

	return account, err.Error
}

func (a *AccountRepository) Create(account *entity.Account) error {

	err := a.DB.Create(&account)

	return err.Error
}
