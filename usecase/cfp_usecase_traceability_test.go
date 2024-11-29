package usecase_test

import (
	"encoding/json"
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
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答(依頼元)
// [x] 1-2. 200: 全項目応答(依頼先)
// [x] 1-3. 200: 必須項目のみ(依頼元)
// [x] 1-4. 200: 必須項目のみ(依頼元)(キーなし)
// [x] 1-5. 200: 必須項目のみ(依頼先)
// [x] 1-6. 200: 必須項目のみ(依頼先)(キーなし)
// [x] 1-7. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetCfp(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	getCfpInputRequesting := f.NewGetCfpInput()
	getCfpInputRequesting.TraceIDs = []uuid.UUID{uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")}
	getCfpInputRequested := f.NewGetCfpInput()
	getCfpInputRequested.TraceIDs = []uuid.UUID{uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5")}

	dsExpectedResAllOwn := `[
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 0.5,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preProduction",
			"dqrType": "preProcessing",
			"dqrValue": {
				"TeR": 3.1,
				"GeR": 3.2,
				"TiR": 3.3
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 0.6,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainProduction",
			"dqrType": "mainProcessing",
			"dqrValue": {
				"TeR": 3.4,
				"GeR": 3.5,
				"TiR": 3.6
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.1,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preComponent",
			"dqrType": "preProcessing",
			"dqrValue": {
				"TeR": 3.1,
				"GeR": 3.2,
				"TiR": 3.3
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.2,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainComponent",
			"dqrType": "mainProcessing",
			"dqrValue": {
				"TeR": 3.4,
				"GeR": 3.5,
				"TiR": 3.6
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.5,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preProductionTotal",
			"dqrType": "preProcessingTotal",
			"dqrValue": {
				"TeR": 4.1,
				"GeR": 4.2,
				"TiR": 4.3
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 2.1,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preComponentTotal",
			"dqrType": "preProcessingTotal",
			"dqrValue": {
				"TeR": 4.1,
				"GeR": 4.2,
				"TiR": 4.3
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.6,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainProductionTotal",
			"dqrType": "mainProcessingTotal",
			"dqrValue": {
				"TeR": 4.4,
				"GeR": 4.5,
				"TiR": 4.6
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 2.2,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainComponentTotal",
			"dqrType": "mainProcessingTotal",
			"dqrValue": {
				"TeR": 4.4,
				"GeR": 4.5,
				"TiR": 4.6
			}
		}
	]`

	dsExpectedResRequireItemOnlyOwn := `[
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 0.5,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preProduction",
			"dqrType": "preProcessing",
			"dqrValue": {
				"TeR": 3.1,
				"GeR": 3.2,
				"TiR": 3.3
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 0.6,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainProduction",
			"dqrType": "mainProcessing",
			"dqrValue": {
				"TeR": 3.4,
				"GeR": 3.5,
				"TiR": 3.6
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.1,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preComponent",
			"dqrType": "preProcessing",
			"dqrValue": {
				"TeR": 3.1,
				"GeR": 3.2,
				"TiR": 3.3
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.2,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainComponent",
			"dqrType": "mainProcessing",
			"dqrValue": {
				"TeR": 3.4,
				"GeR": 3.5,
				"TiR": 3.6
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.5,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preProductionTotal",
			"dqrType": "preProcessingTotal",
			"dqrValue": {
				"TeR": null,
				"GeR": null,
				"TiR": null
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 2.1,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preComponentTotal",
			"dqrType": "preProcessingTotal",
			"dqrValue": {
				"TeR": null,
				"GeR": null,
				"TiR": null
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 1.6,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainProductionTotal",
			"dqrType": "mainProcessingTotal",
			"dqrValue": {
				"TeR": null,
				"GeR": null,
				"TiR": null
			}
		},
		{
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
			"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
			"ghgEmission": 2.2,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainComponentTotal",
			"dqrType": "mainProcessingTotal",
			"dqrValue": {
				"TeR": null,
				"GeR": null,
				"TiR": null
			}
		}
	]`

	dsExpectedResAllSupplier := `[
		{
			"cfpId": null,
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"ghgEmission": 0.1,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preProductionResponse",
			"dqrType": "preProcessingResponse",
			"dqrValue": {
				"TeR": 3.1,
				"GeR": 3.2,
				"TiR": 3.3
			}
		},
		{
			"cfpId": null,
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"ghgEmission": 0.4,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainProductionResponse",
			"dqrType": "mainProcessingResponse",
			"dqrValue": {
				"TeR": 3.4,
				"GeR": 3.5,
				"TiR": 3.6
			}
		}
	]`

	dsExpectedResRequireItemOnlySupplier := `[
		{
			"cfpId": null,
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"ghgEmission": null,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "preProductionResponse",
			"dqrType": "preProcessingResponse",
			"dqrValue": {
				"TeR": null,
				"GeR": null,
				"TiR": null
			}
		},
		{
			"cfpId": null,
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"ghgEmission": null,
			"ghgDeclaredUnit": "kgCO2e/kilogram",
			"cfpType": "mainProductionResponse",
			"dqrType": "mainProcessingResponse",
			"dqrValue": {
				"TeR": null,
				"GeR": null,
				"TiR": null
			}
		}
	]`

	dsExpectedResNoData := `[]`

	tests := []struct {
		name         string
		input        traceability.GetCfpInput
		receiveCfp   string
		receiveTrade string
		expect       string
	}{
		{
			name:         "1-1. 200: 全項目応答(依頼元)",
			input:        getCfpInputRequesting,
			receiveCfp:   f.GetCfp_AllItem(),
			receiveTrade: f.GetTradeRequests_NoData(),
			expect:       dsExpectedResAllOwn,
		},
		{
			name:         "1-2. 200: 全項目応答(依頼先)",
			input:        getCfpInputRequested,
			receiveCfp:   f.GetCfp_NoData(),
			receiveTrade: f.GetTradeRequests_AllItem_NoNext(),
			expect:       dsExpectedResAllSupplier,
		},
		{
			name:         "1-3. 200: 必須項目のみ(依頼元)",
			input:        getCfpInputRequesting,
			receiveCfp:   f.GetCfp_RequireItemOnly(),
			receiveTrade: f.GetTradeRequests_NoData(),
			expect:       dsExpectedResRequireItemOnlyOwn,
		},
		{
			name:         "1-4. 200: 必須項目のみ(依頼元)(キーなし)",
			input:        getCfpInputRequesting,
			receiveCfp:   f.GetCfp_RequireItemOnlyWithUndefined(),
			receiveTrade: f.GetTradeRequests_NoData(),
			expect:       dsExpectedResRequireItemOnlyOwn,
		},
		{
			name:         "1-5. 200: 必須項目のみ(依頼先)",
			input:        getCfpInputRequested,
			receiveCfp:   f.GetCfp_NoData(),
			receiveTrade: f.GetTradeRequests_RequireItemOnlyAnswered(),
			expect:       dsExpectedResRequireItemOnlySupplier,
		},
		{
			name:         "1-6. 200: 必須項目のみ(依頼先)(キーなし)",
			input:        getCfpInputRequested,
			receiveCfp:   f.GetCfp_NoData(),
			receiveTrade: f.GetTradeRequests_RequireItemOnlyAnsweredWithUndefined(),
			expect:       dsExpectedResRequireItemOnlySupplier,
		},
		{
			name:         "1-7. 200: 検索結果なし",
			input:        getCfpInputRequested,
			receiveCfp:   f.GetCfp_NoData(),
			receiveTrade: f.GetTradeRequests_NoData(),
			expect:       dsExpectedResNoData,
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

				getCfpResponse := traceabilityentity.GetCfpResponses{}

				if err := json.Unmarshal([]byte(test.receiveCfp), &getCfpResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}
				getTradeRequestResponse := traceabilityentity.GetTradeRequestsResponse{}

				if err := json.Unmarshal([]byte(test.receiveTrade), &getTradeRequestResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}
				var expected []traceability.CfpModel
				err := json.Unmarshal([]byte(test.expect), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("GetCfp", mock.Anything, mock.Anything).Return(getCfpResponse, nil)
				traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(getTradeRequestResponse, nil)

				usecase := usecase.NewCfpTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, err := usecase.GetCfp(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, expected, actualRes, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-2. 400: データ取得エラー(依頼元)
// [x] 2-3. 400: データ取得エラー(依頼先)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetCfp_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	getCfpInputRequesting := f.NewGetCfpInput()
	getCfpInputRequesting.TraceIDs = []uuid.UUID{uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")}
	getCfpInputRequested := f.NewGetCfpInput()
	getCfpInputRequested.TraceIDs = []uuid.UUID{uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5")}

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name              string
		input             traceability.GetCfpInput
		receiveCfp        string
		receiveCfpError   error
		receiveTrade      string
		receiveTradeError error
		expect            error
	}{
		{
			name:              "2-2. 400: データ取得エラー(依頼元)",
			input:             getCfpInputRequested,
			receiveCfp:        f.GetCfp_NoData(),
			receiveCfpError:   expectedPagingError,
			receiveTrade:      f.GetTradeRequests_NoData(),
			receiveTradeError: nil,
			expect:            expectedPagingError,
		},
		{
			name:              "2-3. 400: データ取得エラー(依頼先)",
			input:             getCfpInputRequesting,
			receiveCfp:        f.GetCfp_AllItem(),
			receiveCfpError:   nil,
			receiveTrade:      f.GetTradeRequests_NoData(),
			receiveTradeError: expectedPagingError,
			expect:            expectedPagingError,
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

				getCfpResponse := traceabilityentity.GetCfpResponses{}
				if err := json.Unmarshal([]byte(test.receiveCfp), &getCfpResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				getTradeRequestResponse := traceabilityentity.GetTradeRequestsResponse{}
				if err := json.Unmarshal([]byte(test.receiveTrade), &getTradeRequestResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("GetCfp", mock.Anything, mock.Anything).Return(getCfpResponse, test.receiveCfpError)
				traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(getTradeRequestResponse, test.receiveTradeError)

				usecase := usecase.NewCfpTraceabilityUsecase(traceabilityRepositoryMock)
				_, err := usecase.GetCfp(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect, err)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// put /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答(依頼元)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutCfp(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"
	traceabilityResponseCfpID := "892262ab-6795-4a97-bf25-d92c512ebb31"
	putCfpInput := f.NewPutCfpInputs2()
	for i, cfp := range putCfpInput {
		cfp.TraceID = "38bdd8a5-76a7-a53d-de12-725707b04a1b"
		cfp.CfpID = &traceabilityResponseCfpID
		putCfpInput[i] = cfp
	}

	tests := []struct {
		name       string
		operatorId string
		input      traceability.PutCfpInputs
		receive    string
	}{
		{
			name:       "1-1. 200: 正常終了",
			operatorId: f.OperatorId,
			input:      putCfpInput,
			receive:    f.PutCfp_AllItem("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
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
				inputModels, _ := test.input.ToModels()

				postCfpResponse := traceabilityentity.PostCfpResponses{}

				if err := json.Unmarshal([]byte(test.receive), &postCfpResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				traceabilityRepositoryMock.On("PostCfp", mock.Anything, mock.Anything).Return(postCfpResponse, common.ResponseHeaders{}, nil)

				usecase := usecase.NewCfpTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, _, err := usecase.PutCfp(c, test.input, test.operatorId)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, inputModels, actualRes, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// put /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: 全項目応答(依頼元)
// [x] 2-2. 400: 全項目応答(依頼元)
// [x] 2-3. 400: 全項目応答(依頼元)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_PutCfp_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	putCfpInputs := f.NewPutCfpInputs2()
	for i, cfp := range putCfpInputs {
		cfp.TraceID = "38bdd8a5-76a7-a53d-de12-725707b04a1b"
		cfp.CfpID = nil
		putCfpInputs[i] = cfp
	}

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name         string
		operatorId   string
		input        traceability.PutCfpInputs
		receive      string
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 400: データ取得エラー(CFP)",
			operatorId:   f.OperatorId,
			input:        putCfpInputs,
			receive:      f.PutCfp_AllItem("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
			receiveError: expectedPagingError,
			expect:       expectedPagingError,
		},
		{
			name:         "2-2. 200: パースエラー(CFP)",
			operatorId:   f.OperatorId,
			input:        putCfpInputs,
			receive:      f.PutCfp_AllItem_InvalidCfp("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
			receiveError: nil,
			expect:       nil,
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

				postCfpResponse := traceabilityentity.PostCfpResponses{}

				if err := json.Unmarshal([]byte(test.receive), &postCfpResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}

				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				if test.receiveError != nil {
					traceabilityRepositoryMock.On("PostCfp", mock.Anything, mock.Anything).Return(traceabilityentity.PostCfpResponses{}, common.ResponseHeaders{}, test.receiveError)
				} else {
					traceabilityRepositoryMock.On("PostCfp", mock.Anything, mock.Anything).Return(postCfpResponse, common.ResponseHeaders{}, nil)
				}

				usecase := usecase.NewCfpTraceabilityUsecase(traceabilityRepositoryMock)
				_, _, err := usecase.PutCfp(c, test.input, test.operatorId)
				assert.Equal(t, err, test.expect, f.AssertMessage)
			},
		)
	}
}
