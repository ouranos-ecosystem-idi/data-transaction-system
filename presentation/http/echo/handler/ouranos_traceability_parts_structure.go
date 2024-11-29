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

// IPartsStructureHandler
// Summary: This is interface which defines PartsStructureHandler.
//
//go:generate mockery --name IPartsStructureHandler --output ../../../../test/mock --case underscore
type IPartsStructureHandler interface {
	// #9 GetPartsStructureItem.
	GetPartsStructureModel(c echo.Context) error
	// #6 PutPartsStructureItem.
	PutPartsStructureModel(c echo.Context) error
}

// partsStructureHandler
// Summary: This is structure which defines partsStructureHandler.
type partsStructureHandler struct {
	partsStructureUsecase usecase.IPartsStructureUsecase
}

// NewPartsStructureHandler
// Summary: This is function to create new partsStructureHandler.
// input: u(usecase.IPartsStructureUsecase) use case interface
// output: (IPartsStructureHandler) handler interface
func NewPartsStructureHandler(u usecase.IPartsStructureUsecase) IPartsStructureHandler {
	return &partsStructureHandler{u}
}

// GetPartsStructureModel
// Summary: This is function which get the request and response partsStructure.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *partsStructureHandler) GetPartsStructureModel(c echo.Context) error {

	var (
		traceID    uuid.UUID
		operatorID string
		err        error
	)
	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID = c.Get("operatorID").(string)

	traceID, err = common.QueryParamUUID(c, "traceId")
	if err != nil {
		logger.Set(c).Warn(err.Error())
		errDetails := common.UnexpectedQueryParameter("traceId")
		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	getPartsStructureInput := traceability.GetPartsStructureInput{
		TraceID:    traceID,
		OperatorID: operatorID,
	}

	getPartsStructure, err := h.partsStructureUsecase.GetPartsStructure(c, getPartsStructureInput)
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
	return c.JSON(http.StatusOK, getPartsStructure)
}

// PutPartsStructureModel
// Summary: This is function which put the request and response partsStructure.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *partsStructureHandler) PutPartsStructureModel(c echo.Context) error {

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method
	operatorID := c.Get("operatorID").(string)

	var putPartsStructureInput traceability.PutPartsStructureInput
	if err := c.Bind(&putPartsStructureInput); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.FormatBindErrMsg(err)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	if err := putPartsStructureInput.Validate(); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	isValidOperatorID := true
	if putPartsStructureInput.ParentPartsInput.OperatorID != operatorID {
		isValidOperatorID = false
	}
	for _, childPartsInput := range *putPartsStructureInput.ChildrenPartsInput {
		if childPartsInput.OperatorID != operatorID {
			isValidOperatorID = false
		}
	}

	if !isValidOperatorID {
		logger.Set(c).Warnf(common.Err403AccessDenied)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceDataspace, common.Err403AccessDenied, operatorID, dataTarget, method))
	}

	res, headers, err := h.partsStructureUsecase.PutPartsStructure(c, putPartsStructureInput)
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
