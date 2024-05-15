package usecase

import (
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// IPartsUsecase
// Summary: This interface defines use cases for the parts.
//
//go:generate mockery --name IPartsUsecase --output ../test/mock --case underscore
type IPartsUsecase interface {
	GetPartsList(c echo.Context, getPlantPartsModel traceability.GetPartsModel) ([]traceability.PartsModel, *string, error)
}
