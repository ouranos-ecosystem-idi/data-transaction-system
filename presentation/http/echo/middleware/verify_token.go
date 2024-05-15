package middleware

import (
	"net/http"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/presentation/http/echo/handler"

	"github.com/labstack/echo/v4"
)

// VerifyToken
// Summary: This is function which verifies the token.
// input: h(handler.AppHandler) handler object
// output: (echo.MiddlewareFunc) echo middleware function
func VerifyToken(h handler.AppHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			dataTarget := c.QueryParam("dataTarget")

			token := common.ExtractBearerToken(c)
			if token == "" {
				logger.Set(c).Error(common.Err401Authentication)

				return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusUnauthorized, common.HTTPErrorSourceAuth, common.Err401Authentication, "", dataTarget, method))
			}

			operatorID, err := h.VerifyToken(c)
			if err != nil {
				logger.Set(c).Error(err.Error())

				return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusUnauthorized, common.HTTPErrorSourceAuth, common.Err401InvalidToken, "", dataTarget, method))
			}
			if operatorID == nil {
				return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusUnauthorized, common.HTTPErrorSourceAuth, common.Err401InvalidToken, "", dataTarget, method))
			}
			// Set operatorID to echo context if token is valid
			c.Set("operatorID", *operatorID)

			return next(c)
		}
	}
}
