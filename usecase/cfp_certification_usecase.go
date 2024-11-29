package usecase

import (
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// ICfpCertificationUsecase
// Summary: This interface defines use cases for the cfp certification.
//
//go:generate mockery --name ICfpCertificationUsecase --output ../test/mock --case underscore
type ICfpCertificationUsecase interface {
	GetCfpCertification(c echo.Context, getCfpCertificationInput traceability.GetCfpCertificationInput) (traceability.CfpCertificationModels, error)
}
