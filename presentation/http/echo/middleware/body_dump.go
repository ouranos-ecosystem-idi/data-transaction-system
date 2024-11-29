package middleware

import (
	"data-spaces-backend/extension/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// DumpHandler
// Summary: This is function which dumps the request and response body
// input: c(echo.Context) context object
// input: reqBody([]byte) request body
// input: resBody([]byte) response body
func dumpHandler(c echo.Context, reqBody, resBody []byte) {

	if zap.S().Level() == zap.DebugLevel {
		logger.Set(c).Debugf(logger.DataSpaceAPILog, c.Request().URL.String(), c.Request().Header, string(reqBody), string(resBody), c.Response().Header().Get("X-Track"))
	} else {
		header := c.Request().Header
		for k := range header {
			if k == "Authorization" {
				header[k] = []string{"Bearer ******"}
			}
		}
		logger.Set(c).Infof(logger.DataSpaceAPILog, c.Request().URL.String(), c.Request().Header, "******", "******", c.Response().Header().Get("X-Track"))
	}
}
