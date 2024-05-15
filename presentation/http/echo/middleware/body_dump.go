package middleware

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"

	"github.com/labstack/echo/v4"
)

// DumpHandler
// Summary: This is function which dumps the request and response body
// input: c(echo.Context) context object
// input: reqBody([]byte) request body
// input: resBody([]byte) response body
func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	var operatorID string
	i := c.Get("operatorID")
	if i != nil {
		operatorID = i.(string)
	}

	header := c.Request().Header
	for k := range header {
		if k == "Authorization" {
			header[k] = []string{"Bearer ******"}
		}
		if k == "apiKey" {
			header[k] = []string{"******"}
		}
	}

	logger.Set(c).Debugf(logger.AccessDebugLog, c.Request().URL.String(), operatorID, c.Request().Header, string(reqBody), string(resBody))
}

// DumpSkipper
// Summary: This is function which skips the dump
// input: c(echo.Context) context object
// output: (bool) boolean value
func dumpSkipper(c echo.Context) bool {
	return !common.IsOutputDump()
}
