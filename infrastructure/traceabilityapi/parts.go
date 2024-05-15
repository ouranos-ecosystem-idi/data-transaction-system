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

// GetParts
// Summary: This function is used to retrieve the results of filtering the part information by traceId.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.GetPartsRequest) api request
// input: limit(int) upper threshold
// output: res(traceabilityentity.GetPartsResponse) api response
// output: (error) Error object
func (r *traceabilityRepository) GetParts(c echo.Context, request traceabilityentity.GetPartsRequest, limit int) (res traceabilityentity.GetPartsResponse, err error) {
	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Get(client.PathParts, token, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.GetPartsResponse{}, err
	}

	if err := json.Unmarshal([]byte(resString), &res); err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceabilityentity.GetPartsResponse{}, err
	}

	if len(res.Parts) > limit {
		res.Next = res.Parts[limit].TraceID
		res.Parts = res.Parts[:limit]
	}

	return res, nil
}
