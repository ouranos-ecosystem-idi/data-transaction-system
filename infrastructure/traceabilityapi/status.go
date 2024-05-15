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
// output: (error) error object
func (r *traceabilityRepository) PostTradeRequestsCancel(c echo.Context, request traceabilityentity.PostTradeRequestsCancelRequest) (traceabilityentity.PostTradeRequestsCancelResponse, error) {
	token := common.ExtractBearerToken(c)

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsCancelResponse{}, err
	}

	resString, err := r.cli.Post(client.PathTradeRequestsCancel, token, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.PostTradeRequestsCancelResponse{}, err
	}
	var res traceabilityentity.PostTradeRequestsCancelResponse
	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsCancelResponse{}, err
	}

	return res, nil
}

// PostTradeRequestsReject
// Summary: This function execute post trade request reject api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.PostTradeRequestsRejectRequest) PostTradeRequestsRejectRequest object
// output: (traceabilityentity.PostTradeRequestsRejectResponse) PostTradeRequestsRejectResponse object
// output: (error) error object
func (r *traceabilityRepository) PostTradeRequestsReject(c echo.Context, request traceabilityentity.PostTradeRequestsRejectRequest) (traceabilityentity.PostTradeRequestsRejectResponse, error) {
	token := common.ExtractBearerToken(c)

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsRejectResponse{}, err
	}

	resString, err := r.cli.Post(client.PathTradeRequestsReject, token, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.PostTradeRequestsRejectResponse{}, err
	}
	var res traceabilityentity.PostTradeRequestsRejectResponse
	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostTradeRequestsRejectResponse{}, err
	}

	return res, nil
}
