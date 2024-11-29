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
// Traceability PostTradeRequestsCancel テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTradeRequestsCancel(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradeRequestsCancelRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.PostTradeRequestsCancelResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.PostTradeRequestsCancelRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				CancelRequests: []traceabilityentity.PostTradeRequestsCancelRequestCancelRequest{
					{
						RequestID: "5185a435-c039-4196-bb34-0ee0c2395478",
					},
				},
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.PutPostTradeRequestsCancelResponse(),
			expect: traceabilityentity.PostTradeRequestsCancelResponse{
				traceabilityentity.PostTradeRequestsCancelResponseCancelRequests{
					RequestID: "5185a435-c039-4196-bb34-0ee0c2395478",
					TradeID:   "a84012cc-73fb-4f9b-9130-59ae546f7092",
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequestsCancel),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsCancelResponse{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080/tradeRequests/cancel", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.PostTradeRequestsCancel(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTradeRequestsCancel テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTradeRequestsCancel_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradeRequestsCancelRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.PostTradeRequestsCancelRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				CancelRequests: []traceabilityentity.PostTradeRequestsCancelRequestCancelRequest{
					{
						RequestID: "5185a435-c039-4196-bb34-0ee0c2395478",
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequestsCancel),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsCancelResponse{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080//tradeRequests/cancel", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.PostTradeRequestsCancel(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTradeRequestsReject テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTradeRequestsReject(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradeRequestsRejectRequest
		receiveCode int
		receiveBody string
		expect      traceabilityentity.PostTradeRequestsRejectResponse
	}{
		{
			name: "1-1: 正常系：正常返却の場合",
			input: traceabilityentity.PostTradeRequestsRejectRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				RejectRequests: []traceabilityentity.PostRejectRequest{
					{
						RequestID:    "5185a435-c039-4196-bb34-0ee0c2395478",
						ReplyMessage: common.StringPtr("A01のCFP値を回答しました"),
					},
				},
			},
			receiveCode: http.StatusOK,
			receiveBody: fixtures.PutPostTradeRequestsRejectResponse(),
			expect: traceabilityentity.PostTradeRequestsRejectResponse{
				traceabilityentity.PostTradeRequestsRejectResponseRejectRequests{
					RequestID: "5185a435-c039-4196-bb34-0ee0c2395478",
					TradeID:   "a84012cc-73fb-4f9b-9130-59ae546f7092",
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequestsReject),
					func(req *http.Request) (*http.Response, error) {
						assert.Equal(t, "ja-JP", req.Header.Get("accept-language"))
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsRejectResponse{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080/tradeRequests/reject", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set("Accept-Language", "ja-JP")
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")
				r := traceabilityapi.NewTraceabilityRepository(cli)
				actual, _, err := r.PostTradeRequestsReject(c, test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Traceability PostTradeRequestsReject テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：503の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Traceability_PostTradeRequestsReject_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		input       traceabilityentity.PostTradeRequestsRejectRequest
		receiveCode int
		receiveBody string
		expect      error
	}{
		{
			name: "2-1: 異常系：503の場合",
			input: traceabilityentity.PostTradeRequestsRejectRequest{
				OperatorID: "b1234567-1234-1234-1234-123456789012",
				RejectRequests: []traceabilityentity.PostRejectRequest{
					{
						RequestID:    "5185a435-c039-4196-bb34-0ee0c2395478",
						ReplyMessage: common.StringPtr("A01のCFP値を回答しました"),
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
				httpmock.RegisterResponder("POST", fmt.Sprintf("=~^%s/%s?.*", "http://localhost:8080", client.PathTradeRequestsReject),
					func(req *http.Request) (*http.Response, error) {
						if test.receiveCode == http.StatusOK {
							response := traceabilityentity.PostTradeRequestsRejectResponse{}
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
				req := httptest.NewRequest("POST", "http://localhost:8080//tradeRequests/reject", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.Set("operatorID", test.input.OperatorID)

				cli := client.NewClient("APIKey", "APIVersion", "http://localhost:8080")

				r := traceabilityapi.NewTraceabilityRepository(cli)
				_, _, err := r.PostTradeRequestsReject(c, test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
