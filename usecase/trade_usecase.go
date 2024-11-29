package usecase

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// ITradeUsecase
// Summary: This is interface which defines TradeUsecase.
//
//go:generate mockery --name ITradeUsecase --output ../test/mock --case underscore
type ITradeUsecase interface {
	// #10 GetTradeRequestList
	GetTradeRequest(c echo.Context, getTradeRequestInput traceability.GetTradeRequestInput) ([]traceability.TradeModel, *string, error)
	// #12 GetTradeResponseList
	GetTradeResponse(c echo.Context, getTradeRequestInput traceability.GetTradeResponseInput) ([]traceability.TradeResponseModel, *string, error)
	// #7 PutTradeRequestItem
	PutTradeRequest(c echo.Context, putTradeRequestInput traceability.PutTradeRequestInput) (traceability.TradeRequestModel, common.ResponseHeaders, error)
	// #13 PutTradeResponseItem
	PutTradeResponse(c echo.Context, putTradeResponseInput traceability.PutTradeResponseInput) (traceability.TradeModel, common.ResponseHeaders, error)
}
