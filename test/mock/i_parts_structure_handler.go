// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IPartsStructureHandler is an autogenerated mock type for the IPartsStructureHandler type
type IPartsStructureHandler struct {
	mock.Mock
}

// GetPartsStructureModel provides a mock function with given fields: c
func (_m *IPartsStructureHandler) GetPartsStructureModel(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetPartsStructureModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutPartsStructureModel provides a mock function with given fields: c
func (_m *IPartsStructureHandler) PutPartsStructureModel(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for PutPartsStructureModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIPartsStructureHandler creates a new instance of IPartsStructureHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPartsStructureHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPartsStructureHandler {
	mock := &IPartsStructureHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
