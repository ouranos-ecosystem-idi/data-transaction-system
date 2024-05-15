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

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 構成部品なし
// [x] 1-4. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetPartsStructure(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	dsResAll := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PartsName:          "B01",
			SupportPartsName:   common.StringPtr("A000001"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr("kilogram"),
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				TraceID:            uuid.MustParse("4d987ed4-f1b0-4bf1-8795-1fdb25300e34"),
				OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
				PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
				PartsName:          "B01001",
				SupportPartsName:   common.StringPtr("B001"),
				TerminatedFlag:     false,
				AmountRequired:     common.Float64Ptr(2.1),
				AmountRequiredUnit: common.StringPtr("kilogram"),
				DeletedAt:          gorm.DeletedAt{Time: time.Now()},
				CreatedAt:          time.Now(),
				CreatedUserId:      "seed",
				UpdatedAt:          time.Now(),
				UpdatedUserId:      "seed",
			},
		},
	}

	dsResRequireOnly := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				TraceID:            uuid.MustParse("4d987ed4-f1b0-4bf1-8795-1fdb25300e34"),
				OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
				PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
				PartsName:          "B01001",
				SupportPartsName:   nil,
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: common.StringPtr(""),
				DeletedAt:          gorm.DeletedAt{Time: time.Now()},
				CreatedAt:          time.Now(),
				CreatedUserId:      "seed",
				UpdatedAt:          time.Now(),
				UpdatedUserId:      "seed",
			},
		},
	}

	dsResNoComponent := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{},
	}

	dsResNoData := traceability.PartsStructureEntity{
		ParentPartsEntity:   nil,
		ChildrenPartsEntity: traceability.PartsModelEntities{},
	}

	dsExpectedResAll := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": "A000001",
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false
		},
		"childrenPartsModel": [
			{
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsName": "B01001",
				"supportPartsName": "B001",
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"amountRequired": 2.1
			}
		]
	}`

	dsExpectedResRequireOnly := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountRequiredUnit": null,
			"terminatedFlag": false,
			"amountRequired": null
		},
		"childrenPartsModel": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsName": "B01001",
				"supportPartsName": null,
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountRequiredUnit": "",
				"terminatedFlag": false,
				"amountRequired": null
			}
		]
	}`

	dsExpectedResNoComponent := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountRequiredUnit": null,
			"terminatedFlag": false,
			"amountRequired": null
		},
		"childrenPartsModel": []
	}`

	dsExpectedResNoData := `{
		"parentPartsModel": null,
		"childrenPartsModel": []
	}`

	tests := []struct {
		name    string
		input   traceability.GetPartsStructureModel
		receive traceability.PartsStructureEntity
		expect  string
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewGetPartsStructureModel(),
			receive: dsResAll,
			expect:  dsExpectedResAll,
		},
		{
			name:    "1-2. 200: 必須項目のみ",
			input:   f.NewGetPartsStructureModel(),
			receive: dsResRequireOnly,
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-3. 200: 構成部品なし",
			input:   f.NewGetPartsStructureModel(),
			receive: dsResNoComponent,
			expect:  dsExpectedResNoComponent,
		},
		{
			name:    "1-4. 200: 検索結果なし",
			input:   f.NewGetPartsStructureModel(),
			receive: dsResNoData,
			expect:  dsExpectedResNoData,
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

				var expected traceability.PartsStructureModel
				err := json.Unmarshal([]byte(test.expect), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetPartsStructure", mock.Anything, mock.Anything).Return(test.receive, nil)

				partsStructureUsecase := usecase.NewPartsStructureDatastoreUsecase(ouranosRepositoryMock)

				actual, err := partsStructureUsecase.GetPartsStructure(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					if test.name == "1-4. 200: 検索結果なし" {
						assert.Equal(t, expected.ParentPartsModel, actual.ParentPartsModel, f.AssertMessage)
					} else {
						assert.Equal(t, expected.ParentPartsModel.OperatorID, actual.ParentPartsModel.OperatorID, f.AssertMessage)
						assert.Equal(t, expected.ParentPartsModel.TraceID, actual.ParentPartsModel.TraceID, f.AssertMessage)
						assert.Equal(t, expected.ParentPartsModel.PartsName, actual.ParentPartsModel.PartsName, f.AssertMessage)
						assert.Equal(t, expected.ParentPartsModel.SupportPartsName, actual.ParentPartsModel.SupportPartsName, f.AssertMessage)
						assert.Equal(t, expected.ParentPartsModel.PlantID, actual.ParentPartsModel.PlantID, f.AssertMessage)
						assert.Equal(t, expected.ParentPartsModel.AmountRequiredUnit, actual.ParentPartsModel.AmountRequiredUnit, f.AssertMessage)
						assert.Equal(t, expected.ParentPartsModel.TerminatedFlag, actual.ParentPartsModel.TerminatedFlag, f.AssertMessage)
					}
					assert.ElementsMatch(t, expected.ChildrenPartsModel, actual.ChildrenPartsModel, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecasedatastore_GetPartsStructure_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	dsResGetError := fmt.Errorf("DB AccessError")

	tests := []struct {
		name         string
		input        traceability.GetPartsStructureModel
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー",
			input:        f.NewGetPartsStructureModel(),
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
				ouranosRepositoryMock.On("GetPartsStructure", mock.Anything, mock.Anything).Return(traceability.PartsStructureEntity{}, test.receiveError)

				partsStructureUsecase := usecase.NewPartsStructureDatastoreUsecase(ouranosRepositoryMock)

				actualRes, err := partsStructureUsecase.GetPartsStructure(c, test.input)
				if assert.Error(t, err) {
					assert.Nil(t, actualRes.ParentPartsModel)
					assert.Nil(t, actualRes.ChildrenPartsModel)
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答
// [x] 1-1. 200: 必須項目のみ
// [x] 1-1. 200: 構成部品なし
// [x] 1-1. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutPartsStructure(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	dsExpectedResAll := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": "A000001",
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false
		},
		"childrenPartsModel": [
			{
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsName": "B01001",
				"supportPartsName": "B001",
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"amountRequired": 2.1
			}
		]
	}`

	tests := []struct {
		name    string
		input   traceability.PartsStructureModel
		receive traceability.PartsStructureEntity
		expect  string
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewPutPartsStructureModel(),
			receive: f.NewPutPartsStructureEntityModel(),
			expect:  dsExpectedResAll,
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

				var expected traceability.PartsStructureModel
				err := json.Unmarshal([]byte(test.expect), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("PutPartsStructure", mock.Anything).Return(test.receive, nil)

				partsStructureUsecase := usecase.NewPartsStructureDatastoreUsecase(ouranosRepositoryMock)

				actual, err := partsStructureUsecase.PutPartsStructure(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, expected.ParentPartsModel.OperatorID, actual.ParentPartsModel.OperatorID, f.AssertMessage)
					assert.Equal(t, expected.ParentPartsModel.TraceID, actual.ParentPartsModel.TraceID, f.AssertMessage)
					assert.Equal(t, expected.ParentPartsModel.PartsName, actual.ParentPartsModel.PartsName, f.AssertMessage)
					assert.Equal(t, expected.ParentPartsModel.SupportPartsName, actual.ParentPartsModel.SupportPartsName, f.AssertMessage)
					assert.Equal(t, expected.ParentPartsModel.PlantID, actual.ParentPartsModel.PlantID, f.AssertMessage)
					assert.Equal(t, expected.ParentPartsModel.AmountRequiredUnit, actual.ParentPartsModel.AmountRequiredUnit, f.AssertMessage)
					assert.Equal(t, expected.ParentPartsModel.TerminatedFlag, actual.ParentPartsModel.TerminatedFlag, f.AssertMessage)
					assert.ElementsMatch(t, expected.ChildrenPartsModel, actual.ChildrenPartsModel, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutPartsStructure_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	dsResPutError := fmt.Errorf("DB AccessError")

	tests := []struct {
		name    string
		input   traceability.PartsStructureModel
		receive error
		expect  error
	}{
		{
			name:    "2-1. 400: データ取得エラー",
			input:   f.NewPutPartsStructureModel(),
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
				c.Set("operatorID", f.OperatorId)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("PutPartsStructure", mock.Anything).Return(traceability.PartsStructureEntity{}, test.receive)

				partsStructureUsecase := usecase.NewPartsStructureDatastoreUsecase(ouranosRepositoryMock)

				_, err := partsStructureUsecase.PutPartsStructure(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
