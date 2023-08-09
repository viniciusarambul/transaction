package repository_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viniciusarambul/transaction/entity"
	"github.com/viniciusarambul/transaction/repository"
	"github.com/viniciusarambul/transaction/repository/repositorytest"
)

func TestOperations(t *testing.T) {
	t.Run("find operation type", func(t *testing.T) {
		mockDb, db := repositorytest.NewGorm(t)
		repo := repository.NewOperationRepository(db)

		mockDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "operations_types" WHERE operation_type_id = $1`)).
			WithArgs(1).
			WillReturnRows(
				mockDb.NewRows([]string{"id", "operation_type", "description", "type"}).AddRow(1, 1, "COMPRA A VISTA", entity.DEBIT),
			)

		operation, err := repo.FindByOperationType(1)

		expected := entity.OperationsTypes{
			ID:            1,
			OperationType: 1,
			Description:   "COMPRA A VISTA",
			Type:          entity.DEBIT,
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, operation)
	})
}
