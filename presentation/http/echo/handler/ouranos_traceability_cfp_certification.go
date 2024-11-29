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

// ICfpCertificationHandler
// Summary: This is interface which defines CfpCertificationHandler
//
//go:generate mockery --name ICfpCertificationHandler --output ../../../../test/mock --case underscore
type ICfpCertificationHandler interface {
	GetCfpCertification(c echo.Context) error
}

// cfpCertificationHandler
// Summary: This is structure which defines cfpCertificationHandler.
type cfpCertificationHandler struct {
	cfpCertificationUsecase usecase.ICfpCertificationUsecase
}

// NewCfpCertificationHandler
// Summary: This is function to create new cfpCertificationHandler.
// input: u(usecase.ICfpCertificationUsecase) use case interface
// output: (ICfpCertificationHandler) handler interface
func NewCfpCertificationHandler(u usecase.ICfpCertificationUsecase) ICfpCertificationHandler {
	return &cfpCertificationHandler{u}
}

// GetCfpCertification
// Summary: This is function which get cfp certification.
// input: c(echo.Context) echo context
// output: (error) error
func (h cfpCertificationHandler) GetCfpCertification(c echo.Context) error {
	dataTarget := c.QueryParam("dataTarget")
	method := c.Request().Method

	operatorID := c.Get("operatorID").(string)
	OperatorUUID, err := uuid.Parse(operatorID)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err401InvalidToken, operatorID, dataTarget, method))
	}

	traceID, err := common.QueryParamUUID(c, "traceId")
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.UnexpectedQueryParameter("traceId")

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}

	getCfpCertificationInput := traceability.GetCfpCertificationInput{
		OperatorID: OperatorUUID,
		TraceID:    traceID,
	}
	res, err := h.cfpCertificationUsecase.GetCfpCertification(c, getCfpCertificationInput)
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
