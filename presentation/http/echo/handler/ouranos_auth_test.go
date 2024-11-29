package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/presentation/http/echo/handler"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase/input"
	"data-spaces-backend/usecase/output"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/authinfo VerifyAPIKey テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectAuth_VerifyAPIKey_Normal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyApiKey"

	tests := []struct {
		name         string
		APIKey       string
		XFF          []string
		receive      output.VerifyAPIKey
		expectStatus int
	}{
		{
			name:   "1-1. 200: 正常終了",
			APIKey: "apiKey",
			XFF:    []string{"client", "app server", "cloud Armor", "ALB"},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    true,
				IsIPAddressValid: true,
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("apiKey", test.APIKey)
				req.Header.Set("X-Forwarded-For", strings.Join(test.XFF, ", "))
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				verifyUsecase := new(mocks.IVerifyUsecase)
				verifyUsecase.On("VerifyAPIKey", mock.Anything).Return(test.receive, nil)
				authHandler := handler.NewAuthHandler(verifyUsecase)

				err := authHandler.VerifyAPIKey(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					verifyUsecase.AssertExpectations(t)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/authinfo VerifyAPIKey テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 500: 検証エラーの場合
// [x] 2-2. 403: 不正APIキーの場合
// [x] 2-3. 403: 不正IPアドレスの場合
// [x] 2-4. 403: APIキー未設定の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectAuth_VerifyAPIKey(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyApiKey"

	tests := []struct {
		name string
		// input        input.VerifyAPIKey
		APIKey       string
		XFF          []string
		receive      output.VerifyAPIKey
		receiveError error
		expectError  string
		expectStatus int
	}{
		{
			name:   "2-1. 500: 検証エラーの場合",
			APIKey: "apikey",
			XFF:    []string{"client", "app server", "cloud Armor", "ALB"},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    true,
				IsIPAddressValid: true,
			},
			receiveError: fmt.Errorf("accessError"),
			expectError:  "code=500, message={[auth] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name:   "2-2. 403: 不正APIキーの場合",
			APIKey: "apikey",
			XFF:    []string{"client", "app server", "cloud Armor", "ALB"},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    false,
				IsIPAddressValid: true,
			},
			expectError:  "code=403, message={[auth] AccessDenied Invalid key",
			expectStatus: http.StatusForbidden,
		},
		{
			name:   "2-3. 403: 不正IPアドレスの場合",
			APIKey: "apikey",
			XFF:    []string{"client", "app server", "cloud Armor", "ALB"},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    true,
				IsIPAddressValid: false,
			},
			expectError:  "code=403, message={[auth] AccessDenied IP address not authorized for this API key",
			expectStatus: http.StatusForbidden,
		},
		{
			name:         "2-4. 403: APIキー未設定の場合",
			APIKey:       "",
			XFF:          []string{"client", "app server", "cloud Armor", "ALB"},
			expectError:  "code=403, message={[auth] AccessDenied You do not have the necessary privileges",
			expectStatus: http.StatusForbidden,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("apiKey", test.APIKey)
				req.Header.Set("X-Forwarded-For", strings.Join(test.XFF, ", "))
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				verifyUsecase := new(mocks.IVerifyUsecase)
				verifyUsecase.On("VerifyAPIKey", mock.Anything).Return(test.receive, test.receiveError)
				authHandler := handler.NewAuthHandler(verifyUsecase)

				err := authHandler.VerifyAPIKey(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/authinfo VerifyToken テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectAuth_VerifyToken_Normal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyToken"

	tests := []struct {
		name         string
		input        input.VerifyToken
		receive      output.VerifyToken
		expect       *string
		expectStatus int
	}{
		{
			name: "1-1. 200: 正常終了",
			input: input.VerifyToken{
				Token: "token",
			},
			receive: output.VerifyToken{
				OperatorID: common.StringPtr("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			},
			expect:       common.StringPtr("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Request().Header.Set("Authorization", "Bearer token")
				verifyUsecase := new(mocks.IVerifyUsecase)
				verifyUsecase.On("VerifyToken", mock.Anything).Return(test.receive, nil)
				authHandler := handler.NewAuthHandler(verifyUsecase)

				res, err := authHandler.VerifyToken(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.Equal(t, test.receive.OperatorID, res)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/authinfo VerifyToken テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: 検証エラーの場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectAuth_VerifyToken(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyToken"

	tests := []struct {
		name         string
		input        input.VerifyToken
		receive      output.VerifyToken
		receiveError error
		expectError  string
	}{
		{
			name: "2-1. 400: 検証エラーの場合",
			input: input.VerifyToken{
				Token: "token",
			},
			receive: output.VerifyToken{
				OperatorID: common.StringPtr("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			},
			receiveError: fmt.Errorf("accessError"),
			expectError:  "code=400, message={[auth] BadRequest Invalid or expired token id",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Request().Header.Set("Authorization", "Bearer token")
				verifyUsecase := new(mocks.IVerifyUsecase)
				verifyUsecase.On("VerifyToken", mock.Anything).Return(test.receive, test.receiveError)
				authHandler := handler.NewAuthHandler(verifyUsecase)

				_, err := authHandler.VerifyToken(c)
				if assert.Error(t, err) {
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}
