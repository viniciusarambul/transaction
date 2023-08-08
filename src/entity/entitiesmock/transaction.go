// Code generated by MockGen. DO NOT EDIT.
// Source: transaction.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/viniciusarambul/transaction/src/entity"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionRepository) Create(transaction *entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTransactionRepositoryMockRecorder) Create(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionRepository)(nil).Create), transaction)
}

// SumTotalBalance mocks base method.
func (m *MockTransactionRepository) SumTotalBalance(accountId int) (entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumTotalBalance", accountId)
	ret0, _ := ret[0].(entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumTotalBalance indicates an expected call of SumTotalBalance.
func (mr *MockTransactionRepositoryMockRecorder) SumTotalBalance(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumTotalBalance", reflect.TypeOf((*MockTransactionRepository)(nil).SumTotalBalance), accountId)
}

// MockTransactionUseCase is a mock of TransactionUseCase interface.
type MockTransactionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionUseCaseMockRecorder
}

// MockTransactionUseCaseMockRecorder is the mock recorder for MockTransactionUseCase.
type MockTransactionUseCaseMockRecorder struct {
	mock *MockTransactionUseCase
}

// NewMockTransactionUseCase creates a new mock instance.
func NewMockTransactionUseCase(ctrl *gomock.Controller) *MockTransactionUseCase {
	mock := &MockTransactionUseCase{ctrl: ctrl}
	mock.recorder = &MockTransactionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionUseCase) EXPECT() *MockTransactionUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionUseCase) Create(transactionInput *entity.TransactionInput) (entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", transactionInput)
	ret0, _ := ret[0].(entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTransactionUseCaseMockRecorder) Create(transactionInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionUseCase)(nil).Create), transactionInput)
}

// MockTransactionHandler is a mock of TransactionHandler interface.
type MockTransactionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionHandlerMockRecorder
}

// MockTransactionHandlerMockRecorder is the mock recorder for MockTransactionHandler.
type MockTransactionHandlerMockRecorder struct {
	mock *MockTransactionHandler
}

// NewMockTransactionHandler creates a new mock instance.
func NewMockTransactionHandler(ctrl *gomock.Controller) *MockTransactionHandler {
	mock := &MockTransactionHandler{ctrl: ctrl}
	mock.recorder = &MockTransactionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionHandler) EXPECT() *MockTransactionHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionHandler) Create(context *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", context)
}

// Create indicates an expected call of Create.
func (mr *MockTransactionHandlerMockRecorder) Create(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionHandler)(nil).Create), context)
}