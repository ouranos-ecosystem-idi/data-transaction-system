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
// input: getStatusModel(traceability.GetStatusModel) GetStatusModel object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func (u *statusTraceabilityUsecase) GetStatus(c echo.Context, getStatusModel traceability.GetStatusModel) ([]traceability.StatusModel, *string, error) {
	switch getStatusModel.StatusTarget {
	case traceability.Request:
		return getRequestStatus(u, c, getStatusModel)
	case traceability.Response:
		return getResponseStatus(u, c, getStatusModel)
	default:
		return getBothStatus(u, c, getStatusModel)
	}
}

// getRequestStatus
// Summary: This function gets a list of request status.
// input: u(*statusTraceabilityUsecase) use case interface
// input: c(echo.Context) echo context
// input: getStatusModel(traceability.GetStatusModel) GetStatusModel object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func getRequestStatus(u *statusTraceabilityUsecase, c echo.Context, getStatusModel traceability.GetStatusModel) ([]traceability.StatusModel, *string, error) {
	req := traceabilityentity.GetTradeRequestsRequest{
		OperatorID: getStatusModel.OperatorID.String(),
		TraceID:    common.UUIDPtrToStringPtr(getStatusModel.TraceID),
		After:      common.UUIDPtrToStringPtr(getStatusModel.After),
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

	if len(response.TradeRequests) > getStatusModel.Limit {
		response.Next = response.TradeRequests[getStatusModel.Limit].Trade.TradeRelation.DownstreamTraceID
		response.TradeRequests = response.TradeRequests[:getStatusModel.Limit]
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
// input: getStatusModel(traceability.GetStatusModel) GetStatusModel object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func getResponseStatus(u *statusTraceabilityUsecase, c echo.Context, getStatusModel traceability.GetStatusModel) ([]traceability.StatusModel, *string, error) {
	req := traceabilityentity.GetTradeRequestsReceivedRequest{
		OperatorID: getStatusModel.OperatorID.String(),
		RequestID:  common.UUIDPtrToStringPtr(getStatusModel.StatusID),
		After:      common.UUIDPtrToStringPtr(getStatusModel.After),
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

	if len(response.TradeRequests) > getStatusModel.Limit {
		response.Next = response.TradeRequests[getStatusModel.Limit].Request.RequestID
		response.TradeRequests = response.TradeRequests[:getStatusModel.Limit]
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
// input: getStatusModel(traceability.GetStatusModel) GetStatusModel object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func getBothStatus(u *statusTraceabilityUsecase, c echo.Context, getStatusModel traceability.GetStatusModel) ([]traceability.StatusModel, *string, error) {
	tradeRecievedReq := traceabilityentity.GetTradeRequestsReceivedRequest{
		OperatorID: getStatusModel.OperatorID.String(),
		RequestID:  common.UUIDPtrToStringPtr(getStatusModel.StatusID),
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
		OperatorID: getStatusModel.OperatorID.String(),
	}

	tradeRequestRes := traceabilityentity.GetTradeRequestsResponse{}
	if !(getStatusModel.StatusID != nil && len(recievedStatusModels) > 0) {
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

	if getStatusModel.StatusID != nil {
		modelsForSort = modelsForSort.FilterByStatusID(*getStatusModel.StatusID)
	}
	res, after := modelsForSort.GetStatusModels(getStatusModel)

	return res, after, nil
}

// PutStatusCancel
// Summary: This function cancel a request status.
// input: c(echo.Context) echo context
// input: statusModel(traceability.StatusModel) StatusModel object
// output: (error) error object
func (u *statusTraceabilityUsecase) PutStatusCancel(c echo.Context, statusModel traceability.StatusModel) error {
	operatorID := c.Get("operatorID").(string)
	req := traceabilityentity.NewPostTradeRequestsCancelRequest(operatorID, statusModel.StatusID.String())
	_, err := u.TraceabilityRepository.PostTradeRequestsCancel(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return err
	}
	return nil
}

// PutStatusReject
// Summary: This function reject a request status.
// input: c(echo.Context) echo context
// input: statusModel(traceability.StatusModel) StatusModel object
// output: (error) error object
func (u *statusTraceabilityUsecase) PutStatusReject(c echo.Context, statusModel traceability.StatusModel) error {
	operatorID := c.Get("operatorID").(string)
	req := traceabilityentity.NewPostTradeRequestsRejectRequest(operatorID, statusModel.StatusID.String(), statusModel.ReplyMessage)
	_, err := u.TraceabilityRepository.PostTradeRequestsReject(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return err
	}
	return nil
}
