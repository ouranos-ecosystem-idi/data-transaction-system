package traceabilityapi

import (
	"encoding/json"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/infrastructure/traceabilityapi/client"

	"github.com/labstack/echo/v4"
)

// GetPartsStructure
// Summary: This function  the partsStructure of a request and response.
// input: request(traceabilityentity.GetPartsStructureRequest) target of the partsStructure
// output: (traceability.PartsStructureEntity) api response
// output: (error) error object
func (r *traceabilityRepository) GetPartsStructures(c echo.Context, request traceabilityentity.GetPartsStructuresRequest) (res traceabilityentity.GetPartsStructuresResponse, err error) {
	headers := map[string]string{}
	headers["Authorization"] = common.ExtractBearerToken(c)
	if lang := common.ExtractAcceptLanguage(c); lang != "" {
		headers["accept-language"] = lang
	}

	resString, err := r.cli.Get(c, client.PathPartsStructures, headers, request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.GetPartsStructuresResponse{}, err
	}

	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.GetPartsStructuresResponse{}, err
	}

	return res, nil
}

// PostPartsStructures
// Summary: This function post the partsStructure of a request and response.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.PostPartsStructuresRequest) target of the partsStructure
// output: (traceabilityentity.PostPartsStructuresResponse) api response
// output: (common.ResponseHeaders) ResponseHeaders object
// output: (error) error object
func (r *traceabilityRepository) PostPartsStructures(c echo.Context, request traceabilityentity.PostPartsStructuresRequest) (traceabilityentity.PostPartsStructuresResponse, common.ResponseHeaders, error) {
	headers := map[string]string{}
	headers["Authorization"] = common.ExtractBearerToken(c)
	if lang := common.ExtractAcceptLanguage(c); lang != "" {
		headers["accept-language"] = lang
	}

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostPartsStructuresResponse{}, common.ResponseHeaders{}, err
	}

	res, err := r.cli.Post(c, client.PathPartsStructures, headers, body)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostPartsStructuresResponse{}, common.ResponseHeaders{}, err
	}
	var response traceabilityentity.PostPartsStructuresResponse
	if err = json.Unmarshal([]byte(res.Body), &response); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostPartsStructuresResponse{}, common.ResponseHeaders{}, err
	}

	return response, res.Headers, nil
}
