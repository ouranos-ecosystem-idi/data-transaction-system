package handler

import (
	"errors"
	"net/http"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ITradeHandler
// Summary: This is interface which defines TradeHandler.
//
//go:generate mockery --name ITradeHandler --output ../../../../test/mock --case underscore
type ITradeHandler interface {
	// #10 GetTradeRequestList
	GetTradeRequest(c echo.Context) error
	// #12 GetTradeResponseList
	GetTradeResponse(c echo.Context) error
	// #7 PutTradeRequestItem.
	PutTradeRequest(c echo.Context) error
	// #13 PutTradeResponseItem.
	PutTradeResponse(c echo.Context) error
}

// tradeHandler
// Summary: This is structure which defines tradeHandler.
type tradeHandler struct {
	tradeUsecase usecase.ITradeUsecase
	host         string
}

// NewTradeHandler
// Summary: This is function to create new tradeHandler.
// input: u(usecase.ITradeUsecase) use case interface
// input: host(string) value of host
// output: (ITradeHandler) handler interface
func NewTradeHandler(u usecase.ITradeUsecase, host string) ITradeHandler {
	return &tradeHandler{u, host}
}

// GetTradeRequest
// Summary: This is function which get a list of trade and response status.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *tradeHandler) GetTradeRequest(c echo.Context) error {

	var defaultLimit int = 100
	var upperTraceIDs int = 50
	var input traceability.GetTradeRequestInput

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	OperatorUUID, err := uuid.Parse(operatorID)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err401InvalidToken, operatorID, dataTarget, method))
	}
	input.OperatorID = OperatorUUID

	// Get request parameters
	limit, err := common.QueryParamIntPtr(c, "limit", defaultLimit)
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	if *limit > defaultLimit {
		logger.Set(c).Warnf(common.LimitUpperError(defaultLimit))
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	if *limit <= 0 {
		logger.Set(c).Warnf(common.LimitLessThanError(0, *limit))
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	input.Limit = *limit

	// Obtain if traceIds is specified in the query parameter
	input.TraceIDs, err = common.QueryParamUUIDs(c, "traceIds")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("traceIds")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	// Error if more than 50 traceIds are specified
	if len(input.TraceIDs) > upperTraceIDs {
		logger.Set(c).Warnf(common.LimitUpperError(upperTraceIDs))
		errDetails := common.UnexpectedQueryParameter("traceIds")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	After, err := common.QueryParamUUIDPtr(c, "after")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("after")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	// Set after only when traceId is not specified
	if len(input.TraceIDs) == 0 {
		input.After = After
	} else {
		// If traceId is specified, set limit to default value
		input.Limit = defaultLimit
	}

	response, afterRes, err := h.tradeUsecase.GetTradeRequest(c, input)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) {
			if customErr.IsWarn() {
				logger.Set(c).Warnf(err.Error())
			} else {
				logger.Set(c).Errorf(err.Error())
			}

			return echo.NewHTTPError(common.HTTPErrorGenerate(int(customErr.Code), customErr.Source, customErr.Message, operatorID, dataTarget, method, *customErr.MessageDetail))
		}
		logger.Set(c).Errorf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, operatorID, dataTarget, method))
	}

	if afterRes != nil {
		c.Response().Header().Set("Link", common.CreateAfterLink(h.host, dataTarget, *afterRes, input))
	}

	common.SetResponseHeader(c, common.ResponseHeaders{})
	return c.JSON(http.StatusOK, response)
}

// GetTradeResponse
// Summary: This is function which get a list of trade and response status.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *tradeHandler) GetTradeResponse(c echo.Context) error {

	var defaultLimit int = 100
	var input traceability.GetTradeResponseInput

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	OperatorUUID, err := uuid.Parse(operatorID)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err401InvalidToken, operatorID, dataTarget, method))
	}
	input.OperatorID = OperatorUUID

	// Get request parameters
	limit, err := common.QueryParamIntPtr(c, "limit", defaultLimit)
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	if *limit > defaultLimit {
		logger.Set(c).Warnf(common.LimitUpperError(defaultLimit))
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	if *limit <= 0 {
		logger.Set(c).Warnf(common.LimitLessThanError(0, *limit))
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	input.Limit = *limit

	input.After, err = common.QueryParamUUIDPtr(c, "after")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("after")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	response, afterRes, err := h.tradeUsecase.GetTradeResponse(c, input)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) {
			if customErr.IsWarn() {
				logger.Set(c).Warnf(err.Error())
			} else {
				logger.Set(c).Errorf(err.Error())
			}

			return echo.NewHTTPError(common.HTTPErrorGenerate(int(customErr.Code), customErr.Source, customErr.Message, operatorID, dataTarget, method, *customErr.MessageDetail))
		}
		logger.Set(c).Errorf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, operatorID, dataTarget, method))
	}

	if afterRes != nil {
		c.Response().Header().Set("Link", common.CreateAfterLink(h.host, dataTarget, *afterRes, input))
	}

	common.SetResponseHeader(c, common.ResponseHeaders{})
	return c.JSON(http.StatusOK, response)
}

// PutTradeRequest
// Summary: This is function which put trade and get response status.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *tradeHandler) PutTradeRequest(c echo.Context) error {

	var putTradeRequestInput traceability.PutTradeRequestInput
	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	// Obtaining an authentication token
	operatorID := c.Get("operatorID").(string)
	// Get request body
	if err := c.Bind(&putTradeRequestInput); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.FormatBindErrMsg(err)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	// Validate the obtained RequestBody
	if err := putTradeRequestInput.Validate(); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}
	// Error if the business ID specified in TradeModel does not specify its own business ID
	if putTradeRequestInput.Trade.DownstreamOperatorID != operatorID {
		logger.Set(c).Warnf(common.Err403AccessDenied)
		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceDataspace, common.Err403AccessDenied, operatorID, dataTarget, method))
	}

	response, headers, err := h.tradeUsecase.PutTradeRequest(c, putTradeRequestInput)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) {
			if customErr.IsWarn() {
				logger.Set(c).Warnf(err.Error())
			} else {
				logger.Set(c).Errorf(err.Error())
			}

			return echo.NewHTTPError(common.HTTPErrorGenerate(int(customErr.Code), customErr.Source, customErr.Message, operatorID, dataTarget, method, *customErr.MessageDetail))
		}
		logger.Set(c).Errorf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, operatorID, dataTarget, method))
	}

	common.SetResponseHeader(c, headers)
	return c.JSON(http.StatusCreated, response)
}

// PutTradeResponse
// Summary: This is function which put trade and get response status.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *tradeHandler) PutTradeResponse(c echo.Context) error {

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	operatorUUID, err := uuid.Parse(operatorID)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err401InvalidToken, operatorID, dataTarget, method))
	}

	tradeID, err := common.QueryParamUUID(c, "tradeId")
	if err != nil {
		logger.Set(c).Warn(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	traceID, err := common.QueryParamUUID(c, "traceId")
	if err != nil {
		logger.Set(c).Warn(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	putTradeResponseInput := traceability.PutTradeResponseInput{
		OperatorID: operatorUUID,
		TradeID:    tradeID,
		TraceID:    traceID,
	}

	response, headers, err := h.tradeUsecase.PutTradeResponse(c, putTradeResponseInput)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) {
			if customErr.IsWarn() {
				logger.Set(c).Warnf(err.Error())
			} else {
				logger.Set(c).Errorf(err.Error())
			}

			return echo.NewHTTPError(common.HTTPErrorGenerate(int(customErr.Code), customErr.Source, customErr.Message, operatorID, dataTarget, method, *customErr.MessageDetail))
		}
		logger.Set(c).Errorf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, operatorID, dataTarget, method))
	}

	common.SetResponseHeader(c, headers)
	return c.JSON(http.StatusCreated, response)
}
