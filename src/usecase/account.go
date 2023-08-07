package usecase

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/viniciusarambul/transaction/src/entity"
)

type AccountUseCase struct {
	accountRepository entity.AccountRepository
	accountPresenter  entity.AccountPresenter
	log               *logrus.Logger
}

func NewAccountUseCase(repository entity.AccountRepository, presenter entity.AccountPresenter, log *logrus.Logger) entity.AccountUseCase {
	return &AccountUseCase{
		accountRepository: repository,
		accountPresenter:  presenter,
		log:               log,
	}
}

func (accountUseCase *AccountUseCase) Create(accountInput entity.AccountInput) (entity.Account, error) {
	accountUseCase.log.Info("Start method Create on AccountUseCase")

	account := entity.Account{
		Document:  accountInput.Document,
		LimitMax:  accountInput.LimitMax,
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
	}

	if account.LimitMax.LessThan(decimal.Zero) {
		return entity.Account{}, errors.New("Limit não pode ser negativo")
	}

	err := accountUseCase.accountRepository.Create(&account)

	if err != nil {
		accountUseCase.log.WithError(err).Error("Error to Create Account")

		return entity.Account{}, err
	}
	accountUseCase.log.WithFields(logrus.Fields{
		"ID":        account.ID,
		"limit_max": account.LimitMax,
	}).Info("Account create successful")

	accountUseCase.log.Info("Finish method Create on AccountUseCase")

	return account, err
}

func (accountUseCase *AccountUseCase) Find(id int) (entity.AccountOutput, error) {
	accountUseCase.log.WithField("id", id).Info("Start method Find on AccountUseCase")

	account, err := accountUseCase.accountRepository.Find(id)
	if err != nil {
		accountUseCase.log.WithError(err).Error("Error to Find Account")
		return entity.AccountOutput{}, err
	}

	accountOutput := accountUseCase.accountPresenter.Output(account)
	accountUseCase.log.Info("Finish method Find on AccountUseCase")

	return accountOutput, nil
}
