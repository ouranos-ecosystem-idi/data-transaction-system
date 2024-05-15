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

// GetCfp
// Summary: This function execute get cfp api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.GetCfpRequest) GetCfpRequest object
// output: (traceabilityentity.GetCfpResponses) GetCfpResponses object
// output: (error) error object
func (r *traceabilityRepository) GetCfp(c echo.Context, request traceabilityentity.GetCfpRequest) (traceabilityentity.GetCfpResponses, error) {
	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Get(client.PathCfp, token, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.GetCfpResponses{}, err
	}

	var res traceabilityentity.GetCfpResponses
	if resString != "[]" {
		if err := json.Unmarshal([]byte(resString), &res); err != nil {
			logger.Set(c).Errorf(err.Error())

			return traceabilityentity.GetCfpResponses{}, err
		}
	}

	return res, nil
}

// PostCfp
// Summary: This function execute post cfp api.
// input: c(echo.Context) echo context
// input: requests(traceabilityentity.PostCfpRequest) PostCfpRequest object
// output: (traceabilityentity.PostCfpResponses) PostCfpResponses object
// output: (error) error object
func (r *traceabilityRepository) PostCfp(c echo.Context, requests traceabilityentity.PostCfpRequest) (traceabilityentity.PostCfpResponses, error) {
	body, err := json.Marshal(requests)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, err
	}

	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Post(client.PathCfp, token, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, err
	}

	var responses traceabilityentity.PostCfpResponses
	if err := json.Unmarshal([]byte(resString), &responses); err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, err
	}

	return responses, nil
}
