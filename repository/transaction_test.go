package repository_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusarambul/transaction/entity"
	"github.com/viniciusarambul/transaction/repository"
	"github.com/viniciusarambul/transaction/repository/repositorytest"
)

func TestTransaction_Create(t *testing.T) {

	t.Run("success create transaction", func(t *testing.T) {
		mockDb, db := repositorytest.NewGorm(t)
		repo := repository.NewTransactionRepository(db)

		trs := &entity.Transaction{
			ID:              1,
			IdempotencyKey:  "123456",
			AccountId:       1,
			OperationTypeId: 2,
			Amount:          decimal.NewFromFloat(100.00),
			EventDate:       time.Now(),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		mockDb.ExpectBegin()

		mockDb.ExpectQuery(`INSERT INTO "transactions" (.+) RETURNING`).
			WithArgs(trs.IdempotencyKey, trs.AccountId, trs.OperationTypeId, trs.Amount, trs.EventDate, trs.CreatedAt, trs.UpdatedAt, trs.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mockDb.ExpectCommit()

		err := repo.Create(trs)
		assert.NoError(t, err)
	})

	t.Run("sum total balance", func(t *testing.T) {
		mockDb, db := repositorytest.NewGorm(t)
		repo := repository.NewTransactionRepository(db)

		mockDb.ExpectQuery(regexp.QuoteMeta(`SELECT sum(amount) as amount FROM "transactions" WHERE account_id = $1`)).
			WithArgs(1).
			WillReturnRows(
				mockDb.NewRows([]string{"account_id", "amount"}).AddRow(1, 100.00),
			)

		balance, err := repo.SumTotalBalance(1)

		expected := entity.Transaction{
			AccountId: 1,
			Amount:    decimal.NewFromFloat(100.00),
		}
		assert.NoError(t, err)
		assert.Equal(t, expected, balance)
	})
}

//expected: entity.Transaction{ID:0, IdempotencyKey:"", AccountId:1, OperationTypeId:0, Amount:decimal.Decimal{value:(*big.Int)(0xc0004679a0), exp:2}, EventDate:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), CreatedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}
//actual  : entity.Transaction{ID:0, IdempotencyKey:"", AccountId:1, OperationTypeId:0, Amount:decimal.Decimal{value:(*big.Int)(0xc000467980), exp:0}, EventDate:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), CreatedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}
