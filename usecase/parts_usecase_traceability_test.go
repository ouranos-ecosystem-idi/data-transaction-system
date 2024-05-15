package usecase_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"net/url"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetParts(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	dsExpectedResAll := `[
		{
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": "A000001",
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false
		}
	]`

	dsExpectedResRequireOnly := `[
		{
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountRequiredUnit": null,
			"terminatedFlag": false
		}
	]`

	expectedResNoData := `[]`

	tests := []struct {
		name        string
		input       traceability.GetPartsModel
		receive     string
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetPartsModel(),
			receive:     f.GetParts_AllItem(nil),
			expectData:  dsExpectedResAll,
			expectAfter: common.StringPtr("2680ed32-19a3-435b-a094-23ff43aaa612"),
		},
		{
			name:        "1-2. 200: 必須項目のみ",
			input:       f.NewGetPartsModel(),
			receive:     f.GetParts_RequireItemOnly(),
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 検索結果なし",
			input:       f.NewGetPartsModel(),
			receive:     f.GetParts_NoData(),
			expectData:  expectedResNoData,
			expectAfter: nil,
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
				c.Set("operatorID", f.OperatorId)

				getPartsResponse := traceabilityentity.GetPartsResponse{}

				if err := json.Unmarshal([]byte(test.receive), &getPartsResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				var expected []traceability.PartsModel
				err := json.Unmarshal([]byte(test.expectData), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("GetParts", mock.Anything, mock.Anything, mock.Anything).Return(getPartsResponse, nil)

				partsUsecase := usecase.NewPartsTraceabilityUsecase(traceabilityRepositoryMock)

				actualRes, actualAfter, err := partsUsecase.GetPartsList(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, expected, actualRes, f.AssertMessage)
					assert.Equal(t, test.expectAfter, actualAfter, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: ページングエラー
// [x] 2-2. 500: 想定外の単位
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetParts_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	expectedInvalidTypeError := fmt.Errorf("unexpected AmountRequiredUnit. get value: kg")

	tests := []struct {
		name         string
		input        traceability.GetPartsModel
		receive      *string
		receiveError *string
		expect       error
	}{
		{
			name:         "2-1. 400: ページングエラー",
			input:        f.NewGetPartsModel(),
			receive:      nil,
			receiveError: common.StringPtr(f.Error_PagingError()),
			expect:       &expectedPagingError,
		},
		{
			name:         "2-2. 400: 想定外の単位",
			input:        f.NewGetPartsModel(),
			receive:      common.StringPtr(f.GetParts_InvalidTypeError()),
			receiveError: nil,
			expect:       expectedInvalidTypeError,
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
				c.Set("operatorID", f.OperatorId)

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				if test.receive != nil {
					getPartsResponse := traceabilityentity.GetPartsResponse{}
					if err := json.Unmarshal([]byte(*test.receive), &getPartsResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetParts", mock.Anything, mock.Anything, mock.Anything).Return(getPartsResponse, nil)
				} else if test.receiveError != nil {
					getPartsResponse := common.ToTracebilityAPIError(*test.receiveError).ToCustomError(400)
					traceabilityRepositoryMock.On("GetParts", mock.Anything, mock.Anything, mock.Anything).Return(traceabilityentity.GetPartsResponse{}, getPartsResponse)
				}

				partsUsecase := usecase.NewPartsTraceabilityUsecase(traceabilityRepositoryMock)

				actualRes, actualAfter, err := partsUsecase.GetPartsList(c, test.input)
				if assert.Error(t, err) {
					assert.Nil(t, actualRes)
					assert.Nil(t, actualAfter)
					if test.name == "2-1. 400: ページングエラー" {
						assert.Equal(t, test.expect.(*common.CustomError).Code, err.(*common.CustomError).Code)
						assert.Equal(t, test.expect.(*common.CustomError).Message, err.(*common.CustomError).Message)
						assert.Equal(t, test.expect.(*common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
						assert.Equal(t, test.expect.(*common.CustomError).Source, err.(*common.CustomError).Source)
					} else if test.name == "2-2. 400: 想定外の単位" {
						assert.Equal(t, test.expect.Error(), err.Error())
					}
				}
			},
		)
	}
}
