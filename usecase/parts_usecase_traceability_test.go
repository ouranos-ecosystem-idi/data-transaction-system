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
// [x] 1-2. 200: nil許容項目がnil
// [x] 1-3. 200: 任意項目が未定義
// [x] 1-4. 200: 検索結果なし
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
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false,
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		}
	]`

	dsExpectedResRequireOnly := `[
		{
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": null,
			"terminatedFlag": false,
			"partsLabelName": null,
			"partsAddInfo1": null,
			"partsAddInfo2": null,
			"partsAddInfo3": null
		}
	]`

	expectedResNoData := `[]`

	tests := []struct {
		name        string
		input       traceability.GetPartsInput
		receive     string
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetPartsInput(),
			receive:     f.GetParts_AllItem(nil),
			expectData:  dsExpectedResAll,
			expectAfter: common.StringPtr("2680ed32-19a3-435b-a094-23ff43aaa612"),
		},
		{
			name:        "1-2. 200: nil許容項目がnil",
			input:       f.NewGetPartsInput(),
			receive:     f.GetParts_RequireItemOnly(),
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 任意項目が未定義",
			input:       f.NewGetPartsInput(),
			receive:     f.GetParts_RequireItemOnlyWithUndefined(),
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-4. 200: 検索結果なし",
			input:       f.NewGetPartsInput(),
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
		input        traceability.GetPartsInput
		receive      *string
		receiveError *string
		expect       error
	}{
		{
			name:         "2-1. 400: ページングエラー",
			input:        f.NewGetPartsInput(),
			receive:      nil,
			receiveError: common.StringPtr(f.Error_PagingError()),
			expect:       &expectedPagingError,
		},
		{
			name:         "2-2. 400: 想定外の単位",
			input:        f.NewGetPartsInput(),
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
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Delete /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_DeleteParts(tt *testing.T) {

	var method = "DELETE"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name    string
		input   traceability.DeletePartsInput
		receive string
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewDeletePartsInput(f.TraceID),
			receive: f.DeleteParts_AllItem(f.TraceID),
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
				c.Set("traceId", f.TraceID)

				receiveRes := traceabilityentity.DeletePartsResponse{}

				if err := json.Unmarshal([]byte(test.receive), &receiveRes); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("DeleteParts", mock.Anything, mock.Anything).Return(receiveRes, common.ResponseHeaders{}, nil)

				partsUsecase := usecase.NewPartsTraceabilityUsecase(traceabilityRepositoryMock)

				_, err := partsUsecase.DeleteParts(c, test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Delete /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: 部品情報使用中エラー(部品構成)
// [x] 2-2. 400: 部品情報使用中エラー(依頼)
// [x] 2-3. 400: 部品情報使用中エラー(受領依頼)
// [x] 2-4. 400: 部品存在チェックエラー
// [x] 2-5. 400: ファイル削除エラー
// [x] 2-6. 400: 認証情報エラー
// [x] 2-7. 503: トレサビエラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_DeleteParts_Abnormal(tt *testing.T) {

	var method = "DELETE"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	expect := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			{
				ErrorCode:        "MSGAECP0014",
				ErrorDescription: "指定された部品は部品構成が存在するため削除できません。",
				RelevantData: &[]string{
					"a84012cc-73fb-4f9b-9130-59ae546f7091",
					"a84012cc-73fb-4f9b-9130-59ae546f7092",
				},
			},
		},
	}

	expect2 := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			{
				ErrorCode:        "MSGAECP0015",
				ErrorDescription: "指定された部品は依頼済みのため削除できません。",
				RelevantData: &[]string{
					"a84012cc-73fb-4f9b-9130-59ae546f7093",
				},
			},
		},
	}

	expect3 := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			{
				ErrorCode:        "MSGAECP0016",
				ErrorDescription: "指定された部品は受領済みの依頼に紐づいているため削除できません。",
				RelevantData: &[]string{
					"a84012cc-73fb-4f9b-9130-59ae546f7094",
					"a84012cc-73fb-4f9b-9130-59ae546f7095",
				},
			},
		},
	}

	expect4 := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			{
				ErrorCode:        "MSGAECP0013",
				ErrorDescription: "指定された部品は存在しません。",
			},
		},
	}

	expect5 := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			{
				ErrorCode:        "MSGAECP0034",
				ErrorDescription: "ファイル削除に失敗しました。",
			},
		},
	}

	expect6 := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			{
				ErrorCode:        "MSGAECO0025",
				ErrorDescription: "認証情報と事業者識別子が一致しません。",
			},
		},
	}

	expect7 := common.TraceabilityAPIErrorDelete{
		Message: common.StringPtr("Service Unavailable"),
	}

	tests := []struct {
		name       string
		input      traceability.DeletePartsInput
		receiveErr *string
		expect     error
	}{
		{
			name:       "2-1. 400: 部品情報使用中エラー(部品構成)",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_BlockingPartsStructureError()),
			expect:     expect.ToCustomError(400),
		},
		{
			name:       "2-2. 400: 部品情報使用中エラー(依頼)",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_BlockingTradeRequestError()),
			expect:     expect2.ToCustomError(400),
		},
		{
			name:       "2-3. 400: 部品情報使用中エラー(受領依頼)",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_BlockingTradeResponseError()),
			expect:     expect3.ToCustomError(400),
		},
		{
			name:       "2-4. 400: 部品存在チェックエラー",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_PartsIdNotFound()),
			expect:     expect4.ToCustomError(400),
		},
		{
			name:       "2-5. 400: ファイル削除エラー",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_FileDeleteError()),
			expect:     expect5.ToCustomError(400),
		},
		{
			name:       "2-6. 400: 認証情報エラー",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_AuthDiffError()),
			expect:     expect6.ToCustomError(400),
		},
		{
			name:       "2-7. 503: ゲートウェイエラー",
			input:      f.NewDeletePartsInput(f.TraceID),
			receiveErr: common.StringPtr(f.Error_GatewayError()),
			expect:     expect7.ToCustomError(503),
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
				c.Set("traceId", f.TraceID)
				receiveErr := common.ToTracebilityAPIErrorDelete(*test.receiveErr).ToCustomError(400)
				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("DeleteParts", mock.Anything, mock.Anything).Return(traceabilityentity.DeletePartsResponse{}, common.ResponseHeaders{}, receiveErr)

				partsUsecase := usecase.NewPartsTraceabilityUsecase(traceabilityRepositoryMock)

				_, err := partsUsecase.DeleteParts(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
