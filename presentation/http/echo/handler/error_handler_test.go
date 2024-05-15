package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"syscall"
	"testing"

	"data-spaces-backend/presentation/http/echo/handler"
	f "data-spaces-backend/test/fixtures"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////////////////////////////////////////////////
// エラーハンドラーテストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 401: 認証エラー：JWT不正の場合
// [x] 1-2. 404: 存在チェックエラー：Not Foundの場合
// [x] 1-3. 400: バリデーションエラー：上記以外のHTTP例外の場合
// [x] 1-4. 500: システムエラー：上記以外のHTTP例外の場合
// [x] 1-5. 503: システムエラー：Postgresシャットダウン例外の場合
// [x] 1-6. 503: システムエラー：Postgresシャットダウン例外の場合
// [x] 1-7. 500: システムエラー：上記以外のPostgres例外の場合
// [x] 1-8. 503: システムエラー：ConnectionRefuseの場合
// [x] 1-9. 500: システムエラー：上記以外の例外の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_CustomHTTPErrorHandler(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1"

	tests := []struct {
		name        string
		input       error
		errorStatus int
	}{
		{
			name:        "1-1. 401: バリデーションエラー：JWT不正の場合",
			input:       middleware.ErrJWTMissing,
			errorStatus: http.StatusUnauthorized,
		},
		{
			name: "1-2. 404: バリデーションエラー：Not Foundの場合",
			input: &echo.HTTPError{
				Code:     404,
				Message:  "Not Found",
				Internal: fmt.Errorf("Not Found"),
			},
			errorStatus: http.StatusNotFound,
		},
		{
			name: "1-3. 400: バリデーションエラー：上記以外のHTTP例外の場合",
			input: &echo.HTTPError{
				Code:     400,
				Message:  "Request Invalid",
				Internal: fmt.Errorf("Request Invalid"),
			},
			errorStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 500: バリデーションエラー：上記以外のHTTP例外の場合",
			input: &echo.HTTPError{
				Code:     500,
				Message:  "Request Invalid",
				Internal: fmt.Errorf("Internal Server Error"),
			},
			errorStatus: http.StatusInternalServerError,
		},
		{
			name: "1-5. 500: システムエラー：Postgresシャットダウン例外の場合",
			input: &pgconn.PgError{
				Severity: "ERROR",
				Code:     "57P01",
				Message:  "Admin Shutdown Error",
			},
			errorStatus: http.StatusServiceUnavailable,
		},
		{
			name: "1-6. 500: システムエラー：Postgresシャットダウン例外の場合",
			input: &pgconn.PgError{
				Severity: "ERROR",
				Code:     "57P02",
				Message:  "Clash Shutdown Error",
			},
			errorStatus: http.StatusServiceUnavailable,
		},
		{
			name: "1-7. 500: システムエラー：上記以外のPostgres例外の場合",
			input: &pgconn.PgError{
				Severity: "ERROR",
				Code:     "57P03",
				Message:  "Another Error",
			},
			errorStatus: http.StatusInternalServerError,
		},
		{
			name:        "1-8. 500: システムエラー：ConnectionRefuseの場合",
			input:       syscall.ECONNREFUSED,
			errorStatus: http.StatusServiceUnavailable,
		},
		{
			name:        "1-9. 500: システムエラー：上記以外の例外の場合",
			input:       fmt.Errorf("Another Error"),
			errorStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()
				q := make(url.Values)
				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				handler.CustomHTTPErrorHandler(test.input, c)
				assert.Equal(t, rec.Code, test.errorStatus)
			},
		)
	}
}
