package router

import (
	"data-spaces-backend/config"
	"data-spaces-backend/presentation/http/echo/handler"
	custom_middleware "data-spaces-backend/presentation/http/echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// SetRouter
// Summary: This is function which sets the router.
// input: e(*echo.Echo) echo
// input: h(handler.AppHandler) handler
// input: config(*config.Config) config
// input: conn(*gorm.DB) gorm database connection
func SetRouter(e *echo.Echo, h handler.AppHandler, config *config.Config, conn *gorm.DB) {
	env := config.Env
	if env == "local" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
		}))
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
		}))
	}

	e.Use(middleware.BodyLimit("25M"))

	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	e.Use(custom_middleware.VerifyAPIKey(h))
	e.Use(custom_middleware.VerifyToken(h))

	e.GET("/api/v1/datatransport", func(c echo.Context) error { return h.GetOuranos(c) })
	e.PUT("/api/v1/datatransport", func(c echo.Context) error { return h.PutOuranos(c) })
}
