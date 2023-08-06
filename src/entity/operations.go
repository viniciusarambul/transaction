package entity

type (
	Operation struct {
		ID            int
		OperationType int
		Description   string
		Type          string
	}
	OperationRepository interface {
		FindByOperationType(id int) (Operation, error)
	}
)
