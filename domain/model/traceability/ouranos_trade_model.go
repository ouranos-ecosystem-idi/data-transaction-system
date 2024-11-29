package traceability

import (
	"fmt"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TradeModel
// Summary: This is structure which defines TradeModel.
// Service: DataSpace
// Router: [GET] /api/v1/authInfo?dataTarget=tradeRequest
// Router: [PUT] /api/v1/authInfo?dataTarget=tradeResponse
// Usage: output
type TradeModel struct {
	TradeID              *uuid.UUID `json:"tradeId"`
	DownstreamOperatorID uuid.UUID  `json:"downstreamOperatorId"`
	UpstreamOperatorID   uuid.UUID  `json:"upstreamOperatorId"`
	DownstreamTraceID    uuid.UUID  `json:"downstreamTraceId"`
	UpstreamTraceID      *uuid.UUID `json:"upstreamTraceId"`
}

// TradeResponseModel
// Summary: This is structure which defines TradeResponseModel.
// Service: DataSpace
// Router: [GET] /api/v1/authInfo?dataTarget=tradeResponse
// Usage: output
type TradeResponseModel struct {
	StatusModel StatusModel `json:"statusModel"`
	TradeModel  TradeModel  `json:"tradeModel"`
	PartsModel  PartsModel  `json:"partsModel"`
}

// TradeEntityModel
// Summary: This is structure which defines TradeEntityModel.
// DBName: trade
type TradeEntityModel struct {
	TradeID              *uuid.UUID     `json:"tradeId" gorm:"type:uuid"`
	DownstreamOperatorID uuid.UUID      `json:"downstreamOperatorId" gorm:"type:uuid;not null"`
	UpstreamOperatorID   *uuid.UUID     `json:"upstreamOperatorId" gorm:"type:uuid;not null"`
	DownstreamTraceID    uuid.UUID      `json:"downstreamTraceId" gorm:"type:uuid;not null"`
	UpstreamTraceID      *uuid.UUID     `json:"upstreamTraceId" gorm:"type:uuid"`
	TradeDate            *string        `json:"tradeDate" gorm:"type:string"`
	DeletedAt            gorm.DeletedAt `json:"deletedAt"`
	CreatedAt            time.Time      `json:"createdAt" gorm:"<-:create "`
	CreatedUserID        string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt            time.Time      `json:"updatedAt"`
	UpdatedUserID        string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
}

// PutTradeRequestInput
// Summary: This is structure which defines PutTradeRequestInput.
// Service: DataSpace
// Router: [PUT] /api/v1/authInfo?dataTarget=tradeRequest
// Usage: input
type PutTradeRequestInput struct {
	Trade  PutTradeInput  `json:"tradeModel"`
	Status PutStatusInput `json:"statusModel"`
}

// PutTradeRequestInput
// Summary: This is structure which defines PutTradeInput.
// Service: DataSpace
// Router: [PUT] /api/v1/authInfo?dataTarget=tradeRequest
// Usage: input
type PutTradeInput struct {
	TradeID              *string `json:"tradeId"`
	DownstreamOperatorID string  `json:"downstreamOperatorId"`
	UpstreamOperatorID   string  `json:"upstreamOperatorId"`
	DownstreamTraceID    string  `json:"downstreamTraceId"`
}

// TradeRequestModel
// Summary: This is structure which defines PutTradeInput.
// Service: DataSpace
// Router: [PUT] /api/v1/authInfo?dataTarget=tradeRequest
// Usage: input
type TradeRequestModel struct {
	TradeModel  TradeModel  `json:"tradeModel" gorm:"-"`
	StatusModel StatusModel `json:"statusModel" gorm:"-"`
}

// TradeRequestEntityModel
// Summary: This is structure which defines TradeRequestEntityModel.
type TradeRequestEntityModel struct {
	TradeEntityModel  TradeEntityModel  `json:"tradeEntityModel" gorm:"-"`
	StatusEntityModel StatusEntityModel `json:"statusEntityModel" gorm:"-"`
}

// PutTradeResponseInput
// Summary: This is structure which defines PutTradeInput.
// Service: DataSpace
// Router: [PUT] /api/v1/authInfo?dataTarget=tradeResponse
// Usage: input
type PutTradeResponseInput struct {
	OperatorID uuid.UUID `json:"operatorId"`
	TradeID    uuid.UUID `json:"tradeId"`
	TraceID    uuid.UUID `json:"traceId"`
}

// TradeModels
// Summary: This is structure which defines TradeModels.
type TradeModels []TradeModel

// TradeEntityModels
// Summary: This is structure which defines TradeEntityModels.
type TradeEntityModels []TradeEntityModel

// GetTradeResponseInput
// Summary: This is structure which defines GetTradeResponseInput.
// Service: DataSpace
// Router: [GET] /api/v1/authInfo?dataTarget=tradeResponse
// Usage: input
type GetTradeResponseInput struct {
	OperatorID uuid.UUID
	Limit      int `json:"limit"`
	After      *uuid.UUID
}

// GetTradeRequestInput
// Summary: This is structure which defines GetTradeResponseInput.
// Service: DataSpace
// Router: [GET] /api/v1/authInfo?dataTarget=tradeRequest
// Usage: input
type GetTradeRequestInput struct {
	OperatorID uuid.UUID
	Limit      int `json:"limit"`
	After      *uuid.UUID
	TraceIDs   []uuid.UUID
}

// Validate
// Summary: This is function which validate value of PutTradeRequestInput.
// output: (error) error object
func (i PutTradeRequestInput) Validate() error {
	var errors []error
	if err := i.validate(); err != nil {
		logger.Set(nil).Warnf(err.Error())
		errors = append(errors, err)
	}

	// PutTradeInput.TradeID, PutStatusInput.StatusID, PutStatusInput.TradeID
	// Update if all values ​​are present, new registration if all values ​​are zero, error otherwise.
	if !(i.Trade.TradeID == nil && i.Status.TradeID == nil && i.Status.StatusID == nil) &&
		!(i.Trade.TradeID != nil && i.Status.TradeID != nil && i.Status.StatusID != nil) {
		logger.Set(nil).Warnf(common.NotHaveValuesError("tradeModel.tradeId", "statusModel.statusId", "statusModel.tradeId"))
		err := fmt.Errorf(common.NotHaveValuesError("tradeModel.tradeId", "statusModel.statusId", "statusModel.tradeId"))
		errors = append(errors, err)
	}

	// An error occurs if the TradeID specified in TradeModel and the TradeID specified in StatusModel do not match.
	// if the TradeID specified in TradeModel and the TradeID specified in StatusModel do not match, An error occurs.
	if i.Trade.TradeID != nil && i.Status.TradeID != nil {
		if *i.Trade.TradeID != *i.Status.TradeID {
			logger.Set(nil).Warnf(common.InconsistentError("tradeModel.tradeId", "statusModel.tradeId"))

			err := fmt.Errorf(common.InconsistentError("tradeModel.tradeId", "statusModel.tradeId"))
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		joinedErr := common.JoinErrors(errors)
		logger.Set(nil).Warnf(joinedErr.Error())
		return common.JoinErrors(errors)
	}

	return nil
}

// validate
// Summary: This is function which validate value of PutTradeRequestInput.
// output: (error) error object
func (i PutTradeRequestInput) validate() error {
	var errors []error
	if err := i.Trade.validate(); err != nil {
		errors = append(errors, err)
	}

	if err := i.Status.validate(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		joinedErr := common.JoinErrors(errors)
		logger.Set(nil).Warnf(joinedErr.Error())
		return common.JoinErrors(errors)
	}
	return nil
}

// ToModel
// Summary: This is function which convert PutTradeRequestInput to TradeRequestModel.
// output: (TradeRequestModel) TradeRequestModel object
func (i PutTradeRequestInput) ToModel() TradeRequestModel {
	var tradeID uuid.UUID
	if i.Trade.TradeID != nil {
		tradeID, _ = uuid.Parse(*i.Trade.TradeID)
	}
	downstreamOperatorID, _ := uuid.Parse(i.Trade.DownstreamOperatorID)
	upstreamOperatorID, _ := uuid.Parse(i.Trade.UpstreamOperatorID)
	downstreamTraceID, _ := uuid.Parse(i.Trade.DownstreamTraceID)
	tradeModel := TradeModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: downstreamOperatorID,
		UpstreamOperatorID:   upstreamOperatorID,
		DownstreamTraceID:    downstreamTraceID,
	}

	var statusID uuid.UUID
	if i.Status.StatusID != nil {
		statusID, _ = uuid.Parse(*i.Status.StatusID)
	}
	requestStatus := RequestStatus{
		CfpResponseStatus: i.Status.PutRequestStatusInput.CfpResponseStatus,
		TradeTreeStatus:   i.Status.PutRequestStatusInput.TradeTreeStatus,
	}
	statusModel := StatusModel{
		StatusID:        statusID,
		TradeID:         tradeID,
		Message:         i.Status.Message,
		ReplyMessage:    i.Status.ReplyMessage,
		RequestStatus:   requestStatus,
		RequestType:     i.Status.RequestType.ToString(),
		ResponseDueDate: common.StringPtr(i.Status.ResponseDueDate),
	}

	return TradeRequestModel{
		TradeModel:  tradeModel,
		StatusModel: statusModel,
	}
}

// validate
// Summary: This is function which validate value of PutTradeInput.
// output: (error) error object
func (i PutTradeInput) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.TradeID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.DownstreamOperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.UpstreamOperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.DownstreamTraceID,
			validation.By(common.StringUUIDValid),
		),
	)
}

// ToModel
// Summary: This is function which convert TradeEntityModel to TradeModel.
// output: (TradeModel) TradeModel object
func (e TradeEntityModel) ToModel() TradeModel {
	return TradeModel{
		TradeID:              e.TradeID,
		DownstreamOperatorID: e.DownstreamOperatorID,
		UpstreamOperatorID:   *e.UpstreamOperatorID,
		DownstreamTraceID:    e.DownstreamTraceID,
		UpstreamTraceID:      e.UpstreamTraceID,
	}
}

// ToModels
// Summary: This is function which convert TradeEntityModels to array of TradeModel.
// output: ([]TradeModel) list of TradeModel
func (es TradeEntityModels) ToModels() []TradeModel {
	rs := []TradeModel{}
	for _, e := range es {
		rs = append(rs, e.ToModel())
	}
	return rs
}

// ToModel
// Summary: This is function which convert TradeRequestEntityModel to array of TradeRequestModel.
// output: (TradeRequestModel) TradeRequestModel object
// output: (error) error object
func (e TradeRequestEntityModel) ToModel(DataSpacesApi string) (TradeRequestModel, error) {
	m, err := e.StatusEntityModel.ToModel(DataSpacesApi)
	if err != nil {
		return TradeRequestModel{}, err
	}
	return TradeRequestModel{
		TradeModel:  e.TradeEntityModel.ToModel(),
		StatusModel: m,
	}, nil
}
