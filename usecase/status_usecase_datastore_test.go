package usecase_test

import (
	"fmt"
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

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetStatus(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	getStatusInput := f.NewGetStatusInput()

	cc, _ := time.Parse("2006-01-02T15:04:05Z", f.CompletedCountModifiedAt)
	tc, _ := time.Parse("2006-01-02T15:04:05Z", f.TradesCountModifiedAt)
	dsResAll := []traceability.StatusEntityModel{
		{
			StatusID:                 uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:                  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			CfpResponseStatus:        traceability.CfpResponseStatusCancel.ToString(),
			TradeTreeStatus:          traceability.TradeTreeStatusUnterminated.ToString(),
			Message:                  common.StringPtr("A01のCFP値を回答ください"),
			ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
			RequestType:              f.RequestType.ToString(),
			ResponseDueDate:          f.ResponseDueDate,
			CompletedCount:           &f.CompletedCount,
			CompletedCountModifiedAt: &cc,
			TradesCount:              &f.TradesCount,
			TradesCountModifiedAt:    &tc,
			DeletedAt:                gorm.DeletedAt{Time: time.Now()},
			CreatedAt:                time.Now(),
			CreatedUserId:            "seed",
			UpdatedAt:                time.Now(),
			UpdatedUserId:            "seed",
		},
	}

	dsResRequireOnly := []traceability.StatusEntityModel{
		{
			StatusID:                 uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:                  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			CfpResponseStatus:        traceability.CfpResponseStatusCancel.ToString(),
			TradeTreeStatus:          traceability.TradeTreeStatusUnterminated.ToString(),
			Message:                  nil,
			ReplyMessage:             nil,
			RequestType:              f.RequestType.ToString(),
			ResponseDueDate:          f.ResponseDueDate,
			CompletedCount:           &f.CompletedCount,
			CompletedCountModifiedAt: &cc,
			TradesCount:              &f.TradesCount,
			TradesCountModifiedAt:    &tc,
			DeletedAt:                gorm.DeletedAt{Time: time.Now()},
			CreatedAt:                time.Now(),
			CreatedUserId:            "seed",
			UpdatedAt:                time.Now(),
			UpdatedUserId:            "seed",
		},
	}

	dsResNoData := []traceability.StatusEntityModel{}

	cfpResponseStatusCancel := traceability.CfpResponseStatusCancel
	tradeTreeStatusUnterminated := traceability.TradeTreeStatusUnterminated

	dsExpectedResAll := []traceability.StatusModel{
		{
			StatusID: uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus:        &cfpResponseStatusCancel,
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

	dsExpectedResRequireOnly := []traceability.StatusModel{
		{
			StatusID: uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus:        &cfpResponseStatusCancel,
				TradeTreeStatus:          &tradeTreeStatusUnterminated,
				CompletedCount:           &f.CompletedCount,
				CompletedCountModifiedAt: &f.CompletedCountModifiedAt,
				TradesCount:              &f.TradesCount,
				TradesCountModifiedAt:    &f.TradesCountModifiedAt,
			},
			Message:         nil,
			ReplyMessage:    nil,
			RequestType:     f.RequestType.ToString(),
			ResponseDueDate: &f.ResponseDueDate,
		},
	}

	expectedResNoData := []traceability.StatusModel{}

	tests := []struct {
		name        string
		input       traceability.GetStatusInput
		receive     traceability.StatusEntityModels
		expectData  traceability.StatusModels
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       getStatusInput,
			receive:     dsResAll,
			expectData:  dsExpectedResAll,
			expectAfter: nil,
		},
		{
			name:        "1-2. 200: 必須項目のみ",
			input:       getStatusInput,
			receive:     dsResRequireOnly,
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 検索結果なし",
			input:       getStatusInput,
			receive:     dsResNoData,
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(test.receive, nil)
				ouranosRepositoryMock.On("CountStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(1, nil)

				usecase := usecase.NewStatusUsecase(ouranosRepositoryMock)

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
// [x] 2-1. 400: データ取得エラー
// [x] 2-2. 400: 件数取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecasedatastore_GetStatus_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	getStatusInput := f.NewGetStatusInput()

	dsResGetError := fmt.Errorf("DB AccessError")

	dsResDataCountGetError := traceability.StatusEntityModels{
		{
			StatusID:          uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
			TradeID:           uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
			CfpResponseStatus: traceability.CfpResponseStatusCancel.ToString(),
			TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
			Message:           common.StringPtr("A01のCFP値を回答ください"),
			ReplyMessage:      common.StringPtr("A01のCFP値を回答しました"),
			RequestType:       f.RequestType.ToString(),
			DeletedAt:         gorm.DeletedAt{Time: time.Now()},
			CreatedAt:         time.Now(),
			CreatedUserId:     "seed",
			UpdatedAt:         time.Now(),
			UpdatedUserId:     "seed",
		},
	}

	tests := []struct {
		name         string
		input        traceability.GetStatusInput
		receive      traceability.StatusEntityModels
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー",
			input:        getStatusInput,
			receive:      nil,
			receiveError: dsResGetError,
			expect:       dsResGetError,
		},
		{
			name:         "2-2. 400: 件数取得エラー",
			input:        getStatusInput,
			receive:      dsResDataCountGetError,
			receiveError: dsResGetError,
			expect:       dsResGetError,
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
					ouranosRepositoryMock.On("GetStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(traceability.StatusEntityModels{}, test.receiveError)
				} else if test.name == "2-2. 400: 件数取得エラー" {
					ouranosRepositoryMock.On("GetStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(test.receive, nil)
					ouranosRepositoryMock.On("CountStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(0, test.receiveError)
				}

				usecase := usecase.NewStatusUsecase(ouranosRepositoryMock)

				_, actualAfter, err := usecase.GetStatus(c, test.input)
				if assert.Error(t, err) {
					assert.Nil(t, actualAfter)
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutStatusCancel(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name  string
		input traceability.PutStatusInput
	}{
		{
			name:  "1-1. 200: 正常終了",
			input: f.NewPutStatusInput(),
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
				ouranosRepositoryMock.On("PutStatusCancel", mock.Anything, mock.Anything).Return(nil)

				usecase := usecase.NewStatusUsecase(ouranosRepositoryMock)

				_, err := usecase.PutStatusCancel(c, test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutStatusCancel_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive error
	}{
		{
			name:    "1-1. 400: データ取得エラー",
			input:   f.NewPutStatusInput(),
			receive: fmt.Errorf("DB AccessError"),
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
				ouranosRepositoryMock.On("PutStatusCancel", mock.Anything, mock.Anything).Return(test.receive)

				usecase := usecase.NewStatusUsecase(ouranosRepositoryMock)

				_, err := usecase.PutStatusCancel(c, test.input)
				assert.Error(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutStatusReject(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	dsRes := traceability.StatusEntityModel{
		StatusID:          uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
		TradeID:           uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
		CfpResponseStatus: traceability.CfpResponseStatusCancel.ToString(),
		TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
		Message:           common.StringPtr("A01のCFP値を回答ください"),
		ReplyMessage:      common.StringPtr("A01のCFP値を回答しました"),
		RequestType:       f.RequestType.ToString(),
		DeletedAt:         gorm.DeletedAt{Time: time.Now()},
		CreatedAt:         time.Now(),
		CreatedUserId:     "seed",
		UpdatedAt:         time.Now(),
		UpdatedUserId:     "seed",
	}

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive traceability.StatusEntityModel
	}{
		{
			name:    "1-1. 200: 正常終了",
			input:   f.NewPutStatusInput(),
			receive: dsRes,
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
				ouranosRepositoryMock.On("PutStatusReject", mock.Anything, mock.Anything, mock.Anything).Return(test.receive, nil)

				usecase := usecase.NewStatusUsecase(ouranosRepositoryMock)

				_, err := usecase.PutStatusReject(c, test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutStatusReject_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "status"

	tests := []struct {
		name    string
		input   traceability.PutStatusInput
		receive error
	}{
		{
			name:    "1-1. 400: データ取得エラー",
			input:   f.NewPutStatusInput(),
			receive: fmt.Errorf("DB AccessError"),
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
				ouranosRepositoryMock.On("PutStatusReject", mock.Anything, mock.Anything, mock.Anything).Return(traceability.StatusEntityModel{}, test.receive)

				usecase := usecase.NewStatusUsecase(ouranosRepositoryMock)

				_, err := usecase.PutStatusReject(c, test.input)
				assert.Error(t, err)
			},
		)
	}
}
