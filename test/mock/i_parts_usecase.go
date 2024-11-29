// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	common "data-spaces-backend/domain/common"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	traceability "data-spaces-backend/domain/model/traceability"
)

// IPartsUsecase is an autogenerated mock type for the IPartsUsecase type
type IPartsUsecase struct {
	mock.Mock
}

// DeleteParts provides a mock function with given fields: c, getPartsInput
func (_m *IPartsUsecase) DeleteParts(c echo.Context, getPartsInput traceability.DeletePartsInput) (common.ResponseHeaders, error) {
	ret := _m.Called(c, getPartsInput)

	if len(ret) == 0 {
		panic("no return value specified for DeleteParts")
	}

	var r0 common.ResponseHeaders
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.DeletePartsInput) (common.ResponseHeaders, error)); ok {
		return rf(c, getPartsInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.DeletePartsInput) common.ResponseHeaders); ok {
		r0 = rf(c, getPartsInput)
	} else {
		r0 = ret.Get(0).(common.ResponseHeaders)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.DeletePartsInput) error); ok {
		r1 = rf(c, getPartsInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPartsList provides a mock function with given fields: c, getPartsInput
func (_m *IPartsUsecase) GetPartsList(c echo.Context, getPartsInput traceability.GetPartsInput) ([]traceability.PartsModel, *string, error) {
	ret := _m.Called(c, getPartsInput)

	if len(ret) == 0 {
		panic("no return value specified for GetPartsList")
	}

	var r0 []traceability.PartsModel
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetPartsInput) ([]traceability.PartsModel, *string, error)); ok {
		return rf(c, getPartsInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetPartsInput) []traceability.PartsModel); ok {
		r0 = rf(c, getPartsInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]traceability.PartsModel)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.GetPartsInput) *string); ok {
		r1 = rf(c, getPartsInput)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(echo.Context, traceability.GetPartsInput) error); ok {
		r2 = rf(c, getPartsInput)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewIPartsUsecase creates a new instance of IPartsUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPartsUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPartsUsecase {
	mock := &IPartsUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
