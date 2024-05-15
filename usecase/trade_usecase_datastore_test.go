package usecase_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// TestProjectUsecaseDatastore_GetTradeRequest
// Summary: This is normal test class which confirm the operation of API #10 GetTradeRequest.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 回答済(必須項目のみ)
// [x] 1-3. 200: 未回答(必須項目のみ)
// [x] 1-4. 200: 検索結果なし
func TestProjectUsecaseDatastore_GetTradeRequest(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	upstreamOperatorID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	upstreamTraceID := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	dsResAll := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      &upstreamTraceID,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}

	dsResAnsweredRequireOnly := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      &upstreamTraceID,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}

	dsResAnsweringRequireOnly := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      nil,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}

	dsResNoData := traceability.TradeEntityModels{}

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
		receive     traceability.TradeEntityModels
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetTradeRequestInput(),
			receive:     dsResAll,
			expectData:  dsExpectedResAll,
			expectAfter: nil,
		},
		{
			name:        "1-2. 200: 回答済(必須項目のみ)",
			input:       f.NewGetTradeRequestInput(),
			receive:     dsResAnsweredRequireOnly,
			expectData:  dsExpectedResAnsweredRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 未回答(必須項目のみ)",
			input:       f.NewGetTradeRequestInput(),
			receive:     dsResAnsweringRequireOnly,
			expectData:  dsExpectedResAnsweringRequireOnly,
			expectAfter: nil,
		},

		{
			name:        "1-4. 200: 検索結果なし",
			input:       f.NewGetTradeRequestInput(),
			receive:     dsResNoData,
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

				var expected []traceability.TradeModel
				err := json.Unmarshal([]byte(test.expectData), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetTradeRequest", mock.Anything, mock.Anything, mock.Anything).Return(test.receive, nil)
				ouranosRepositoryMock.On("CountTradeRequest", mock.Anything).Return(1, nil)

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				actualRes, actualAfter, err := tradeUsecase.GetTradeRequest(c, test.input)
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

// TestProjectUsecaseDatastore_GetTradeRequest_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #10 GetTradeRequest.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 2-1. 400: データ取得エラー
// [x] 2-1. 400: 件数取得エラー
func TestProjectUsecaseDatastore_GetTradeRequest_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	dsResGetError := fmt.Errorf("DB AccessError")

	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	upstreamOperatorID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	upstreamTraceID := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	dsResDataCountGetError := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      &upstreamTraceID,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}

	dsResCountGetError := fmt.Errorf("DB AccessError")

	tests := []struct {
		name         string
		input        traceability.GetTradeRequestInput
		receive      traceability.TradeEntityModels
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー",
			input:        f.NewGetTradeRequestInput(),
			receive:      nil,
			receiveError: dsResGetError,
			expect:       dsResGetError,
		},
		{
			name:         "2-2. 400: 件数取得エラー",
			input:        f.NewGetTradeRequestInput(),
			receive:      dsResDataCountGetError,
			receiveError: dsResCountGetError,
			expect:       dsResCountGetError,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				if test.name == "2-1. 400: データ取得エラー" {
					ouranosRepositoryMock.On("GetTradeRequest", mock.Anything, mock.Anything, mock.Anything).Return(traceability.TradeEntityModels{}, test.receiveError)
				} else if test.name == "2-2. 400: 件数取得エラー" {
					ouranosRepositoryMock.On("GetTradeRequest", mock.Anything, mock.Anything, mock.Anything).Return(test.receive, nil)
					ouranosRepositoryMock.On("CountTradeRequest", mock.Anything).Return(1, test.receiveError)
				}

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				_, _, err := tradeUsecase.GetTradeRequest(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// TestProjectUsecaseDatastore_PutTradeRequest
// Summary: This is normal test class which confirm the operation of API #7 PutTradeRequest.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答(新規)
// [x] 1-1. 200: 全項目応答(更新)
func TestProjectUsecaseDatastore_PutTradeRequest(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	tradeID := uuid.MustParse(f.TradeID)
	upstreamOperatorID := uuid.MustParse(f.OperatorID)
	upstreamTraceID := uuid.MustParse(f.TraceID2)

	dsResData := traceability.TradeRequestEntityModel{
		TradeEntityModel: traceability.TradeEntityModel{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse(f.OperatorID),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse(f.TraceID),
			UpstreamTraceID:      &upstreamTraceID,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
		StatusEntityModel: traceability.StatusEntityModel{
			StatusID:          uuid.MustParse(f.StatusID),
			TradeID:           tradeID,
			CfpResponseStatus: traceability.CfpResponseStatusPending.ToString(),
			TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
			Message:           &f.TradeRequestMessage,
			ReplyMessage:      nil,
			RequestType:       traceability.RequestTypeCFP.ToString(),
			DeletedAt:         gorm.DeletedAt{Time: time.Now()},
			CreatedAt:         time.Now(),
			CreatedUserId:     "seed",
			UpdatedAt:         time.Now(),
			UpdatedUserId:     "seed",
		},
	}

	expect := traceability.TradeRequestModel{
		TradeModel: traceability.TradeModel{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse(f.OperatorID),
			UpstreamOperatorID:   upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse(f.TraceID),
			UpstreamTraceID:      &upstreamTraceID,
		},
		StatusModel: traceability.StatusModel{
			StatusID: uuid.MustParse(f.StatusID),
			TradeID:  tradeID,
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus: traceability.CfpResponseStatusPending,
				TradeTreeStatus:   traceability.TradeTreeStatusUnterminated,
			},
			Message:      &f.TradeRequestMessage,
			ReplyMessage: nil,
			RequestType:  traceability.RequestTypeCFP.ToString(),
		},
	}

	tests := []struct {
		name    string
		input   traceability.TradeRequestModel
		receive traceability.TradeRequestEntityModel
		expect  traceability.TradeRequestModel
	}{
		{
			name:    "1-1. 200: 全項目応答(新規)",
			input:   f.NewPutTradeRequestModel(true),
			receive: dsResData,
			expect:  expect,
		},
		{
			name:    "1-2. 200: 全項目応答(更新)",
			input:   f.NewPutTradeRequestModel(false),
			receive: dsResData,
			expect:  expect,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("PutTradeRequest", mock.Anything, mock.Anything).Return(test.receive, nil)

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				actualRes, err := tradeUsecase.PutTradeRequest(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					//assert.Equal(t, test.expect.StatusModel.StatusID, actualRes.StatusModel.StatusID, f.AssertMessage)
					//assert.Equal(t, test.expect.StatusModel.TradeID, actualRes.StatusModel.TradeID, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.RequestStatus, actualRes.StatusModel.RequestStatus, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.Message, actualRes.StatusModel.Message, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.ReplyMessage, actualRes.StatusModel.ReplyMessage, f.AssertMessage)
					assert.Equal(t, test.expect.StatusModel.RequestType, actualRes.StatusModel.RequestType, f.AssertMessage)
					//assert.Equal(t, test.expect.TradeModel.TradeID, actualRes.TradeModel.TradeID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.DownstreamOperatorID, actualRes.TradeModel.DownstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.DownstreamTraceID, actualRes.TradeModel.DownstreamTraceID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.UpstreamOperatorID, actualRes.TradeModel.UpstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.TradeModel.UpstreamTraceID, actualRes.TradeModel.UpstreamTraceID, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecaseDatastore_PutTradeRequest_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #7 PutTradeRequest.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 2-1. 400: データ更新エラー
func TestProjectUsecaseDatastore_PutTradeRequest_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeRequest"

	dsResPutError := fmt.Errorf("DB AccessError")

	tests := []struct {
		name    string
		input   traceability.TradeRequestModel
		receive error
		expect  error
	}{
		{
			name:    "2-1. 400: データ更新エラー",
			input:   f.NewPutTradeRequestModel(false),
			receive: dsResPutError,
			expect:  dsResPutError,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("PutTradeRequest", mock.Anything, mock.Anything).Return(traceability.TradeRequestEntityModel{}, test.receive)

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				_, err := tradeUsecase.PutTradeRequest(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// TestProjectUsecaseDatastore_GetTradeResponse
// Summary: This is normal test class which confirm the operation of API #12 GetTradeResponse.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 検索結果なし
func TestProjectUsecaseDatastore_GetTradeResponse(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	upstreamOperatorID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	upstreamTraceID := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	plantID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResAllTrade := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      &upstreamTraceID,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}
	dsResAllStatus := traceability.StatusEntityModel{
		StatusID:          uuid.MustParse(f.StatusID),
		TradeID:           tradeID,
		CfpResponseStatus: traceability.CfpResponseStatusPending.ToString(),
		TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
		Message:           &f.TradeRequestMessage,
		ReplyMessage:      &f.TradeRequestMessage,
		RequestType:       f.RequestType.ToString(),
		DeletedAt:         gorm.DeletedAt{Time: time.Now()},
		CreatedAt:         time.Now(),
		CreatedUserId:     "seed",
		UpdatedAt:         time.Now(),
		UpdatedUserId:     "seed",
	}
	dsResAllParts := traceability.PartsModelEntity{
		TraceID:            uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		OperatorID:         uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PartsName:          "B01",
		SupportPartsName:   common.StringPtr("B01001"),
		TerminatedFlag:     true,
		AmountRequired:     common.Float64Ptr(1),
		AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
		DeletedAt:          gorm.DeletedAt{Time: time.Now()},
		CreatedAt:          time.Now(),
		CreatedUserId:      "seed",
		UpdatedAt:          time.Now(),
		UpdatedUserId:      "seed",
	}

	dsResRequireOnlyTrade := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      nil,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}
	dsResRequireOnlyStatus := traceability.StatusEntityModel{
		StatusID:          uuid.MustParse(f.StatusID),
		TradeID:           tradeID,
		CfpResponseStatus: traceability.CfpResponseStatusPending.ToString(),
		TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
		Message:           nil,
		ReplyMessage:      nil,
		RequestType:       f.RequestType.ToString(),
		DeletedAt:         gorm.DeletedAt{Time: time.Now()},
		CreatedAt:         time.Now(),
		CreatedUserId:     "seed",
		UpdatedAt:         time.Now(),
		UpdatedUserId:     "seed",
	}
	dsResRequireOnlyParts := traceability.PartsModelEntity{
		TraceID:            uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		OperatorID:         uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PartsName:          "B01",
		SupportPartsName:   nil,
		TerminatedFlag:     true,
		AmountRequired:     common.Float64Ptr(1),
		AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
		DeletedAt:          gorm.DeletedAt{Time: time.Now()},
		CreatedAt:          time.Now(),
		CreatedUserId:      "seed",
		UpdatedAt:          time.Now(),
		UpdatedUserId:      "seed",
	}

	dsResNoDataTrade := traceability.TradeEntityModels{}
	dsResNoDataStatus := traceability.StatusEntityModel{}
	dsResNoDataParts := traceability.PartsModelEntity{}

	dsExpectedResAll := []traceability.TradeResponseModel{
		{
			StatusModel: traceability.StatusModel{
				StatusID: uuid.MustParse(f.StatusID),
				TradeID:  tradeID,
				RequestStatus: traceability.RequestStatus{
					CfpResponseStatus: traceability.CfpResponseStatusPending,
					TradeTreeStatus:   traceability.TradeTreeStatusUnterminated,
				},
				ReplyMessage: nil,
				Message:      &f.TradeRequestMessage,
				RequestType:  f.RequestType.ToString(),
			},
			TradeModel: traceability.TradeModel{
				TradeID:              &tradeID,
				DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
				UpstreamOperatorID:   upstreamOperatorID,
				DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				UpstreamTraceID:      &upstreamTraceID,
			},
			PartsModel: traceability.PartsModel{
				TraceID:            uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				OperatorID:         uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
				PlantID:            &plantID,
				PartsName:          "B01",
				SupportPartsName:   common.StringPtr("B01001"),
				TerminatedFlag:     true,
				AmountRequired:     nil,
				AmountRequiredUnit: &amountRequiredUnit,
			},
		},
	}

	dsExpectedRequireOnly := []traceability.TradeResponseModel{
		{
			StatusModel: traceability.StatusModel{
				StatusID: uuid.MustParse(f.StatusID),
				TradeID:  tradeID,
				RequestStatus: traceability.RequestStatus{
					CfpResponseStatus: traceability.CfpResponseStatusPending,
					TradeTreeStatus:   traceability.TradeTreeStatusUnterminated,
				},
				ReplyMessage: nil,
				Message:      nil,
				RequestType:  f.RequestType.ToString(),
			},
			TradeModel: traceability.TradeModel{
				TradeID:              &tradeID,
				DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
				UpstreamOperatorID:   upstreamOperatorID,
				DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				UpstreamTraceID:      nil,
			},
			PartsModel: traceability.PartsModel{
				TraceID:            uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				OperatorID:         uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
				PlantID:            &plantID,
				PartsName:          "B01",
				SupportPartsName:   nil,
				TerminatedFlag:     true,
				AmountRequired:     nil,
				AmountRequiredUnit: &amountRequiredUnit,
			},
		},
	}

	dsExpectedResNoData := []traceability.TradeResponseModel{}

	tests := []struct {
		name          string
		input         traceability.GetTradeResponseInput
		receiveTrade  traceability.TradeEntityModels
		receiveStatus traceability.StatusEntityModel
		receiveParts  traceability.PartsModelEntity
		expectData    []traceability.TradeResponseModel
		expectAfter   *string
	}{
		{
			name:          "1-1. 200: 全項目応答",
			input:         f.NewGetTradeResponseInput(),
			receiveTrade:  dsResAllTrade,
			receiveStatus: dsResAllStatus,
			receiveParts:  dsResAllParts,
			expectData:    dsExpectedResAll,
			expectAfter:   nil,
		},
		{
			name:          "1-2. 200: 必須項目のみ",
			input:         f.NewGetTradeResponseInput(),
			receiveTrade:  dsResRequireOnlyTrade,
			receiveStatus: dsResRequireOnlyStatus,
			receiveParts:  dsResRequireOnlyParts,
			expectData:    dsExpectedRequireOnly,
			expectAfter:   nil,
		},
		{
			name:          "1-3. 200: 検索結果なし",
			input:         f.NewGetTradeResponseInput(),
			receiveTrade:  dsResNoDataTrade,
			receiveStatus: dsResNoDataStatus,
			receiveParts:  dsResNoDataParts,
			expectData:    dsExpectedResNoData,
			expectAfter:   nil,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetTradeResponse", mock.Anything, mock.Anything).Return(test.receiveTrade, nil)
				ouranosRepositoryMock.On("CountTradeResponse", mock.Anything).Return(1, nil)
				ouranosRepositoryMock.On("GetStatusByTradeID", mock.Anything).Return(test.receiveStatus, nil)
				ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, nil)

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				actualRes, actualAfter, err := tradeUsecase.GetTradeResponse(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, test.expectData, actualRes, f.AssertMessage)
					assert.Nil(t, actualAfter)
				}
			},
		)
	}
}

// TestProjectUsecaseDatastore_GetTradeResponse_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #12 GetTradeResponse.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 2-1. 400: データ取得エラー(Trade)
// [x] 2-2. 400: データ取得エラー(Trade_Count)
// [x] 2-3. 400: データ取得エラー(Status)
// [x] 2-4. 400: データ取得エラー(Parts)
func TestProjectUsecaseDatastore_GetTradeResponse_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	dsResGetError := fmt.Errorf("DB AccessError")
	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	upstreamOperatorID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	upstreamTraceID := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	dsResDataErrorTrade := traceability.TradeEntityModels{
		{
			TradeID:              &tradeID,
			DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
			UpstreamOperatorID:   &upstreamOperatorID,
			DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
			UpstreamTraceID:      &upstreamTraceID,
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{Time: time.Now()},
			CreatedAt:            time.Now(),
			CreatedUserID:        "seed",
			UpdatedAt:            time.Now(),
			UpdatedUserID:        "seed",
		},
	}
	dsResDataErrorStatus := traceability.StatusEntityModel{}
	dsResDataErrorParts := traceability.PartsModelEntity{}

	tests := []struct {
		name               string
		input              traceability.GetTradeResponseInput
		receiveTrade       traceability.TradeEntityModels
		receiveTradeError  error
		receiveCountError  error
		receiveStatus      traceability.StatusEntityModel
		receiveStatusError error
		receiveParts       traceability.PartsModelEntity
		receivePartsError  error
		expectData         error
		expectAfter        *string
	}{
		{
			name:               "2-1. 400: データ取得エラー(Trade)",
			input:              f.NewGetTradeResponseInput(),
			receiveTrade:       dsResDataErrorTrade,
			receiveTradeError:  dsResGetError,
			receiveCountError:  nil,
			receiveStatus:      dsResDataErrorStatus,
			receiveStatusError: nil,
			receiveParts:       dsResDataErrorParts,
			receivePartsError:  nil,
			expectData:         dsResGetError,
			expectAfter:        nil,
		},
		{
			name:               "2-2. 400: データ取得エラー(Trade_Count)",
			input:              f.NewGetTradeResponseInput(),
			receiveTrade:       dsResDataErrorTrade,
			receiveTradeError:  nil,
			receiveCountError:  dsResGetError,
			receiveStatus:      dsResDataErrorStatus,
			receiveStatusError: nil,
			receiveParts:       dsResDataErrorParts,
			receivePartsError:  nil,
			expectData:         dsResGetError,
			expectAfter:        nil,
		},
		{
			name:               "2-3. 400: データ取得エラー(Status)",
			input:              f.NewGetTradeResponseInput(),
			receiveTrade:       dsResDataErrorTrade,
			receiveTradeError:  nil,
			receiveCountError:  nil,
			receiveStatus:      dsResDataErrorStatus,
			receiveStatusError: dsResGetError,
			receiveParts:       dsResDataErrorParts,
			receivePartsError:  nil,
			expectData:         dsResGetError,
			expectAfter:        nil,
		},
		{
			name:               "2-4. 400: データ取得エラー(Parts)",
			input:              f.NewGetTradeResponseInput(),
			receiveTrade:       dsResDataErrorTrade,
			receiveTradeError:  nil,
			receiveCountError:  nil,
			receiveStatus:      dsResDataErrorStatus,
			receiveStatusError: nil,
			receiveParts:       dsResDataErrorParts,
			receivePartsError:  dsResGetError,
			expectData:         dsResGetError,
			expectAfter:        nil,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetTradeResponse", mock.Anything, mock.Anything).Return(test.receiveTrade, test.receiveTradeError)
				ouranosRepositoryMock.On("CountTradeResponse", mock.Anything).Return(1, test.receiveCountError)
				ouranosRepositoryMock.On("GetStatusByTradeID", mock.Anything).Return(test.receiveStatus, test.receiveStatusError)
				ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, test.receivePartsError)

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				_, actualAfter, err := tradeUsecase.GetTradeResponse(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expectData, err, f.AssertMessage)
					assert.Nil(t, actualAfter)
				}
			},
		)
	}
}

// TestProjectUsecaseDatastore_PutTradeResponse
// Summary: This is normal test class which confirm the operation of API #13 PutTradeResponse.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系(入力TraceId該当あり)
// [x] 1-2. 200: 正常系(入力TraceId該当なし)
func TestProjectUsecaseDatastore_PutTradeResponse(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	upstreamOperatorID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	upstreamTraceID := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	CFPID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResCFP := traceability.CfpEntityModel{
		CfpID:              &CFPID,
		TraceID:            upstreamTraceID,
		GhgEmission:        &f.GhgEmission,
		GhgDeclaredUnit:    f.GhgDeclaredUnit,
		CfpCertificateList: f.CfpCertificateList,
		CfpType:            f.CfpType.ToString(),
		DqrType:            traceability.DqrPreProcessingResponse.ToString(),
		TeR:                &f.TeR,
		GeR:                &f.GeR,
		TiR:                &f.TiR,
		DeletedAt:          gorm.DeletedAt{Time: time.Now()},
		CreatedAt:          time.Now(),
		CreatedUserId:      "seed",
		UpdatedAt:          time.Now(),
		UpdatedUserId:      "seed",
	}
	dsResParts := traceability.PartsModelEntity{
		TraceID:            uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		OperatorID:         uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PartsName:          "B01",
		SupportPartsName:   common.StringPtr("B01001"),
		TerminatedFlag:     true,
		AmountRequired:     common.Float64Ptr(1),
		AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
		DeletedAt:          gorm.DeletedAt{Time: time.Now()},
		CreatedAt:          time.Now(),
		CreatedUserId:      "seed",
		UpdatedAt:          time.Now(),
		UpdatedUserId:      "seed",
	}
	dsResTrade := traceability.TradeEntityModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		UpstreamOperatorID:   &upstreamOperatorID,
		DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		UpstreamTraceID:      &upstreamTraceID,
		TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
		DeletedAt:            gorm.DeletedAt{Time: time.Now()},
		CreatedAt:            time.Now(),
		CreatedUserID:        "seed",
		UpdatedAt:            time.Now(),
		UpdatedUserID:        "seed",
	}

	dsExpectedRes := traceability.TradeModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		UpstreamOperatorID:   upstreamOperatorID,
		DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		UpstreamTraceID:      &upstreamTraceID,
	}

	tests := []struct {
		name            string
		input           traceability.PutTradeResponseInput
		receiveCFP      traceability.CfpEntityModel
		receiveCFPError error
		receiveParts    traceability.PartsModelEntity
		receiveTrade    traceability.TradeEntityModel
		expectData      traceability.TradeModel
	}{
		{
			name:            "1-1. 200: 正常系(入力TraceId該当あり)",
			input:           f.NewPutTradeResponseInput(),
			receiveCFP:      dsResCFP,
			receiveParts:    dsResParts,
			receiveCFPError: nil,
			receiveTrade:    dsResTrade,
			expectData:      dsExpectedRes,
		},
		{
			name:            "1-2. 200: 正常系(入力TraceId該当なし)",
			input:           f.NewPutTradeResponseInput(),
			receiveCFP:      dsResCFP,
			receiveCFPError: gorm.ErrRecordNotFound,
			receiveParts:    dsResParts,
			receiveTrade:    dsResTrade,
			expectData:      dsExpectedRes,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				if test.name == "1-2. 200: 正常系(入力TraceIdなし)" {
					ouranosRepositoryMock.On("GetCFPInformation", f.TraceId).Return(test.receiveCFP, test.receiveCFPError)
					ouranosRepositoryMock.On("GetTrade", mock.Anything).Return(test.receiveTrade, nil)
					ouranosRepositoryMock.On("GetCFPInformation", "087aaa4b-8974-4a0a-9c11-b2e66ed468c5").Return(test.receiveCFP, nil)
					ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, nil)
					ouranosRepositoryMock.On("PutTradeResponse", mock.Anything, mock.Anything).Return(test.receiveTrade, nil)
				} else {
					ouranosRepositoryMock.On("GetCFPInformation", mock.Anything).Return(test.receiveCFP, nil)
					ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, nil)
					ouranosRepositoryMock.On("PutTradeResponse", mock.Anything, mock.Anything).Return(test.receiveTrade, nil)
				}

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				actualRes, err := tradeUsecase.PutTradeResponse(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expectData.UpstreamTraceID, actualRes.UpstreamTraceID, f.AssertMessage)
					assert.Equal(t, test.expectData.UpstreamOperatorID, actualRes.UpstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expectData.DownstreamTraceID, actualRes.DownstreamTraceID, f.AssertMessage)
					assert.Equal(t, test.expectData.DownstreamOperatorID, actualRes.DownstreamOperatorID, f.AssertMessage)
					assert.Equal(t, test.expectData.TradeID, actualRes.TradeID, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecaseDatastore_PutTradeResponse_Abnormal
// Summary: This is abnormal test class which confirm the operation of API #13 PutTradeResponse.
// Target: trade_usecase_datastore_impl.go
// TestPattern:
// [x] 2-1. 400: データ取得エラー(CFP)
// [x] 2-2. 400: データ取得エラー(Parts)
// [x] 2-3. 400: データ更新エラー
func TestProjectUsecaseDatastore_PutTradeResponse_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "tradeResponse"

	dsResPutError := fmt.Errorf("DB AccessError")
	tradeID := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	upstreamOperatorID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	upstreamTraceID := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	CFPID := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResCFP := traceability.CfpEntityModel{
		CfpID:              &CFPID,
		TraceID:            upstreamTraceID,
		GhgEmission:        &f.GhgEmission,
		GhgDeclaredUnit:    f.GhgDeclaredUnit,
		CfpCertificateList: f.CfpCertificateList,
		CfpType:            f.CfpType.ToString(),
		DqrType:            traceability.DqrPreProcessingResponse.ToString(),
		TeR:                &f.TeR,
		GeR:                &f.GeR,
		TiR:                &f.TiR,
		DeletedAt:          gorm.DeletedAt{Time: time.Now()},
		CreatedAt:          time.Now(),
		CreatedUserId:      "seed",
		UpdatedAt:          time.Now(),
		UpdatedUserId:      "seed",
	}
	dsResParts := traceability.PartsModelEntity{
		TraceID:            uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		OperatorID:         uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PartsName:          "B01",
		SupportPartsName:   common.StringPtr("B01001"),
		TerminatedFlag:     true,
		AmountRequired:     common.Float64Ptr(1),
		AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
		DeletedAt:          gorm.DeletedAt{Time: time.Now()},
		CreatedAt:          time.Now(),
		CreatedUserId:      "seed",
		UpdatedAt:          time.Now(),
		UpdatedUserId:      "seed",
	}
	dsResTrade := traceability.TradeEntityModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: uuid.MustParse("e03cc699-7234-31ed-86be-cc18c92208e5"),
		UpstreamOperatorID:   &upstreamOperatorID,
		DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		UpstreamTraceID:      &upstreamTraceID,
		TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
		DeletedAt:            gorm.DeletedAt{Time: time.Now()},
		CreatedAt:            time.Now(),
		CreatedUserID:        "seed",
		UpdatedAt:            time.Now(),
		UpdatedUserID:        "seed",
	}

	tests := []struct {
		name              string
		input             traceability.PutTradeResponseInput
		receiveCFP        traceability.CfpEntityModel
		receiveCFPError   error
		receiveParts      traceability.PartsModelEntity
		receivePartsError error
		receiveTrade      traceability.TradeEntityModel
		receiveTradeError error
		expectData        error
	}{
		{
			name:              "2-1. 400: データ取得エラー(CFP)",
			input:             f.NewPutTradeResponseInput(),
			receiveCFP:        dsResCFP,
			receiveCFPError:   dsResPutError,
			receiveParts:      dsResParts,
			receivePartsError: nil,
			receiveTrade:      dsResTrade,
			receiveTradeError: nil,
			expectData:        dsResPutError,
		},
		{
			name:              "2-2. 400: データ取得エラー(Parts)",
			input:             f.NewPutTradeResponseInput(),
			receiveCFP:        dsResCFP,
			receiveCFPError:   nil,
			receiveParts:      dsResParts,
			receivePartsError: dsResPutError,
			receiveTrade:      dsResTrade,
			receiveTradeError: nil,
			expectData:        dsResPutError,
		},
		{
			name:              "2-3. 400: データ更新エラー",
			input:             f.NewPutTradeResponseInput(),
			receiveCFP:        dsResCFP,
			receiveCFPError:   nil,
			receiveParts:      dsResParts,
			receivePartsError: nil,
			receiveTrade:      dsResTrade,
			receiveTradeError: dsResPutError,
			expectData:        dsResPutError,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetCFPInformation", mock.Anything).Return(test.receiveCFP, test.receiveCFPError)
				ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, test.receivePartsError)
				ouranosRepositoryMock.On("PutTradeResponse", mock.Anything, mock.Anything).Return(test.receiveTrade, test.receiveTradeError)

				tradeUsecase := usecase.NewTradeUsecase(ouranosRepositoryMock)
				_, err := tradeUsecase.PutTradeResponse(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expectData, err, f.AssertMessage)
				}
			},
		)
	}
}
