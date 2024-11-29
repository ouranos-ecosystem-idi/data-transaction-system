// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	common "data-spaces-backend/domain/common"

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

	if len(ret) == 0 {
		panic("no return value specified for GetTradeRequest")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for GetTradeResponse")
	}

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

// PutTradeRequest provides a mock function with given fields: c, putTradeRequestInput
func (_m *ITradeUsecase) PutTradeRequest(c echo.Context, putTradeRequestInput traceability.PutTradeRequestInput) (traceability.TradeRequestModel, common.ResponseHeaders, error) {
	ret := _m.Called(c, putTradeRequestInput)

	if len(ret) == 0 {
		panic("no return value specified for PutTradeRequest")
	}

	var r0 traceability.TradeRequestModel
	var r1 common.ResponseHeaders
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.PutTradeRequestInput) (traceability.TradeRequestModel, common.ResponseHeaders, error)); ok {
		return rf(c, putTradeRequestInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.PutTradeRequestInput) traceability.TradeRequestModel); ok {
		r0 = rf(c, putTradeRequestInput)
	} else {
		r0 = ret.Get(0).(traceability.TradeRequestModel)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.PutTradeRequestInput) common.ResponseHeaders); ok {
		r1 = rf(c, putTradeRequestInput)
	} else {
		r1 = ret.Get(1).(common.ResponseHeaders)
	}

	if rf, ok := ret.Get(2).(func(echo.Context, traceability.PutTradeRequestInput) error); ok {
		r2 = rf(c, putTradeRequestInput)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PutTradeResponse provides a mock function with given fields: c, putTradeResponseInput
func (_m *ITradeUsecase) PutTradeResponse(c echo.Context, putTradeResponseInput traceability.PutTradeResponseInput) (traceability.TradeModel, common.ResponseHeaders, error) {
	ret := _m.Called(c, putTradeResponseInput)

	if len(ret) == 0 {
		panic("no return value specified for PutTradeResponse")
	}

	var r0 traceability.TradeModel
	var r1 common.ResponseHeaders
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.PutTradeResponseInput) (traceability.TradeModel, common.ResponseHeaders, error)); ok {
		return rf(c, putTradeResponseInput)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, traceability.PutTradeResponseInput) traceability.TradeModel); ok {
		r0 = rf(c, putTradeResponseInput)
	} else {
		r0 = ret.Get(0).(traceability.TradeModel)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, traceability.PutTradeResponseInput) common.ResponseHeaders); ok {
		r1 = rf(c, putTradeResponseInput)
	} else {
		r1 = ret.Get(1).(common.ResponseHeaders)
	}

	if rf, ok := ret.Get(2).(func(echo.Context, traceability.PutTradeResponseInput) error); ok {
		r2 = rf(c, putTradeResponseInput)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewITradeUsecase creates a new instance of ITradeUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITradeUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITradeUsecase {
	mock := &ITradeUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
