package traceabilityentity

import (
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
)

// PostTradeRequestsCancelRequest
// Summary: This is structure which defines PostTradeRequestsCancelRequest.
// Service: Traceability
// Router: [POST] /tradeRequests/cancel
// Usage: input
type PostTradeRequestsCancelRequest struct {
	OperatorID     string                                        `json:"operatorId"`
	CancelRequests []PostTradeRequestsCancelRequestCancelRequest `json:"cancelRequests"`
}

// PostTradeRequestsCancelRequestCancelRequest
// Summary: This is structure which defines PostTradeRequestsCancelRequestCancelRequest.
type PostTradeRequestsCancelRequestCancelRequest struct {
	RequestID string `json:"requestId"`
}

// NewPostTradeRequestsCancelRequest
// Summary: This is function to create new PostTradeRequestsCancelRequest.
// input: operatorID(string) ID of the operator
// input: statusID(string) ID of the status
// output: (PostTradeRequestsCancelRequest) PostTradeRequestsCancelRequest object
func NewPostTradeRequestsCancelRequest(operatorID string, statusID string) PostTradeRequestsCancelRequest {
	requests := PostTradeRequestsCancelRequest{
		OperatorID: operatorID,
		CancelRequests: []PostTradeRequestsCancelRequestCancelRequest{
			{
				RequestID: statusID,
			},
		},
	}

	return requests
}

// PostTradeRequestsRejectRequest
// Summary: This is structure which defines PostTradeRequestsRejectRequest.
// Service: Traceability
// Router: [POST] /tradeRequests/reject
// Usage: input
type PostTradeRequestsRejectRequest struct {
	OperatorID     string              `json:"operatorId"`
	RejectRequests []PostRejectRequest `json:"rejectRequests"`
}

// PostRejectRequest
// Summary: This is structure which defines PostRejectRequest.
type PostRejectRequest struct {
	RequestID    string  `json:"requestId"`
	ReplyMessage *string `json:"replyMessage"`
}

// NewPostTradeRequestsRejectRequest
// Summary: This is function to create new PostTradeRequestsRejectRequest.
// input: operatorID(string) ID of the operator
// input: statusID(string) ID of the status
// input: replyMessage(*string) reply message
// output: (PostTradeRequestsRejectRequest) PostTradeRequestsRejectRequest object
func NewPostTradeRequestsRejectRequest(operatorID string, statusID string, replyMessage *string) PostTradeRequestsRejectRequest {
	requests := PostTradeRequestsRejectRequest{
		OperatorID: operatorID,
		RejectRequests: []PostRejectRequest{
			{
				RequestID:    statusID,
				ReplyMessage: replyMessage,
			},
		},
	}

	return requests
}

// PostTradeRequestsCancelResponse
// Summary: This is a type that defines a list of PostTradeRequestsCancelResponse.
// Service: Traceability
// Router: [POST] /tradeRequests/cancel
// Usage: output
type PostTradeRequestsCancelResponse []PostTradeRequestsCancelResponseCancelRequests

// PostTradeRequestsCancelResponse
// Summary: This is structure which defines PostTradeRequestsCancelResponse.
type PostTradeRequestsCancelResponseCancelRequests struct {
	RequestID string `json:"requestId"`
	TradeID   string `json:"tradeId"`
}

// ToModel
// Summary: This is function to convert PostTradeRequestsCancelResponse to traceability.StatusModel.
// output: (traceability.StatusModel) StatusModel object
// output: (error) error object
func (r PostTradeRequestsCancelResponse) ToModel() (traceability.StatusModel, error) {
	response := r[0]
	requestID, err := uuid.Parse(response.RequestID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, err
	}
	tradeID, err := uuid.Parse(response.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, err
	}

	model := traceability.StatusModel{
		StatusID: requestID,
		TradeID:  tradeID,
	}

	return model, nil
}

// PostTradeRequestsRejectResponse
// Summary: This is a type that defines a list of PostTradeRequestsRejectResponseRejectRequests.
// Service: Traceability
// Router: [POST] /tradeRequests/reject
// Usage: output
type PostTradeRequestsRejectResponse []PostTradeRequestsRejectResponseRejectRequests

// PostTradeRequestsRejectResponseRejectRequests
// Summary: This is structure which defines PostTradeRequestsRejectResponseRejectRequests.
type PostTradeRequestsRejectResponseRejectRequests struct {
	RequestID string `json:"requestId"`
	TradeID   string `json:"tradeId"`
}

// ToModel
// Summary: This is function to convert PostTradeRequestsRejectResponse to traceability.StatusModel.
// output: (traceability.StatusModel) StatusModel object
// output: (error) error object
func (r PostTradeRequestsRejectResponse) ToModel() (traceability.StatusModel, error) {
	response := r[0]
	requestID, err := uuid.Parse(response.RequestID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, err
	}
	tradeID, err := uuid.Parse(response.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, err
	}

	model := traceability.StatusModel{
		StatusID: requestID,
		TradeID:  tradeID,
	}

	return model, nil
}
