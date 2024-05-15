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

// GetCfpCertifications
// Summary: This function execute get cfp certifications api.
// input: c(echo.Context) echo context
// input: request(traceabilityentity.GetCfpCertificationsRequest) api request
// output: (traceabilityentity.GetCfpCertificationsResponse) api response
// output: (error) error object
func (r *traceabilityRepository) GetCfpCertifications(c echo.Context, request traceabilityentity.GetCfpCertificationsRequest) (traceabilityentity.GetCfpCertificationsResponse, error) {
	token := common.ExtractBearerToken(c)

	resString, err := r.cli.Get(client.PathCfpCertifications, token, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceabilityentity.GetCfpCertificationsResponse{}, err
	}

	var res traceabilityentity.GetCfpCertificationsResponse
	if resString != "[]" {
		if err := json.Unmarshal([]byte(resString), &res); err != nil {
			logger.Set(c).Errorf(err.Error())

			return traceabilityentity.GetCfpCertificationsResponse{}, err
		}
	}

	return res, nil
}
