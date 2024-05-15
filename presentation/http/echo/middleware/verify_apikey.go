package middleware

import (
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/presentation/http/echo/handler"

	"github.com/labstack/echo/v4"
)

// VerifyAPIKey
// Summary: This is function which verifies the API key.
// input: h(handler.AppHandler) handler object
// output: (echo.MiddlewareFunc) echo middleware function
func VerifyAPIKey(h handler.AppHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := h.VerifyAPIKey(c)
			if err != nil {
				logger.Set(c).Error(err.Error())

				return err
			}
			return next(c)
		}
	}
}
