package handler

import (
	"data-spaces-backend/domain/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	HealthCheckHandler interface {
		HealthCheck(c echo.Context) error
	}

	healthCheckHandler struct {
	}
)

// NewHealthCheckHandler
// Summary: This is function to create new healthCheckHandler.
// output: (HealthCheckHandler) handler interface
func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}

// HealthCheck
// Summary: This is function which performs a helth check.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *healthCheckHandler) HealthCheck(c echo.Context) error {
	healthCheckResponse := common.HealthCheckResponse{
		IsSystemHealthy: true,
	}
	return c.JSON(http.StatusOK, healthCheckResponse)
}
