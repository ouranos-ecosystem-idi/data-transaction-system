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
// Traceability GetPartsStructures テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetPartsStructures(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetPartsStructuresRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.GetPartsStructuresResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.GetPartsStructuresRequest{
				OperatorID:    "b1234567-1234-1234-1234-123456789012",
				ParentTraceID: "2680ed32-19a3-435b-a094-23ff43aaa611",
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetPartsStructure_AllItem(),
			expect: traceabilityentity.GetPartsStructuresResponse{
				Parent: &traceabilityentity.GetPartsStructuresResponseParent{
					TraceID:          "2680ed32-19a3-435b-a094-23ff43aaa611",
					PartsItem:        "B01",
					SupportPartsItem: common.StringPtr("A000001"),
					PlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
					OperatorID:       "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					AmountUnitName:   common.StringPtr("kilogram"),
					EndFlag:          false,
					PartsLabelName:   common.StringPtr("PartsB"),
					PartsAddInfo1:    common.StringPtr("Ver3.0"),
					PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
					PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
				},
				Children: []traceabilityentity.GetPartsStructuresResponseChildren{
					{
						PartsStructureID: "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
						TraceID:          "1c2f37f5-25b9-dea5-346a-7b88035f2553",
						PartsItem:        "B01001",
						SupportPartsItem: common.StringPtr("B001"),
						PlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						OperatorID:       "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						AmountUnitName:   common.StringPtr("kilogram"),
						EndFlag:          false,
						Amount:           common.Float64Ptr(2.1),
						PartsLabelName:   common.StringPtr("PartsB"),
						PartsAddInfo1:    common.StringPtr("Ver3.0"),
						PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
						PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathPartsStructures),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.GetPartsStructuresResponse{}
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
				req := httptest.NewRequest("GET", "http://localhost:8080/partsStructures", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, err := r.GetPartsStructures(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability GetPartsStructures テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_GetPartsStructures_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.GetPartsStructuresRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.GetPartsStructuresRequest{
				OperatorID:    "b1234567-1234-1234-1234-123456789012",
				ParentTraceID: "2680ed32-19a3-435b-a094-23ff43aaa611",
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
				httpmock.RegisterResponder("GET", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathPartsStructures),
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
				_, err := r.GetPartsStructures(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostPartsStructures テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostPartsStructures(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostPartsStructuresRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.PostPartsStructuresResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.PostPartsStructuresRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				Parent: traceabilityentity.PostPartsStructuresRequestParent{
					PartsItem:        "B01",
					SupportPartsItem: common.StringPtr("A000001"),
					PlantID:          "Plant1",
					OperatorID:       "b1234567-1234-1234-1234-123456789012",
					AmountUnitName:   common.StringPtr("kilogram"),
					EndFlag:          false,
					PartsLabelName:   common.StringPtr("PartsB"),
					PartsAddInfo1:    common.StringPtr("Ver3.0"),
					PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
					PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
				},
				Children: []traceabilityentity.PostPartsStructuresRequestChild{
					{
						PartsItem:        "B01001",
						SupportPartsItem: common.StringPtr("B001"),
						PlantID:          "b1234567-1234-1234-1234-123456789012",
						OperatorID:       "b1234567-1234-1234-1234-123456789012",
						AmountUnitName:   common.StringPtr("kilogram"),
						EndFlag:          false,
						Amount:           common.Float64Ptr(2.1),
						PartsLabelName:   common.StringPtr("PartsB"),
						PartsAddInfo1:    common.StringPtr("Ver3.0"),
						PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
						PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
					},
				},
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.GetPartsStructure_AllItem(),
			expect: traceabilityentity.PostPartsStructuresResponse{
				Parent: traceabilityentity.PostPartsStructuresResponseParent{
					TraceID:          "2680ed32-19a3-435b-a094-23ff43aaa611",
					PartsItem:        "B01",
					SupportPartsItem: common.StringPtr("A000001"),
					PartsLabelName:   common.StringPtr("PartsB"),
					PartsAddInfo1:    common.StringPtr("Ver3.0"),
					PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
					PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
				},
				Children: []traceabilityentity.PostPartsStructuresResponseChild{
					{
						PartsStructureID: "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
						TraceID:          "1c2f37f5-25b9-dea5-346a-7b88035f2553",
						PartsItem:        "B01001",
						PlantID:          "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						SupportPartsItem: common.StringPtr("B001"),
						PartsLabelName:   common.StringPtr("PartsB"),
						PartsAddInfo1:    common.StringPtr("Ver3.0"),
						PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
						PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathPartsStructures),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostPartsStructuresResponse{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080/partsStructures", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.PostPartsStructures(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostPartsStructures テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostPartsStructures_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostPartsStructuresRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.PostPartsStructuresRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				Parent: traceabilityentity.PostPartsStructuresRequestParent{
					PartsItem:        "B01",
					SupportPartsItem: common.StringPtr("A000001"),
					PlantID:          "Plant1",
					OperatorID:       "b1234567-1234-1234-1234-123456789012",
					AmountUnitName:   common.StringPtr("kilogram"),
					EndFlag:          false,
					PartsLabelName:   common.StringPtr("PartsB"),
					PartsAddInfo1:    common.StringPtr("Ver3.0"),
					PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
					PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
				},
				Children: []traceabilityentity.PostPartsStructuresRequestChild{
					{
						PartsItem:        "B01001",
						SupportPartsItem: common.StringPtr("B001"),
						PlantID:          "b1234567-1234-1234-1234-123456789012",
						OperatorID:       "b1234567-1234-1234-1234-123456789012",
						AmountUnitName:   common.StringPtr("kilogram"),
						EndFlag:          false,
						Amount:           common.Float64Ptr(2.1),
						PartsLabelName:   common.StringPtr("PartsB"),
						PartsAddInfo1:    common.StringPtr("Ver3.0"),
						PartsAddInfo2:    common.StringPtr("2024-12-01-2024-12-31"),
						PartsAddInfo3:    common.StringPtr("任意の情報が入ります"),
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathPartsStructures),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostPartsStructuresResponse{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080/partsStructures", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.PostPartsStructures(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
