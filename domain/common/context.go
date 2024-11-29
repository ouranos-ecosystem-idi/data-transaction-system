package common

import (
	"fmt"
	"strconv"
	"strings"

	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ExtractBearerToken
// Summary: This is function which extracts the bearer token from the request header with key "Authorization".
// input: c(echo.Context) echo context
// output: (string) Bearer token
func ExtractBearerToken(c echo.Context) string {
	authHeader := c.Request().Header.Get("Authorization")
	var token string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}
	return token
}

// ExtractAcceptLanguage
// Summary: This is function which extracts the accept language from the request header with key "Accept-Language".
// input: c(echo.Context) echo context
// output: (string) Accept language
func ExtractAcceptLanguage(c echo.Context) string {
	return c.Request().Header.Get("Accept-Language")
}

// QueryParamPtr
// Summary: This is function which extracts the query parameter value from the request context and returns the pointer to the value.
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// output: (*string) pointer to the query parameter value
func QueryParamPtr(c echo.Context, paramName string) *string {
	val := c.QueryParam(paramName)
	if val == "" {
		return nil
	}
	return &val
}

// QueryParamIntPtr
// Summary: This is function which extracts the query parameter value from the request context and returns the pointer to the integer value.
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// input: defaultValue(...int) default value
// output: (*int) pointer to the integer value
// output: (error) error object
func QueryParamIntPtr(c echo.Context, paramName string, defaultValue ...int) (*int, error) {
	val := c.QueryParam(paramName)
	if val == "" {
		if len(defaultValue) > 0 {
			return &defaultValue[0], nil
		}
		return nil, nil
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return nil, err
	}
	return &valInt, nil
}

// QueryParamBoolPtr
// Summary: This is function which extracts the query parameter value from the request context and returns the pointer to the boolean value.
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// output: (*bool) pointer to the boolean value
// output: (error) error object
func QueryParamBoolPtr(c echo.Context, paramName string) (*bool, error) {
	val := c.QueryParam(paramName)
	if val == "" {
		return nil, nil
	}
	valBool, err := strconv.ParseBool(val)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return nil, err
	}
	return &valBool, nil
}

// QueryParamUUID
// Summary: This is function which extracts the query parameter value from the request context and returns the UUID value.
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// output: (uuid.UUID) UUID value
// output: (error) error object
func QueryParamUUID(c echo.Context, paramName string) (uuid.UUID, error) {
	paramValue := c.QueryParam(paramName)
	if paramValue == "" {
		logger.Set(c).Warnf(UnexpectedQueryParameter(paramName))

		return uuid.Nil, fmt.Errorf(UnexpectedQueryParameter(paramName))
	}

	if len(paramValue) != 36 {
		logger.Set(c).Warnf(UnexpectedQueryParameter(paramName))

		return uuid.Nil, fmt.Errorf(UnexpectedQueryParameter(paramName))
	}

	parsedUUID, err := uuid.Parse(paramValue)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return uuid.Nil, fmt.Errorf(UnexpectedQueryParameter(paramName))
	}

	return parsedUUID, nil
}

// QueryParamUUIDPtr
// Summary: This is function which extracts the query parameter value from the request context and returns the pointer to the UUID value.
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// output: (*uuid.UUID) pointer to the UUID value
// output: (error) error object
func QueryParamUUIDPtr(c echo.Context, paramName string) (*uuid.UUID, error) {
	val := c.QueryParam(paramName)
	if val == "" {
		return nil, nil
	}

	if len(val) != 36 {
		logger.Set(c).Warnf(UnexpectedQueryParameter(paramName))

		return nil, fmt.Errorf(UnexpectedQueryParameter(paramName))
	}

	parsedUUID, err := uuid.Parse(val)
	if err != nil {
		logger.Set(c).Warnf(err.Error())

		return nil, err
	}

	return &parsedUUID, nil
}

// QueryParamUUIDs
// Summary: This is function which extracts the query parameter value from the request context and returns the UUID slice.
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// output: ([]uuid.UUID) UUID slice
// output: (error) error object
func QueryParamUUIDs(c echo.Context, paramName string) ([]uuid.UUID, error) {
	val := c.QueryParam(paramName)
	if val == "" {
		return nil, nil
	}

	var parsedUUIDs []uuid.UUID
	for _, v := range strings.Split(val, ",") {
		if len(v) != 36 {
			logger.Set(c).Warnf(UnexpectedQueryParameter(paramName))

			return nil, fmt.Errorf(UnexpectedQueryParameter(paramName))
		}

		parsedUUID, err := uuid.Parse(v)
		if err != nil {
			logger.Set(c).Warnf(err.Error())

			return nil, err
		}
		parsedUUIDs = append(parsedUUIDs, parsedUUID)
	}

	return parsedUUIDs, nil

}

// QueryParamExists
// Summary: This is function which checks whether the query parameter exists
// input: c(echo.Context) echo context
// input: paramName(string) query parameter name
// output: (bool) true if the query parameter exists, false otherwise
func QueryParamExists(c echo.Context, paramName string) bool {
	queryParams := c.QueryParams()
	_, ok := queryParams[paramName]
	return ok
}
