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

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答(依頼)
// [x] 1-2. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含む)
// [x] 1-3. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含まない)
// [x] 1-4. 200: 全項目応答(受領依頼)
// [x] 1-5. 200: 全項目応答(受領依頼)(トレサビレスポンスにnullを含む)
// [x] 1-6. 200: 全項目応答(受領依頼)(トレサビレスポンスにnullを含まない)
// [x] 1-7. 200: 全項目応答(両方)
// [x] 1-8. 200: 回答済(必須項目のみ)(依頼)
// [x] 1-9. 200: 回答済(必須項目のみ)(依頼)(キーなし)
// [x] 1-10. 200: 未回答(必須項目のみ)(依頼)
// [x] 1-11. 200: 未回答(必須項目のみ)(依頼)(キーなし)
// [x] 1-12. 200: 未回答(必須項目のみ)(受領依頼)
// [x] 1-13. 200: 未回答(必須項目のみ)(受領依頼)(キーなし)
// [x] 1-14. 200: 必須項目のみ(両方)
// [x] 1-15. 200: 検索結果なし(依頼)
// [x] 1-16. 200: 検索結果なし(受領依頼)
// [x] 1-17. 200: 検索結果なし(両方)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetStatus(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	getStatusInput := f.NewGetStatusInput()
	getStatusInputRequest := f.NewGetStatusInput()
	getStatusInputRequest.StatusTarget = traceability.Request
	getStatusInputResponse := f.NewGetStatusInput()
	getStatusInputResponse.StatusTarget = traceability.Response

	cfpResponseStatusComplete := traceability.CfpResponseStatusComplete
	tradeTreeStatusUnterminated := traceability.TradeTreeStatusUnterminated
	expectedResAll := []traceability.StatusModel{
		{
			StatusID: uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus:        &cfpResponseStatusComplete,
				TradeTreeStatus:          &tradeTreeStatusUnterminated,
				CompletedCount:           &f.CompletedCount,
				CompletedCountModifiedAt: &f.CompletedCountModifiedAt,
				TradesCount:              &f.TradesCount,
				TradesCountModifiedAt:    &f.TradesCountModifiedAt,
			},
			Message:         common.StringPtr("A01のCFP値を回答ください"),
			ReplyMessage:    common.StringPtr("A01のCFP値を回答しました"),
			RequestType:     f.RequestType.ToString(),
			ResponseDueDate: &f.ResponseDueDate,
		},
	}
	expectedResAllWithNull := []traceability.StatusModel{
		{
			StatusID: uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus:        &cfpResponseStatusComplete,
				TradeTreeStatus:          &tradeTreeStatusUnterminated,
				CompletedCount:           nil,
				CompletedCountModifiedAt: nil,
				TradesCount:              nil,
				TradesCountModifiedAt:    nil,
			},
			Message:         common.StringPtr("A01のCFP値を回答ください"),
			ReplyMessage:    common.StringPtr("A01のCFP値を回答しました"),
			RequestType:     f.RequestType.ToString(),
			ResponseDueDate: nil,
		},
	}

	expectedResRequireOnly := []traceability.StatusModel{
		{
			StatusID: uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus:        &cfpResponseStatusComplete,
				TradeTreeStatus:          &tradeTreeStatusUnterminated,
				CompletedCount:           &f.CompletedCount,
				CompletedCountModifiedAt: &f.CompletedCountModifiedAt,
				TradesCount:              &f.TradesCount,
				TradesCountModifiedAt:    &f.TradesCountModifiedAt,
			},
			Message:         common.StringPtr(""),
			ReplyMessage:    nil,
			RequestType:     f.RequestType.ToString(),
			ResponseDueDate: &f.ResponseDueDate,
		},
	}

	expectedResNoData := []traceability.StatusModel{}

	req := traceability.Request
	res := traceability.Response
	tests := []struct {
		name         string
		statusTarget *traceability.StatusTarget
		input        traceability.GetStatusInput
		receiveReq   string
		receiveRes   string
		expectData   []traceability.StatusModel
		expectAfter  *string
	}{
		{
			name:         "1-1. 200: 全項目応答(依頼)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_AllItem(),
			expectData:   expectedResAll,
			expectAfter:  common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:         "1-2. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含む)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_AllItem_WithNull(),
			expectData:   expectedResAllWithNull,
			expectAfter:  common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:         "1-3. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含まない)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_AllItem_WithUndefined(),
			expectData:   expectedResAllWithNull,
			expectAfter:  common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:         "1-4. 200: 全項目応答(受領依頼)",
			statusTarget: &res,
			input:        getStatusInputResponse,
			receiveRes:   f.GetTradeRequestsReceived_AllItem(),
			expectData:   expectedResAll,
			expectAfter:  common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:         "1-5. 200: 全項目応答(受領依頼)(トレサビレスポンスにnullを含む)",
			statusTarget: &res,
			input:        getStatusInputResponse,
			receiveRes:   f.GetTradeRequestsReceived_AllItem_WithNull(),
			expectData:   expectedResAllWithNull,
			expectAfter:  common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:         "1-6. 200: 全項目応答(受領依頼)(トレサビレスポンスにnullを含まない)",
			statusTarget: &res,
			input:        getStatusInputResponse,
			receiveRes:   f.GetTradeRequestsReceived_AllItem_WithUndefined(),
			expectData:   expectedResAllWithNull,
			expectAfter:  common.StringPtr("026ad6a0-a689-4b8c-8a14-7304b817096d"),
		},
		{
			name:         "1-7. 200: 全項目応答(両方)",
			statusTarget: nil,
			input:        getStatusInput,
			receiveReq:   f.GetTradeRequests_AllItem_NoNext(),
			receiveRes:   f.GetTradeRequestsReceived_AllItem_NoNext(),
			expectData:   expectedResAll,
			expectAfter:  nil,
		},
		{
			name:         "1-8. 200: 回答済(必須項目のみ)(依頼)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_RequireItemOnlyAnswered(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-9. 200: 回答済(必須項目のみ)(キーなし)(依頼)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_RequireItemOnlyAnsweredWithUndefined(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-10. 200: 未回答(必須項目のみ)(依頼)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_RequireItemOnlyAnswering(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-11. 200: 未回答(必須項目のみ)(キーなし)(依頼)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_RequireItemOnlyAnsweringWithUndefined(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-12. 200: 未回答(受領依頼)",
			statusTarget: &res,
			input:        getStatusInputResponse,
			receiveRes:   f.GetTradeRequestsReceived_RequireItemOnly(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-13. 200: 未回答(受領依頼)(キーなし)",
			statusTarget: &res,
			input:        getStatusInputResponse,
			receiveRes:   f.GetTradeRequestsReceived_RequireItemOnlyWithUndefined(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-14. 200: 必須項目のみ(両方)",
			statusTarget: nil,
			input:        getStatusInput,
			receiveReq:   f.GetTradeRequests_RequireItemOnlyAnswered(),
			receiveRes:   f.GetTradeRequestsReceived_RequireItemOnly(),
			expectData:   expectedResRequireOnly,
			expectAfter:  nil,
		},
		{
			name:         "1-15. 200: 検索結果なし(依頼)",
			statusTarget: &req,
			input:        getStatusInputRequest,
			receiveReq:   f.GetTradeRequests_NoData(),
			expectData:   expectedResNoData,
			expectAfter:  nil,
		},
		{
			name:         "1-16. 200: 検索結果なし(受領依頼)",
			statusTarget: &res,
			input:        getStatusInputResponse,
			receiveRes:   f.GetTradeRequestsReceived_NoData(),
			expectData:   expectedResNoData,
			expectAfter:  nil,
		},
		{
			name:         "1-17. 200: 検索結果なし(両方)",
			statusTarget: nil,
			input:        getStatusInput,
			receiveReq:   f.GetTradeRequests_NoData(),
			receiveRes:   f.GetTradeRequestsReceived_NoData(),
			expectData:   expectedResNoData,
			expectAfter:  nil,
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
				if test.statusTarget == nil {
					getTradeRequestsReceivedResponse := traceabilityentity.GetTradeRequestsReceivedResponse{}
					if err := json.Unmarshal([]byte(test.receiveRes), &getTradeRequestsReceivedResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					getTradeRequestsResponse := traceabilityentity.GetTradeRequestsResponse{}
					if err := json.Unmarshal([]byte(test.receiveReq), &getTradeRequestsResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(getTradeRequestsResponse, nil)

					traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(getTradeRequestsReceivedResponse, nil)
				} else if *test.statusTarget == traceability.Request {
					getTradeRequestsResponse := traceabilityentity.GetTradeRequestsResponse{}
					if err := json.Unmarshal([]byte(test.receiveReq), &getTradeRequestsResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(getTradeRequestsResponse, nil)
				} else if *test.statusTarget == traceability.Response {
					getTradeRequestsReceivedResponse := traceabilityentity.GetTradeRequestsReceivedResponse{}
					if err := json.Unmarshal([]byte(test.receiveRes), &getTradeRequestsReceivedResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(getTradeRequestsReceivedResponse, nil)
				}

				usecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, actualAfter, err := usecase.GetStatus(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, test.expectData, actualRes, f.AssertMessage)
					assert.Equal(t, test.expectAfter, actualAfter, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: ページングエラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetStatus_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	getStatusInput := f.NewGetStatusInput()
	getStatusInputRequest := f.NewGetStatusInput()
	getStatusInputRequest.StatusTarget = traceability.Request
	getStatusInputResponse := f.NewGetStatusInput()
	getStatusInputResponse.StatusTarget = traceability.Response

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	req := traceability.Request
	res := traceability.Response
	tests := []struct {
		name            string
		statusTarget    *traceability.StatusTarget
		input           traceability.GetStatusInput
		receiveReqError error
		receiveResError error
		expect          error
	}{
		{
			name:            "2-1. 400: ページングエラー",
			statusTarget:    &req,
			input:           getStatusInputRequest,
			receiveReqError: expectedPagingError,
			receiveResError: nil,
			expect:          expectedPagingError,
		},
		{
			name:            "2-2. 400: ページングエラー",
			statusTarget:    &res,
			input:           getStatusInputResponse,
			receiveReqError: nil,
			receiveResError: expectedPagingError,
			expect:          expectedPagingError,
		},
		{
			name:            "2-3. 400: ページングエラー",
			statusTarget:    nil,
			input:           getStatusInput,
			receiveReqError: nil,
			receiveResError: expectedPagingError,
			expect:          expectedPagingError,
		},
		{
			name:            "2-4. 400: ページングエラー",
			statusTarget:    nil,
			input:           getStatusInput,
			receiveReqError: expectedPagingError,
			receiveResError: nil,
			expect:          expectedPagingError,
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
				if test.statusTarget == nil {
					traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsResponse{}, test.receiveReqError)
					traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsReceivedResponse{}, test.receiveResError)
				} else if *test.statusTarget == traceability.Request {
					traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsResponse{}, test.receiveReqError)
				} else if *test.statusTarget == traceability.Response {
					traceabilityRepositoryMock.On("GetTradeRequestsReceived", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsReceivedResponse{}, test.receiveResError)
				}

				usecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, actualAfter, err := usecase.GetStatus(c, test.input)
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

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutStatusCancel(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive string
		expect  error
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewPutStatusInput(),
			receive: f.PutPostTradeRequestsCancelResponse(),
			expect:  nil,
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
				postTradeRequestsRejectResponse := traceabilityentity.PostTradeRequestsCancelResponse{}

				if err := json.Unmarshal([]byte(test.receive), &postTradeRequestsRejectResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}
				traceabilityRepositoryMock.On("PostTradeRequestsCancel", mock.Anything, mock.Anything).Return(postTradeRequestsRejectResponse, common.ResponseHeaders{}, nil)

				usecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := usecase.PutStatusCancel(c, test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutStatusCancel_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive error
		expect  error
	}{
		{
			name:    "2-1. 400: データ取得エラー",
			input:   f.NewPutStatusInput(),
			receive: fmt.Errorf("Trade not found"),
			expect:  fmt.Errorf("Trade not found"),
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
				traceabilityRepositoryMock.On("PostTradeRequestsCancel", mock.Anything, mock.Anything).Return(traceabilityentity.PostTradeRequestsCancelResponse{}, common.ResponseHeaders{}, test.receive)

				usecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := usecase.PutStatusCancel(c, test.input)
				assert.Error(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutStatusReject(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive string
		expect  error
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewPutStatusInput(),
			receive: f.PutPostTradeRequestsRejectResponse(),
			expect:  nil,
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
				postTradeRequestsRejectResponse := traceabilityentity.PostTradeRequestsRejectResponse{}

				if err := json.Unmarshal([]byte(test.receive), &postTradeRequestsRejectResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}
				traceabilityRepositoryMock.On("PostTradeRequestsReject", mock.Anything, mock.Anything).Return(postTradeRequestsRejectResponse, common.ResponseHeaders{}, nil)

				usecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := usecase.PutStatusReject(c, test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutStatusReject_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive error
		expect  error
	}{
		{
			name:    "2-1. 400: データ取得エラー",
			input:   f.NewPutStatusInput(),
			receive: fmt.Errorf("Trade not found"),
			expect:  fmt.Errorf("Trade not found"),
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
				traceabilityRepositoryMock.On("PostTradeRequestsReject", mock.Anything, mock.Anything).Return(traceabilityentity.PostTradeRequestsRejectResponse{}, common.ResponseHeaders{}, test.receive)

				usecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := usecase.PutStatusReject(c, test.input)
				assert.Error(t, err)
			},
		)
	}
}
