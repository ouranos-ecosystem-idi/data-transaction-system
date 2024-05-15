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
// input: getPartsStructureModel(traceability.GetPartsStructureModel) target of the partsStructure
// output: (traceability.PartsStructureEntity) api response
// output: (error) error object
func (r *traceabilityRepository) GetPartsStructures(c echo.Context, request traceabilityentity.GetPartsStructuresRequest) (res traceabilityentity.GetPartsStructuresResponse, err error) {
	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Get(client.PathPartsStructures, token, request)
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
// output: (error) error object
func (r *traceabilityRepository) PostPartsStructures(c echo.Context, request traceabilityentity.PostPartsStructuresRequest) (traceabilityentity.PostPartsStructuresResponse, error) {
	token := common.ExtractBearerToken(c)

	body, err := json.Marshal(request)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostPartsStructuresResponse{}, err
	}

	resString, err := r.cli.Post(client.PathPartsStructures, token, body)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostPartsStructuresResponse{}, err
	}
	var response traceabilityentity.PostPartsStructuresResponse
	if err = json.Unmarshal([]byte(resString), &response); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.PostPartsStructuresResponse{}, err
	}

	return response, nil
}
