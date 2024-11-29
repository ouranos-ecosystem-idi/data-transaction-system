package usecase

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// IPartsStructureUsecase
// Summary: This interface defines use cases for the partsStructure.
//
//go:generate mockery --name IPartsStructureUsecase --output ../test/mock --case underscore
type IPartsStructureUsecase interface {
	GetPartsStructure(c echo.Context, getPartsStructureModel traceability.GetPartsStructureInput) (traceability.PartsStructureModel, error)
	PutPartsStructure(c echo.Context, putpartsStructureInput traceability.PutPartsStructureInput) (traceability.PartsStructureModel, common.ResponseHeaders, error)
}
