// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	traceability "data-spaces-backend/domain/model/traceability"
)

// IPartsUsecase is an autogenerated mock type for the IPartsUsecase type
type IPartsUsecase struct {
	mock.Mock
}

// GetPartsList provides a mock function with given fields: c, getPlantPartsModel
func (_m *IPartsUsecase) GetPartsList(c echo.Context, getPlantPartsModel traceability.GetPartsModel) ([]traceability.PartsModel, *string, error) {
	ret := _m.Called(c, getPlantPartsModel)

	var r0 []traceability.PartsModel
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetPartsModel) ([]traceability.PartsModel, *string, error)); ok {
		return rf(c, getPlantPartsModel)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetPartsModel) []traceability.PartsModel); ok {
		r0 = rf(c, getPlantPartsModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]traceability.PartsModel)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.GetPartsModel) *string); ok {
		r1 = rf(c, getPlantPartsModel)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(echo.Context, traceability.GetPartsModel) error); ok {
		r2 = rf(c, getPlantPartsModel)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewIPartsUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPartsUsecase creates a new instance of IPartsUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPartsUsecase(t mockConstructorTestingTNewIPartsUsecase) *IPartsUsecase {
	mock := &IPartsUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}