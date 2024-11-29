package usecase

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// IPartsUsecase
// Summary: This interface defines use cases for the parts.
//
//go:generate mockery --name IPartsUsecase --output ../test/mock --case underscore
type IPartsUsecase interface {
	GetPartsList(c echo.Context, getPartsInput traceability.GetPartsInput) ([]traceability.PartsModel, *string, error)
	DeleteParts(c echo.Context, getPartsInput traceability.DeletePartsInput) (common.ResponseHeaders, error)
}
