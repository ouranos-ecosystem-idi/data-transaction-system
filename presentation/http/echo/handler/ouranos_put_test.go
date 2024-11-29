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
	testhelper "data-spaces-backend/test/test_helper"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/ テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常系：partsStructureの場合
// [x] 1-2. 200: 正常系：partsの場合
// [x] 1-3. 200: 正常系：tradeRequestの場合
// [x] 1-4. 200: 正常系：tradeResponseの場合
// [x] 1-5. 200: 正常系：cfpの場合
// [x] 1-6. 200: 正常系：statusの場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_Put_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
	}{
		{
			name: "1-1. 200: 正常系：partsStructureの場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
			},
		},
		{
			name: "1-2. 200: 正常系：partsの場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "parts")
			},
		},
		{
			name: "1-3. 200: 正常系：tradeRequestの場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "tradeRequest")
			},
		},
		{
			name: "1-4. 200: 正常系：tradeResponseの場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "tradeResponse")
			},
		},
		{
			name: "1-5. 200: 正常系：cfpの場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "cfp")
			},
		},
		{
			name: "1-6. 200: 正常系：statusの場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "status")
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
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsStructureHandler := new(mocks.IPartsStructureHandler)
				partsStructureHandler.On("PutPartsStructureModel", mock.Anything).Return(nil)
				partsHandler := new(mocks.IPartsHandler)
				partsHandler.On("PutPartsModel", mock.Anything).Return(nil)
				tradeHandler := new(mocks.ITradeHandler)
				tradeHandler.On("PutTradeRequest", mock.Anything).Return(nil)
				tradeHandler.On("PutTradeResponse", mock.Anything).Return(nil)
				cfpHandler := new(mocks.ICfpHandler)
				cfpHandler.On("PutCfp", mock.Anything).Return(nil)
				cfpCertificationHandler := new(mocks.ICfpCertificationHandler)
				statusHandler := new(mocks.IStatusHandler)
				statusHandler.On("PutStatus", mock.Anything).Return(nil)
				h := handler.NewOuranosHandler(cfpHandler, cfpCertificationHandler, partsHandler, partsStructureHandler, tradeHandler, statusHandler)
				err := h.PutOuranos(c)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/ テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：dataTargetの値が未指定の場合
// [x] 1-2. 400: バリデーションエラー：dataTargetの値が不正の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_Put(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：dataTargetが含まれない場合",
			modifyQueryParams: func(q url.Values) {
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, dataTarget: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：dataTargetがoperator以外の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "hoge")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, dataTarget: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)
				host := ""

				h := testhelper.NewMockHandler(host)

				err := h.PutOuranos(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					fmt.Println(err)
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}
