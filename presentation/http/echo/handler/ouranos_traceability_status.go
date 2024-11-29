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

// IStatusHandler
// Summary: This is interface which defines StatusHandler.
//
//go:generate mockery --name IStatusHandler --output ../../../../test/mock --case underscore
type IStatusHandler interface {
	// #11 GetStatusList.
	GetStatus(c echo.Context) error
	// #16 PutStatusItem.
	PutStatus(c echo.Context) error
}

// statusHandler
// Summary: This is structure which defines statusHandler.
type statusHandler struct {
	statusUsecase usecase.IStatusUsecase
	host          string
}

// NewStatusHandler
// Summary: This is function to create new statusHandler.
// input: u(usecase.IStatusUsecase) use case interface
// input: host(string) host name
// output: (IStatusHandler) handler interface
func NewStatusHandler(u usecase.IStatusUsecase, host string) IStatusHandler {
	return &statusHandler{u, host}
}

// GetStatus
// Summary: This is function which get a list of request and response status.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *statusHandler) GetStatus(c echo.Context) error {
	var defaultLimit int = 100
	var input traceability.GetStatusInput

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorId := c.Get("operatorID").(string)
	OperatorUUID, err := uuid.Parse(operatorId)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err401InvalidToken, operatorId, dataTarget, method))
	}
	input.OperatorID = OperatorUUID

	limit, err := common.QueryParamIntPtr(c, "limit", defaultLimit)
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	if *limit > defaultLimit {
		logger.Set(c).Warnf(common.LimitUpperError(*limit))
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	if *limit <= 0 {
		logger.Set(c).Warnf(common.LimitLessThanError(0, *limit))
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	input.Limit = *limit

	statusTarget := c.QueryParam("statusTarget")
	StatusTarget, err := traceability.NewStatusTarget(statusTarget)
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("statusTarget")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	input.StatusTarget = StatusTarget

	statusID, err := common.QueryParamUUIDPtr(c, "statusId")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("statusId")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	if input.StatusTarget == traceability.StatusTarget("") || input.StatusTarget == traceability.Response {
		input.StatusID = statusID
	}

	traceID, err := common.QueryParamUUIDPtr(c, "traceId")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("traceId")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	if input.StatusTarget == traceability.Request {
		input.TraceID = traceID
	}

	after, err := common.QueryParamUUIDPtr(c, "after")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("after")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorId, dataTarget, method, errDetails))
	}
	if input.StatusID == nil && input.TraceID == nil {
		input.After = after
	}

	response, afterRes, err := h.statusUsecase.GetStatus(c, input)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) {
			if customErr.IsWarn() {
				logger.Set(c).Warnf(err.Error())
			} else {
				logger.Set(c).Errorf(err.Error())
			}

			return echo.NewHTTPError(common.HTTPErrorGenerate(int(customErr.Code), customErr.Source, customErr.Message, operatorId, dataTarget, method, *customErr.MessageDetail))
		}
		logger.Set(c).Errorf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, operatorId, dataTarget, method))
	}

	if afterRes != nil {
		c.Response().Header().Set("Link", common.CreateAfterLink(h.host, dataTarget, *afterRes, input))
	}

	common.SetResponseHeader(c, common.ResponseHeaders{})
	return c.JSON(http.StatusOK, response)
}

// PutStatus
// Summary: This is function to update the status to Cancel or Reject.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *statusHandler) PutStatus(c echo.Context) error {
	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	var input traceability.PutStatusInput
	if err := c.Bind(&input); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.FormatBindErrMsg(err)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	if err := input.ValidateForCancelOrReject(); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	var headers common.ResponseHeaders
	var err error
	if input.IsCfpRequestStatusCancel() {
		headers, err = h.statusUsecase.PutStatusCancel(c, input)
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
	}

	if input.IsCfpRequestStatusReject() {
		headers, err = h.statusUsecase.PutStatusReject(c, input)
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
	}

	common.SetResponseHeader(c, headers)
	return c.JSON(http.StatusCreated, common.EmptyBody{})
}
