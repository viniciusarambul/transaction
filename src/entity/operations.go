//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=operations.go -destination=entitiesmock/operations.go .
package entity

const (
	CREDIT = "CREDIT"
	DEBIT  = "DEBIT"
)

type (
	OperationsTypes struct {
		ID            int
		OperationType int
		Description   string
		Type          string
	}
	OperationRepository interface {
		FindByOperationType(id int) (OperationsTypes, error)
	}
)
