package handler

import (
	"data-spaces-backend/usecase"

	"github.com/labstack/echo/v4"
)

type (
	AuthHandler interface {
		VerifyAPIKey(c echo.Context) error
		VerifyToken(c echo.Context) (*string, error)
	}

	authHandler struct {
		VerifyUsecase usecase.IVerifyUsecase
	}
)

func NewAuthHandler(
	verifyUsecase usecase.IVerifyUsecase,
) AuthHandler {
	return &authHandler{
		VerifyUsecase: verifyUsecase,
	}
}
