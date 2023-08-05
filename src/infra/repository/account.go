package repository

import (
	"github.com/viniciusarambul/transaction/src/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func (a *AccountRepository) Fund(id int) (entity.Account, error) {
	//var account entity.Account

	//err := a.DB.First(&account, id)

}
