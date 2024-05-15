package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"data-spaces-backend/presentation/http/echo/handler"
	f "data-spaces-backend/test/fixtures"
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
		input        input.VerifyAPIKey
		receive      output.VerifyAPIKey
		expectStatus int
	}{
		{
			name: "1-1. 200: 正常終了",
			input: input.VerifyAPIKey{
				APIKey:    "apikey",
				IPAddress: "127.0.0.1",
			},
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
				req.Header.Set("apiKey", test.input.APIKey)
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
		name         string
		input        input.VerifyAPIKey
		receive      output.VerifyAPIKey
		receiveError error
		expectError  string
	}{
		{
			name: "2-1. 500: 検証エラーの場合",
			input: input.VerifyAPIKey{
				APIKey:    "apikey",
				IPAddress: "127.0.0.1",
			},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    true,
				IsIPAddressValid: true,
			},
			receiveError: fmt.Errorf("accessError"),
			expectError:  "code=500, message={[auth] InternalServerError Unexpected error occurred",
		},
		{
			name: "2-2. 403: 不正APIキーの場合",
			input: input.VerifyAPIKey{
				APIKey:    "apikey",
				IPAddress: "127.0.0.1",
			},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    false,
				IsIPAddressValid: true,
			},
			expectError: "code=403, message={[auth] AccessDenied Invalid key",
		},
		{
			name: "2-3. 403: 不正IPアドレスの場合",
			input: input.VerifyAPIKey{
				APIKey:    "apikey",
				IPAddress: "127.0.0.1",
			},
			receive: output.VerifyAPIKey{
				IsAPIKeyValid:    true,
				IsIPAddressValid: false,
			},
			expectError: "code=403, message={[auth] AccessDenied IP address not authorized for this API key",
		},
		{
			name: "2-4. 403: APIキー未設定の場合",
			input: input.VerifyAPIKey{
				APIKey:    "",
				IPAddress: "127.0.0.1",
			},
			expectError: "code=403, message={[auth] AccessDenied You do not have the necessary privileges",
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
				req.Header.Set("apiKey", test.input.APIKey)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				verifyUsecase := new(mocks.IVerifyUsecase)
				verifyUsecase.On("VerifyAPIKey", mock.Anything).Return(test.receive, test.receiveError)
				authHandler := handler.NewAuthHandler(verifyUsecase)

				err := authHandler.VerifyAPIKey(c)
				if assert.Error(t, err) {
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
				OperatorID: &f.OperatorId,
			},
			expect:       &f.OperatorId,
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
				OperatorID: &f.OperatorId,
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
