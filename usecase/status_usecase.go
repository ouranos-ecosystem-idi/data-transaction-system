package usecase

import (
	"data-spaces-backend/domain/model/traceability"

	"github.com/labstack/echo/v4"
)

// IStatusUsecase
// Summary: This interface defines use cases for the status.
//
//go:generate mockery --name IStatusUsecase --output ../test/mock --case underscore
type IStatusUsecase interface {
	GetStatus(c echo.Context, getStatusModel traceability.GetStatusModel) ([]traceability.StatusModel, *string, error)
	PutStatusCancel(c echo.Context, statusModel traceability.StatusModel) error
	PutStatusReject(c echo.Context, statusModel traceability.StatusModel) error
}
