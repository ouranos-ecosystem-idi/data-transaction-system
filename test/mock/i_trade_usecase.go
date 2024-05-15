// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	traceability "data-spaces-backend/domain/model/traceability"
)

// ITradeUsecase is an autogenerated mock type for the ITradeUsecase type
type ITradeUsecase struct {
	mock.Mock
}

// GetTradeRequest provides a mock function with given fields: c, getTradeRequestInput
func (_m *ITradeUsecase) GetTradeRequest(c echo.Context, getTradeRequestInput traceability.GetTradeRequestInput) ([]traceability.TradeModel, *string, error) {
	ret := _m.Called(c, getTradeRequestInput)

	var r0 []traceability.TradeModel
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetTradeRequestInput) ([]traceability.TradeModel, *string, error)); ok {
		return rf(c, getTradeRequestInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetTradeRequestInput) []traceability.TradeModel); ok {
		r0 = rf(c, getTradeRequestInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]traceability.TradeModel)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.GetTradeRequestInput) *string); ok {
		r1 = rf(c, getTradeRequestInput)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(echo.Context, traceability.GetTradeRequestInput) error); ok {
		r2 = rf(c, getTradeRequestInput)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTradeResponse provides a mock function with given fields: c, getTradeRequestInput
func (_m *ITradeUsecase) GetTradeResponse(c echo.Context, getTradeRequestInput traceability.GetTradeResponseInput) ([]traceability.TradeResponseModel, *string, error) {
	ret := _m.Called(c, getTradeRequestInput)

	var r0 []traceability.TradeResponseModel
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetTradeResponseInput) ([]traceability.TradeResponseModel, *string, error)); ok {
		return rf(c, getTradeRequestInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.GetTradeResponseInput) []traceability.TradeResponseModel); ok {
		r0 = rf(c, getTradeRequestInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]traceability.TradeResponseModel)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.GetTradeResponseInput) *string); ok {
		r1 = rf(c, getTradeRequestInput)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(echo.Context, traceability.GetTradeResponseInput) error); ok {
		r2 = rf(c, getTradeRequestInput)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PutTradeRequest provides a mock function with given fields: c, tradeRequestModel
func (_m *ITradeUsecase) PutTradeRequest(c echo.Context, tradeRequestModel traceability.TradeRequestModel) (traceability.TradeRequestModel, error) {
	ret := _m.Called(c, tradeRequestModel)

	var r0 traceability.TradeRequestModel
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.TradeRequestModel) (traceability.TradeRequestModel, error)); ok {
		return rf(c, tradeRequestModel)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.TradeRequestModel) traceability.TradeRequestModel); ok {
		r0 = rf(c, tradeRequestModel)
	} else {
		r0 = ret.Get(0).(traceability.TradeRequestModel)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.TradeRequestModel) error); ok {
		r1 = rf(c, tradeRequestModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutTradeResponse provides a mock function with given fields: c, putTradeResponseInput
func (_m *ITradeUsecase) PutTradeResponse(c echo.Context, putTradeResponseInput traceability.PutTradeResponseInput) (traceability.TradeModel, error) {
	ret := _m.Called(c, putTradeResponseInput)

	var r0 traceability.TradeModel
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.PutTradeResponseInput) (traceability.TradeModel, error)); ok {
		return rf(c, putTradeResponseInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.PutTradeResponseInput) traceability.TradeModel); ok {
		r0 = rf(c, putTradeResponseInput)
	} else {
		r0 = ret.Get(0).(traceability.TradeModel)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.PutTradeResponseInput) error); ok {
		r1 = rf(c, putTradeResponseInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewITradeUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewITradeUsecase creates a new instance of ITradeUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITradeUsecase(t mockConstructorTestingTNewITradeUsecase) *ITradeUsecase {
	mock := &ITradeUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}