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

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestProjectUsecaseTraceability_GetTradeRequest
// Summary: This is normal test class which confirm the operation of API #10 GetTradeRequest.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 全項目応答(トレサビレスポンスにnullを含む)
// [x] 1-3. 200: 全項目応答(トレサビレスポンスに未定義項目を含む)
// [x] 1-4. 200: 回答済(必須項目のみ)
// [x] 1-5. 200: 回答済(任意項目が未定義)
// [x] 1-6. 200: 未回答(必須項目のみ)
// [x] 1-7. 200: 未回答(任意項目が未定義)
// [x] 1-8. 200: 検索結果なし
func TestProjectUsecaseTraceability_GetTradeRequest(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	dsExpectedResAll := `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
		}
	]`

	dsExpectedResAnsweredRequireOnly := `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
		}
	]`

	dsExpectedResAnsweringRequireOnly := `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"upstreamTraceId": null,
			"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
		}
	]`

	dsExpectedResNoData := `[]`

	tests := []struct {
		name        string
		input       traceability.GetTradeRequestInput
		receive     string
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_AllItem(),
			expectData:  dsExpectedResAll,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:        "1-2. 200: 全項目応答(トレサビレスポンスにnullを含む)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_AllItem_WithNull(),
			expectData:  dsExpectedResAll,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:        "1-3. 200: 全項目応答(トレサビレスポンスに未定義項目を含む)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_AllItem_WithUndefined(),
			expectData:  dsExpectedResAll,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:        "1-4. 200: 回答済(必須項目のみ)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_RequireItemOnlyAnswered(),
			expectData:  dsExpectedResAnsweredRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-5. 200: 回答済(任意項目が未定義)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_RequireItemOnlyAnsweredWithUndefined(),
			expectData:  dsExpectedResAnsweredRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-6. 200: 未回答(必須項目のみ)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_RequireItemOnlyAnswering(),
			expectData:  dsExpectedResAnsweringRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-7. 200: 未回答(任意項目が未定義)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_RequireItemOnlyAnsweringWithUndefined(),
			expectData:  dsExpectedResAnsweringRequireOnly,
			expectAfter: nil,
		},

		{
			name:        "1-8. 200: 検索結果なし",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_NoData(),
			expectData:  dsExpectedResNoData,
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

				getTradeRequestsResponse := traceabilityentity.GetTradeRequestsResponse{}

				if err := json.Unmarshal([]byte(test.receive), &getTradeRequestsResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				var expected []traceability.TradeModel
				err := json.Unmarshal([]byte(test.expectData), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(getTradeRequestsResponse, nil)

				partsUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, actualAfter, err := partsUsecase.GetTradeRequest(c, test.input)
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

// TestProjectUsecaseTraceability_GetTradeRequest_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #10 GetTradeRequest.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 2-1. 400: ページングエラー
func TestProjectUsecaseTraceability_GetTradeRequest_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name    string
		input   traceability.GetTradeRequestInput
		receive string
		expect  error
	}{
		{
			name:    "2-1. 400: ページングエラー",
			input:   f.NewGetTradeRequestInput(),
			receive: f.Error_PagingError(),
			expect:  &expectedPagingError,
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
				getTradeRequestsResponse := common.ToTracebilityAPIError(test.receive).ToCustomError(400)
				traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsResponse{}, getTradeRequestsResponse)

				partsUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, actualAfter, err := partsUsecase.GetTradeRequest(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Nil(t, actualRes)
					assert.Nil(t, actualAfter)
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}

// TestProjectUsecaseTraceability_GetTradeResponse
// Summary: This is normal test class which confirm the operation of API #10 GetTradeRequest.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 全項目応答(トレサビレスポンスにnullを含む)
// [x] 1-3. 200: 全項目応答(トレサビレスポンスに未定義項目を含む)
// [x] 1-4. 200: 必須項目のみ
// [x] 1-5. 200: 任意項目が未定義
// [x] 1-6. 200: 検索結果なし
// [x] 1-7. 200: 項目長最大値
func TestProjectUsecaseTraceability_GetTradeResponse(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	dsExpectedResAll := `[
		{
			"statusModel": {
				"statusId": "5185a435-c039-4196-bb34-0ee0c2395478",
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"requestStatus": {
					"cfpResponseStatus": "COMPLETED",
					"tradeTreeStatus": "UNTERMINATED",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z",
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"message": "A01のCFP値を回答ください",
				"replyMessage": "A01のCFP値を回答しました",
				"requestType": "CFP",
				"responseDueDate": "2024-12-31"
			},
			"tradeModel": {
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
			},
			"partsModel": {
				"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"partsName": "B01",
				"supportPartsName": "B0100",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		}
	]`

	dsExpectedResAllWithNull := `[
		{
			"statusModel": {
				"statusId": "5185a435-c039-4196-bb34-0ee0c2395478",
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"requestStatus": {
					"cfpResponseStatus": "COMPLETED",
					"tradeTreeStatus": "UNTERMINATED",
					"completedCount": null,
					"completedCountModifiedAt": null,
					"tradesCount": null,
					"tradesCountModifiedAt": null
				},
				"message": "A01のCFP値を回答ください",
				"replyMessage": "A01のCFP値を回答しました",
				"requestType": "CFP",
				"responseDueDate": null
			},
			"tradeModel": {
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
			},
			"partsModel": {
				"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"partsName": "B01",
				"supportPartsName": "B0100",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"partsLabelName": null,
				"partsAddInfo1": null,
				"partsAddInfo2": null,
				"partsAddInfo3": null
			}
		}
	]`

	dsExpectedResRequireOnly := `[
		{
			"statusModel": {
				"statusId": "5185a435-c039-4196-bb34-0ee0c2395478",
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"requestStatus": {
					"cfpResponseStatus": "COMPLETED",
					"tradeTreeStatus": "UNTERMINATED",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z",
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"message": "",
				"replyMessage": null,
				"requestType": "CFP",
				"responseDueDate": "2024-12-31"
			},
			"tradeModel": {
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"upstreamTraceId": null,
				"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
			},
			"partsModel": {
				"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"partsName": "B01",
				"supportPartsName": "",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"partsLabelName": null,
				"partsAddInfo1": null,
				"partsAddInfo2": null,
				"partsAddInfo3": null
			}
		}
	]`

	dsExpectedResAllWithMaxLength := `[
		{
			"statusModel": {
				"statusId": "5185a435-c039-4196-bb34-0ee0c2395478",
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"requestStatus": {
					"cfpResponseStatus": "COMPLETED",
					"tradeTreeStatus": "UNTERMINATED",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z",
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"message": "１０００文字ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
				"replyMessage": "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
				"requestType": "CFP",
				"responseDueDate": "2024-12-31"
			},
			"tradeModel": {
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698"
			},
			"partsModel": {
				"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"partsName": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
				"supportPartsName": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"partsLabelName": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
				"partsAddInfo1": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
				"partsAddInfo2": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
				"partsAddInfo3": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ"
			}
		}
	]`

	dsExpectedResNoData := `[]`

	tests := []struct {
		name        string
		input       traceability.GetTradeResponseInput
		receive     string
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_AllItem(),
			expectData:  dsExpectedResAll,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:        "1-2. 200: 全項目応答(トレサビレスポンスにnullを含む)",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_AllItem_WithNull(),
			expectData:  dsExpectedResAllWithNull,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:        "1-3. 200: 全項目応答(トレサビレスポンスに未定義項目を含む)",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_AllItem_WithUndefined(),
			expectData:  dsExpectedResAllWithNull,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:        "1-4. 200: 必須項目のみ",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_RequireItemOnly(),
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-5. 200: 任意項目が未定義",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_RequireItemOnlyWithUndefined(),
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-6. 200: 検索結果なし",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_NoData(),
			expectData:  dsExpectedResNoData,
			expectAfter: nil,
		},
		{
			name:        "1-7. 200: 項目長最大値",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_AllItem_MaxLength(),
			expectData:  dsExpectedResAllWithMaxLength,
			expectAfter: common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
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
				c.Set("operatorID", f.OperatorID)

				getTradeRequestsReceivedResponse := traceabilityentity.GetTradeRequestsReceivedResponse{}

				if err := json.Unmarshal([]byte(test.receive), &getTradeRequestsReceivedResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				var expected []traceability.TradeResponseModel
				err := json.Unmarshal([]byte(test.expectData), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(getTradeRequestsReceivedResponse, nil)

				usecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, actualAfter, err := usecase.GetTradeResponse(c, test.input)
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

// TestProjectUsecaseTraceability_GetTradeResponse_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #12 GetTradeResponse.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 2-1. 400: ページングエラー
func TestProjectUsecaseTraceability_GetTradeResponse_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name    string
		input   traceability.GetTradeResponseInput
		receive string
		expect  error
	}{
		{
			name:    "2-1. 400: ページングエラー",
			input:   f.NewGetTradeResponseInput(),
			receive: f.Error_PagingError(),
			expect:  &expectedPagingError,
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
				c.Set("operatorID", f.OperatorID)

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				getTradeRequestsReceivedResponse := common.ToTracebilityAPIError(test.receive).ToCustomError(400)
				traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsReceivedResponse{}, getTradeRequestsReceivedResponse)

				partsUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, actualAfter, err := partsUsecase.GetTradeResponse(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Nil(t, actualRes)
					assert.Nil(t, actualAfter)
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}

// TestProjectUsecaseTraceability_PutTradeRequest
// Summary: This is normal test class which confirm the operation of API #12 GetTradeResponse.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答
func TestProjectUsecaseTraceability_PutTradeRequest(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	res := f.NewPutTradeRequestModel()
	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	res.StatusModel.StatusID = uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478")
	res.StatusModel.TradeID = tradeID
	res.TradeModel.TradeID = &tradeID

	tests := []struct {
		name    string
		input   traceability.PutTradeRequestInput
		receive string
		expect  traceability.TradeRequestModel
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewPutTradeRequestInput(),
			receive: f.PutTradeRequests(),
			expect:  res,
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
				c.Set("operatorID", f.OperatorID)

				putTradeRequestsResponse := traceabilityentity.PostTradeRequestsResponses{}

				if err := json.Unmarshal([]byte(test.receive), &putTradeRequestsResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("PostTradeRequests", mock.Anything, mock.Anything).Return(putTradeRequestsResponse, common.ResponseHeaders{}, nil)

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, _, err := tradeUsecase.PutTradeRequest(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.StatusModel, actualRes.StatusModel, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel, actualRes.TradeModel, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecaseTraceability_PutTradeRequest_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #7 PutTradeRequest.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 2-1. 400: トレース識別子存在エラー
func TestProjectUsecaseTraceability_PutTradeRequest_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	expectedExistError := common.CustomError{
		Code:          400,
		Message:       "リクエストパラメータのトレース識別子に、存在しない部品が含まれています。",
		MessageDetail: common.StringPtr("MSGAECI0005"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name    string
		input   traceability.PutTradeRequestInput
		receive string
		expect  error
	}{
		{
			name:    "2-1. 400: トレース識別子存在エラー",
			input:   f.NewPutTradeRequestInput(),
			receive: f.Error_TraceIdNotFound(),
			expect:  &expectedExistError,
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
				c.Set("operatorID", f.OperatorID)

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				getTradeResponse := common.ToTracebilityAPIError(test.receive).ToCustomError(400)
				traceabilityRepositoryMock.On("PostTradeRequests", mock.Anything, mock.Anything).Return(traceabilityentity.PostTradeRequestsResponses{}, common.ResponseHeaders{}, getTradeResponse)

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				_, _, err := tradeUsecase.PutTradeRequest(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}

// TestProjectUsecaseTraceability_PutTradeResponse
// Summary: This is normal test class which confirm the operation of API #13 PutTradeResponse.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 全項目応答(トレサビレスポンスにnullを含む)
// [x] 1-3. 200: 全項目応答(トレサビレスポンスに未定義項目を含む)
func TestProjectUsecaseTraceability_PutTradeResponse(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	res := f.NewPutTradeResponseModel()

	tests := []struct {
		name         string
		input        traceability.PutTradeResponseInput
		inputFunc    func() traceability.PutTradeResponseInput
		receiveTrRes string
		receiveTrRec string
		expect       traceability.TradeModel
	}{
		{
			name:  "1-1. 200: 全項目応答",
			input: f.NewPutTradeResponseInput(),
			inputFunc: func() traceability.PutTradeResponseInput {
				return f.NewPutTradeResponseInput()
			},
			receiveTrRes: f.PostTrades(),
			receiveTrRec: f.GetTradeRequestsReceived_AllItem(),
			expect:       res,
		},
		{
			name:  "1-2. 200: 全項目応答(トレサビレスポンスにnullを含む)",
			input: f.NewPutTradeResponseInput(),
			inputFunc: func() traceability.PutTradeResponseInput {
				return f.NewPutTradeResponseInput()
			},
			receiveTrRes: f.PostTrades(),
			receiveTrRec: f.GetTradeRequestsReceived_AllItem_WithNull(),
			expect:       res,
		},
		{
			name:  "1-3. 200: 全項目応答(トレサビレスポンスに未定義項目を含む)",
			input: f.NewPutTradeResponseInput(),
			inputFunc: func() traceability.PutTradeResponseInput {
				return f.NewPutTradeResponseInput()
			},
			receiveTrRes: f.PostTrades(),
			receiveTrRec: f.GetTradeRequestsReceived_AllItem_WithUndefined(),
			expect:       res,
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
				c.Set("operatorID", f.OperatorID)

				postTradesResponse := traceabilityentity.PostTradesResponse{}
				if err := json.Unmarshal([]byte(test.receiveTrRes), &postTradesResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				getTradesResponse := traceabilityentity.GetTradeRequestsReceivedResponse{}
				if err := json.Unmarshal([]byte(test.receiveTrRec), &getTradesResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("PostTrades", mock.Anything, mock.Anything).Return(postTradesResponse, common.ResponseHeaders{}, nil)
				traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(getTradesResponse, nil)

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, _, err := tradeUsecase.PutTradeResponse(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect, actualRes, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecaseTraceability_PutTradeResponse_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #13 PutTradeResponse.
// Target: trade_usecase_traceability_impl.go
// TestPattern:
// [x] 2-1. 400: ページングエラー(PUT)
// [x] 2-2. 400: ページングエラー(取得)
// [x] 2-3. 400: 対象データなし
func TestProjectUsecaseTraceability_PutTradeResponse_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	expectedPagingNotFound := fmt.Errorf("Trade not found")

	tests := []struct {
		name              string
		input             traceability.PutTradeResponseInput
		receiveTrRes      *string
		receiveTrResError *string
		receiveTrRec      *string
		receiveTrRecError *string
		expect            error
	}{
		{
			name:              "2-1. 400: ページングエラー(PUT)",
			input:             f.NewPutTradeResponseInput(),
			receiveTrRes:      nil,
			receiveTrResError: common.StringPtr(f.Error_PagingError()),
			receiveTrRec:      nil,
			receiveTrRecError: nil,
			expect:            &expectedPagingError,
		},
		{
			name:              "2-2. 400: ページングエラー(取得)",
			input:             f.NewPutTradeResponseInput(),
			receiveTrRes:      common.StringPtr(f.PostTrades()),
			receiveTrResError: nil,
			receiveTrRec:      nil,
			receiveTrRecError: common.StringPtr(f.Error_PagingError()),
			expect:            &expectedPagingError,
		},
		{
			name:              "2-3. 400: 対象データなし",
			input:             f.NewPutTradeResponseInput(),
			receiveTrRes:      common.StringPtr(f.PostTrades()),
			receiveTrResError: nil,
			receiveTrRec:      common.StringPtr(f.GetTradeRequestsReceived_NoData()),
			receiveTrRecError: nil,
			expect:            expectedPagingNotFound,
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
				c.Set("operatorID", f.OperatorID)

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				if test.receiveTrRes != nil {
					postTradesResponse := traceabilityentity.PostTradesResponse{}
					if err := json.Unmarshal([]byte(*test.receiveTrRes), &postTradesResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("PostTrades", mock.Anything, mock.Anything).Return(postTradesResponse, common.ResponseHeaders{}, nil)
					if test.receiveTrRec != nil {
						getTradesResponse := traceabilityentity.GetTradeRequestsReceivedResponse{}
						if err := json.Unmarshal([]byte(*test.receiveTrRec), &getTradesResponse); err != nil {
							log.Fatalf(f.UnmarshalMockFailureMessage, err)
						}
						traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(getTradesResponse, nil)
					} else {
						getTradesResponse := common.ToTracebilityAPIError(*test.receiveTrRecError).ToCustomError(400)
						if err := json.Unmarshal([]byte(*test.receiveTrRecError), &getTradesResponse); err != nil {
							log.Fatalf(f.UnmarshalMockFailureMessage, err)
						}
						traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsReceivedResponse{}, getTradesResponse)
					}
				} else {
					postTradesResponse := common.ToTracebilityAPIError(*test.receiveTrResError).ToCustomError(400)
					if err := json.Unmarshal([]byte(*test.receiveTrResError), &postTradesResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("PostTrades", mock.Anything, mock.Anything).Return(traceabilityentity.PostTradesResponse{}, common.ResponseHeaders{}, postTradesResponse)
				}

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				_, _, err := tradeUsecase.PutTradeResponse(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}
