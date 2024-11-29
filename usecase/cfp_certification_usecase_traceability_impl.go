package usecase

import (
	"errors"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/labstack/echo/v4"
)

// cfpCertificationTraceabilityUsecase
// Summary: This struct defines traceability use cases for the cfp certification.
type cfpCertificationTraceabilityUsecase struct {
	TraceabilityRepository repository.TraceabilityRepository
}

// NewCfpCertificationTraceabilityUsecase
// Summary: This function creates a new cfpCertificationTraceabilityUsecase.
// input: r(repository.TraceabilityRepository) traceability repository
// output: (ICfpCertificationUsecase) cfp certification use case interface
func NewCfpCertificationTraceabilityUsecase(r repository.TraceabilityRepository) ICfpCertificationUsecase {
	return &cfpCertificationTraceabilityUsecase{r}
}

// GetCfpCertification
// Summary: This function gets cfp certification.
// input: c(echo.Context) echo context
// input: getCfpCertificationInput(traceability.GetCfpCertificationInput) GetCfpCertificationInput object
// output: (traceability.CfpCertificationModels) CfpCertificationModels object
// output: (error) error object
func (u *cfpCertificationTraceabilityUsecase) GetCfpCertification(c echo.Context, getCfpCertificationInput traceability.GetCfpCertificationInput) (traceability.CfpCertificationModels, error) {
	getCfpCertificationsRequest := traceabilityentity.GetCfpCertificationsRequest{
		OperatorID: getCfpCertificationInput.OperatorID.String(),
		TraceID:    getCfpCertificationInput.TraceID.String(),
	}

	getCfpCertificationsResponse, err := u.TraceabilityRepository.GetCfpCertifications(c, getCfpCertificationsRequest)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}
		return nil, err
	}

	if len(getCfpCertificationsResponse) > 0 {
		return getCfpCertificationsResponse.ToModels(), nil
	}

	getTradeRequestsForCfpCertificationsRequest := traceabilityentity.GetTradeRequestsRequest{
		OperatorID: getCfpCertificationInput.OperatorID.String(),
		TraceID:    common.StringPtr(getCfpCertificationInput.TraceID.String()),
	}

	getTradeRequestsResponse, err := u.TraceabilityRepository.GetTradeRequests(c, getTradeRequestsForCfpCertificationsRequest)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}
		return nil, err
	}
	return getTradeRequestsResponse.ToCertificationModels()
}
