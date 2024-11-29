package usecase

import (
	"errors"
	"fmt"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// tradeTraceabilityUsecase
// Summary: This is structure which defines tradeTraceabilityUsecase.
type tradeTraceabilityUsecase struct {
	TraceabilityRepository repository.TraceabilityRepository
}

// NewTradeTraceabilityUsecase
// Summary: This is function to create new TradeTraceabilityUsecase.
// input: r(repository.TraceabilityRepository) repository interface
// output: (ITradeUsecase) usecase interface
func NewTradeTraceabilityUsecase(r repository.TraceabilityRepository) ITradeUsecase {
	return &tradeTraceabilityUsecase{r}
}

// GetTradeRequest
// Summary: This is function which get a list of trade and response status.
// input: c(echo.Context) echo context
// input: input(traceability.GetTradeRequestInput) GetTradeRequestInput object
// output: ([]traceability.TradeModel) list of TradeModel
// output: (*string) next id
// output: (error) error object
func (u *tradeTraceabilityUsecase) GetTradeRequest(c echo.Context, getTradeRequestInput traceability.GetTradeRequestInput) ([]traceability.TradeModel, *string, error) {
	request := traceabilityentity.GetTradeRequestsRequest{
		OperatorID: getTradeRequestInput.OperatorID.String(),
		TraceID:    common.JoinUUIDsAsPtr(getTradeRequestInput.TraceIDs, ","),
		After:      common.UUIDPtrToStringPtr(getTradeRequestInput.After),
	}

	response, err := u.TraceabilityRepository.GetTradeRequests(c, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, nil, err
	}

	if len(response.TradeRequests) > getTradeRequestInput.Limit {
		response.Next = response.TradeRequests[getTradeRequestInput.Limit].Trade.TradeRelation.DownstreamTraceID
		response.TradeRequests = response.TradeRequests[:getTradeRequestInput.Limit]
	}

	tradeModels, err := response.ToTradeModels(c)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	return tradeModels, response.GetNextPtr(), nil
}

// GetTradeResponse
// Summary: This is function which get a list of trade and response status.
// input: c(echo.Context) echo context
// input: input(traceability.GetTradeResponseInput) GetTradeResponseInput object
// output: ([]traceability.TradeResponseModel) list of TradeResponseModel
// output: (*string) next id
// output: (error) error object
func (u *tradeTraceabilityUsecase) GetTradeResponse(c echo.Context, getTradeResponseInput traceability.GetTradeResponseInput) ([]traceability.TradeResponseModel, *string, error) {
	request := traceabilityentity.GetTradeRequestsReceivedRequest{
		OperatorID: getTradeResponseInput.OperatorID.String(),
		After:      common.UUIDPtrToStringPtr(getTradeResponseInput.After),
	}

	response, err := u.TraceabilityRepository.GetTradeRequestsReceived(c, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, nil, err
	}

	// If the number of items exceeds the limit, delete the excess and set next.
	if len(response.TradeRequests) > getTradeResponseInput.Limit {
		response.Next = response.TradeRequests[getTradeResponseInput.Limit].Request.RequestID
		response.TradeRequests = response.TradeRequests[:getTradeResponseInput.Limit]
	}
	tradeResponses, err := response.ToTradeResponseModels(getTradeResponseInput.OperatorID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	return tradeResponses, response.GetNextPtr(), nil
}

// PutTradeRequest
// Summary: This is function which put trade request with TradeRequestModel.
// input: c(echo.Context) echo context
// input: putTradeRequestInput(traceability.PutTradeRequestInput) PutTradeRequestInput object
// output: (traceability.TradeRequestModel) TradeRequestModel object
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *tradeTraceabilityUsecase) PutTradeRequest(c echo.Context, putTradeRequestInput traceability.PutTradeRequestInput) (traceability.TradeRequestModel, common.ResponseHeaders, error) {
	tradeRequestModel := putTradeRequestInput.ToModel()

	req := traceabilityentity.NewPostTradeRequestRequestFromModel(tradeRequestModel)

	res, headers, err := u.TraceabilityRepository.PostTradeRequests(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceability.TradeRequestModel{}, common.ResponseHeaders{}, err
	}

	if len(res) != 1 {
		logger.Set(c).Errorf(common.UnexpectedResponse("Traceability"))

		return traceability.TradeRequestModel{}, common.ResponseHeaders{}, fmt.Errorf(common.UnexpectedResponse("Traceability"))
	}
	resElem := res[0]

	StatusID, err := uuid.Parse(resElem.RequestID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.TradeRequestModel{}, common.ResponseHeaders{}, err
	}
	tradeRequestModel.StatusModel.StatusID = StatusID

	tradeID, err := uuid.Parse(resElem.TradeID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.TradeRequestModel{}, common.ResponseHeaders{}, err
	}
	tradeRequestModel.TradeModel.TradeID = &tradeID
	tradeRequestModel.StatusModel.TradeID = tradeID

	return tradeRequestModel, headers, nil
}

// PutTradeResponse
// Summary: This is function which put trade response with PutTradeResponse.
// input: c(echo.Context) echo context
// input: putTradeResponseInput(traceability.PutTradeResponseInput) PutTradeResponseInput object
// output: (traceability.TradeModel) TradeModel object
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *tradeTraceabilityUsecase) PutTradeResponse(c echo.Context, putTradeResponseInput traceability.PutTradeResponseInput) (traceability.TradeModel, common.ResponseHeaders, error) {
	tradesRequest := traceabilityentity.PostTradesRequest{
		OperatorID: putTradeResponseInput.OperatorID.String(),
		TradeID:    putTradeResponseInput.TradeID.String(),
		TraceID:    putTradeResponseInput.TraceID.String(),
	}

	_, headers, err := u.TraceabilityRepository.PostTrades(c, tradesRequest)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceability.TradeModel{}, common.ResponseHeaders{}, err
	}

	// Obtain the response value using the receipt request information search
	tradeRequestsReceivedRequest := traceabilityentity.GetTradeRequestsReceivedRequest{
		OperatorID: putTradeResponseInput.OperatorID.String(),
	}
	tradeRequestsReceivedResponse, err := u.TraceabilityRepository.GetTradeRequestsReceived(c, tradeRequestsReceivedRequest)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return traceability.TradeModel{}, common.ResponseHeaders{}, err
	}

	tradeModel, err := tradeRequestsReceivedResponse.ExtractModelByTradeID(putTradeResponseInput.TradeID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.TradeModel{}, common.ResponseHeaders{}, err
	}

	if tradeModel == (traceability.TradeModel{}) {
		logger.Set(c).Errorf(common.NotFoundError("Trade"))

		return traceability.TradeModel{}, common.ResponseHeaders{}, fmt.Errorf(common.NotFoundError("Trade"))
	}

	// NOTE: Since the response does not include UpstreamOperatorID, set the request information.
	tradeModel.UpstreamOperatorID = putTradeResponseInput.OperatorID

	// NOTE: Due to asynchronous processing, there may be cases where UpstreamTraceID is not set in Get immediately after registration, so set the request information.
	tradeModel.UpstreamTraceID = &putTradeResponseInput.TraceID

	return tradeModel, headers, nil
}
