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
// [x] 1-2. 200: 回答済(必須項目のみ)
// [x] 1-3. 200: 未回答(必須項目のみ)
// [x] 1-4. 200: 検索結果なし
func TestProjectUsecaseTraceability_GetTradeRequest(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	dsExpectedResAll := `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"downstreamOperatorId": "e03cc699-7234-31ed-86be-cc18c92208e5"
		}
	]`

	dsExpectedResAnsweredRequireOnly := `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"downstreamOperatorId": "e03cc699-7234-31ed-86be-cc18c92208e5"
		}
	]`

	dsExpectedResAnsweringRequireOnly := `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"upstreamTraceId": null,
			"downstreamOperatorId": "e03cc699-7234-31ed-86be-cc18c92208e5"
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
			name:        "1-2. 200: 回答済(必須項目のみ)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_RequireItemOnlyAnswered(),
			expectData:  dsExpectedResAnsweredRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 未回答(必須項目のみ)",
			input:       f.NewGetTradeRequestInput(),
			receive:     f.GetTradeRequests_RequireItemOnlyAnswering(),
			expectData:  dsExpectedResAnsweringRequireOnly,
			expectAfter: nil,
		},

		{
			name:        "1-4. 200: 検索結果なし",
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
			expect:  expectedPagingError,
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
					assert.Equal(t, test.expect.(common.CustomError).Code, err.(*common.CustomError).Code)
					assert.Equal(t, test.expect.(common.CustomError).Message, err.(*common.CustomError).Message)
					assert.Equal(t, test.expect.(common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
					assert.Equal(t, test.expect.(common.CustomError).Source, err.(*common.CustomError).Source)
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
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 検索結果なし
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
					"tradeTreeStatus": "UNTERMINATED"
				},
				"message": "A01のCFP値を回答ください",
				"replyMessage": "A01のCFP値を回答しました",
				"requestType": "CFP"
			},
			"tradeModel": {
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"upstreamOperatorId": "e03cc699-7234-31ed-86be-cc18c92208e5",
				"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"downstreamOperatorId": "b1234567-1234-1234-1234-123456789012"
			},
			"partsModel": {
				"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"partsName": "B01",
				"supportPartsName": "B0100",
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false
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
					"tradeTreeStatus": "UNTERMINATED"
				},
				"message": "",
				"replyMessage": "",
				"requestType": "CFP"
			},
			"tradeModel": {
				"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
				"upstreamOperatorId": "e03cc699-7234-31ed-86be-cc18c92208e5",
				"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"upstreamTraceId": null,
				"downstreamOperatorId": "b1234567-1234-1234-1234-123456789012"
			},
			"partsModel": {
				"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				"partsName": "B01",
				"supportPartsName": "",
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false
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
			name:        "1-2. 200: 必須項目のみ",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_RequireItemOnly(),
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},

		{
			name:        "1-3. 200: 検索結果なし",
			input:       f.NewGetTradeResponseInput(),
			receive:     f.GetTradeRequestsReceived_NoData(),
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
			expect:  expectedPagingError,
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
					assert.Equal(t, test.expect.(common.CustomError).Code, err.(*common.CustomError).Code)
					assert.Equal(t, test.expect.(common.CustomError).Message, err.(*common.CustomError).Message)
					assert.Equal(t, test.expect.(common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
					assert.Equal(t, test.expect.(common.CustomError).Source, err.(*common.CustomError).Source)
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

	res := f.NewPutTradeRequestModel(false)
	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	res.StatusModel.StatusID = uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478")
	res.StatusModel.TradeID = tradeID
	res.TradeModel.TradeID = &tradeID

	tests := []struct {
		name    string
		input   traceability.TradeRequestModel
		receive string
		expect  traceability.TradeRequestModel
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewPutTradeRequestModel(false),
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
				traceabilityRepositoryMock.On("PostTradeRequests", mock.Anything, mock.Anything).Return(putTradeRequestsResponse, nil)

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, err := tradeUsecase.PutTradeRequest(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.StatusModel.StatusID, actualRes.StatusModel.StatusID, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.TradeID, actualRes.StatusModel.TradeID, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.RequestStatus, actualRes.StatusModel.RequestStatus, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.Message, actualRes.StatusModel.Message, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.ReplyMessage, actualRes.StatusModel.ReplyMessage, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.RequestType, actualRes.StatusModel.RequestType, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.TradeID, actualRes.TradeModel.TradeID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.DownstreamOperatorID, actualRes.TradeModel.DownstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.DownstreamTraceID, actualRes.TradeModel.DownstreamTraceID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.UpstreamOperatorID, actualRes.TradeModel.UpstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.UpstreamTraceID, actualRes.TradeModel.UpstreamTraceID, f.AssertMessage)
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
		input   traceability.TradeRequestModel
		receive string
		expect  error
	}{
		{
			name:    "2-1. 400: トレース識別子存在エラー",
			input:   f.NewPutTradeRequestModel(false),
			receive: f.Error_TraceIdNotFound(),
			expect:  expectedExistError,
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
				traceabilityRepositoryMock.On("PostTradeRequests", mock.Anything, mock.Anything).Return(traceabilityentity.PostTradeRequestsResponses{}, getTradeResponse)

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := tradeUsecase.PutTradeRequest(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.(common.CustomError).Code, err.(*common.CustomError).Code)
					assert.Equal(t, test.expect.(common.CustomError).Message, err.(*common.CustomError).Message)
					assert.Equal(t, test.expect.(common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
					assert.Equal(t, test.expect.(common.CustomError).Source, err.(*common.CustomError).Source)
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
func TestProjectUsecaseTraceability_PutTradeResponse(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	res := f.NewPutTradeResponseModel()

	tests := []struct {
		name         string
		input        traceability.PutTradeResponseInput
		receiveTrRes string
		receiveTrRec string
		expect       traceability.TradeModel
	}{
		{
			name:         "1-1. 200: 全項目応答",
			input:        f.PutTradeResponseInput,
			receiveTrRes: f.PostTrades(),
			receiveTrRec: f.GetTradeRequestsReceived_AllItem(),
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
				traceabilityRepositoryMock.On("PostTrades", mock.Anything, mock.Anything).Return(postTradesResponse, nil)
				traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(getTradesResponse, nil)

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, err := tradeUsecase.PutTradeResponse(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.TradeID, actualRes.TradeID, f.AssertMessage)
					assert.Equal(t, test.expect.DownstreamOperatorID, actualRes.DownstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.DownstreamTraceID, actualRes.DownstreamTraceID, f.AssertMessage)
					assert.Equal(t, test.expect.UpstreamOperatorID, actualRes.UpstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.UpstreamTraceID, actualRes.UpstreamTraceID, f.AssertMessage)
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
			input:             f.PutTradeResponseInput,
			receiveTrRes:      nil,
			receiveTrResError: common.StringPtr(f.Error_PagingError()),
			receiveTrRec:      nil,
			receiveTrRecError: nil,
			expect:            expectedPagingError,
		},
		{
			name:              "2-2. 400: ページングエラー(取得)",
			input:             f.PutTradeResponseInput,
			receiveTrRes:      common.StringPtr(f.PostTrades()),
			receiveTrResError: nil,
			receiveTrRec:      nil,
			receiveTrRecError: common.StringPtr(f.Error_PagingError()),
			expect:            expectedPagingError,
		},
		{
			name:              "2-3. 400: 対象データなし",
			input:             f.PutTradeResponseInput,
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
					traceabilityRepositoryMock.On("PostTrades", mock.Anything, mock.Anything).Return(postTradesResponse, nil)
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
					traceabilityRepositoryMock.On("PostTrades", mock.Anything, mock.Anything).Return(traceabilityentity.PostTradesResponse{}, postTradesResponse)
				}

				tradeUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := tradeUsecase.PutTradeResponse(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					if test.name == "2-3. 400: 対象データなし" {
						assert.Equal(t, test.expect.Error(), err.Error())
					} else {
						assert.Equal(t, test.expect.(common.CustomError).Code, err.(*common.CustomError).Code)
						assert.Equal(t, test.expect.(common.CustomError).Message, err.(*common.CustomError).Message)
						assert.Equal(t, test.expect.(common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
						assert.Equal(t, test.expect.(common.CustomError).Source, err.(*common.CustomError).Source)
					}
				}
			},
		)
	}
}
