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
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetParts(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	plantId := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResAll := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   common.StringPtr("A000001"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr(amountRequiredUnit.ToString()),
		},
	}

	dsResRequireOnly := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
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
		receive     traceability.PartsModelEntities
		expectData  string
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 全項目応答",
			input:       f.NewGetPartsModel(),
			receive:     dsResAll,
			expectData:  dsExpectedResAll,
			expectAfter: nil,
		},
		{
			name:        "1-2. 200: 必須項目のみ",
			input:       f.NewGetPartsModel(),
			receive:     dsResRequireOnly,
			expectData:  dsExpectedResRequireOnly,
			expectAfter: nil,
		},
		{
			name:        "1-3. 200: 検索結果なし",
			input:       f.NewGetPartsModel(),
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

	plantId := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	dsResDataCountGetError := traceability.PartsModelEntities{
		{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
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
		input        traceability.GetPartsModel
		receive      traceability.PartsModelEntities
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー",
			input:        f.NewGetPartsModel(),
			receive:      nil,
			receiveError: dsResGetError,
			expect:       dsResGetError,
		},
		{
			name:         "2-2. 400: 件数取得エラー",
			input:        f.NewGetPartsModel(),
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
