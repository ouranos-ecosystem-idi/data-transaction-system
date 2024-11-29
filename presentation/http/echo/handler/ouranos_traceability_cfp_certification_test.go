package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/presentation/http/echo/handler"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfpCertification 正常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常系
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetCfpCertification_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport/cfpCertification"
	var dataTarget = "cfpCertification"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		expectStatus      int
	}{
		{
			name: "1-1. 200: 正常系",
			modifyQueryParams: func(q url.Values) {
				// q.Set("traceId", "d17833fe-22b7-4a4a-b097-bc3f2150c9a6")
				q.Set("traceId", "38bdd8a5-76a7-a53d-de12-725707b04a1b")
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			operatorUUID, _ := uuid.Parse(f.OperatorId)
			traceId, _ := uuid.Parse(f.TraceId)
			input := traceability.GetCfpCertificationInput{
				OperatorID: operatorUUID,
				TraceID:    traceId,
			}

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)
			test.modifyQueryParams(q)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			cfpCertificationUsecase := new(mocks.ICfpCertificationUsecase)
			cfpCertificationHandler := handler.NewCfpCertificationHandler(cfpCertificationUsecase)
			cfpCertificationUsecase.On("GetCfpCertification", c, input).Return(traceability.CfpCertificationModels{}, nil)

			// エラーが発生しないことを確認
			if assert.NoError(t, cfpCertificationHandler.GetCfpCertification(c)) {
				// ステータスコードが期待通りであることを確認
				assert.Equal(t, test.expectStatus, rec.Code)
				// モックの呼び出しが期待通りであることを確認
				cfpCertificationUsecase.AssertExpectations(t)
			}

			// レスポンスヘッダにX-Trackが含まれているかチェック
			_, ok := rec.Header()["X-Track"]
			assert.True(t, ok, "Header should have 'X-Track' key")
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfpCertification 異常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：traceIdが含まれない場合
// [x] 1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合
// [x] 1-3. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-4. 500: システムエラー：取得処理エラー
// [x] 1-5. 500: システムエラー：取得処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetCfpCertification_Abnormal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport/cfpCertification"
	var dataTarget = "cfpCertification"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		modifyContexts    func(c echo.Context)
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdが含まれない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "invalid")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "d17833fe-22b7-4a4a-b097-bc3f2150c9a6")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalid")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid or expired token",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "d17833fe-22b7-4a4a-b097-bc3f2150c9a6")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-5. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "d17833fe-22b7-4a4a-b097-bc3f2150c9a6")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-6. 503: システムエラー：トレサビエラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "d17833fe-22b7-4a4a-b097-bc3f2150c9a6")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:      common.ToTracebilityAPIError(f.Error_MaintenanceError()).ToCustomError(503),
			expectError:  "code=503, message={[traceability] ServiceUnavailable The service is currently undergoing maintenance. We apologize for any inconvenience. MSGXXXXYYYY",
			expectStatus: http.StatusServiceUnavailable,
		},
		{
			name: "1-7. 503: システムエラー：トレサビエラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", "d17833fe-22b7-4a4a-b097-bc3f2150c9a6")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:      common.ToTracebilityAPIError(f.Error_GatewayError()).ToCustomError(503),
			expectError:  "code=503, message={[traceability] ServiceUnavailable  Service Unavailable",
			expectStatus: http.StatusServiceUnavailable,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)
			test.modifyQueryParams(q)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			test.modifyContexts(c)

			cfpCertificationUsecase := new(mocks.ICfpCertificationUsecase)
			cfpCertificationUsecase.On("GetCfpCertification", mock.Anything, mock.Anything).Return(traceability.CfpCertificationModels{}, test.receive)
			cfpCertificationHandler := handler.NewCfpCertificationHandler(cfpCertificationUsecase)

			err := cfpCertificationHandler.GetCfpCertification(c)
			e.HTTPErrorHandler(err, c)
			// エラーが返されることを確認
			if assert.Error(t, err) {
				// ステータスコードが期待通りであることを確認
				assert.Equal(t, test.expectStatus, rec.Code)
				// エラーメッセージが期待通りであることを確認
				assert.ErrorContains(t, err, test.expectError)
			}
		})
	}
}
