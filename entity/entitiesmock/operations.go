// Code generated by MockGen. DO NOT EDIT.
// Source: operations.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/viniciusarambul/transaction/entity"
)

// MockOperationRepository is a mock of OperationRepository interface.
type MockOperationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOperationRepositoryMockRecorder
}

// MockOperationRepositoryMockRecorder is the mock recorder for MockOperationRepository.
type MockOperationRepositoryMockRecorder struct {
	mock *MockOperationRepository
}

// NewMockOperationRepository creates a new mock instance.
func NewMockOperationRepository(ctrl *gomock.Controller) *MockOperationRepository {
	mock := &MockOperationRepository{ctrl: ctrl}
	mock.recorder = &MockOperationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationRepository) EXPECT() *MockOperationRepositoryMockRecorder {
	return m.recorder
}

// FindByOperationType mocks base method.
func (m *MockOperationRepository) FindByOperationType(id int) (entity.OperationsTypes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOperationType", id)
	ret0, _ := ret[0].(entity.OperationsTypes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOperationType indicates an expected call of FindByOperationType.
func (mr *MockOperationRepositoryMockRecorder) FindByOperationType(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOperationType", reflect.TypeOf((*MockOperationRepository)(nil).FindByOperationType), id)
}