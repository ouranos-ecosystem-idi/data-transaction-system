package usecase

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// ICfpUsecase
// Summary: This interface defines use cases for the cfp.
//
//go:generate mockery --name ICfpUsecase --output ../test/mock --case underscore
type ICfpUsecase interface {
	GetCfp(c echo.Context, getCfpInput traceability.GetCfpInput) ([]traceability.CfpModel, error)
	PutCfp(c echo.Context, putCfpInputs traceability.PutCfpInputs, operatorID string) ([]traceability.CfpModel, common.ResponseHeaders, error)
}
