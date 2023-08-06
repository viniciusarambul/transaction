package usecase

import (
	"time"

	"github.com/viniciusarambul/transaction/src/entity"
)

type TransactionUseCase struct {
	transactiontRepository entity.TransactionRepository
	operationRepository    entity.OperationRepository
}

func NewTransactionUseCase(transactionRepository entity.TransactionRepository, operationRepository entity.OperationRepository) entity.TransactionUseCase {
	return &TransactionUseCase{
		transactiontRepository: transactionRepository,
		operationRepository:    operationRepository,
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
