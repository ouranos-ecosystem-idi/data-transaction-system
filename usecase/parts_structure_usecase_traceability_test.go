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
// Get /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答
// [x] 1-2. 200: 必須項目のみ
// [x] 1-3. 200: 構成部品なし
// [x] 1-4. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetPartsStructure(tt *testing.T) {

	var method = "GET"
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
		receive string
		expect  string
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewGetPartsStructureModel(),
			receive: f.GetPartsStructure_AllItem(),
			expect:  dsExpectedResAll,
		},
		{
			name:    "1-2. 200: 必須項目のみ",
			input:   f.NewGetPartsStructureModel(),
			receive: f.GetPartsStructure_RequireItemOnly(),
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-3. 200: 構成部品なし",
			input:   f.NewGetPartsStructureModel(),
			receive: f.GetPartsStructure_NoComponent(),
			expect:  dsExpectedResNoComponent,
		},
		{
			name:    "1-4. 200: 検索結果なし",
			input:   f.NewGetPartsStructureModel(),
			receive: f.GetPartsStructure_NoData(),
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

				getPartsStructuresResponse := traceabilityentity.GetPartsStructuresResponse{}

				if err := json.Unmarshal([]byte(test.receive), &getPartsStructuresResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				var expected traceability.PartsStructureModel
				err := json.Unmarshal([]byte(test.expect), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("GetPartsStructures", mock.Anything, mock.Anything).Return(getPartsStructuresResponse, nil)

				partsStructureUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepositoryMock)

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
// [x] 2-1. 400: ページングエラー
// [x] 2-2. 500: 想定外の単位
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetPartsStructure_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	expectedInvalidTypeError := fmt.Errorf("unexpected AmountRequiredUnit. get value: kg")

	tests := []struct {
		name         string
		input        traceability.GetPartsStructureModel
		receive      *string
		receiveError *string
		expect       error
	}{
		{
			name:         "2-1. 400: ページングエラー",
			input:        f.NewGetPartsStructureModel(),
			receive:      nil,
			receiveError: common.StringPtr(f.Error_PagingError()),
			expect:       &expectedPagingError,
		},
		{
			name:         "2-2. 400: 想定外の単位",
			input:        f.NewGetPartsStructureModel(),
			receive:      common.StringPtr(f.GetPartsStructure_InvalidTypeError()),
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
					getPartsStructuresResponse := traceabilityentity.GetPartsStructuresResponse{}
					if err := json.Unmarshal([]byte(*test.receive), &getPartsStructuresResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetPartsStructures", mock.Anything, mock.Anything).Return(getPartsStructuresResponse, nil)
				} else if test.receiveError != nil {
					getPartsStructuresResponse := common.ToTracebilityAPIError(*test.receiveError).ToCustomError(400)
					traceabilityRepositoryMock.On("GetPartsStructures", mock.Anything, mock.Anything).Return(traceabilityentity.GetPartsStructuresResponse{}, getPartsStructuresResponse)
				}

				partsStructureUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepositoryMock)

				actualRes, err := partsStructureUsecase.GetPartsStructure(c, test.input)
				if assert.Error(t, err) {
					assert.Nil(t, actualRes.ParentPartsModel)
					assert.Nil(t, actualRes.ChildrenPartsModel)
					if test.name == "2-1. 400: ページングエラー" {
						assert.Equal(t, test.expect.(*common.CustomError).Code, err.(*common.CustomError).Code)
						assert.Equal(t, test.expect.(*common.CustomError).Message, err.(*common.CustomError).Message)
						assert.Equal(t, test.expect.(*common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
						assert.Equal(t, test.expect.(*common.CustomError).Source, err.(*common.CustomError).Source)
					} else if test.name == "2-2. 400: 想定外の単位" {
						assert.Equal(t, test.expect.Error(), err.Error())
					}
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常系
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutPartsStructure(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	dsExpectedRes := `{
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
		receive string
		expect  string
	}{
		{
			name:    "1-1. 200: 正常系",
			input:   f.NewPutPartsStructureModel(),
			receive: f.PutPartsStructure(),
			expect:  dsExpectedRes,
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

				postPartsStructuresResponse := traceabilityentity.PostPartsStructuresResponse{}

				if err := json.Unmarshal([]byte(test.receive), &postPartsStructuresResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				var expected traceability.PartsStructureModel
				err := json.Unmarshal([]byte(test.expect), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("PostPartsStructures", mock.Anything, mock.Anything).Return(postPartsStructuresResponse, nil)

				partsStructureUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepositoryMock)

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
// [x] 2-1. 400: 存在チェックエラー（事業者）
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutPartsStructure_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	expectedExistError := common.CustomError{
		Code:          400,
		Message:       "存在しない事業者識別子が使用されています。",
		MessageDetail: common.StringPtr("MSGAECP0004"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name    string
		input   traceability.PartsStructureModel
		receive *string
		expect  error
	}{
		{
			name:    "2-1. 400: 存在チェックエラー（事業者）",
			input:   f.NewPutPartsStructureModel(),
			receive: common.StringPtr(f.Error_OperatorIdNotFound()),
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
				c.Set("operatorID", f.OperatorId)

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				postPartsStructuresResponse := common.ToTracebilityAPIError(*test.receive).ToCustomError(400)
				traceabilityRepositoryMock.On("PostPartsStructures", mock.Anything, mock.Anything).Return(traceabilityentity.PostPartsStructuresResponse{}, postPartsStructuresResponse)

				partsStructureUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepositoryMock)

				_, err := partsStructureUsecase.PutPartsStructure(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.(*common.CustomError).Code, err.(*common.CustomError).Code)
					assert.Equal(t, test.expect.(*common.CustomError).Message, err.(*common.CustomError).Message)
					assert.Equal(t, test.expect.(*common.CustomError).MessageDetail, err.(*common.CustomError).MessageDetail)
					assert.Equal(t, test.expect.(*common.CustomError).Source, err.(*common.CustomError).Source)
				}
			},
		)
	}
}
