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
// Cfp GetCfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetCfp(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetCfpRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.GetCfpResponses
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.GetCfpRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    "2680ed32-19a3-435b-a094-23ff43aaa611",
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetCfp_AllItem(),
			expect: traceabilityentity.GetCfpResponses{
				traceabilityentity.GetCfpResponse{
					Cfp: traceabilityentity.GetCfpResponseCfp{
						CfpID:                           "892262ab-6795-4a97-bf25-d92c512ebb31",
						TraceID:                         "38bdd8a5-76a7-a53d-de12-725707b04a1b",
						PreProcessingOwnEmissions:       0.5,
						MainProductionOwnEmissions:      0.6,
						PreProcessingSupplierEmissions:  1.1,
						MainProductionSupplierEmissions: 1.2,
						EmissionsUnitName:               "kgCO2e/kilogram",
						CfpComment:                      "部品B01001のCFPのCFP情報コメント",
						Dqr: traceabilityentity.Dqr{
							PreProcessingTeR:  common.Float64Ptr(3.1),
							PreProcessingGeR:  common.Float64Ptr(3.2),
							PreProcessingTiR:  common.Float64Ptr(3.3),
							MainProductionTeR: common.Float64Ptr(3.4),
							MainProductionGeR: common.Float64Ptr(3.5),
							MainProductionTiR: common.Float64Ptr(3.6),
						},
						ParentFlag: true,
					},
					TotalCfp: traceabilityentity.GetCfpResponseTotalCfp{
						TotalPreProcessingOwnOriginatedEmissions:       common.Float64Ptr(1.5),
						TotalMainProductionOwnOriginatedEmissions:      common.Float64Ptr(1.6),
						TotalPreProcessingSupplierOriginatedEmissions:  common.Float64Ptr(2.1),
						TotalMainProductionSupplierOriginatedEmissions: common.Float64Ptr(2.2),
						TotalEmissionsUnitName:                         common.StringPtr("kgCO2e/kilogram"),
						TotalDqr: traceabilityentity.Dqr{
							PreProcessingTeR:  common.Float64Ptr(4.1),
							PreProcessingGeR:  common.Float64Ptr(4.2),
							PreProcessingTiR:  common.Float64Ptr(4.3),
							MainProductionTeR: common.Float64Ptr(4.4),
							MainProductionGeR: common.Float64Ptr(4.5),
							MainProductionTiR: common.Float64Ptr(4.6),
						},
					},
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathCfp),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetCfpResponses{}
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
				req := httptest.NewRequest("GET", "http://localhost:8080/cfp", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, err := r.GetCfp(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect), len(actual))
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp GetCfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetCfp_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetCfpRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.GetCfpRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    "2680ed32-19a3-435b-a094-23ff43aaa611",
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathCfp),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetCfpResponses{}
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
				req := httptest.NewRequest("GET", "http://localhost:8080/cfp", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, err := r.GetCfp(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp PostCfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostCfp(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostCfpRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.PostCfpResponses
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.PostCfpRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				Cfp: traceabilityentity.PostCfpRequests{
					traceabilityentity.PostCfpRequestCfp{
						CfpID:                                     common.StringPtr("892262ab-6795-4a97-bf25-d92c512ebb31"),
						TraceID:                                   "38bdd8a5-76a7-a53d-de12-725707b04a1b",
						PreProcessingOwnOriginatedEmissions:       common.Float64Ptr(0.5),
						MainProductionOwnOriginatedEmissions:      common.Float64Ptr(0.6),
						PreProcessingSupplierOriginatedEmissions:  common.Float64Ptr(1.1),
						MainProductionSupplierOriginatedEmissions: common.Float64Ptr(1.2),
						EmissionsUnitName:                         "kgCO2e/kilogram",
						CfpComment:                                common.StringPtr("部品B01001のCFPのCFP情報コメント"),
						Dqr: &traceabilityentity.PostCfpRequestCfpDqr{
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
			receiveCode: http.StatusOK,
			receiveBody: fixtures.PutCfp_AllItem("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
			expect: traceabilityentity.PostCfpResponses{
				traceabilityentity.PostCfpResponse{
					TraceID: "38bdd8a5-76a7-a53d-de12-725707b04a1b",
					CfpID:   "892262ab-6795-4a97-bf25-d92c512ebb31",
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathCfp),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostCfpResponses{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080/cfp", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.PostCfp(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}

}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp PostCfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostCfp_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostCfpRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.PostCfpRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				Cfp: traceabilityentity.PostCfpRequests{
					traceabilityentity.PostCfpRequestCfp{
						CfpID:                                     common.StringPtr("892262ab-6795-4a97-bf25-d92c512ebb31"),
						TraceID:                                   "38bdd8a5-76a7-a53d-de12-725707b04a1b",
						PreProcessingOwnOriginatedEmissions:       common.Float64Ptr(0.5),
						MainProductionOwnOriginatedEmissions:      common.Float64Ptr(0.6),
						PreProcessingSupplierOriginatedEmissions:  common.Float64Ptr(1.1),
						MainProductionSupplierOriginatedEmissions: common.Float64Ptr(1.2),
						EmissionsUnitName:                         "kgCO2e/kilogram",
						CfpComment:                                common.StringPtr("部品B01001のCFPのCFP情報コメント"),
						Dqr: &traceabilityentity.PostCfpRequestCfpDqr{
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathCfp),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostCfpResponses{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080/cfp", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.PostCfp(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
