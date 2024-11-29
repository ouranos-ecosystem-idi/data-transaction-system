package usecase_test

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/authentication"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase"
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
func TestProjectAuth_VerifyAPIKey(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyApiKey"

	tests := []struct {
		name    string
		input   input.VerifyAPIKey
		receive authentication.VeriryAPIKeyResponse
		expect  output.VerifyAPIKey
	}{
		{
			name: "1-1. 200: 正常終了",
			input: input.VerifyAPIKey{
				APIKey:    "apikey",
				IPAddress: "127.0.0.1",
			},
			receive: authentication.VeriryAPIKeyResponse{
				IsAPIKeyValid:    true,
				IsIPAddressValid: true,
			},
			expect: output.VerifyAPIKey{
				IsAPIKeyValid:    true,
				IsIPAddressValid: true,
			},
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

				authAPIRepositoryMock := new(mocks.AuthAPIRepository)
				authAPIRepositoryMock.On("VerifyAPIKey", mock.Anything).Return(test.receive, nil)
				usecase := usecase.NewVerifyUsecase(authAPIRepositoryMock)

				actual, err := usecase.VerifyAPIKey(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect.IsAPIKeyValid, actual.IsAPIKeyValid, f.AssertMessage)
					assert.Equal(t, test.expect.IsIPAddressValid, actual.IsIPAddressValid, f.AssertMessage)
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
func TestProjectAuth_VerifyToken(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyToken"

	tests := []struct {
		name    string
		input   input.VerifyToken
		receive authentication.VeriryTokenResponse
		expect  output.VerifyToken
	}{
		{
			name: "1-1. 200: 正常終了",
			input: input.VerifyToken{
				Token: "token",
			},
			receive: authentication.VeriryTokenResponse{
				OperatorID: common.StringPtr("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			},
			expect: output.VerifyToken{
				OperatorID: common.StringPtr("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			},
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

				authAPIRepositoryMock := new(mocks.AuthAPIRepository)
				authAPIRepositoryMock.On("VerifyToken", mock.Anything).Return(test.receive, nil)
				usecase := usecase.NewVerifyUsecase(authAPIRepositoryMock)

				actual, err := usecase.VerifyToken(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect.OperatorID, actual.OperatorID, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/authinfo VerifyAPIKey テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 500: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectAuth_VerifyAPIKey_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "VerifyApiKey"

	tests := []struct {
		name    string
		input   input.VerifyAPIKey
		receive error
		expect  error
	}{
		{
			name: "2-1. 500: データ取得エラー",
			input: input.VerifyAPIKey{
				APIKey:    "apikey",
				IPAddress: "127.0.0.1",
			},
			receive: fmt.Errorf("AccessError"),
			expect:  fmt.Errorf("AccessError"),
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

				authAPIRepositoryMock := new(mocks.AuthAPIRepository)
				authAPIRepositoryMock.On("VerifyAPIKey", mock.Anything).Return(authentication.VeriryAPIKeyResponse{}, test.receive)
				usecase := usecase.NewVerifyUsecase(authAPIRepositoryMock)

				_, err := usecase.VerifyAPIKey(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect, err, f.AssertMessage)
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
func TestProjectAuth_VerifyToken_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/auth/v1/authInfo"
	var dataTarget = "verifyToken"

	tests := []struct {
		name    string
		input   input.VerifyToken
		receive error
		expect  error
	}{
		{
			name: "2-1. 500: データ取得エラー",
			input: input.VerifyToken{
				Token: "token",
			},
			receive: fmt.Errorf("AccessError"),
			expect:  fmt.Errorf("AccessError"),
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

				authAPIRepositoryMock := new(mocks.AuthAPIRepository)
				authAPIRepositoryMock.On("VerifyToken", mock.Anything).Return(authentication.VeriryTokenResponse{}, test.receive)
				usecase := usecase.NewVerifyUsecase(authAPIRepositoryMock)

				_, err := usecase.VerifyToken(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect, err, f.AssertMessage)
				}
			},
		)
	}
}
