package handler

import (
	"errors"
	"net/http"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/usecase"

	"github.com/labstack/echo/v4"
)

// IPartsHandler
// Summary: This is interface which defines PartsHandler.
//
//go:generate mockery --name IPartsHandler --output ../../../../test/mock --case underscore
type IPartsHandler interface {
	// GetPartsModel
	// Summary: This is function which defines #8 GetPartsList.
	GetPartsModel(c echo.Context) error
	// PutPartsModel
	// Summary: This is function which defines #5 PutPartsItem.
	PutPartsModel(c echo.Context) error
	// DeletePartsModel
	// Summary: This is function which defines #19 DeletePartsItem.
	DeletePartsModel(c echo.Context) error
}

// partsHandler
// Summary: This is structure which defines partsHandler.
type partsHandler struct {
	partsUsecase          usecase.IPartsUsecase
	partsStructureUsecase usecase.IPartsStructureUsecase
	host                  string
}

// NewPartsHandler
// Summary: This is function to create new partsHandler.
// input: pu(usecase.IPartsUsecase) use case interface
// input: psu(usecase.IPartsStructureUsecase) use case interface
// input: host(string) host name
// output: (IPartsHandler) handler interface
func NewPartsHandler(pu usecase.IPartsUsecase, psu usecase.IPartsStructureUsecase, host string) IPartsHandler {
	return &partsHandler{pu, psu, host}
}

// GetPartsModel
// Summary: This is function which get a list of parts.
// input: c(echo.Context) echo context
// output: (error) Error object
func (h *partsHandler) GetPartsModel(c echo.Context) error {

	var defaultLimit int = 100
	var input traceability.GetPartsInput

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	input.OperatorID = operatorID

	limit, err := common.QueryParamIntPtr(c, "limit", defaultLimit)
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("limit")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	if *limit > defaultLimit {
		logger.Set(c).Warnf(common.LimitUpperError(*limit))
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

	traceID, err := common.QueryParamUUIDPtr(c, "traceId")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("traceId")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	input.TraceID = common.UUIDPtrToStringPtr(traceID)

	input.PartsName = common.QueryParamPtr(c, "partsName")

	plantID, err := common.QueryParamUUIDPtr(c, "plantId")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("plantId")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	input.PlantID = common.UUIDPtrToStringPtr(plantID)

	input.ParentFlag, err = common.QueryParamBoolPtr(c, "parentFlag")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("parentFlag")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	res, afterRes, err := h.partsUsecase.GetPartsList(c, input)
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
	return c.JSON(http.StatusOK, res)
}

// PutPartsModel
// Summary: This is function which update parts information.
// input: c(echo.Context) echo context
// output: (error) Error object
func (h *partsHandler) PutPartsModel(c echo.Context) error {

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	var putPartsInput traceability.PutPartsInput
	if err := c.Bind(&putPartsInput); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.FormatBindErrMsg(err)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}

	if err := putPartsInput.Validate(); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400Validation, operatorID, dataTarget, method, errDetails))
	}
	if operatorID != putPartsInput.OperatorID {
		logger.Set(c).Warnf(common.Err403AccessDenied)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceDataspace, common.Err403AccessDenied, operatorID, dataTarget, method))
	}

	putPartsStructureInput := traceability.PutPartsStructureInput{
		ParentPartsInput: &putPartsInput,
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

	putPartsResponse := res.ParentPartsModel

	common.SetResponseHeader(c, headers)
	return c.JSON(http.StatusCreated, putPartsResponse)
}

// DeletePartsModel
// Summary: This is function which delete parts information.
// input: c(echo.Context) echo context
// output: (error) Error object
func (h *partsHandler) DeletePartsModel(c echo.Context) error {

	var deletePartsInput traceability.DeletePartsInput

	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	traceID, err := common.QueryParamUUID(c, "traceId")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("traceId")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
	deletePartsInput.TraceID = traceID.String()

	headers, err := h.partsUsecase.DeleteParts(c, deletePartsInput)
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

	return c.NoContent(http.StatusNoContent)
}
