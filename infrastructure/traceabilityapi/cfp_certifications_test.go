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
// Traceability GetCfpCertifications テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetCfpCertifications(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetCfpCertificationsRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.GetCfpCertificationsResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.GetCfpCertificationsRequest{
				OperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				TraceID:    "2680ed32-19a3-435b-a094-23ff43aaa611",
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetCfpCertifications_AllItem(),
			expect: traceabilityentity.GetCfpCertificationsResponse{
				traceabilityentity.GetCfpCertificationsResponseCfpCertification{
					CfpCertificationID:          "a84012cc-73fb-4f9b-9130-59ae546f7092",
					TraceID:                     "2680ed32-19a3-435b-a094-23ff43aaa611",
					CfpCertificationDescription: "B01のCFP証明書説明。",
					CfpCertificationFileInfo: &[]traceabilityentity.CfpCertificationFileInfo{
						{
							OperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
							FileID:     "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
							FileName:   "B01_CFP.pdf",
						},
					},
					CreatedAt: common.StringPtr("2024-01-01T00:00:00Z"),
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathCfpCertifications),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetCfpCertificationsResponse{}
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
				req := httptest.NewRequest("GET", fmt.Sprintf("%s/%s", "http://localhost:8080", client.PathCfpCertifications), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, err := r.GetCfpCertifications(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect), len(actual))
					for i, data := range test.expect {
						assert.Equal(t, data.CfpCertificationID, actual[i].CfpCertificationID)
						assert.Equal(t, data.TraceID, actual[i].TraceID)
						assert.Equal(t, data.CfpCertificationDescription, actual[i].CfpCertificationDescription)
						assert.Equal(t, data.CreatedAt, actual[i].CreatedAt)
						assert.ElementsMatch(t, *data.CfpCertificationFileInfo, *actual[i].CfpCertificationFileInfo)
					}
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetCfpCertifications テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetCfpCertifications_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetCfpCertificationsRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.GetCfpCertificationsRequest{
				OperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("%s/cfpCertifications?operatorId=%s&traceId=%s", "http://localhost:8080", test.input.OperatorID, test.input.TraceID),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetCfpCertificationsResponse{}
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
				req := httptest.NewRequest("GET", "http://localhost:8080/cfpCertifications", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, err := r.GetCfpCertifications(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
