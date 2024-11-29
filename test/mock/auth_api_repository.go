// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	authentication "data-spaces-backend/domain/model/authentication"

	mock "github.com/stretchr/testify/mock"

	repository "data-spaces-backend/domain/repository"
)

// AuthAPIRepository is an autogenerated mock type for the AuthAPIRepository type
type AuthAPIRepository struct {
	mock.Mock
}

// VerifyAPIKey provides a mock function with given fields: request
func (_m *AuthAPIRepository) VerifyAPIKey(request repository.VerifyAPIKeyBody) (authentication.VeriryAPIKeyResponse, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for VerifyAPIKey")
	}

	var r0 authentication.VeriryAPIKeyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.VerifyAPIKeyBody) (authentication.VeriryAPIKeyResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(repository.VerifyAPIKeyBody) authentication.VeriryAPIKeyResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(authentication.VeriryAPIKeyResponse)
	}

	if rf, ok := ret.Get(1).(func(repository.VerifyAPIKeyBody) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: request
func (_m *AuthAPIRepository) VerifyToken(request repository.VerifyTokenBody) (authentication.VeriryTokenResponse, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 authentication.VeriryTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.VerifyTokenBody) (authentication.VeriryTokenResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(repository.VerifyTokenBody) authentication.VeriryTokenResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(authentication.VeriryTokenResponse)
	}

	if rf, ok := ret.Get(1).(func(repository.VerifyTokenBody) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthAPIRepository creates a new instance of AuthAPIRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthAPIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthAPIRepository {
	mock := &AuthAPIRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
