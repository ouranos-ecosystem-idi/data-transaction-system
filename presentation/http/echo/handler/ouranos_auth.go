package handler

import (
	"net/http"
	"os"
	"strings"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/usecase/input"

	"github.com/labstack/echo/v4"
)

// VerifyAPIKey
// Summary: This is function which verifies the API key.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *authHandler) VerifyAPIKey(c echo.Context) error {
	method := c.Request().Method
	dataTarget := c.QueryParam("dataTarget")

	env := os.Getenv("GO_ENV")
	var ip string
	if env == "local" {
		ip = c.RealIP()
	} else {
		xff := c.Request().Header.Get("X-Forwarded-For") // X-Forwarded-For: APP_IP, ALB_IP
		ips := strings.Split(xff, ", ")
		if len(ips) < 2 {
			return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceAuth, common.Err403IPNotAuthorizedForKey, "", "", method))
		}
		ip = ips[len(ips)-2]
	}

	apiKey := c.Request().Header.Get("apiKey")
	if apiKey == "" {
		logger.Set(c).Errorf(common.Err403AccessDenied)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceAuth, common.Err403AccessDenied, "", dataTarget, method))
	}

	input := input.VerifyAPIKey{
		APIKey:    apiKey,
		IPAddress: ip,
	}
	res, err := h.VerifyUsecase.VerifyAPIKey(input)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceAuth, common.Err500Unexpected, "", dataTarget, method))
	}
	if !res.IsAPIKeyValid {
		logger.Set(c).Errorf(common.Err403AccessDenied)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceAuth, common.Err403InvalidKey, "", dataTarget, method))
	}
	if !res.IsIPAddressValid {
		logger.Set(c).Errorf(common.Err403IPNotAuthorizedForKey)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusForbidden, common.HTTPErrorSourceAuth, common.Err403IPNotAuthorizedForKey, "", dataTarget, method))
	}
	return nil
}

// VerifyToken
// Summary: This is function which verifies the token.
// input: c(echo.Context) echo context
// output: (*string) operator ID
// output: (error) error object
func (h *authHandler) VerifyToken(c echo.Context) (*string, error) {
	method := c.Request().Method

	token := common.ExtractBearerToken(c)
	input := input.VerifyToken{
		Token: token,
	}

	output, err := h.VerifyUsecase.VerifyToken(input)
	if err != nil {
		return nil, echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err401InvalidToken, "", "", method))
	}

	return output.OperatorID, nil
}
