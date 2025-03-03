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
// [x] 1-2. 200: nil許容項目がnil
// [x] 1-3. 200: 任意項目が未定義
// [x] 1-4. 200: 構成部品なし
// [x] 1-5. 200: 検索結果なし
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
				"amountRequiredUnit": "",
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
			"partsLabelName": null,
			"partsAddInfo1": null,
			"partsAddInfo2": null,
			"partsAddInfo3": null
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
		receive string
		expect  string
	}{
		{
			name:    "1-1. 200: 全項目応答",
			input:   f.NewGetPartsStructureInput(),
			receive: f.GetPartsStructure_AllItem(),
			expect:  dsExpectedResAll,
		},
		{
			name:    "1-2. 200: nil許容項目がnil",
			input:   f.NewGetPartsStructureInput(),
			receive: f.GetPartsStructure_RequireItemOnly(),
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-3. 200: 任意項目が未定義",
			input:   f.NewGetPartsStructureInput(),
			receive: f.GetPartsStructure_RequireItemOnlyWithUndefined(),
			expect:  dsExpectedResRequireOnly,
		},
		{
			name:    "1-4. 200: 構成部品なし",
			input:   f.NewGetPartsStructureInput(),
			receive: f.GetPartsStructure_NoComponent(),
			expect:  dsExpectedResNoComponent,
		},
		{
			name:    "1-5. 200: 検索結果なし",
			input:   f.NewGetPartsStructureInput(),
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
		input        traceability.GetPartsStructureInput
		receive      *string
		receiveError *string
		expect       error
	}{
		{
			name:         "2-1. 400: ページングエラー",
			input:        f.NewGetPartsStructureInput(),
			receive:      nil,
			receiveError: common.StringPtr(f.Error_PagingError()),
			expect:       &expectedPagingError,
		},
		{
			name:         "2-2. 400: 想定外の単位",
			input:        f.NewGetPartsStructureInput(),
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
					assert.Equal(t, test.expect, err)
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
			"traceId": "d17833fe-22b7-4a4a-b097-bc3f2150c9a6",
			"partsName": "PartsA-002123",
			"supportPartsName": "modelA",
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountRequiredUnit": "liter",
			"terminatedFlag": false,
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"childrenPartsModel": [
			{
				"traceId": "06c9b015-4225-ba30-1ed3-6faf02cb3fe6",
				"partsName": "PartsA-002123",
				"supportPartsName": "modelA",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountRequiredUnit": "liter",
				"terminatedFlag": false,
				"amountRequired": 1,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		]
	}`

	tests := []struct {
		name    string
		input   traceability.PutPartsStructureInput
		receive string
		expect  string
	}{
		{
			name:    "1-1. 200: 正常系",
			input:   f.NewPutPartsStructureInput(),
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
				traceabilityRepositoryMock.On("PostPartsStructures", mock.Anything, mock.Anything).Return(postPartsStructuresResponse, common.ResponseHeaders{}, nil)

				partsStructureUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepositoryMock)

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
		input   traceability.PutPartsStructureInput
		receive *string
		expect  error
	}{
		{
			name:    "2-1. 400: 存在チェックエラー（事業者）",
			input:   f.NewPutPartsStructureInput(),
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
				traceabilityRepositoryMock.On("PostPartsStructures", mock.Anything, mock.Anything).Return(traceabilityentity.PostPartsStructuresResponse{}, common.ResponseHeaders{}, postPartsStructuresResponse)

				partsStructureUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepositoryMock)

				_, _, err := partsStructureUsecase.PutPartsStructure(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}
