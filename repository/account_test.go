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

func TestAccount_Create(t *testing.T) {
	t.Run("success create account", func(t *testing.T) {
		mockDb, db := repositorytest.NewGorm(t)
		repo := repository.NewAccountRepository(db)

		acc := &entity.Account{
			ID:        1,
			Document:  "123456",
			LimitMax:  decimal.NewFromFloat(100.00),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockDb.ExpectBegin()

		mockDb.ExpectQuery(`INSERT INTO "accounts" (.+) RETURNING`).
			WithArgs(acc.Document, acc.LimitMax, acc.CreatedAt, acc.UpdatedAt, acc.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mockDb.ExpectCommit()

		err := repo.Create(acc)
		assert.NoError(t, err)
	})

	t.Run("find account", func(t *testing.T) {
		mockDb, db := repositorytest.NewGorm(t)
		repo := repository.NewAccountRepository(db)

		mockDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" LIMIT 1`)).
			WithArgs(1).
			WillReturnRows(
				mockDb.NewRows([]string{"id", "document"}).AddRow(1, "123456"),
			)

		account, err := repo.Find(1)

		expected := entity.Account{
			ID:       1,
			Document: "123456",
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, account)
	})
}
