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
	headers := map[string]string{}
	headers["Authorization"] = common.ExtractBearerToken(c)
	if lang := common.ExtractAcceptLanguage(c); lang != "" {
		headers["accept-language"] = lang
	}

	resString, err := r.cli.Get(c, client.PathCfp, headers, request)
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
// output: (common.ResponseHeaders) ResponseHeaders object
// output: (error) error object
func (r *traceabilityRepository) PostCfp(c echo.Context, requests traceabilityentity.PostCfpRequest) (traceabilityentity.PostCfpResponses, common.ResponseHeaders, error) {
	body, err := json.Marshal(requests)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, common.ResponseHeaders{}, err
	}

	headers := map[string]string{}
	headers["Authorization"] = common.ExtractBearerToken(c)
	if lang := common.ExtractAcceptLanguage(c); lang != "" {
		headers["accept-language"] = lang
	}

	res, err := r.cli.Post(c, client.PathCfp, headers, body)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, common.ResponseHeaders{}, err
	}

	var responses traceabilityentity.PostCfpResponses
	if err := json.Unmarshal([]byte(res.Body), &responses); err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, common.ResponseHeaders{}, err
	}

	return responses, res.Headers, nil
}
