package usecase

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/viniciusarambul/transaction/src/entity"
)

type TransactionUseCase struct {
	transactiontRepository entity.TransactionRepository
	operationRepository    entity.OperationRepository
	accountRepository      entity.AccountRepository
}

func NewTransactionUseCase(transactionRepository entity.TransactionRepository, operationRepository entity.OperationRepository, accountRepository entity.AccountRepository) entity.TransactionUseCase {
	return &TransactionUseCase{
		transactiontRepository: transactionRepository,
		operationRepository:    operationRepository,
		accountRepository:      accountRepository,
	}
}

func (transactionUseCase *TransactionUseCase) Create(transactionInput *entity.TransactionInput, idempotency string) (entity.Transaction, error) {
	transactionBeforeVerify := entity.Transaction{
		IdempotencyKey:  idempotency,
		AccountId:       transactionInput.AccountId,
		OperationTypeId: transactionInput.OperationTypeId,
		Amount:          transactionInput.Amount,
		EventDate:       time.Now().Local().UTC(),
	}

	transaction, err := transactionUseCase.verifyOperationCondition(&transactionBeforeVerify)

	if err != nil {
		return entity.Transaction{}, err
	}

	limit, err := transactionUseCase.accountRepository.Find(transaction.AccountId)

	if err != nil {
		return entity.Transaction{}, err
	}

	totalBalance, err := transactionUseCase.transactiontRepository.SumTotalBalance(transaction.AccountId)

	if err != nil {
		return entity.Transaction{}, err
	}

	if transactionUseCase.verifyBalanceToTransaction(transaction, limit.LimitMax, totalBalance.Amount) {
		return entity.Transaction{}, errors.New("Sem crédito suficiente para esta transação")
	}

	err = transactionUseCase.transactiontRepository.Create(transaction)

	if err != nil {
		return entity.Transaction{}, err
	}

	return transactionBeforeVerify, err
}

func (transactionUseCase *TransactionUseCase) verifyOperationCondition(transaction *entity.Transaction) (*entity.Transaction, error) {
	operation, err := transactionUseCase.operationRepository.FindByOperationType(transaction.OperationTypeId)

	if err != nil {
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
