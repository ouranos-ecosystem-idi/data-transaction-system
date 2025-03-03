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
// [x] 1-2. 200: nil許容項目がnil
// [x] 1-3. 200: 任意項目が未定義
// [x] 1-4. 200: 構成部品なし
// [x] 1-5. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetPartsStructure(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	dsResAll := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
			PartsName:          "B01",
			SupportPartsName:   common.StringPtr("A000001"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr("kilogram"),
			PartsLabelName:     common.StringPtr("PartsA"),
			PartsAddInfo1:      common.StringPtr("Ver2.0"),
			PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
			PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				TraceID:            uuid.MustParse("1c2f37f5-25b9-dea5-346a-7b88035f2553"),
				OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
				PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
				PartsName:          "B01001",
				SupportPartsName:   common.StringPtr("B001"),
				TerminatedFlag:     false,
				AmountRequired:     common.Float64Ptr(2.1),
				AmountRequiredUnit: common.StringPtr("kilogram"),
				PartsLabelName:     common.StringPtr("PartsB"),
				PartsAddInfo1:      common.StringPtr("Ver2.0"),
				PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
				PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
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
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			PartsLabelName:     nil,
			PartsAddInfo1:      nil,
			PartsAddInfo2:      nil,
			PartsAddInfo3:      nil,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				TraceID:            uuid.MustParse("1c2f37f5-25b9-dea5-346a-7b88035f2553"),
				OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
				PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
				PartsName:          "B01001",
				SupportPartsName:   nil,
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: nil,
				PartsLabelName:     nil,
				PartsAddInfo1:      nil,
				PartsAddInfo2:      nil,
				PartsAddInfo3:      nil,
				DeletedAt:          gorm.DeletedAt{Time: time.Now()},
				CreatedAt:          time.Now(),
				CreatedUserId:      "seed",
				UpdatedAt:          time.Now(),
				UpdatedUserId:      "seed",
			},
		},
	}

	dsResRequireOnlyWithUndefined := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
			PartsName:          "B01",
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
				TraceID:            uuid.MustParse("1c2f37f5-25b9-dea5-346a-7b88035f2553"),
				OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
				PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
				PartsName:          "B01001",
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: nil,
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
			OperatorID:         uuid.MustParse("f99c9546-e76e-9f15-35b2-abb9c9b21698"),
			PlantID:            uuid.MustParse("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			PartsLabelName:     common.StringPtr("PartsA"),
			PartsAddInfo1:      common.StringPtr("Ver2.0"),
			PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
			PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
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
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false,
			"partsLabelName": "PartsA",
			"partsAddInfo1": "Ver2.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"childrenPartsModel": [
			{
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsName": "B01001",
				"supportPartsName": "B001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"amountRequired": 2.1,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver2.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		]
	}`

	dsExpectedResRequireOnly := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": null,
			"terminatedFlag": false,
			"amountRequired": null,
			"partsLabelName": null,
			"partsAddInfo1": null,
			"partsAddInfo2": null,
			"partsAddInfo3": null
		},
		"childrenPartsModel": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsName": "B01001",
				"supportPartsName": null,
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": null,
				"terminatedFlag": false,
				"amountRequired": null,
				"partsLabelName": null,
				"partsAddInfo1": null,
				"partsAddInfo2": null,
				"partsAddInfo3": null
			}
		]
	}`

	dsExpectedResNoComponent := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": null,
			"terminatedFlag": false,
			"amountRequired": null,
			"partsLabelName": "PartsA",
			"partsAddInfo1": "Ver2.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"childrenPartsModel": []
	}`

	dsExpectedResNoData := `{
		"parentPartsModel": null,
		"childrenPartsModel": []
	}`

	tests := []struct {
		name    string
		input   traceability.GetPartsStructureInput
		receive traceability.PartsStructureEntity
		expect  string
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewGetPartsStructureInput(),
			receive: dsResAll,
			expect:  dsExpectedResAll,
		},
		{
			name:    "1-2. 200: nil許容項目がnil",
			input:   f.NewGetPartsStructureInput(),
			receive: dsResRequireOnly,
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-3. 200: 任意項目が未定義",
			input:   f.NewGetPartsStructureInput(),
			receive: dsResRequireOnlyWithUndefined,
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-4. 200: 構成部品なし",
			input:   f.NewGetPartsStructureInput(),
			receive: dsResNoComponent,
			expect:  dsExpectedResNoComponent,
		},
		{
			name:    "1-5. 200: 検索結果なし",
			input:   f.NewGetPartsStructureInput(),
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
					assert.Equal(t, expected.ParentPartsModel, actual.ParentPartsModel, f.AssertMessage)
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
		input        traceability.GetPartsStructureInput
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー",
			input:        f.NewGetPartsStructureInput(),
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
// PUT /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答(新規)
// [x] 1-2. 200: 全項目応答(更新)
// [x] 1-3. 200: nil許容項目がnil
// [x] 1-4. 200: 任意項目が未定義
// [x] 1-5. 200: 構成部品なし
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
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": "kilogram",
			"terminatedFlag": false,
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"childrenPartsModel": [
			{
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsName": "B01001",
				"supportPartsName": "B001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "kilogram",
				"terminatedFlag": false,
				"amountRequired": 2.1,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		]
	}`

	dsExpectedResRequireOnly := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": null,
			"terminatedFlag": false,
			"amountRequired": null,
			"partsLabelName": null,
			"partsAddInfo1": null,
			"partsAddInfo2": null,
			"partsAddInfo3": null
		},
		"childrenPartsModel": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsName": "B01001",
				"supportPartsName": null,
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": null,
				"terminatedFlag": false,
				"amountRequired": null,
				"partsLabelName": null,
				"partsAddInfo1": null,
				"partsAddInfo2": null,
				"partsAddInfo3": null
			}
		]
	}`

	dsExpectedResNoComponent := `{
		"parentPartsModel": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsName": "B01",
			"supportPartsName": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": "liter",
			"terminatedFlag": false,
			"amountRequired": null,
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"childrenPartsModel": []
	}`

	tests := []struct {
		name    string
		input   traceability.PutPartsStructureInput
		receive traceability.PartsStructureEntity
		expect  string
	}{
		{
			name:    "1-1. 200: 全項目応答(新規)",
			input:   f.NewPutPartsStructureInput_Insert(),
			receive: f.NewPutPartsStructureEntityModel(),
			expect:  dsExpectedResAll,
		},
		{
			name:    "1-2. 200: 全項目応答(更新)",
			input:   f.NewPutPartsStructureInput(),
			receive: f.NewPutPartsStructureEntityModel(),
			expect:  dsExpectedResAll,
		},
		{
			name:    "1-3. 200: nil許容項目がnil",
			input:   f.NewPutPartsStructureInput_RequiredOnly(),
			receive: f.NewPutPartsStructureEntityModel_RequiredOnly(),
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-4. 200: 任意項目が未定義",
			input:   f.NewPutPartsStructureInput_RequiredOnlyWithUndefined(),
			receive: f.NewPutPartsStructureEntityModel_RequiredOnly(),
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-5. 200: 構成部品なし",
			input:   f.NewPutPartsStructureInput_NoComponent(),
			receive: f.NewPutPartsStructureEntityModel_NoComponent(),
			expect:  dsExpectedResNoComponent,
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

				actual, _, err := partsStructureUsecase.PutPartsStructure(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, expected.ParentPartsModel, actual.ParentPartsModel, f.AssertMessage)
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
		input   traceability.PutPartsStructureInput
		receive error
		expect  error
	}{
		{
			name:    "2-1. 400: データ取得エラー",
			input:   f.NewPutPartsStructureInput(),
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

				_, _, err := partsStructureUsecase.PutPartsStructure(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
