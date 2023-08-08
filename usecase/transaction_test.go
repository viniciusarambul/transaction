package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusarambul/transaction/entity"
	mock_entity "github.com/viniciusarambul/transaction/entity/entitiesmock"
	"github.com/viniciusarambul/transaction/usecase"
	mock_utils "github.com/viniciusarambul/transaction/utils/utilsmock"
)

func Test_TransactionUseCaseMethodCreate(t *testing.T) {
	now := time.Now()
	t.Run("success create transaction", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		transactionRepository := mock_entity.NewMockTransactionRepository(ctrl)
		operationRepository := mock_entity.NewMockOperationRepository(ctrl)
		clock := mock_utils.NewMockClock(ctrl)
		log := logrus.New()

		clock.EXPECT().Now().Return(now)

		input := entity.TransactionInput{
			AccountId:       1,
			OperationTypeId: 4,
			Amount:          decimal.NewFromFloat(100.00),
			IdempotencyKey:  "123456",
		}

		transaction := entity.Transaction{
			IdempotencyKey:  input.IdempotencyKey,
			AccountId:       input.AccountId,
			OperationTypeId: input.OperationTypeId,
			Amount:          input.Amount,
			EventDate:       now,
		}

		operation := entity.OperationsTypes{
			OperationType: 4,
			Description:   "PAGAMENTO",
			Type:          entity.CREDIT,
		}

		account := entity.Account{
			ID:       1,
			Document: "123456",
			LimitMax: decimal.NewFromFloat(200.00),
		}

		balanceTransaction := entity.Transaction{
			Amount: decimal.NewFromFloat(100.00),
		}

		operationRepository.EXPECT().FindByOperationType(transaction.OperationTypeId).Return(operation, nil)
		accountRepository.EXPECT().Find(transaction.AccountId).Return(account, nil)
		transactionRepository.EXPECT().SumTotalBalance(transaction.AccountId).Return(balanceTransaction, nil)
		transactionRepository.EXPECT().Create(&transaction).Return(nil)

		uc := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

		_, err := uc.Create(&input)

		assert.NoError(t, err)
	})

	t.Run("Repository error FindByOperationType", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		transactionRepository := mock_entity.NewMockTransactionRepository(ctrl)
		operationRepository := mock_entity.NewMockOperationRepository(ctrl)
		clock := mock_utils.NewMockClock(ctrl)
		log := logrus.New()

		clock.EXPECT().Now().Return(now)

		input := entity.TransactionInput{
			AccountId:       1,
			OperationTypeId: 4,
			Amount:          decimal.NewFromFloat(100.00),
			IdempotencyKey:  "123456",
		}

		transaction := entity.Transaction{
			IdempotencyKey:  input.IdempotencyKey,
			AccountId:       input.AccountId,
			OperationTypeId: input.OperationTypeId,
			Amount:          input.Amount,
			EventDate:       now,
		}

		operation := entity.OperationsTypes{}

		operationRepository.EXPECT().FindByOperationType(transaction.OperationTypeId).Return(operation, errors.New("test error"))

		uc := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

		_, err := uc.Create(&input)

		assert.EqualError(t, err, "test error")
	})

	t.Run("AccountRepository error find account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		transactionRepository := mock_entity.NewMockTransactionRepository(ctrl)
		operationRepository := mock_entity.NewMockOperationRepository(ctrl)
		clock := mock_utils.NewMockClock(ctrl)
		log := logrus.New()

		clock.EXPECT().Now().Return(now)

		input := entity.TransactionInput{
			AccountId:       1,
			OperationTypeId: 4,
			Amount:          decimal.NewFromFloat(100.00),
			IdempotencyKey:  "123456",
		}

		transaction := entity.Transaction{
			IdempotencyKey:  input.IdempotencyKey,
			AccountId:       input.AccountId,
			OperationTypeId: input.OperationTypeId,
			Amount:          input.Amount,
			EventDate:       now,
		}

		operation := entity.OperationsTypes{
			OperationType: 4,
			Description:   "PAGAMENTO",
			Type:          entity.CREDIT,
		}

		account := entity.Account{
			ID:       1,
			Document: "123456",
			LimitMax: decimal.NewFromFloat(200.00),
		}

		operationRepository.EXPECT().FindByOperationType(transaction.OperationTypeId).Return(operation, nil)
		accountRepository.EXPECT().Find(transaction.AccountId).Return(account, errors.New("test error"))

		uc := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

		_, err := uc.Create(&input)

		assert.EqualError(t, err, "test error")
	})

	t.Run("TransactionRepository error sum total balance", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		transactionRepository := mock_entity.NewMockTransactionRepository(ctrl)
		operationRepository := mock_entity.NewMockOperationRepository(ctrl)
		clock := mock_utils.NewMockClock(ctrl)
		log := logrus.New()

		clock.EXPECT().Now().Return(now)

		input := entity.TransactionInput{
			AccountId:       1,
			OperationTypeId: 4,
			Amount:          decimal.NewFromFloat(100.00),
			IdempotencyKey:  "123456",
		}

		transaction := entity.Transaction{
			IdempotencyKey:  input.IdempotencyKey,
			AccountId:       input.AccountId,
			OperationTypeId: input.OperationTypeId,
			Amount:          input.Amount,
			EventDate:       now,
		}

		operation := entity.OperationsTypes{
			OperationType: 4,
			Description:   "PAGAMENTO",
			Type:          entity.CREDIT,
		}

		account := entity.Account{
			ID:       1,
			Document: "123456",
			LimitMax: decimal.NewFromFloat(200.00),
		}

		balanceTransaction := entity.Transaction{
			Amount: decimal.NewFromFloat(100.00),
		}

		operationRepository.EXPECT().FindByOperationType(transaction.OperationTypeId).Return(operation, nil)
		accountRepository.EXPECT().Find(transaction.AccountId).Return(account, nil)
		transactionRepository.EXPECT().SumTotalBalance(transaction.AccountId).Return(balanceTransaction, errors.New("test error"))

		uc := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

		_, err := uc.Create(&input)

		assert.EqualError(t, err, "test error")
	})

	t.Run("No balance for transaction", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		transactionRepository := mock_entity.NewMockTransactionRepository(ctrl)
		operationRepository := mock_entity.NewMockOperationRepository(ctrl)
		clock := mock_utils.NewMockClock(ctrl)
		log := logrus.New()

		clock.EXPECT().Now().Return(now)

		input := entity.TransactionInput{
			AccountId:       1,
			OperationTypeId: 2,
			Amount:          decimal.NewFromFloat(301.00),
			IdempotencyKey:  "123456",
		}

		transaction := entity.Transaction{
			IdempotencyKey:  input.IdempotencyKey,
			AccountId:       input.AccountId,
			OperationTypeId: input.OperationTypeId,
			Amount:          input.Amount,
			EventDate:       now,
		}

		operation := entity.OperationsTypes{
			OperationType: 2,
			Description:   "COMPRA PARCELADA",
			Type:          entity.DEBIT,
		}

		account := entity.Account{
			ID:       1,
			Document: "123456",
			LimitMax: decimal.NewFromFloat(200.00),
		}

		balanceTransaction := entity.Transaction{
			Amount: decimal.NewFromFloat(100.00),
		}

		operationRepository.EXPECT().FindByOperationType(transaction.OperationTypeId).Return(operation, nil)
		accountRepository.EXPECT().Find(transaction.AccountId).Return(account, nil)
		transactionRepository.EXPECT().SumTotalBalance(transaction.AccountId).Return(balanceTransaction, nil)

		uc := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

		_, err := uc.Create(&input)

		assert.Error(t, err, "Sem crédito suficiente para esta transação")
	})

	t.Run("Error repository create transaction", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepository := mock_entity.NewMockAccountRepository(ctrl)
		transactionRepository := mock_entity.NewMockTransactionRepository(ctrl)
		operationRepository := mock_entity.NewMockOperationRepository(ctrl)
		clock := mock_utils.NewMockClock(ctrl)
		log := logrus.New()

		clock.EXPECT().Now().Return(now)

		input := entity.TransactionInput{
			AccountId:       1,
			OperationTypeId: 4,
			Amount:          decimal.NewFromFloat(100.00),
			IdempotencyKey:  "123456",
		}

		transaction := entity.Transaction{
			IdempotencyKey:  input.IdempotencyKey,
			AccountId:       input.AccountId,
			OperationTypeId: input.OperationTypeId,
			Amount:          input.Amount,
			EventDate:       now,
		}

		operation := entity.OperationsTypes{
			OperationType: 4,
			Description:   "PAGAMENTO",
			Type:          entity.CREDIT,
		}

		account := entity.Account{
			ID:       1,
			Document: "123456",
			LimitMax: decimal.NewFromFloat(200.00),
		}

		balanceTransaction := entity.Transaction{
			Amount: decimal.NewFromFloat(100.00),
		}

		operationRepository.EXPECT().FindByOperationType(transaction.OperationTypeId).Return(operation, nil)
		accountRepository.EXPECT().Find(transaction.AccountId).Return(account, nil)
		transactionRepository.EXPECT().SumTotalBalance(transaction.AccountId).Return(balanceTransaction, nil)
		transactionRepository.EXPECT().Create(&transaction).Return(errors.New("test error"))

		uc := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

		_, err := uc.Create(&input)

		assert.EqualError(t, err, "test error")
	})
}
