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
