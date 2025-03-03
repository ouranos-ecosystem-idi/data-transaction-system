package traceabilityapi_test

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/infrastructure/traceabilityapi"
	"data-spaces-backend/infrastructure/traceabilityapi/client"
	"data-spaces-backend/test/fixtures"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetTradeRequests テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// [x] 1-2: 正常系：正常返却の場合(トレサビレスポンスがnullのフィールドを含む場合)
// [x] 1-3: 正常系：正常返却の場合(トレサビレスポンスが未定義項目を含む場合)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetTradeRequests(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetTradeRequestsRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.GetTradeRequestsResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.GetTradeRequestsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    common.StringPtr("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				After:      common.StringPtr(""),
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetTradeRequests_AllItem(),
			expect: traceabilityentity.GetTradeRequestsResponse{
				TradeRequests: []traceabilityentity.GetTradeRequestsResponseTradeRequest{
					{
						Request: traceabilityentity.GetTradeRequestsResponseRequest{
							RequestID:                "5185a435-c039-4196-bb34-0ee0c2395478",
							RequestType:              "CFP",
							RequestStatus:            "COMPLETED",
							RequestedToOperatorID:    "f99c9546-e76e-9f15-35b2-abb9c9b21698",
							RequestedAt:              "2024-02-14T15:25:35Z",
							RequestMessage:           "A01のCFP値を回答ください",
							ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
							ResponseDueDate:          &fixtures.ResponseDueDate,
							CompletedCount:           &fixtures.CompletedCount,
							CompletedCountModifiedAt: &fixtures.CompletedCountModifiedAt,
						},
						Trade: traceabilityentity.GetTradeRequestsResponseTrade{
							TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
							TreeStatus: "UNTERMINATED",
							Downstream: traceabilityentity.GetTradeRequestsResponseTradeDownstream{
								DownstreamAmountUnitName: "kilogram",
							},
							TradeRelation: traceabilityentity.GetTradeRequestsResponseTradeRelation{
								UpstreamOperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
								DownstreamTraceID:  "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
								UpstreamTraceID:    common.StringPtr("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
							},
							TradesCount:           &fixtures.TradesCount,
							TradesCountModifiedAt: &fixtures.TradesCountModifiedAt,
						},
						Response: &traceabilityentity.GetTradeRequestsResponseResponse{
							ResponseID:                      "b26e3bd3-7443-4f23-8cce-2056052f0452",
							ResponseType:                    "CFP",
							ResponsedAt:                     "2024-02-16T17:46:21Z",
							ResponsePreProcessingEmissions:  common.Float64Ptr(0.1),
							ResponseMainProductionEmissions: common.Float64Ptr(0.4),
							EmissionsUnitName:               "kgCO2e/kilogram",
							CFPCertificationFileInfo: []traceabilityentity.GetTradeRequestsResponseCFPCertificationFileInfo{
								{
									FileID:   "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
									FileName: "B01_CFP.pdf",
								},
							},
							ResponseDqr: traceabilityentity.GetTradeRequestsResponseResponseDqr{
								PreProcessingTeR:  common.Float64Ptr(3.1),
								PreProcessingGeR:  common.Float64Ptr(3.2),
								PreProcessingTiR:  common.Float64Ptr(3.3),
								MainProductionTeR: common.Float64Ptr(3.4),
								MainProductionGeR: common.Float64Ptr(3.5),
								MainProductionTiR: common.Float64Ptr(3.6),
							},
						},
					},
				},
				Next: "026ad6a0-a689-4b8c-8a14-7304b817096d",
			},
		},
		{
			name: "1-2: 正常系：正常返却の場合(トレサビレスポンスがnullのフィールドを含む場合)",
			input: traceabilityentity.GetTradeRequestsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    common.StringPtr("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				After:      common.StringPtr(""),
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetTradeRequests_AllItem_WithNull(),
			expect: traceabilityentity.GetTradeRequestsResponse{
				TradeRequests: []traceabilityentity.GetTradeRequestsResponseTradeRequest{
					{
						Request: traceabilityentity.GetTradeRequestsResponseRequest{
							RequestID:                "5185a435-c039-4196-bb34-0ee0c2395478",
							RequestType:              "CFP",
							RequestStatus:            "COMPLETED",
							RequestedToOperatorID:    "f99c9546-e76e-9f15-35b2-abb9c9b21698",
							RequestedAt:              "2024-02-14T15:25:35Z",
							RequestMessage:           "A01のCFP値を回答ください",
							ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
							ResponseDueDate:          nil,
							CompletedCount:           nil,
							CompletedCountModifiedAt: nil,
						},
						Trade: traceabilityentity.GetTradeRequestsResponseTrade{
							TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
							TreeStatus: "UNTERMINATED",
							Downstream: traceabilityentity.GetTradeRequestsResponseTradeDownstream{
								DownstreamAmountUnitName: "kilogram",
							},
							TradeRelation: traceabilityentity.GetTradeRequestsResponseTradeRelation{
								UpstreamOperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
								DownstreamTraceID:  "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
								UpstreamTraceID:    common.StringPtr("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
							},
							TradesCount:           nil,
							TradesCountModifiedAt: nil,
						},
						Response: &traceabilityentity.GetTradeRequestsResponseResponse{
							ResponseID:                      "b26e3bd3-7443-4f23-8cce-2056052f0452",
							ResponseType:                    "CFP",
							ResponsedAt:                     "2024-02-16T17:46:21Z",
							ResponsePreProcessingEmissions:  common.Float64Ptr(0.1),
							ResponseMainProductionEmissions: common.Float64Ptr(0.4),
							EmissionsUnitName:               "kgCO2e/kilogram",
							CFPCertificationFileInfo: []traceabilityentity.GetTradeRequestsResponseCFPCertificationFileInfo{
								{
									FileID:   "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
									FileName: "B01_CFP.pdf",
								},
							},
							ResponseDqr: traceabilityentity.GetTradeRequestsResponseResponseDqr{
								PreProcessingTeR:  common.Float64Ptr(3.1),
								PreProcessingGeR:  common.Float64Ptr(3.2),
								PreProcessingTiR:  common.Float64Ptr(3.3),
								MainProductionTeR: common.Float64Ptr(3.4),
								MainProductionGeR: common.Float64Ptr(3.5),
								MainProductionTiR: common.Float64Ptr(3.6),
							},
						},
					},
				},
				Next: "026ad6a0-a689-4b8c-8a14-7304b817096d",
			},
		},
		{
			name: "1-3: 正常系：正常返却の場合(トレサビレスポンスが未定義項目を含む場合)",
			input: traceabilityentity.GetTradeRequestsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    common.StringPtr("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				After:      common.StringPtr(""),
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetTradeRequests_AllItem_WithUndefined(),
			expect: traceabilityentity.GetTradeRequestsResponse{
				TradeRequests: []traceabilityentity.GetTradeRequestsResponseTradeRequest{
					{
						Request: traceabilityentity.GetTradeRequestsResponseRequest{
							RequestID:                "5185a435-c039-4196-bb34-0ee0c2395478",
							RequestType:              "CFP",
							RequestStatus:            "COMPLETED",
							RequestedToOperatorID:    "f99c9546-e76e-9f15-35b2-abb9c9b21698",
							RequestedAt:              "2024-02-14T15:25:35Z",
							RequestMessage:           "A01のCFP値を回答ください",
							ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
							ResponseDueDate:          nil,
							CompletedCount:           nil,
							CompletedCountModifiedAt: nil,
						},
						Trade: traceabilityentity.GetTradeRequestsResponseTrade{
							TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
							TreeStatus: "UNTERMINATED",
							Downstream: traceabilityentity.GetTradeRequestsResponseTradeDownstream{
								DownstreamAmountUnitName: "kilogram",
							},
							TradeRelation: traceabilityentity.GetTradeRequestsResponseTradeRelation{
								UpstreamOperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
								DownstreamTraceID:  "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
								UpstreamTraceID:    common.StringPtr("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
							},
							TradesCount:           nil,
							TradesCountModifiedAt: nil,
						},
						Response: &traceabilityentity.GetTradeRequestsResponseResponse{
							ResponseID:                      "b26e3bd3-7443-4f23-8cce-2056052f0452",
							ResponseType:                    "CFP",
							ResponsedAt:                     "2024-02-16T17:46:21Z",
							ResponsePreProcessingEmissions:  common.Float64Ptr(0.1),
							ResponseMainProductionEmissions: common.Float64Ptr(0.4),
							EmissionsUnitName:               "kgCO2e/kilogram",
							CFPCertificationFileInfo: []traceabilityentity.GetTradeRequestsResponseCFPCertificationFileInfo{
								{
									FileID:   "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
									FileName: "B01_CFP.pdf",
								},
							},
							ResponseDqr: traceabilityentity.GetTradeRequestsResponseResponseDqr{
								PreProcessingTeR:  common.Float64Ptr(3.1),
								PreProcessingGeR:  common.Float64Ptr(3.2),
								PreProcessingTiR:  common.Float64Ptr(3.3),
								MainProductionTeR: common.Float64Ptr(3.4),
								MainProductionGeR: common.Float64Ptr(3.5),
								MainProductionTiR: common.Float64Ptr(3.6),
							},
						},
					},
				},
				Next: "026ad6a0-a689-4b8c-8a14-7304b817096d",
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequests),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetTradeRequestsResponse{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/tradeRequests", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, err := r.GetTradeRequests(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
					assert.Equal(t, test.expect.Next, actual.Next)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetTradeRequests テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetTradeRequests_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetTradeRequestsRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.GetTradeRequestsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    common.StringPtr("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
				After:      common.StringPtr(""),
			},
			receiveCode: http.StatusServiceUnavailable,
			receiveBody: fixtures.Error_MaintenanceError(),
			expect:      common.NewCustomError(http.StatusServiceUnavailable, "The service is currently undergoing maintenance. We apologize for any inconvenience.", common.StringPtr("MSGXXXXYYYY"), common.HTTPErrorSourceTraceability),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequests),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetTradeRequestsResponse{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080//tradeRequests", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, err := r.GetTradeRequests(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetTradeRequestsReceived テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// [x] 1-2: 正常系：正常返却の場合(トレサビレスポンスにnullを含む)
// [x] 1-3: 正常系：正常返却の場合(トレサビレスポンスに未定義項目を含む)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetTradeRequestsReceived(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetTradeRequestsReceivedRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.GetTradeRequestsReceivedResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.GetTradeRequestsReceivedRequest{
				OperatorID:        "b1234567-1234-1234-1234-123456789012",
				RequestID:         common.StringPtr("5185a435-c039-4196-bb34-0ee0c2395478"),
				RequestedDateFrom: &fixtures.DummyTime,
				RequestedDateTo:   &fixtures.DummyTime,
				After:             common.StringPtr(""),
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetTradeRequestsReceived_AllItem(),
			expect: traceabilityentity.GetTradeRequestsReceivedResponse{
				TradeRequests: []traceabilityentity.GetTradeRequestsReceivedResponseTradeRequest{
					{
						Request: traceabilityentity.GetTradeRequestsReceivedResponseRequest{
							RequestID:                "5185a435-c039-4196-bb34-0ee0c2395478",
							RequestType:              "CFP",
							RequestStatus:            "COMPLETED",
							RequestedFromOperatorID:  "b1234567-1234-1234-1234-123456789012",
							RequestedAt:              "2024-02-14T15:25:35Z",
							RequestMessage:           "A01のCFP値を回答ください",
							ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
							ResponseDueDate:          &fixtures.ResponseDueDate,
							CompletedCount:           &fixtures.CompletedCount,
							CompletedCountModifiedAt: &fixtures.CompletedCountModifiedAt,
						},
						Trade: traceabilityentity.GetTradeRequestsReceivedResponseTrade{
							TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
							TreeStatus: "UNTERMINATED",
							Downstream: traceabilityentity.GetTradeRequestsReceivedResponseTradeDownstream{
								DownstreamPartsItem:        "B01",
								DownstreamSupportPartsItem: "B0100",
								DownstreamPlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
								DownstreamAmountUnitName:   "kilogram",
								DownstreamPartsLabelName:   &fixtures.PartsLabelName,
								DownstreamPartsAddInfo1:    &fixtures.PartsAddInfo1,
								DownstreamPartsAddInfo2:    &fixtures.PartsAddInfo2,
								DownstreamPartsAddInfo3:    &fixtures.PartsAddInfo3,
							},
							TradeRelation: traceabilityentity.GetTradeRequestsReceivedResponseTradeRelation{
								DownstreamOperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
								DownstreamTraceID:    "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
								UpstreamTraceID:      common.StringPtr("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
							},
							TradesCount:           &fixtures.TradesCount,
							TradesCountModifiedAt: &fixtures.TradesCountModifiedAt,
						},
					},
				},
				Next: "026ad6a0-a689-4b8c-8a14-7304b817096d",
			},
		},
		{
			name: "1-2: 正常系：正常返却の場合(トレサビレスポンスにnullを含む)",
			input: traceabilityentity.GetTradeRequestsReceivedRequest{
				OperatorID:        "b1234567-1234-1234-1234-123456789012",
				RequestID:         common.StringPtr("5185a435-c039-4196-bb34-0ee0c2395478"),
				RequestedDateFrom: &fixtures.DummyTime,
				RequestedDateTo:   &fixtures.DummyTime,
				After:             common.StringPtr(""),
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetTradeRequestsReceived_AllItem_WithNull(),
			expect: traceabilityentity.GetTradeRequestsReceivedResponse{
				TradeRequests: []traceabilityentity.GetTradeRequestsReceivedResponseTradeRequest{
					{
						Request: traceabilityentity.GetTradeRequestsReceivedResponseRequest{
							RequestID:                "5185a435-c039-4196-bb34-0ee0c2395478",
							RequestType:              "CFP",
							RequestStatus:            "COMPLETED",
							RequestedFromOperatorID:  "b1234567-1234-1234-1234-123456789012",
							RequestedAt:              "2024-02-14T15:25:35Z",
							RequestMessage:           "A01のCFP値を回答ください",
							ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
							ResponseDueDate:          nil,
							CompletedCount:           nil,
							CompletedCountModifiedAt: nil,
						},
						Trade: traceabilityentity.GetTradeRequestsReceivedResponseTrade{
							TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
							TreeStatus: "UNTERMINATED",
							Downstream: traceabilityentity.GetTradeRequestsReceivedResponseTradeDownstream{
								DownstreamPartsItem:        "B01",
								DownstreamSupportPartsItem: "B0100",
								DownstreamPlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
								DownstreamAmountUnitName:   "kilogram",
								DownstreamPartsLabelName:   nil,
								DownstreamPartsAddInfo1:    nil,
								DownstreamPartsAddInfo2:    nil,
								DownstreamPartsAddInfo3:    nil,
							},
							TradeRelation: traceabilityentity.GetTradeRequestsReceivedResponseTradeRelation{
								DownstreamOperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
								DownstreamTraceID:    "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
								UpstreamTraceID:      common.StringPtr("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
							},
							TradesCount:           nil,
							TradesCountModifiedAt: nil,
						},
					},
				},
				Next: "026ad6a0-a689-4b8c-8a14-7304b817096d",
			},
		},
		{
			name: "1-3: 正常系：正常返却の場合(トレサビレスポンスに未定義項目を含む)",
			input: traceabilityentity.GetTradeRequestsReceivedRequest{
				OperatorID:        "b1234567-1234-1234-1234-123456789012",
				RequestID:         common.StringPtr("5185a435-c039-4196-bb34-0ee0c2395478"),
				RequestedDateFrom: &fixtures.DummyTime,
				RequestedDateTo:   &fixtures.DummyTime,
				After:             common.StringPtr(""),
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetTradeRequestsReceived_AllItem_WithUndefined(),
			expect: traceabilityentity.GetTradeRequestsReceivedResponse{
				TradeRequests: []traceabilityentity.GetTradeRequestsReceivedResponseTradeRequest{
					{
						Request: traceabilityentity.GetTradeRequestsReceivedResponseRequest{
							RequestID:                "5185a435-c039-4196-bb34-0ee0c2395478",
							RequestType:              "CFP",
							RequestStatus:            "COMPLETED",
							RequestedFromOperatorID:  "b1234567-1234-1234-1234-123456789012",
							RequestedAt:              "2024-02-14T15:25:35Z",
							RequestMessage:           "A01のCFP値を回答ください",
							ReplyMessage:             common.StringPtr("A01のCFP値を回答しました"),
							ResponseDueDate:          nil,
							CompletedCount:           nil,
							CompletedCountModifiedAt: nil,
						},
						Trade: traceabilityentity.GetTradeRequestsReceivedResponseTrade{
							TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
							TreeStatus: "UNTERMINATED",
							Downstream: traceabilityentity.GetTradeRequestsReceivedResponseTradeDownstream{
								DownstreamPartsItem:        "B01",
								DownstreamSupportPartsItem: "B0100",
								DownstreamPlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
								DownstreamAmountUnitName:   "kilogram",
								DownstreamPartsLabelName:   nil,
								DownstreamPartsAddInfo1:    nil,
								DownstreamPartsAddInfo2:    nil,
								DownstreamPartsAddInfo3:    nil,
							},
							TradeRelation: traceabilityentity.GetTradeRequestsReceivedResponseTradeRelation{
								DownstreamOperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
								DownstreamTraceID:    "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
								UpstreamTraceID:      common.StringPtr("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
							},
							TradesCount:           nil,
							TradesCountModifiedAt: nil,
						},
					},
				},
				Next: "026ad6a0-a689-4b8c-8a14-7304b817096d",
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequestsRecieved),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetTradeRequestsReceivedResponse{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/tradeRequestsReceived", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, err := r.GetTradeRequestsReceived(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
					assert.Equal(t, test.expect.Next, actual.Next)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetTradeRequestsReceived テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetTradeRequestsReceived_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetTradeRequestsReceivedRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 503の場合",
			input: traceabilityentity.GetTradeRequestsReceivedRequest{
				OperatorID:        "b1234567-1234-1234-1234-123456789012",
				RequestID:         common.StringPtr("5185a435-c039-4196-bb34-0ee0c2395478"),
				RequestedDateFrom: &fixtures.DummyTime,
				RequestedDateTo:   &fixtures.DummyTime,
				After:             common.StringPtr(""),
			},
			receiveCode: http.StatusServiceUnavailable,
			receiveBody: fixtures.Error_MaintenanceError(),
			expect:      common.NewCustomError(http.StatusServiceUnavailable, "The service is currently undergoing maintenance. We apologize for any inconvenience.", common.StringPtr("MSGXXXXYYYY"), common.HTTPErrorSourceTraceability),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequestsRecieved),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetTradeRequestsReceivedResponse{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080//tradeRequestsReceived", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, err := r.GetTradeRequestsReceived(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTradeRequests テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTradeRequests(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradeRequestsRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.PostTradeRequestsResponses
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.PostTradeRequestsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TradeRequests: []traceabilityentity.PostTradeRequestsRequestTradeRequest{
					{
						DownstreamTraceID:  "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						UpstreamOperatorID: "38bdd8a5-76a7-a53d-de12-725707b04a1b",
						RequestType:        "CFP",
						RequestMessage:     nil,
					},
				},
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.PutTradeRequests(),
			expect: traceabilityentity.PostTradeRequestsResponses{
				traceabilityentity.PostTradeRequestsResponse{
					TradeID:           "a84012cc-73fb-4f9b-9130-59ae546f7092",
					RequestID:         "5185a435-c039-4196-bb34-0ee0c2395478",
					DownstreamTraceID: "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequests),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsResponses{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/tradeRequests", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.PostTradeRequests(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTradeRequests テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTradeRequests_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradeRequestsRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.PostTradeRequestsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TradeRequests: []traceabilityentity.PostTradeRequestsRequestTradeRequest{
					{
						DownstreamTraceID:  "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						UpstreamOperatorID: "38bdd8a5-76a7-a53d-de12-725707b04a1b",
						RequestType:        "CFP",
						RequestMessage:     nil,
					},
				},
			},
			receiveCode: http.StatusServiceUnavailable,
			receiveBody: fixtures.Error_MaintenanceError(),
			expect:      common.NewCustomError(http.StatusServiceUnavailable, "The service is currently undergoing maintenance. We apologize for any inconvenience.", common.StringPtr("MSGXXXXYYYY"), common.HTTPErrorSourceTraceability),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequests),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsResponses{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080//tradeRequests", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.PostTradeRequests(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTrades テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTrades(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradesRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.PostTradesResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.PostTradesRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
				TraceID:    "",
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.PostTrades(),
			expect: traceabilityentity.PostTradesResponse{
				TradeID: "a84012cc-73fb-4f9b-9130-59ae546f7092",
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTrades),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradesResponse{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/trades", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.PostTrades(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTrades テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTrades_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradesRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.PostTradesRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TradeID:    "a84012cc-73fb-4f9b-9130-59ae546f7092",
				TraceID:    "",
			},
			receiveCode: http.StatusServiceUnavailable,
			receiveBody: fixtures.Error_MaintenanceError(),
			expect:      common.NewCustomError(http.StatusServiceUnavailable, "The service is currently undergoing maintenance. We apologize for any inconvenience.", common.StringPtr("MSGXXXXYYYY"), common.HTTPErrorSourceTraceability),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTrades),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsResponses{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						} else {
							response := common.TraceabilityAPIError{}
							if err := json.Unmarshal([]byte(test.receiveBody), &response); err != nil {
								assert.Fail(t, "Unmarshal Error")
							}
							return httpmock.NewJsonResponse(test.receiveCode, response)
						}
					})

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080//trades", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.PostTrades(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
