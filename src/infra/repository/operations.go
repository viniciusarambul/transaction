package repository

import (
	"github.com/viniciusarambul/transaction/src/entity"
	"gorm.io/gorm"
)

type OperationRepository struct {
	DB *gorm.DB
}

func NewOperationRepository(DB *gorm.DB) entity.OperationRepository {
	return &OperationRepository{DB}
}

func (o *OperationRepository) FindByOperationType(id int) (entity.OperationsTypes, error) {
	var operation entity.OperationsTypes

	err := o.DB.Where("operation_type_id = ?", id).Find(&operation)

	return operation, err.Error
}
