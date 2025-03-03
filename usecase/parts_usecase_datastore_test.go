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
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase"

	"github.com/google/uuid"
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
func TestProjectUsecaseDatastore_GetParts(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	plantId := uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResAll := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   common.StringPtr("A000001"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
			PartsLabelName:     common.StringPtr("PartsA"),
			PartsAddInfo1:      common.StringPtr("Ver2.0"),
			PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
			PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
		},
	}

	dsResRequireOnly := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			PartsLabelName:     nil,
			PartsAddInfo1:      nil,
			PartsAddInfo2:      nil,
			PartsAddInfo3:      nil,
		},
	}

	dsResRequireOnlyWithUndefined := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
		},
	}

	dsResNoData := traceability.PartsModelEntities{}

	dsExpectedResAll := `[
		{
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": "A000001",
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false,
			"partsLabelName": "PartsA",
			"partsAddInfo1": "Ver2.0",
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
		receive     traceability.PartsModelEntities
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetPartsInput(),
			receive:     dsResAll,
			expectData:  dsExpectedResAll,
			expectAfter: nil,
		},
		{
			name:        "1-2. 200: nil許容項目がnil",
			input:       f.NewGetPartsInput(),
			receive:     dsResRequireOnly,
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 任意項目が未定義",
			input:       f.NewGetPartsInput(),
			receive:     dsResRequireOnlyWithUndefined,
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-4. 200: 検索結果なし",
			input:       f.NewGetPartsInput(),
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

				var expected []traceability.PartsModel
				err := json.Unmarshal([]byte(test.expectData), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("ListParts", mock.Anything).Return(test.receive, nil)
				ouranosRepositoryMock.On("CountPartsList", mock.Anything).Return(1, nil)

				partsUsecase := usecase.NewPartsUsecase(ouranosRepositoryMock)

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
// [x] 2-1. 400: データ取得エラー
// [x] 2-2. 400: 件数取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecasedatastore_GetParts_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	dsResGetError := fmt.Errorf("DB AccessError")

	plantId := uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResDataCountGetError := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   common.StringPtr("A000001"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
		},
	}

	dsResCountGetError := fmt.Errorf("DB AccessError")

	tests := []struct {
		name         string
		input        traceability.GetPartsInput
		receive      traceability.PartsModelEntities
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー",
			input:        f.NewGetPartsInput(),
			receive:      nil,
			receiveError: dsResGetError,
			expect:       dsResGetError,
		},
		{
			name:         "2-2. 400: 件数取得エラー",
			input:        f.NewGetPartsInput(),
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
					ouranosRepositoryMock.On("ListParts", mock.Anything).Return(traceability.PartsModelEntities{}, test.receiveError)
				} else if test.name == "2-2. 400: 件数取得エラー" {
					ouranosRepositoryMock.On("ListParts", mock.Anything).Return(test.receive, nil)
					ouranosRepositoryMock.On("CountPartsList", mock.Anything).Return(0, test.receiveError)
				}

				partsUsecase := usecase.NewPartsUsecase(ouranosRepositoryMock)

				actualRes, actualAfter, err := partsUsecase.GetPartsList(c, test.input)
				if assert.Error(t, err) {
					assert.Nil(t, actualRes)
					assert.Nil(t, actualAfter)
					assert.Equal(t, test.expect.Error(), err.Error())
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
func TestProjectUsecaseDatastore_DeleteParts(tt *testing.T) {

	var method = "DELETE"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name              string
		input             traceability.DeletePartsInput
		receiveParts      traceability.PartsModelEntity
		receiveParents    traceability.PartsStructureEntityModels
		receiveChildren   traceability.PartsStructureEntityModels
		receiveUpstream   traceability.TradeEntityModels
		receiveDownstream traceability.TradeEntityModels
	}{
		{
			name:              "1-1. 200: 全項目応答",
			input:             f.NewDeletePartsInput(f.TraceID),
			receiveParts:      f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:    traceability.PartsStructureEntityModels{},
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
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

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, nil)
				ouranosRepositoryMock.On("ListParentPartsStructureByTraceId", mock.Anything).Return(test.receiveParents, nil)
				ouranosRepositoryMock.On("ListChildPartsStructureByTraceId", mock.Anything).Return(test.receiveChildren, nil)
				ouranosRepositoryMock.On("ListTradeByDownstreamTraceID", mock.Anything).Return(test.receiveDownstream, nil)
				ouranosRepositoryMock.On("ListTradeByUpstreamTraceID", mock.Anything).Return(test.receiveUpstream, nil)
				ouranosRepositoryMock.On("DeletePartsWithCFP", mock.Anything).Return(nil)

				partsUsecase := usecase.NewPartsUsecase(ouranosRepositoryMock)

				_, err := partsUsecase.DeleteParts(c, test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: 部品取得エラー
// [x] 2-2. 400: 事業者不一致エラー
// [x] 2-3. 400: 親部品取得エラー
// [x] 2-4. 400: 親部品未削除
// [x] 2-5. 400: 子部品取得エラー
// [x] 2-6. 400: 子部品未削除
// [x] 2-7. 400: 取引関係依頼取得エラー
// [x] 2-8. 400: 取引関係依頼済
// [x] 2-9. 400: 取引関係回答取得エラー
// [x] 2-10. 400: 取引関係回答済
// [x] 2-11. 400: 部品削除エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_DeleteParts_Abnormal(tt *testing.T) {

	var method = "DELETE"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name                 string
		input                traceability.DeletePartsInput
		receiveParts         traceability.PartsModelEntity
		receivePartsErr      error
		receiveParents       traceability.PartsStructureEntityModels
		receiveParentsErr    error
		receiveChildren      traceability.PartsStructureEntityModels
		receiveChildrenErr   error
		receiveDownstream    traceability.TradeEntityModels
		receiveDownstreamErr error
		receiveUpstream      traceability.TradeEntityModels
		receiveUpstreamErr   error
		receiveDeleteErr     error
		expect               error
	}{
		{
			name:              "2-1. 400: 部品取得エラー",
			input:             f.NewDeletePartsInput(f.TraceID),
			receiveParts:      f.GetPartsModelEntity(f.TraceID, true),
			receivePartsErr:   fmt.Errorf("DB AccessError"),
			receiveParents:    traceability.PartsStructureEntityModels{},
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			expect:            f.NewTraceabilityError("MSGAECP0013", "指定された部品は存在しません。", []uuid.UUID{}),
		},
		{
			name:              "2-2. 400: 事業者不一致エラー",
			input:             f.NewDeletePartsInput(f.TraceID),
			receiveParts:      f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:    traceability.PartsStructureEntityModels{},
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			expect:            f.NewTraceabilityError("MSGAECP0025", "認証情報と事業者識別子が一致しません。", []uuid.UUID{}),
		},
		{
			name:              "2-3. 400: 親部品取得エラー",
			input:             f.NewDeletePartsInput(f.TraceID),
			receiveParts:      f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:    traceability.PartsStructureEntityModels{},
			receiveParentsErr: fmt.Errorf("DB AccessError"),
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			expect:            fmt.Errorf("DB AccessError"),
		},
		{
			name:         "2-4. 400: 親部品未削除",
			input:        f.NewDeletePartsInput(f.TraceID),
			receiveParts: f.GetPartsModelEntity(f.TraceID, true),
			receiveParents: traceability.PartsStructureEntityModels{
				f.GetPartsStructureEntityModel(),
			},
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			expect:            f.NewTraceabilityError("MSGAECP0014", "指定された部品は部品構成が存在するため削除できません。", []uuid.UUID{uuid.MustParse(f.TraceID)}),
		},
		{
			name:               "2-5. 400: 子部品取得エラー",
			input:              f.NewDeletePartsInput(f.TraceID),
			receiveParts:       f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:     traceability.PartsStructureEntityModels{},
			receiveChildren:    traceability.PartsStructureEntityModels{},
			receiveChildrenErr: fmt.Errorf("DB AccessError"),
			receiveDownstream:  traceability.TradeEntityModels{},
			receiveUpstream:    traceability.TradeEntityModels{},
			expect:             fmt.Errorf("DB AccessError"),
		},
		{
			name:           "2-6. 400: 子部品未削除",
			input:          f.NewDeletePartsInput(f.TraceID),
			receiveParts:   f.GetPartsModelEntity(f.TraceID, true),
			receiveParents: traceability.PartsStructureEntityModels{},
			receiveChildren: traceability.PartsStructureEntityModels{
				f.GetPartsStructureEntityModel(),
			},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			expect:            f.NewTraceabilityError("MSGAECP0014", "指定された部品は部品構成が存在するため削除できません。", []uuid.UUID{uuid.MustParse(f.TraceIDChild)}),
		},
		{
			name:                 "2-7. 400: 取引関係依頼取得エラー",
			input:                f.NewDeletePartsInput(f.TraceID),
			receiveParts:         f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:       traceability.PartsStructureEntityModels{},
			receiveChildren:      traceability.PartsStructureEntityModels{},
			receiveDownstream:    traceability.TradeEntityModels{},
			receiveDownstreamErr: fmt.Errorf("DB AccessError"),
			receiveUpstream:      traceability.TradeEntityModels{},
			expect:               fmt.Errorf("DB AccessError"),
		},
		{
			name:            "2-8. 400: 取引関係依頼済",
			input:           f.NewDeletePartsInput(f.TraceID),
			receiveParts:    f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:  traceability.PartsStructureEntityModels{},
			receiveChildren: traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{
				f.GetTradeEntityModel(),
			},
			receiveUpstream: traceability.TradeEntityModels{},
			expect:          f.NewTraceabilityError("MSGAECP0015", "指定された部品は依頼済みのため削除できません。", []uuid.UUID{uuid.MustParse(f.TradeID)}),
		},
		{
			name:               "2-9. 400: 取引関係回答取得エラー",
			input:              f.NewDeletePartsInput(f.TraceID),
			receiveParts:       f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:     traceability.PartsStructureEntityModels{},
			receiveChildren:    traceability.PartsStructureEntityModels{},
			receiveDownstream:  traceability.TradeEntityModels{},
			receiveUpstream:    traceability.TradeEntityModels{},
			receiveUpstreamErr: fmt.Errorf("DB AccessError"),
			expect:             fmt.Errorf("DB AccessError"),
		},
		{
			name:              "2-10. 400: 取引関係回答済",
			input:             f.NewDeletePartsInput(f.TraceID),
			receiveParts:      f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:    traceability.PartsStructureEntityModels{},
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream: traceability.TradeEntityModels{
				f.GetTradeEntityModel(),
			},
			expect: f.NewTraceabilityError("MSGAECP0016", "指定された部品は受領済みの依頼に紐づいているため削除できません。", []uuid.UUID{uuid.MustParse(f.TradeID)}),
		},
		{
			name:              "2-11. 400: 部品削除エラー",
			input:             f.NewDeletePartsInput(f.TraceID),
			receiveParts:      f.GetPartsModelEntity(f.TraceID, true),
			receiveParents:    traceability.PartsStructureEntityModels{},
			receiveChildren:   traceability.PartsStructureEntityModels{},
			receiveDownstream: traceability.TradeEntityModels{},
			receiveUpstream:   traceability.TradeEntityModels{},
			receiveDeleteErr:  fmt.Errorf("DB AccessError"),
			expect:            fmt.Errorf("DB AccessError"),
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
				if test.name == "2-2. 400: 事業者不一致エラー" {
					c.Set("operatorID", f.OperatorID2)
				} else {
					c.Set("operatorID", f.OperatorID)
				}
				c.Set("traceId", f.TraceID)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(test.receiveParts, test.receivePartsErr)
				ouranosRepositoryMock.On("ListParentPartsStructureByTraceId", mock.Anything).Return(test.receiveParents, test.receiveParentsErr)
				ouranosRepositoryMock.On("ListChildPartsStructureByTraceId", mock.Anything).Return(test.receiveChildren, test.receiveChildrenErr)
				ouranosRepositoryMock.On("ListTradeByDownstreamTraceID", mock.Anything).Return(test.receiveDownstream, test.receiveDownstreamErr)
				ouranosRepositoryMock.On("ListTradeByUpstreamTraceID", mock.Anything).Return(test.receiveUpstream, test.receiveUpstreamErr)
				ouranosRepositoryMock.On("DeletePartsWithCFP", mock.Anything).Return(test.receiveDeleteErr)

				partsUsecase := usecase.NewPartsUsecase(ouranosRepositoryMock)

				_, err := partsUsecase.DeleteParts(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
