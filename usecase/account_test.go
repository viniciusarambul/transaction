package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusarambul/transaction/entity"
	mock_entity "github.com/viniciusarambul/transaction/entity/entitiesmock"
	"github.com/viniciusarambul/transaction/usecase"
)

func Test_AccountUseCaseMethodCreate(t *testing.T) {
	t.Run("success create account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		accountPresenter := mock_entity.NewMockAccountPresenter(ctrl)
		log := logrus.New()

		input := entity.AccountInput{
			Document: "1234567892",
			LimitMax: decimal.NewFromFloat(100.00),
		}

		account := entity.Account{
			Document: input.Document,
			LimitMax: input.LimitMax,
		}

		accountRepository.EXPECT().Create(&account).Return(nil)

		uc := usecase.NewAccountUseCase(accountRepository, accountPresenter, log)

		_, err := uc.Create(input)

		assert.NoError(t, err)
	})

	t.Run("error to create account with limit negative", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		accountPresenter := mock_entity.NewMockAccountPresenter(ctrl)
		log := logrus.New()

		input := entity.AccountInput{
			Document: "1234567892",
			LimitMax: decimal.NewFromFloat(100.00).Neg(),
		}

		uc := usecase.NewAccountUseCase(accountRepository, accountPresenter, log)

		_, err := uc.Create(input)

		assert.EqualError(t, err, "Limit não pode ser negativo")
	})

	t.Run("repository error to create account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		accountPresenter := mock_entity.NewMockAccountPresenter(ctrl)
		log := logrus.New()

		input := entity.AccountInput{
			Document: "1234567892",
			LimitMax: decimal.NewFromFloat(100.00),
		}

		account := entity.Account{
			Document: input.Document,
			LimitMax: input.LimitMax,
		}

		accountRepository.EXPECT().Create(&account).Return(errors.New("test error"))

		uc := usecase.NewAccountUseCase(accountRepository, accountPresenter, log)

		_, err := uc.Create(input)

		assert.EqualError(t, err, "test error")
	})
}

func Test_AccountUseCaseMethodFind(t *testing.T) {
	t.Run("success find account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		accountPresenter := mock_entity.NewMockAccountPresenter(ctrl)
		log := logrus.New()

		id := 1

		account := entity.Account{
			ID:       1,
			Document: "123456",
			LimitMax: decimal.NewFromFloat(100.00),
		}

		output := entity.AccountOutput{
			ID:       account.ID,
			Document: account.Document,
			LimitMax: account.LimitMax,
		}

		accountRepository.EXPECT().Find(id).Return(account, nil)
		accountPresenter.EXPECT().Output(account).Return(output)

		uc := usecase.NewAccountUseCase(accountRepository, accountPresenter, log)

		_, err := uc.Find(id)

		assert.NoError(t, err)
	})

	t.Run("repository error find account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		accountPresenter := mock_entity.NewMockAccountPresenter(ctrl)
		log := logrus.New()

		id := 1

		accountRepository.EXPECT().Find(id).Return(entity.Account{}, errors.New("test error"))

		uc := usecase.NewAccountUseCase(accountRepository, accountPresenter, log)

		_, err := uc.Find(id)

		assert.EqualError(t, err, "test error")
	})
}
