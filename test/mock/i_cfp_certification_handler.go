// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ICfpCertificationHandler is an autogenerated mock type for the ICfpCertificationHandler type
type ICfpCertificationHandler struct {
	mock.Mock
}

// GetCfpCertification provides a mock function with given fields: c
func (_m *ICfpCertificationHandler) GetCfpCertification(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetCfpCertification")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewICfpCertificationHandler creates a new instance of ICfpCertificationHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICfpCertificationHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICfpCertificationHandler {
	mock := &ICfpCertificationHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}