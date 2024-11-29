// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	traceability "data-spaces-backend/domain/model/traceability"
)

// ICfpCertificationUsecase is an autogenerated mock type for the ICfpCertificationUsecase type
type ICfpCertificationUsecase struct {
	mock.Mock
}

// GetCfpCertification provides a mock function with given fields: c, getCfpCertificationInput
func (_m *ICfpCertificationUsecase) GetCfpCertification(c echo.Context, getCfpCertificationInput traceability.GetCfpCertificationInput) (traceability.CfpCertificationModels, error) {
	ret := _m.Called(c, getCfpCertificationInput)

	if len(ret) == 0 {
		panic("no return value specified for GetCfpCertification")
	}

	var r0 traceability.CfpCertificationModels
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetCfpCertificationInput) (traceability.CfpCertificationModels, error)); ok {
		return rf(c, getCfpCertificationInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetCfpCertificationInput) traceability.CfpCertificationModels); ok {
		r0 = rf(c, getCfpCertificationInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(traceability.CfpCertificationModels)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.GetCfpCertificationInput) error); ok {
		r1 = rf(c, getCfpCertificationInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICfpCertificationUsecase creates a new instance of ICfpCertificationUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICfpCertificationUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICfpCertificationUsecase {
	mock := &ICfpCertificationUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
