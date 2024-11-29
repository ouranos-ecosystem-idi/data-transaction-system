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

// PostTradeRequestsCancel
// Summary: This function execute post trade request cancel api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.PostTradeRequestsCancelRequest) PostTradeRequestsCancelRequest object
// output: (traceabilityentity.PostTradeRequestsCancelResponse) PostTradeRequestsCancelResponse object
// output: (common.ResponseHeaders) ResponseHeaders object
// output: (error) error object
func (r *traceabilityRepository) PostTradeRequestsCancel(c echo.Context, request traceabilityentity.PostTradeRequestsCancelRequest) (traceabilityentity.PostTradeRequestsCancelResponse, common.ResponseHeaders, error) {
	headers := map[string]string{}
	headers["Authorization"] = common.ExtractBearerToken(c)
	if lang := common.ExtractAcceptLanguage(c); lang != "" {
		headers["accept-language"] = lang
	}

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsCancelResponse{}, common.ResponseHeaders{}, err
	}

	res, err := r.cli.Post(c, client.PathTradeRequestsCancel, headers, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.PostTradeRequestsCancelResponse{}, common.ResponseHeaders{}, err
	}
	var response traceabilityentity.PostTradeRequestsCancelResponse
	if err := json.Unmarshal([]byte(res.Body), &response); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsCancelResponse{}, common.ResponseHeaders{}, err
	}

	return response, res.Headers, nil
}

// PostTradeRequestsReject
// Summary: This function execute post trade request reject api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.PostTradeRequestsRejectRequest) PostTradeRequestsRejectRequest object
// output: (traceabilityentity.PostTradeRequestsRejectResponse) PostTradeRequestsRejectResponse object
// output: (common.ResponseHeaders) ResponseHeaders object
// output: (error) error object
func (r *traceabilityRepository) PostTradeRequestsReject(c echo.Context, request traceabilityentity.PostTradeRequestsRejectRequest) (traceabilityentity.PostTradeRequestsRejectResponse, common.ResponseHeaders, error) {
	headers := map[string]string{}
	headers["Authorization"] = common.ExtractBearerToken(c)
	if lang := common.ExtractAcceptLanguage(c); lang != "" {
		headers["accept-language"] = lang
	}

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsRejectResponse{}, common.ResponseHeaders{}, err
	}

	res, err := r.cli.Post(c, client.PathTradeRequestsReject, headers, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.PostTradeRequestsRejectResponse{}, common.ResponseHeaders{}, err
	}
	var response traceabilityentity.PostTradeRequestsRejectResponse
	if err := json.Unmarshal([]byte(res.Body), &response); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsRejectResponse{}, common.ResponseHeaders{}, err
	}

	return response, res.Headers, nil
}
