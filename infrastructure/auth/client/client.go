package client

import (
	"fmt"
	"io"
	"net/http"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"
)

const (
	PathSystemAuthAPIKey = "api/v1/systemAuth/apiKey"
	PathSystemToken      = "api/v1/systemAuth/token"
)

// Client
// Summary: This is structure which defines Client.
type Client struct {
	httpClient    *http.Client
	apiBaseURL    string
	commonHeaders map[string]string
}

// NewClient
// Summary: This is function which creates new Client.
// input: apiKey(string) API key
// input: apiBaseURL(string) API base URL
// output: (*Client) Client pointer
func NewClient(apiKey string, apiBaseURL string) *Client {
	return &Client{
		httpClient: &http.Client{},
		apiBaseURL: apiBaseURL,
		commonHeaders: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"apiKey":       apiKey,
		},
	}
}

// Post
// Summary: This is function which sends POST request.
// input: path(string) path
// input: body(io.Reader) body
// output: (string) response body
// output: (error) error object
func (c *Client) Post(path string, body io.Reader) (string, error) {
	endPointURL := fmt.Sprintf("%v/%v", c.apiBaseURL, path)

	req, err := http.NewRequest("POST", endPointURL, body)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}

	for key, value := range c.commonHeaders {
		req.Header.Set(key, value)
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

	responseBodyStr := string(responseBody)
	if resp.StatusCode != http.StatusOK {
		var commonErr *common.CustomError
		if apiErr := common.ToAuthAPIError(responseBodyStr); apiErr != nil {
			commonErr = apiErr.ToCustomError(resp.StatusCode)
		} else {
			commonErr = common.NewCustomError(common.CustomErrorCode500, "Internal Server Error", nil, common.HTTPErrorSourceAuth)
		}
		return "", commonErr
	}
	return responseBodyStr, nil
}
