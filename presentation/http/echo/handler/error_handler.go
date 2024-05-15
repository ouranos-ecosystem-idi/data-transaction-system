package handler

import (
	"errors"
	"net/http"
	"syscall"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CustomHTTPErrorHandler
// Summary: This is function which handles the custom HTTP error.
// input: err(error) error object
// input: c(echo.Context) echo context
func CustomHTTPErrorHandler(err error, c echo.Context) {

	method := c.Request().Method
	dataTarget := c.QueryParam("dataTarget")
	// JWT token does not exist
	if errors.Is(err, middleware.ErrJWTMissing) {
		c.Error(echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusUnauthorized, common.HTTPErrorSourceAuth, common.Err401Authentication, "", dataTarget, method)))

		return
	}
	if errors.Is(err, syscall.ECONNREFUSED) {
		c.Error(echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusServiceUnavailable, common.HTTPErrorSourceDataspace, common.Err503OuterService, "", dataTarget, method)))

		return
	}
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Code == http.StatusNotFound && he.Message == "Not Found" {
			c.Error(echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusNotFound, common.HTTPErrorSourceDataspace, common.Err404EndpointNotFound, "", dataTarget, method)))

			return
		}
		if he.Code >= 400 && he.Code < 500 {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}
	} else if pe, ok := err.(*pgconn.PgError); ok {
		if pe.Code == common.PgErrorAdminShutdown || pe.Code == common.PgErrorCrashShutdown {
			c.Error(echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusServiceUnavailable, common.HTTPErrorSourceDataspace, common.Err503OuterService, "", dataTarget, method)))

			return
		} else {
			c.Error(echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, "", dataTarget, method)))

			return
		}
	} else {
		c.Error(echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceDataspace, common.Err500Unexpected, "", dataTarget, method)))

		return
	}

	c.Echo().DefaultHTTPErrorHandler(err, c)
}
