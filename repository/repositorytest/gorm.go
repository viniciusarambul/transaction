package repositorytest

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}

	mock.MatchExpectationsInOrder(true)

	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return mock, gormdb
}
