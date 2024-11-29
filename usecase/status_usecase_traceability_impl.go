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

// statusTraceabilityUsecase
// Summary: This struct defines traceability use cases for the status.
type statusTraceabilityUsecase struct {
	TraceabilityRepository repository.TraceabilityRepository
}

// NewStatusTraceabilityUsecase
// Summary: This function creates a new statusTraceabilityUsecase.
// input: r(repository.TraceabilityRepository) traceability repository
// output: (IStatusUsecase) status use case interface
func NewStatusTraceabilityUsecase(r repository.TraceabilityRepository) IStatusUsecase {
	return &statusTraceabilityUsecase{r}
}

// GetStatus
// Summary: This function gets a list of request and response status.
// input: c(echo.Context) echo context
// input: getStatusInput(traceability.GetStatusInput) GetStatusInput object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func (u *statusTraceabilityUsecase) GetStatus(c echo.Context, getStatusInput traceability.GetStatusInput) ([]traceability.StatusModel, *string, error) {
	switch getStatusInput.StatusTarget {
	case traceability.Request:
		return getRequestStatus(u, c, getStatusInput)
	case traceability.Response:
		return getResponseStatus(u, c, getStatusInput)
	default:
		return getBothStatus(u, c, getStatusInput)
	}
}

// getRequestStatus
// Summary: This function gets a list of request status.
// input: u(*statusTraceabilityUsecase) use case interface
// input: c(echo.Context) echo context
// input: getStatusInput(traceability.GetStatusInput) GetStatusInput object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func getRequestStatus(u *statusTraceabilityUsecase, c echo.Context, getStatusInput traceability.GetStatusInput) ([]traceability.StatusModel, *string, error) {
	req := traceabilityentity.GetTradeRequestsRequest{
		OperatorID: getStatusInput.OperatorID.String(),
		TraceID:    common.UUIDPtrToStringPtr(getStatusInput.TraceID),
		After:      common.UUIDPtrToStringPtr(getStatusInput.After),
	}

	response, err := u.TraceabilityRepository.GetTradeRequests(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, nil, err
	}

	if len(response.TradeRequests) > getStatusInput.Limit {
		response.Next = response.TradeRequests[getStatusInput.Limit].Trade.TradeRelation.DownstreamTraceID
		response.TradeRequests = response.TradeRequests[:getStatusInput.Limit]
	}

	ms, err := response.ToStatusModels()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	return ms, response.GetNextPtr(), nil
}

// getResponseStatus
// Summary: This function gets a list of response status.
// input: u(*statusTraceabilityUsecase) use case interface
// input: c(echo.Context) echo context
// input: getStatusInput(traceability.GetStatusInput) GetStatusInput object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func getResponseStatus(u *statusTraceabilityUsecase, c echo.Context, getStatusInput traceability.GetStatusInput) ([]traceability.StatusModel, *string, error) {
	req := traceabilityentity.GetTradeRequestsReceivedRequest{
		OperatorID: getStatusInput.OperatorID.String(),
		RequestID:  common.UUIDPtrToStringPtr(getStatusInput.StatusID),
		After:      common.UUIDPtrToStringPtr(getStatusInput.After),
	}

	response, err := u.TraceabilityRepository.GetTradeRequestsReceived(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, nil, err
	}

	if len(response.TradeRequests) > getStatusInput.Limit {
		response.Next = response.TradeRequests[getStatusInput.Limit].Request.RequestID
		response.TradeRequests = response.TradeRequests[:getStatusInput.Limit]
	}

	models, err := response.ToStatusModels()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	return models, response.GetNextPtr(), nil
}

// getBothStatus
// Summary: This function gets a list of both request and response status.
// input: u(*statusTraceabilityUsecase) use case interface
// input: c(echo.Context) echo context
// input: getStatusInputs(traceability.GetStatusInputs) GetStatusInputs object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func getBothStatus(u *statusTraceabilityUsecase, c echo.Context, getStatusInput traceability.GetStatusInput) ([]traceability.StatusModel, *string, error) {
	tradeRecievedReq := traceabilityentity.GetTradeRequestsReceivedRequest{
		OperatorID: getStatusInput.OperatorID.String(),
		RequestID:  common.UUIDPtrToStringPtr(getStatusInput.StatusID),
	}

	hasRecievedNext := true
	tradeRecievedRes := traceabilityentity.GetTradeRequestsReceivedResponse{}
	for hasRecievedNext {
		TradeRecievedRes, err := u.TraceabilityRepository.GetTradeRequestsReceived(c, tradeRecievedReq)
		if err != nil {
			var customErr *common.CustomError
			if errors.As(err, &customErr) && customErr.IsWarn() {
				logger.Set(c).Warnf(err.Error())
			} else {
				logger.Set(c).Errorf(err.Error())
			}

			return nil, nil, err
		}
		tradeRecievedRes.TradeRequests = append(tradeRecievedRes.TradeRequests, TradeRecievedRes.TradeRequests...)
		next := TradeRecievedRes.GetNextPtr()
		hasRecievedNext = next != nil
		if hasRecievedNext {
			tradeRecievedReq.After = next
		}
	}

	recievedStatusModels, err := tradeRecievedRes.ToStatusModelsForSort()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	req := traceabilityentity.GetTradeRequestsRequest{
		OperatorID: getStatusInput.OperatorID.String(),
	}

	tradeRequestRes := traceabilityentity.GetTradeRequestsResponse{}
	if !(getStatusInput.StatusID != nil && len(recievedStatusModels) > 0) {
		hasRequestNext := true
		for hasRequestNext {
			TradeRes, err := u.TraceabilityRepository.GetTradeRequests(c, req)
			if err != nil {
				var customErr *common.CustomError
				if errors.As(err, &customErr) && customErr.IsWarn() {
					logger.Set(c).Warnf(err.Error())
				} else {
					logger.Set(c).Errorf(err.Error())
				}

				return nil, nil, err
			}
			tradeRequestRes.TradeRequests = append(TradeRes.TradeRequests, tradeRequestRes.TradeRequests...)
			next := TradeRes.GetNextPtr()
			hasRequestNext = next != nil
			if hasRequestNext {
				req.After = next
			}
		}
	}

	requestStatusModels, err := tradeRequestRes.ToStatusModelsForSort()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	modelsForSort := make(traceability.StatusModelsForSort, 0, len(requestStatusModels)+len(recievedStatusModels))
	modelsForSort = append(modelsForSort, requestStatusModels...)
	modelsForSort = append(modelsForSort, recievedStatusModels...)
	modelsForSort = modelsForSort.SortByRequestedAt()

	if getStatusInput.StatusID != nil {
		modelsForSort = modelsForSort.FilterByStatusID(*getStatusInput.StatusID)
	}
	res, after := modelsForSort.GetStatusModels(getStatusInput)

	return res, after, nil
}

// PutStatusCancel
// Summary: This function cancel a request status.
// input: c(echo.Context) echo context
// input: putStatusInput(traceability.PutStatusInput) PutStatusInput object
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *statusTraceabilityUsecase) PutStatusCancel(c echo.Context, putStatusInput traceability.PutStatusInput) (common.ResponseHeaders, error) {
	statusModel, err := putStatusInput.ToModel()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}

	operatorID := c.Get("operatorID").(string)
	req := traceabilityentity.NewPostTradeRequestsCancelRequest(operatorID, statusModel.StatusID.String())
	_, headers, err := u.TraceabilityRepository.PostTradeRequestsCancel(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return common.ResponseHeaders{}, err
	}
	return headers, nil
}

// PutStatusReject
// Summary: This function reject a request status.
// input: c(echo.Context) echo context
// input: putStatusInput(traceability.PutStatusInput) PutStatusInput object
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *statusTraceabilityUsecase) PutStatusReject(c echo.Context, putStatusInput traceability.PutStatusInput) (common.ResponseHeaders, error) {
	statusModel, err := putStatusInput.ToModel()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}

	operatorID := c.Get("operatorID").(string)
	req := traceabilityentity.NewPostTradeRequestsRejectRequest(operatorID, statusModel.StatusID.String(), statusModel.ReplyMessage)
	_, headers, err := u.TraceabilityRepository.PostTradeRequestsReject(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return common.ResponseHeaders{}, err
	}
	return headers, nil
}
