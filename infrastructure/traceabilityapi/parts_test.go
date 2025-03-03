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
// Traceability GetParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetParts(tt *testing.T) {

	parentFlag := true
	tests := []struct {
		name        string
		input       traceabilityentity.GetPartsRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.GetPartsResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.GetPartsRequest{
				OperatorID:       "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				TraceID:          common.StringPtr("2680ed32-19a3-435b-a094-23ff43aaa611"),
				PartsItem:        common.StringPtr("B01"),
				SupportPartsItem: common.StringPtr("A000001"),
				PlantID:          common.StringPtr("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
				ParentFlag:       &parentFlag,
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetParts_AllItem(nil),
			expect: traceabilityentity.GetPartsResponse{
				Parts: []traceabilityentity.GetPartsResponseParts{
					{
						TraceID:          "2680ed32-19a3-435b-a094-23ff43aaa611",
						PartsItem:        "B01",
						SupportPartsItem: common.StringPtr("A000001"),
						PlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						OperatorID:       "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						AmountUnitName:   common.StringPtr("kilogram"),
						EndFlag:          false,
						ParentFlag:       parentFlag,
						PartsLabelName:   common.StringPtr("PartsB"),
						PartsAddInfo1:    common.StringPtr("Ver3.0"),
						PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
						PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
					},
				},
				Next: "2680ed32-19a3-435b-a094-23ff43aaa612",
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathParts),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetPartsResponse{}
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
				req := httptest.NewRequest("GET", "http://localhost:8080/parts", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, err := r.GetParts(c, test.input, 1)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect.Parts), len(actual.Parts))
					assert.ElementsMatch(t, test.expect.Parts, actual.Parts)
					assert.Equal(t, test.expect.Next, actual.Next)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetParts_Abnormal(tt *testing.T) {

	parentFlag := true
	tests := []struct {
		name        string
		input       traceabilityentity.GetPartsRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.GetPartsRequest{
				OperatorID:       "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				TraceID:          common.StringPtr("2680ed32-19a3-435b-a094-23ff43aaa611"),
				PartsItem:        common.StringPtr("B01"),
				SupportPartsItem: common.StringPtr("A000001"),
				PlantID:          common.StringPtr("eedf264e-cace-4414-8bd3-e10ce1c090e0"),
				ParentFlag:       &parentFlag,
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathParts),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetPartsResponse{}
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
				req := httptest.NewRequest("GET", "http://localhost:8080/parts", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, err := r.GetParts(c, test.input, 1)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability DeleteParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_DeleteParts(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.DeletePartsRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.DeletePartsResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.DeletePartsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    "2680ed32-19a3-435b-a094-23ff43aaa611",
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.DeleteParts_AllItem("2680ed32-19a3-435b-a094-23ff43aaa611"),
			expect: traceabilityentity.DeletePartsResponse{
				TraceID: "2680ed32-19a3-435b-a094-23ff43aaa611",
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
				httpmock.RegisterResponder("DELETE", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathParts),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.DeletePartsResponse{}
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
				req := httptest.NewRequest("DELETE", "http://localhost:8080/parts", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.DeleteParts(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability DeleteParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// [x] 2-2. 異常系：400の場合
// [x] 2-3. 異常系：400の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_DeleteParts_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.DeletePartsRequest
		receiveCode int
		receiveBody string
		expect      error
		skip        bool
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.DeletePartsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    "2680ed32-19a3-435b-a094-23ff43aaa611",
			},
			receiveCode: http.StatusServiceUnavailable,
			receiveBody: fixtures.Error_MaintenanceError(),
			expect:      common.NewCustomError(http.StatusServiceUnavailable, "The service is currently undergoing maintenance. We apologize for any inconvenience.", common.StringPtr("MSGXXXXYYYY"), common.HTTPErrorSourceTraceability),
			// スタブ実装中は試験不可のため、テストをスキップする
			skip: true,
		},
		{
			name: "2-2: 異常系：400の場合",
			input: traceabilityentity.DeletePartsRequest{
				OperatorID: "",
				TraceID:    "2680ed32-19a3-435b-a094-23ff43aaa611",
			},
			receiveCode: http.StatusBadRequest,
			receiveBody: fixtures.Error_RequireError("事業者識別子"),
			expect:      common.NewCustomError(http.StatusBadRequest, "事業者識別子は必須項目です。", common.StringPtr("MSGAECO0001"), common.HTTPErrorSourceTraceability),
			skip:        false,
		},
		{
			name: "2-3: 異常系：400の場合",
			input: traceabilityentity.DeletePartsRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				TraceID:    "",
			},
			receiveCode: http.StatusBadRequest,
			receiveBody: fixtures.Error_RequireError("トレース識別子"),
			expect:      common.NewCustomError(http.StatusServiceUnavailable, "トレース識別子は必須項目です。", common.StringPtr("MSGAECO0001"), common.HTTPErrorSourceTraceability),
			skip:        false,
		},
	}

	for _, test := range tests {
		if test.skip {
			continue
		}
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {

				httpmock.Activate()
				defer httpmock.DeactivateAndReset()
				httpmock.RegisterResponder("DELETE", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathParts),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.DeletePartsResponse{}
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
				req := httptest.NewRequest("DELETE", "http://localhost:8080/parts", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.DeleteParts(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
