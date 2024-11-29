package common

import "github.com/labstack/echo/v4"

const ResponseHeaderXTrack = "X-Track"

type ResponseHeaders struct {
	XTrack string
}

func SetResponseHeader(c echo.Context, headers ResponseHeaders) {
	c.Response().Header()[ResponseHeaderXTrack] = []string{headers.XTrack}
}
