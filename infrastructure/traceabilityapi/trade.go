package traceabilityapi

import (
	"encoding/json"
	"errors"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/infrastructure/traceabilityapi/client"

	"github.com/labstack/echo/v4"
)

// GetTradeRequests
// Summary: This function execute get trade requests api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.GetTradeRequestsRequest) api request
// output: (traceabilityentity.GetTradeRequestsResponse) api response
// output: (error) error object
func (r *traceabilityRepository) GetTradeRequests(c echo.Context, request traceabilityentity.GetTradeRequestsRequest) (traceabilityentity.GetTradeRequestsResponse, error) {
	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Get(client.PathTradeRequests, token, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.GetTradeRequestsResponse{}, err
	}
	var res traceabilityentity.GetTradeRequestsResponse
	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.GetTradeRequestsResponse{}, err
	}

	return res, nil
}

// PostTradeRequests
// Summary: This function execute post trade request api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.PostTradeRequestsRequest) api request
// output: (traceabilityentity.PostTradeRequestsResponses) api response
// output: (error) error object
func (r *traceabilityRepository) PostTradeRequests(c echo.Context, request traceabilityentity.PostTradeRequestsRequest) (traceabilityentity.PostTradeRequestsResponses, error) {
	token := common.ExtractBearerToken(c)

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, err
	}

	resString, err := r.cli.Post(client.PathTradeRequests, token, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, err
	}

	var response traceabilityentity.PostTradeRequestsResponses
	if err = json.Unmarshal([]byte(resString), &response); err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, err
	}

	return response, nil

}

// GetTradeRequestsReceived
// Summary: This function execute get trade requests received api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.GetTradeRequestsReceivedRequest) api request
// output: (traceabilityentity.GetTradeRequestsReceivedResponse) api response
// output: (error) error object
func (r *traceabilityRepository) GetTradeRequestsReceived(c echo.Context, request traceabilityentity.GetTradeRequestsReceivedRequest) (traceabilityentity.GetTradeRequestsReceivedResponse, error) {
	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Get(client.PathTradeRequestsRecieved, token, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.GetTradeRequestsReceivedResponse{}, err
	}
	var res traceabilityentity.GetTradeRequestsReceivedResponse
	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.GetTradeRequestsReceivedResponse{}, err
	}

	return res, nil
}

// PostTrades
// Summary: This function execute post trade api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.PostTradesRequest) api request
// output: (traceabilityentity.PostTradesResponse) api response
// output: (error) error object
func (r *traceabilityRepository) PostTrades(c echo.Context, request traceabilityentity.PostTradesRequest) (traceabilityentity.PostTradesResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradesResponse{}, err
	}

	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Post(client.PathTrades, token, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.PostTradesResponse{}, err
	}

	var res traceabilityentity.PostTradesResponse
	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradesResponse{}, err
	}

	return res, nil
}
