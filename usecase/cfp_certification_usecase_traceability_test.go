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

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/trade テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 全項目応答(証明書)
// [x] 1-2. 200: 全項目応答(依頼)
// [x] 1-3. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含む)
// [x] 1-4. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含まない)
// [x] 1-5. 200: 必須項目のみ(証明書)
// [x] 1-6. 200: 必須項目のみ(証明書)(キーなし)
// [x] 1-7. 200: 回答済(必須項目のみ)(依頼)
// [x] 1-8. 200: 回答済(必須項目のみ)(キーなし)(依頼)
// [x] 1-9. 200: 未回答(必須項目のみ)(依頼)
// [x] 1-10. 200: 未回答(必須項目のみ)(キーなし)(依頼)
// [x] 1-11. 200: 検索結果なし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetCfpCertification(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfpCertification"

	dsExpectedResAll := `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"cfpCertificationDescription": "B01のCFP証明書説明。",
			"cfpCertificationFileInfo": [
				{
					"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
					"fileName": "B01_CFP.pdf"
				}
			]
		}
	]`

	dsExpectedResAllForTradeRequest := `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"cfpCertificationDescription": null,
			"cfpCertificationFileInfo": [
				{
					"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
					"fileName": "B01_CFP.pdf"
				}
			]
		}
	]`

	dsExpectedResRequireOnly := `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"cfpCertificationDescription": "",
			"cfpCertificationFileInfo": [
				{
					"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
					"fileName": "B01_CFP.pdf"
				}
			]
		}
	]`

	dsExpectedResRequireOnlyAnsweredForTradeRequest := `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			"cfpCertificationDescription": null,
			"cfpCertificationFileInfo": null
		}
	]`

	dsExpectedResRequireOnlyAnsweringForTradeRequest := `[]`

	dsExpectedResNoData := `[]`

	tests := []struct {
		name          string
		input         traceability.GetCfpCertificationInput
		receive_cert  string
		receive_trade string
		expect        string
	}{
		{
			name:         "1-1. 200: 全項目応答(証明書)",
			input:        f.NewGetCfpCertificationInput(),
			receive_cert: f.GetCfpCertifications_AllItem(),
			expect:       dsExpectedResAll,
		},
		{
			name:          "1-2. 200: 全項目応答(依頼)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_AllItem(),
			expect:        dsExpectedResAllForTradeRequest,
		},
		{
			name:          "1-3. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含む)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_AllItem_WithNull(),
			expect:        dsExpectedResAllForTradeRequest,
		},
		{
			name:          "1-4. 200: 全項目応答(依頼)(トレサビレスポンスにnullを含まない)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_AllItem_WithUndefined(),
			expect:        dsExpectedResAllForTradeRequest,
		},
		{
			name:         "1-5. 200: 必須項目のみ(証明書)",
			input:        f.NewGetCfpCertificationInput(),
			receive_cert: f.GetCfpCertifications_RequireItemOnly(),
			expect:       dsExpectedResRequireOnly,
		},
		{
			name:         "1-6. 200: 必須項目のみ(証明書)(キーなし)",
			input:        f.NewGetCfpCertificationInput(),
			receive_cert: f.GetCfpCertifications_RequireItemOnlyWithUndefined(),
			expect:       dsExpectedResRequireOnly,
		},
		{
			name:          "1-7. 200: 回答済(必須項目のみ)(依頼)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_RequireItemOnlyAnswered(),
			expect:        dsExpectedResRequireOnlyAnsweredForTradeRequest,
		},
		{
			name:          "1-8. 200: 回答済(必須項目のみ)(キーなし)(依頼)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_RequireItemOnlyAnsweredWithUndefined(),
			expect:        dsExpectedResRequireOnlyAnsweredForTradeRequest,
		},
		{
			name:          "1-9. 200: 未回答(必須項目のみ)(依頼)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_RequireItemOnlyAnswering(),
			expect:        dsExpectedResRequireOnlyAnsweringForTradeRequest,
		},
		{
			name:          "1-10. 200: 未回答(必須項目のみ)(キーなし)(依頼)",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_RequireItemOnlyAnsweringWithUndefined(),
			expect:        dsExpectedResRequireOnlyAnsweringForTradeRequest,
		},
		{
			name:          "1-11. 200: 検索結果なし",
			input:         f.NewGetCfpCertificationInput(),
			receive_cert:  f.GetCfpCertifications_NoData(),
			receive_trade: f.GetTradeRequests_NoData(),
			expect:        dsExpectedResNoData,
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
				getCfpCertificationsResponse := traceabilityentity.GetCfpCertificationsResponse{}

				if err := json.Unmarshal([]byte(test.receive_cert), &getCfpCertificationsResponse); err != nil {
					log.Fatalf(f.UnmarshalMockFailureMessage, err)
				}
				traceabilityRepositoryMock.On("GetCfpCertifications", mock.Anything, mock.Anything).Return(getCfpCertificationsResponse, nil)

				if test.receive_trade != "" {
					getTradeRequestsResponse := traceabilityentity.GetTradeRequestsResponse{}
					if err := json.Unmarshal([]byte(test.receive_trade), &getTradeRequestsResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(getTradeRequestsResponse, nil)
				}

				var expected traceability.CfpCertificationModels
				err := json.Unmarshal([]byte(test.expect), &expected)
				if err != nil {
					log.Fatalf(f.UnmarshalExpectFailureMessage, err)
				}

				usecase := usecase.NewCfpCertificationTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, err := usecase.GetCfpCertification(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					if len(expected) == 0 {
						assert.Equal(t, len(expected), len(actualRes), f.AssertMessage)
					} else {
						for idx, expectItem := range expected {
							assert.Equal(t, expectItem.CfpCertificationID, actualRes[idx].CfpCertificationID, f.AssertMessage)
							assert.Equal(t, expectItem.CfpCertificationDescription, actualRes[idx].CfpCertificationDescription, f.AssertMessage)
							if expectItem.CfpCertificationFileInfo == nil {
								assert.Nil(t, actualRes[idx].CfpCertificationFileInfo, f.AssertMessage)
							} else {
								assert.ElementsMatch(t, *expectItem.CfpCertificationFileInfo, *actualRes[idx].CfpCertificationFileInfo, f.AssertMessage)
							}
						}
					}
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/trade テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー(証明書)
// [x] 2-2. 400: データ取得エラー(依頼)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseTraceability_GetCfpCertification_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfpCertification"

	expectedPagingError := common.CustomError{
		Code:          400,
		Message:       "指定した識別子は存在しません",
		MessageDetail: common.StringPtr("MSGAECO0020"),
		Source:        common.HTTPErrorSourceTraceability,
	}

	tests := []struct {
		name              string
		input             traceability.GetCfpCertificationInput
		receiveCert       string
		receiveCertError  error
		receiveTradeError error
		expect            error
	}{
		{
			name:             "2-1. 400: データ取得エラー(証明書)",
			input:            f.NewGetCfpCertificationInput(),
			receiveCertError: expectedPagingError,
			expect:           expectedPagingError,
		},
		{
			name:              "2-2. 400: データ取得エラー(依頼)",
			input:             f.NewGetCfpCertificationInput(),
			receiveCert:       f.GetCfpCertifications_NoData(),
			receiveCertError:  nil,
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
				traceabilityRepositoryMock := new(mocks.TraceabilityRepository)
				if test.receiveCertError != nil {
					traceabilityRepositoryMock.On("GetCfpCertifications", mock.Anything, mock.Anything).Return(traceabilityentity.GetCfpCertificationsResponse{}, test.receiveCertError)
				} else {
					getCfpCertificationsResponse := traceabilityentity.GetCfpCertificationsResponse{}
					if err := json.Unmarshal([]byte(test.receiveCert), &getCfpCertificationsResponse); err != nil {
						log.Fatalf(f.UnmarshalMockFailureMessage, err)
					}
					traceabilityRepositoryMock.On("GetCfpCertifications", mock.Anything, mock.Anything).Return(getCfpCertificationsResponse, nil)
				}

				if test.receiveTradeError != nil {
					traceabilityRepositoryMock.On("GetTradeRequests", mock.Anything, mock.Anything).Return(traceabilityentity.GetTradeRequestsResponse{}, test.receiveTradeError)
				}

				usecase := usecase.NewCfpCertificationTraceabilityUsecase(traceabilityRepositoryMock)
				actualRes, err := usecase.GetCfpCertification(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Nil(t, actualRes)
					assert.Equal(t, test.expect, err.(common.CustomError))
				}
			},
		)
	}
}
