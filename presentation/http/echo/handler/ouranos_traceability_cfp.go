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

// ICfpHandler
// Summary: This is interface which defines CfpHandler
//
//go:generate mockery --name ICfpHandler --output ../../../../test/mock --case underscore
type ICfpHandler interface {
	// #14 GetCfpList
	GetCfp(c echo.Context) error
	// #15 PutCfpList
	PutCfp(c echo.Context) error
}

// cfpHandler
// Summary: This is structure which defines cfpHandler.
type cfpHandler struct {
	cfpUsecase usecase.ICfpUsecase
}

// NewCfpHandler
// Summary: This is function to create new cfpHandler.
// input: u(usecase.ICfpUsecase) use case interface
// output: (ICfpHandler) handler interface
func NewCfpHandler(u usecase.ICfpUsecase) ICfpHandler {
	return &cfpHandler{u}
}

// GetCfp
// Summary: This is function which get a list of cfp.
// input: c(echo.Context) echo context
// output: (error) error object
func (h cfpHandler) GetCfp(c echo.Context) error {
	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	OperatorUUID, err := uuid.Parse(operatorID)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err401InvalidToken, operatorID, dataTarget, method))
	}

	traceIDs, err := common.QueryParamUUIDs(c, "traceIds")
	if err != nil || traceIDs == nil {
		errDetails := common.UnexpectedQueryParameter("traceIds")
		logger.Set(c).Warnf(errDetails)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	if len(traceIDs) > 50 {
		logger.Set(c).Warnf(common.TraceIDsUpperLimitError(len(traceIDs)))
		errDetails := common.UnexpectedQueryParameter("traceIds")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	getCfpInput := traceability.GetCfpInput{
		OperatorID: OperatorUUID,
		TraceIDs:   traceIDs,
	}

	res, err := h.cfpUsecase.GetCfp(c, getCfpInput)
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
	common.SetResponseHeader(c, common.ResponseHeaders{})
	return c.JSON(http.StatusOK, res)
}

// PutCfp
// Summary: This is function which put a list of cfp.
// input: c(echo.Context) echo context
// output: (error) error object
func (h cfpHandler) PutCfp(c echo.Context) error {
	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)

	var input traceability.PutCfpInputs
	if err := c.Bind(&input); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.FormatBindErrMsg(err)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	if err := input.Validate(); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	res, headers, err := h.cfpUsecase.PutCfp(c, input, operatorID)
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
	return c.JSON(http.StatusCreated, res)
}
