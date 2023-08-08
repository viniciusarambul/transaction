package usecase

import (
	"errors"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/viniciusarambul/transaction/src/entity"
	"github.com/viniciusarambul/transaction/src/pkg"
)

type TransactionUseCase struct {
	transactiontRepository entity.TransactionRepository
	operationRepository    entity.OperationRepository
	accountRepository      entity.AccountRepository
	log                    *logrus.Logger
	clock                  pkg.Clock
}

func NewTransactionUseCase(transactionRepository entity.TransactionRepository, operationRepository entity.OperationRepository, accountRepository entity.AccountRepository, log *logrus.Logger, clock pkg.Clock) entity.TransactionUseCase {
	return &TransactionUseCase{
		transactiontRepository: transactionRepository,
		operationRepository:    operationRepository,
		accountRepository:      accountRepository,
		log:                    log,
		clock:                  clock,
	}
}

func (transactionUseCase *TransactionUseCase) Create(transactionInput *entity.TransactionInput) (entity.Transaction, error) {

	transactionUseCase.log.Info("Start method Create on TransactionUseCase")

	transactionBeforeVerify := entity.Transaction{
		IdempotencyKey:  transactionInput.IdempotencyKey,
		AccountId:       transactionInput.AccountId,
		OperationTypeId: transactionInput.OperationTypeId,
		Amount:          transactionInput.Amount,
		EventDate:       transactionUseCase.clock.Now(),
	}

	transaction, err := transactionUseCase.verifyOperationCondition(&transactionBeforeVerify)

	if err != nil {
		transactionUseCase.log.WithError(err).Error("Error to verifyOperationCondition")

		return entity.Transaction{}, err
	}

	limit, err := transactionUseCase.accountRepository.Find(transaction.AccountId)

	if err != nil {
		transactionUseCase.log.WithError(err).Error("Error to Find AccountID")

		return entity.Transaction{}, err
	}

	totalBalance, err := transactionUseCase.transactiontRepository.SumTotalBalance(transaction.AccountId)

	if err != nil {
		transactionUseCase.log.WithError(err).Error("Error to SumTotalBalance")

		return entity.Transaction{}, err
	}

	if transactionUseCase.verifyBalanceToTransaction(transaction, limit.LimitMax, totalBalance.Amount) {
		transactionUseCase.log.WithError(err).Error("Sem crédito suficiente para esta transação")
		return entity.Transaction{}, errors.New("Sem crédito suficiente para esta transação")
	}

	err = transactionUseCase.transactiontRepository.Create(transaction)

	if err != nil {
		transactionUseCase.log.WithError(err).Error("Erro ao criar transação")
		return entity.Transaction{}, err
	}

	transactionUseCase.log.WithFields(logrus.Fields{
		"idempotency_key":   transaction.IdempotencyKey,
		"account_id":        transaction.AccountId,
		"operation_type_id": transaction.OperationTypeId,
		"amount":            transaction.Amount,
	}).Info("Transaction create successful")

	transactionUseCase.log.Info("End method Create on TransactionUseCase")

	return transactionBeforeVerify, err
}

func (transactionUseCase *TransactionUseCase) verifyOperationCondition(transaction *entity.Transaction) (*entity.Transaction, error) {
	operation, err := transactionUseCase.operationRepository.FindByOperationType(transaction.OperationTypeId)

	if err != nil {
		transactionUseCase.log.WithError(err).Error("Erro method FindByOperationType")
		return nil, err
	}

	if operation.Type == entity.DEBIT {
		transaction.Amount = transaction.Amount.Neg()
	}

	return transaction, nil
}

func (transactionUseCase *TransactionUseCase) verifyBalanceToTransaction(transaction *entity.Transaction, limit decimal.Decimal, totalBalance decimal.Decimal) bool {

	return limit.Neg().GreaterThan(totalBalance.Add(transaction.Amount))

}
