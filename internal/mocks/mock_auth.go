// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "b2b/m/internal/services/auth/models"
	models0 "b2b/m/internal/services/company/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthRepository)(nil).CreateUser), ctx, user)
}

// CreateUserSession mocks base method.
func (m *MockAuthRepository) CreateUserSession(ctx context.Context, userID int64, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSession", ctx, userID, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserSession indicates an expected call of CreateUserSession.
func (mr *MockAuthRepositoryMockRecorder) CreateUserSession(ctx, userID, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSession", reflect.TypeOf((*MockAuthRepository)(nil).CreateUserSession), ctx, userID, hash)
}

// FastRegistration mocks base method.
func (m *MockAuthRepository) FastRegistration(ctx context.Context, newCompany *models0.Company, user *models.User, post string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FastRegistration", ctx, newCompany, user, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// FastRegistration indicates an expected call of FastRegistration.
func (mr *MockAuthRepositoryMockRecorder) FastRegistration(ctx, newCompany, user, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FastRegistration", reflect.TypeOf((*MockAuthRepository)(nil).FastRegistration), ctx, newCompany, user, post)
}

// GetCompanyUserLink mocks base method.
func (m *MockAuthRepository) GetCompanyUserLink(ctx context.Context, userId, companyId int64) (*models0.CompaniesUsersLink, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyUserLink", ctx, userId, companyId)
	ret0, _ := ret[0].(*models0.CompaniesUsersLink)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyUserLink indicates an expected call of GetCompanyUserLink.
func (mr *MockAuthRepositoryMockRecorder) GetCompanyUserLink(ctx, userId, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyUserLink", reflect.TypeOf((*MockAuthRepository)(nil).GetCompanyUserLink), ctx, userId, companyId)
}

// GetUserByEmail mocks base method.
func (m *MockAuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockAuthRepositoryMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockAuthRepository)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockAuthRepository) GetUserByID(ctx context.Context, ID int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, ID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAuthRepositoryMockRecorder) GetUserByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAuthRepository)(nil).GetUserByID), ctx, ID)
}

// GetUserInfo mocks base method.
func (m *MockAuthRepository) GetUserInfo(ctx context.Context, id int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockAuthRepositoryMockRecorder) GetUserInfo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockAuthRepository)(nil).GetUserInfo), ctx, id)
}

// GetUsersCompany mocks base method.
func (m *MockAuthRepository) GetUsersCompany(ctx context.Context, userId int64) (*models0.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersCompany", ctx, userId)
	ret0, _ := ret[0].(*models0.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersCompany indicates an expected call of GetUsersCompany.
func (mr *MockAuthRepositoryMockRecorder) GetUsersCompany(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersCompany", reflect.TypeOf((*MockAuthRepository)(nil).GetUsersCompany), ctx, userId)
}

// RemoveUserSession mocks base method.
func (m *MockAuthRepository) RemoveUserSession(ctx context.Context, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserSession", ctx, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserSession indicates an expected call of RemoveUserSession.
func (mr *MockAuthRepositoryMockRecorder) RemoveUserSession(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserSession", reflect.TypeOf((*MockAuthRepository)(nil).RemoveUserSession), ctx, hash)
}

// UpdateUser mocks base method.
func (m *MockAuthRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAuthRepositoryMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAuthRepository)(nil).UpdateUser), ctx, user)
}

// ValidateUserSession mocks base method.
func (m *MockAuthRepository) ValidateUserSession(ctx context.Context, hash string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUserSession", ctx, hash)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUserSession indicates an expected call of ValidateUserSession.
func (mr *MockAuthRepositoryMockRecorder) ValidateUserSession(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUserSession", reflect.TypeOf((*MockAuthRepository)(nil).ValidateUserSession), ctx, hash)
}
