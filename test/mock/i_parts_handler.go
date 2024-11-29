// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IPartsHandler is an autogenerated mock type for the IPartsHandler type
type IPartsHandler struct {
	mock.Mock
}

// DeletePartsModel provides a mock function with given fields: c
func (_m *IPartsHandler) DeletePartsModel(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for DeletePartsModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPartsModel provides a mock function with given fields: c
func (_m *IPartsHandler) GetPartsModel(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetPartsModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutPartsModel provides a mock function with given fields: c
func (_m *IPartsHandler) PutPartsModel(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for PutPartsModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIPartsHandler creates a new instance of IPartsHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPartsHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPartsHandler {
	mock := &IPartsHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
