package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewMiddleware
// Summary: This is function which creates new Middleware.
// input: e(*echo.Echo) echo
func NewMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: dumpHandler,
	}))
}
