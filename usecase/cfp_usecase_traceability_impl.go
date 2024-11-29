package usecase

import (
	"errors"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// cfpTraceabilityUsecase
// Summary: This struct defines traceability use cases for the cfp.
type cfpTraceabilityUsecase struct {
	TraceabilityRepository repository.TraceabilityRepository
}

// NewCfpTraceabilityUsecase
// Summary: This function creates a new cfpTraceabilityUsecase.
// input: r(repository.TraceabilityRepository) traceability repository
// output: (ICfpUsecase) cfp use case interface
func NewCfpTraceabilityUsecase(r repository.TraceabilityRepository) ICfpUsecase {
	return &cfpTraceabilityUsecase{r}
}

// GetCfp
// Summary: This function gets a list of cfp.
// input: c(echo.Context) echo context
// input: getCfpInput(traceability.GetCfpInput) GetCfpInput object
// output: ([]traceability.CfpModel) list of CfpModel
// output: (error) error object
func (u *cfpTraceabilityUsecase) GetCfp(c echo.Context, getCfpInput traceability.GetCfpInput) ([]traceability.CfpModel, error) {
	traceIDsStr := common.JoinUUIDs(getCfpInput.TraceIDs, ",")
	request := traceabilityentity.GetCfpRequest{
		OperatorID: getCfpInput.OperatorID.String(),
		TraceID:    traceIDsStr,
	}

	// Get a list of CFP information for owned parts
	response, err := u.TraceabilityRepository.GetCfp(c, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, err
	}
	cfpModels, err := response.ToModels()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, err
	}

	// Get a list of CFP information for parts for which a response is being requested
	tradeRequestsRequest := traceabilityentity.GetTradeRequestsRequest{
		OperatorID: getCfpInput.OperatorID.String(),
		TraceID:    &traceIDsStr,
	}
	tradeRequestsResponse, err := u.TraceabilityRepository.GetTradeRequests(c, tradeRequestsRequest)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, err
	}
	requestCfpModels, err := tradeRequestsResponse.ToCfpModels()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, err
	}

	if len(requestCfpModels) > 0 {
		cfpModels = append(cfpModels, requestCfpModels...)
	}

	sortedCfpModels := cfpModels.SortCfpModelsByTraceIDs(getCfpInput.TraceIDs)

	return sortedCfpModels, nil
}

// PutCfp
// Summary: This function puts a list of cfp.
// input: c(echo.Context) echo context
// input: putCfpInputs(traceability.PutCfpInputs) PutCfpInputs object
// input: operatorID(string) ID of the operator
// output: ([]traceability.CfpModel) list of CfpModel
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *cfpTraceabilityUsecase) PutCfp(c echo.Context, putCfpInputs traceability.PutCfpInputs, operatorID string) ([]traceability.CfpModel, common.ResponseHeaders, error) {
	cfpModels, err := putCfpInputs.ToModels()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return nil, common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}

	request, err := traceabilityentity.NewPostCfpRequestFromModel(cfpModels, operatorID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, common.ResponseHeaders{}, nil
	}

	res, headers, err := u.TraceabilityRepository.PostCfp(c, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, common.ResponseHeaders{}, err
	}
	cfpID, err := uuid.Parse(res.GetCfpID())
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, common.ResponseHeaders{}, nil
	}

	cfpModels.SetCfpID(cfpID)

	return cfpModels, headers, nil
}
