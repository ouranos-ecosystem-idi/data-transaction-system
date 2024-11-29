package traceability

import (
	"encoding/json"
	"fmt"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StatusModel
// Summary: This is structure which defines StatusModel.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=status
// Usage: output
type StatusModel struct {
	StatusID        uuid.UUID     `json:"statusId"`
	TradeID         uuid.UUID     `json:"tradeId"`
	RequestStatus   RequestStatus `json:"requestStatus"`
	Message         *string       `json:"message"`
	ReplyMessage    *string       `json:"replyMessage"`
	RequestType     string        `json:"requestType"`
	ResponseDueDate *string       `json:"responseDueDate"`
}

// StatusModelForSort
// Summary: This is structure which defines StatusModelForSort.
type StatusModelForSort struct {
	StatusModel StatusModel
	RequestedAt time.Time
}

// StatusModelsForSort
// Summary: This is a type that defines a list of StatusModelForSort.
type StatusModelsForSort []StatusModelForSort

// SortByRequestedAt
// Summary: This function sorts the StatusModel in descending order of requested time.
// output: (StatusModelsForSort) sorted StatusModelsForSort
func (ms StatusModelsForSort) SortByRequestedAt() StatusModelsForSort {
	for i := 0; i < len(ms); i++ {
		for j := i + 1; j < len(ms); j++ {
			if ms[i].RequestedAt.Before(ms[j].RequestedAt) {
				ms[i], ms[j] = ms[j], ms[i]
			}
		}
	}

	return ms
}

// FilterByStatusID
// Summary: This is the function to filter StatusModel by statusID.
// input: statusID(uuid.UUID) statusID to filter
// output: (StatusModelsForSort) filtered StatusModelsForSort
func (ms StatusModelsForSort) FilterByStatusID(statusID uuid.UUID) StatusModelsForSort {
	r := []StatusModelForSort{}
	for _, m := range ms {
		if m.StatusModel.StatusID == statusID {
			r = append(r, m)
		}
	}
	return r
}

// GetStatusModels
// Summary: This is the function to get StatusModels.
// input: getStatusInput(GetStatusInput) GetStatusInput object
// output: ([]StatusModel) list of StatusModel
// output: (*string) next id
func (ms StatusModelsForSort) GetStatusModels(getStatusInput GetStatusInput) ([]StatusModel, *string) {
	afterIndex := 0
	if getStatusInput.After != nil {
		for i, m := range ms {
			if m.StatusModel.StatusID.String() == getStatusInput.After.String() {
				afterIndex = i
				break
			}
		}
	}

	var after *string
	nextIndex := afterIndex + getStatusInput.Limit
	if len(ms) > nextIndex {
		after = common.StringPtr(ms[nextIndex].StatusModel.StatusID.String())
	}

	lastIndex := afterIndex + getStatusInput.Limit
	if len(ms) >= lastIndex {
		ms = ms[afterIndex:lastIndex]
	} else {
		ms = ms[afterIndex:]
	}

	return ms.ToStatusModels(), after
}

// ToStatusModels
// Summary: This is the function to convert StatusModelsForSort to StatusModels.
// output: ([]StatusModel) list of StatusModel
func (ms StatusModelsForSort) ToStatusModels() []StatusModel {
	r := []StatusModel{}
	for _, m := range ms {
		r = append(r, m.StatusModel)
	}
	return r
}

// RequestStatus
// Summary: This is structure which defines RequestStatus.
type RequestStatus struct {
	CfpResponseStatus        *CfpResponseStatus `json:"cfpResponseStatus,omitempty"`
	TradeTreeStatus          *TradeTreeStatus   `json:"tradeTreeStatus,omitempty"`
	CompletedCount           *int               `json:"completedCount"`
	CompletedCountModifiedAt *string            `json:"completedCountModifiedAt"`
	TradesCount              *int               `json:"tradesCount"`
	TradesCountModifiedAt    *string            `json:"tradesCountModifiedAt"`
}

// ToString
// Summary: This is the function to convert RequestStatus to string.
// output: (string) converted to string
func (r RequestStatus) ToString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// StatusEntityModel
// Summary: This is structure which defines StatusEntityModel.
// DBName: request_status
type StatusEntityModel struct {
	StatusID                 uuid.UUID      `json:"statusId" gorm:"type:uuid" validate:"required" example:"d9a38406-cae2-4679-b052-15a75f5531f0"`
	TradeID                  uuid.UUID      `json:"tradeId" gorm:"type:uuid;not null" validate:"required" example:"d9a38406-cae2-4679-b052-15a75f5531f1"`
	CfpResponseStatus        string         `json:"cfpResponseStatus" gorm:"type:string" validate:"required"`
	TradeTreeStatus          string         `json:"tradeTreeStatus" gorm:"type:string" validate:"required"`
	Message                  *string        `json:"message" gorm:"type:string" validate:"required" example:"回答依頼時のメッセージが入ります" maxLength:"1000"`
	ReplyMessage             *string        `json:"replyMessage" gorm:"type:string" validate:"required" example:"回答時のメッセージが入ります" maxLength:"1000"`
	RequestType              string         `json:"requestType" gorm:"type:string" validate:"required" example:"batteryRequestStatus" maxLength:"256"`
	ResponseDueDate          string         `json:"responseDueDate" gorm:"type:string" example:"2024-01-01" maxLength:"10"`
	CompletedCount           *int           `json:"completedCount"`
	CompletedCountModifiedAt *time.Time     `json:"completedCountModifiedAt"`
	TradesCount              *int           `json:"tradesCount"`
	TradesCountModifiedAt    *time.Time     `json:"tradesCountModifiedAt"`
	DeletedAt                gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`
	CreatedAt                time.Time      `json:"createdAt" gorm:"<-:create " swaggerignore:"true"`
	CreatedUserId            string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create" swaggerignore:"true"`
	UpdatedAt                time.Time      `json:"updatedAt" swaggerignore:"true" `
	UpdatedUserId            string         `json:"updatedUserId" gorm:"type:varchar(256);not null" swaggerignore:"true"`
}

// StatusModels
// Summary: This is a type that defines a list of StatusModel.
type StatusModels []StatusModel

// StatusEntityModels
// Summary: This is a type that defines a list of StatusEntityModel.
type StatusEntityModels []StatusEntityModel

// StatusTarget
// Summary: This is enum which defines StatusTarget.
type StatusTarget string

const (
	Request  StatusTarget = "REQUEST"
	Response StatusTarget = "RESPONSE"
)

// ToString
// Summary: This is the function to convert StatusTarget to string.
// output: (string) converted to string
func (e StatusTarget) ToString() string {
	return string(e)
}

// NewStatusTarget
// Summary: This is the function to create new StatusTarget.
// input: s(string) StatusTarget string
// output: (StatusTarget) StatusTarget
// output: (error) error object
func NewStatusTarget(s string) (StatusTarget, error) {
	switch s {
	case Request.ToString():
		return Request, nil
	case Response.ToString():
		return Response, nil
	case "":
		return StatusTarget(""), nil
	default:
		return StatusTarget(""), fmt.Errorf(common.UnexpectedEnumError("StatusTarget", s))
	}
}

// GetStatusInput
// Summary: This is structure which defines GetStatusInput.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=status
// Usage: input
// NOTE: Give json tags only for parameters required for the link on the next page.
type GetStatusInput struct {
	OperatorID   uuid.UUID
	Limit        int `json:"limit"`
	After        *uuid.UUID
	StatusTarget StatusTarget `json:"statusTarget"`
	StatusID     *uuid.UUID
	TraceID      *uuid.UUID
}

// RequestType
// Summary: This is enum which defines RequestType.
type RequestType string

// ToString
// Summary: This is the function to convert RequestType to string.
// output: (string) converted to string
func (e RequestType) ToString() string {
	return string(e)
}

// CfResponseStatus
// Summary: This is enum which defines CfpResponseStatus.
type CfpResponseStatus string

// ToString
// Summary: This is the function to convert CfpResponseStatus to string.
// output: (string) converted to string
func (e CfpResponseStatus) ToString() string {
	return string(e)
}

// NewCfpResponseStatus
// Summary: This is the function to create new CfpResponseStatus.
// input: s(string) CfpResponseStatus string
// output: (CfpResponseStatus) CfpResponseStatus
// output: (error) error object
func NewCfpResponseStatus(s string) (CfpResponseStatus, error) {
	switch s {
	case CfpResponseStatusPending.ToString():
		return CfpResponseStatusPending, nil
	case CfpResponseStatusComplete.ToString():
		return CfpResponseStatusComplete, nil
	case CfpResponseStatusReject.ToString():
		return CfpResponseStatusReject, nil
	case CfpResponseStatusCancel.ToString():
		return CfpResponseStatusCancel, nil
	default:
		return CfpResponseStatusUnknown, fmt.Errorf(common.UnexpectedEnumError("CfpResponseStatus", s))
	}
}

// TradeTreeStatus
// Summary: This is enum which defines TradeTreeStatus.
type TradeTreeStatus string

// ToString
// Summary: This is the function to convert TradeTreeStatus to string.
// output: (string) converted to string
func (e TradeTreeStatus) ToString() string {
	return string(e)
}

// NewTradeTreeStatus
// Summary: This is the function to create new TradeTreeStatus.
// input: s(string) TradeTreeStatus string
// output: (TradeTreeStatus) TradeTreeStatus
// output: (error) error object
func NewTradeTreeStatus(s string) (TradeTreeStatus, error) {
	switch s {
	case TradeTreeStatusTerminated.ToString():
		return TradeTreeStatusTerminated, nil
	case TradeTreeStatusUnterminated.ToString():
		return TradeTreeStatusUnterminated, nil
	default:
		return TradeTreeStatusUnknown, fmt.Errorf(common.UnexpectedEnumError("TradeTreeStatus", s))
	}
}

const (
	CfpResponseStatusPending  CfpResponseStatus = "NOT_COMPLETED"
	CfpResponseStatusComplete CfpResponseStatus = "COMPLETED"
	CfpResponseStatusReject   CfpResponseStatus = "REJECT"
	CfpResponseStatusCancel   CfpResponseStatus = "CANCEL"
	CfpResponseStatusUnknown  CfpResponseStatus = "unknown"

	TradeTreeStatusUnterminated TradeTreeStatus = "UNTERMINATED"
	TradeTreeStatusTerminated   TradeTreeStatus = "TERMINATED"
	TradeTreeStatusUnknown      TradeTreeStatus = "unknown"

	RequestTypeCFP RequestType = "CFP"
)

// PutStatusInput
// Summary: This is structure which defines PutStatusInput.
// Service: Dataspace
// Router: [PUT] /api/v1/datatransport?dataTarget=status
// Usage: input
type PutStatusInput struct {
	StatusID              *string               `json:"statusId"`
	TradeID               *string               `json:"tradeId"`
	RequestType           RequestType           `json:"requestType"`
	Message               *string               `json:"message"`
	ReplyMessage          *string               `json:"replyMessage"`
	ResponseDueDate       string                `json:"responseDueDate"`
	PutRequestStatusInput PutRequestStatusInput `json:"requestStatus"`
}

// PutRequestStatusInput
// Summary: This is structure which defines PutRequestStatusInput.
type PutRequestStatusInput struct {
	CfpResponseStatus *CfpResponseStatus `json:"cfpResponseStatus"`
	TradeTreeStatus   *TradeTreeStatus   `json:"tradeTreeStatus"`
}

// ValidateForCancelOrReject
// Summary: This is the function to validate PutStatusInput in the event of cancellation or rejection.
// output: (error) error object
func (i PutStatusInput) ValidateForCancelOrReject() error {
	if err := i.validateForCancelOrReject(); err != nil {
		logger.Set(nil).Errorf(err.Error())

		return err
	}

	return nil
}

// validateForCancelOrReject
// Summary: This is the function to validate PutStatusInput in the event of cancellation or rejection.
// output: (error) error object
func (i PutStatusInput) validateForCancelOrReject() error {
	errors := []error{}
	err := validation.ValidateStruct(&i,
		validation.Field(
			&i.StatusID,
			validation.Required,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.TradeID,
			validation.Required,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.ReplyMessage,
			validation.RuneLength(0, 1000),
		),
	)
	if err != nil {
		errors = append(errors, err)
	}

	var requestStatusErr error
	cfpResponseStatusStr := i.PutRequestStatusInput.CfpResponseStatus.ToString()
	cfpResponseStatus, _ := NewCfpResponseStatus(cfpResponseStatusStr)
	if cfpResponseStatus != CfpResponseStatusCancel && cfpResponseStatus != CfpResponseStatusReject {
		requestStatusErr = fmt.Errorf("requestStatus: (cfpResponseStatus: %v)", fmt.Errorf(common.InvalidEnumError(cfpResponseStatusStr)))
		errors = append(errors, requestStatusErr)
	}

	if len(errors) > 0 {
		if requestStatusErr != nil {
			return common.JoinErrors(errors)
		} else {
			return err
		}
	}

	return nil
}

// validate
// Summary: This is the function to validate PutStatusInput.
// output: (error) error object
func (i PutStatusInput) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.StatusID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.TradeID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.Message,
			validation.RuneLength(0, 1000),
		),
		validation.Field(
			&i.RequestType,
			validation.Required,
			validation.By(EnumRequestTypeValid),
		),
		validation.Field(
			&i.ResponseDueDate,
			validation.Required,
			validation.Date("2006-01-02"),
		),
	)
}

// ToModel
// Summary: This is the function to convert PutStatusInput to StatusModel.
// output: (StatusModel) converted StatusModel
func (i PutStatusInput) ToModel() (StatusModel, error) {
	var err error
	var statusID uuid.UUID
	if i.StatusID != nil {
		statusID, err = uuid.Parse(*i.StatusID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return StatusModel{}, err
		}
	}
	var tradeId uuid.UUID
	if i.TradeID != nil {
		tradeId, err = uuid.Parse(*i.StatusID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return StatusModel{}, err
		}

	}
	statusModel := StatusModel{
		StatusID:     statusID,
		TradeID:      tradeId,
		Message:      i.Message,
		ReplyMessage: i.ReplyMessage,
		RequestType:  i.RequestType.ToString(),
		RequestStatus: RequestStatus{
			CfpResponseStatus: i.PutRequestStatusInput.CfpResponseStatus,
			TradeTreeStatus:   i.PutRequestStatusInput.TradeTreeStatus,
		},
	}
	return statusModel, nil

}

// IsCfpRequestStatusCancel
// Summary: This is the function to check if the CfpResponseStatus is cancel.
// output: (bool) true if CfpResponseStatus is cancel
func (i PutStatusInput) IsCfpRequestStatusCancel() bool {
	return *i.PutRequestStatusInput.CfpResponseStatus == CfpResponseStatusCancel
}

// IsCfpRequestStatusReject
// Summary: This is the function to check if the CfpResponseStatus is reject.
// output: (bool) true if CfpResponseStatus is reject
func (i PutStatusInput) IsCfpRequestStatusReject() bool {
	return *i.PutRequestStatusInput.CfpResponseStatus == CfpResponseStatusReject
}

const (
	PathTradeRequest  = "tradeRequest"
	PathTradeResponse = "tradeResponse"
	PathStatus        = "status"
)

// ToModel
// Summary: This is the function to convert StatusEntityModel to StatusModel.
// output: (StatusModel) converted StatusModel
func (e StatusEntityModel) ToModel(DataSpacesApi string) (StatusModel, error) {
	cfpResponseStatus, err := NewCfpResponseStatus(e.CfpResponseStatus)
	if err != nil {
		return StatusModel{}, err
	}
	tradeTreeStatus, err := NewTradeTreeStatus(e.TradeTreeStatus)
	if err != nil {
		return StatusModel{}, err
	}
	var cc *string
	if e.CompletedCountModifiedAt != nil {
		cc = common.StringPtr(common.GenerateTime(*e.CompletedCountModifiedAt))
	}
	var tc *string
	if e.TradesCountModifiedAt != nil {
		tc = common.StringPtr(common.GenerateTime(*e.TradesCountModifiedAt))
	}
	var responseDueDate string
	if len(e.ResponseDueDate) > 10 {
		responseDueDate = e.ResponseDueDate[:10]
	} else {
		responseDueDate = e.ResponseDueDate
	}
	var requestStatus RequestStatus
	if DataSpacesApi == PathTradeRequest {
		requestStatus = RequestStatus{
			CompletedCount:           nil,
			CompletedCountModifiedAt: nil,
			TradesCount:              nil,
			TradesCountModifiedAt:    nil,
		}
	} else {
		requestStatus = RequestStatus{
			CfpResponseStatus:        &cfpResponseStatus,
			TradeTreeStatus:          &tradeTreeStatus,
			CompletedCount:           e.CompletedCount,
			CompletedCountModifiedAt: cc,
			TradesCount:              e.TradesCount,
			TradesCountModifiedAt:    tc,
		}
	}
	return StatusModel{
		StatusID:        e.StatusID,
		TradeID:         e.TradeID,
		RequestStatus:   requestStatus,
		Message:         e.Message,
		ReplyMessage:    e.ReplyMessage,
		RequestType:     e.RequestType,
		ResponseDueDate: &responseDueDate,
	}, nil
}

// ToModels
// Summary: This is the function to convert StatusEntityModels to StatusModels.
// output: ([]StatusModel) list of StatusModel
// output: (error) error object
func (es StatusEntityModels) ToModels(DataSpacesApi string) ([]StatusModel, error) {
	ms := make(StatusModels, len(es))
	for i, e := range es {
		m, err := e.ToModel(DataSpacesApi)
		if err != nil {
			return nil, err
		}
		ms[i] = m
	}
	return ms, nil
}
