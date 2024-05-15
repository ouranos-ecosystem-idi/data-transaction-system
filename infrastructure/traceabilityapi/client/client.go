package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"

	"go.uber.org/zap"
)

const (
	PathParts                 = "parts"
	PathPartsStructures       = "partsStructures"
	PathTrades                = "trades"
	PathTradeRequests         = "tradeRequests"
	PathTradeRequestsCancel   = "tradeRequests/cancel"
	PathTradeRequestsReject   = "tradeRequests/reject"
	PathTradeRequestsRecieved = "tradeRequestsReceived"
	PathCfp                   = "cfp"
	PathCfpCertifications     = "cfpCertifications"
	PathCfpCertificationFiles = "cfpCertificationFiles"
)

// Client
// Summary: This is structure which defines Client.
type Client struct {
	httpClient    *http.Client
	apiBaseURL    string
	commonHeaders map[string]string
}

// NewClient
// Summary: This is function which is used to get the new client
// input: apiKey(string) API Key
// input: apiVersion(string) API Version
// input: apiBaseURL(string) API Base URL
// output: (*Client) pointer of Client struct
func NewClient(apiKey string, apiVersion string, apiBaseURL string) *Client {
	return &Client{
		httpClient: &http.Client{},
		apiBaseURL: apiBaseURL,
		commonHeaders: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"x-api-key":    apiKey,
			"Api-Version":  apiVersion,
		},
	}
}

type QueryParams interface{}

// Get
// Summary: This is function which is used to get the data from the API
// input: path(string) Path
// input: authorization(string) Authorization
// input: params(QueryParams) Query Params
// output: (string) Response Body
// output: (error) error object
func (c *Client) Get(path, authorization string, params QueryParams) (string, error) {
	endPointURL := fmt.Sprintf("%v/%v", c.apiBaseURL, path)

	url := buildGetURL(endPointURL, params)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}

	for key, value := range c.commonHeaders {
		req.Header.Set(key, value)
	}

	req.Header.Set("Authorization", authorization)

	logger.Set(nil).Infof(logger.AccessInfoLog, url)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}
	bodyDump(false, url, req.Header, nil, body)

	bodyStr := string(body)
	if resp.StatusCode != http.StatusOK {
		var commonErr *common.CustomError
		if apiErr := common.ToTracebilityAPIError(bodyStr); apiErr != nil {
			commonErr = apiErr.ToCustomError(resp.StatusCode)
		} else {
			commonErr = common.NewCustomError(common.CustomErrorCode500, "Internal Server Error", nil, common.HTTPErrorSourceTraceability)
		}
		return "", commonErr
	}

	return bodyStr, nil
}

// buildGetURL
// Summary: This is function which is used to build the get URL
// input: endPointURL(string) End Point URL
// input: params(QueryParams) Query Params
// output: (string) URL
// output: (error) error object
func buildGetURL(endPointURL string, params QueryParams) string {
	queryValues := url.Values{}
	if params != nil {
		v := reflect.ValueOf(params)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			formatTag := field.Tag.Get("format")
			value := v.Field(i)

			if field.Type.Kind() == reflect.Ptr && value.IsNil() {
				continue
			}

			if field.Type.Kind() == reflect.String && value.String() == "" {
				continue
			}

			if field.Type.Kind() == reflect.Ptr {
				value = value.Elem()
			}

			var stringValue string
			if field.Type == reflect.TypeOf(time.Time{}) || (field.Type.Kind() == reflect.Ptr && field.Type.Elem() == reflect.TypeOf(time.Time{})) {
				timeValue := value.Interface().(time.Time)
				if formatTag != "" {
					stringValue = timeValue.Format(formatTag)
				} else {
					stringValue = timeValue.Format(time.RFC3339)
				}
			} else {
				stringValue = fmt.Sprintf("%v", value.Interface())
			}

			if jsonTag != "" {
				queryValues.Add(jsonTag, stringValue)
			} else {
				queryValues.Add(field.Name, stringValue)
			}
		}
	}

	if len(queryValues) > 0 {
		endPointURL += "?" + queryValues.Encode()
	}

	return endPointURL
}

// Post
// Summary: This is function which is used to post the data to the API
// input: path(string) Path
// input: authorization(string) Authorization
// input: body([]byte) Body
// output: (string) Response Body
// output: (error) error object
func (c *Client) Post(path string, authorization string, body []byte) (string, error) {
	endPointURL := fmt.Sprintf("%v/%v", c.apiBaseURL, path)

	r := bytes.NewBuffer(body)
	reqBody := io.TeeReader(r, new(bytes.Buffer))

	req, err := http.NewRequest("POST", endPointURL, reqBody)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}

	for key, value := range c.commonHeaders {
		req.Header.Set(key, value)
	}

	req.Header.Set("Authorization", authorization)

	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}

	logger.Set(nil).Infof(logger.AccessInfoLog, endPointURL)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}
	bodyDump(false, endPointURL, req.Header, body, responseBody)

	responseBodyStr := string(responseBody)
	if resp.StatusCode != http.StatusOK {
		zap.S().Errorf("TraceabilityAPI Error, URL: %v, Status: %v, Header, %v, Body: %v", req.URL, resp.Status, resp.Header, responseBodyStr)
		var commonErr *common.CustomError
		if apiErr := common.ToTracebilityAPIError(responseBodyStr); apiErr != nil {
			commonErr = apiErr.ToCustomError(resp.StatusCode)
		} else {
			commonErr = common.NewCustomError(common.CustomErrorCode500, "Internal Server Error", nil, common.HTTPErrorSourceTraceability)
		}
		return "", commonErr
	}

	return responseBodyStr, nil
}

// bodyDump
// Summary: This is function which is used to dump the body
// input: isReq(bool) Is Request
// input: path(string) Path
// input: header(http.Header) Header
// input: reqBody([]byte) Request Body
func bodyDump(isReq bool, path string, header http.Header, reqBody []byte, resBody []byte) {
	if common.IsOutputDump() {
		for k := range header {
			if k == "Authorization" {
				header[k] = []string{"Bearer ******"}
			}
			if k == "x-api-key" {
				header[k] = []string{"******"}
			}
		}
		logger.Set(nil).Debugf(logger.TraceabilityAPIDebugLog, path, header, string(reqBody), string(resBody))
	}
}
