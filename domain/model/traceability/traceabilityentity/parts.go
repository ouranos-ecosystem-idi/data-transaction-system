package traceabilityentity

import (
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
)

// GetPartsRequest
// Summary: This is structure which defines GetPartsRequest.
// Service: Traceability
// Router: [GET] /parts
// Usage: input
type GetPartsRequest struct {
	OperatorID       string  `json:"operatorId"`
	TraceID          *string `json:"traceId"`
	PartsItem        *string `json:"partsItem"`
	SupportPartsItem *string `json:"supportPartsItem"`
	PlantID          *string `json:"plantId"`
	ParentFlag       *bool   `json:"parentFlag"`
	After            *string `json:"after"`
}

// GetPartsResponse
// Summary: This is a type that defines a GetPartsResponse.
// Service: Traceability
// Router: [GET] /parts
// Usage: output
type GetPartsResponse struct {
	Parts []GetPartsResponseParts `json:"parts"`
	Next  string                  `json:"next"`
}

// GetPartsResponseParts
// Summary: This is structure which defines GetPartsResponseParts.
type GetPartsResponseParts struct {
	TraceID          string  `json:"traceId"`
	PartsItem        string  `json:"partsItem"`
	SupportPartsItem *string `json:"supportPartsItem"`
	PlantID          string  `json:"plantId"`
	OperatorID       string  `json:"operatorId"`
	AmountUnitName   *string `json:"amountUnitName"`
	EndFlag          bool    `json:"endFlag"`
	ParentFlag       bool    `json:"parentFlag"`
	PartsLabelName   *string `json:"partsLabelName"`
	PartsAddInfo1    *string `json:"partsAddInfo1"`
	PartsAddInfo2    *string `json:"partsAddInfo2"`
	PartsAddInfo3    *string `json:"partsAddInfo3"`
}

// DeletePartsRequest
// Summary: This is structure which defines DeletePartsRequest.
// Service: Traceability
// Router: [DELETE] /parts
// Usage: input
type DeletePartsRequest struct {
	OperatorID string `json:"operatorId"`
	TraceID    string `json:"traceId"`
}

// DeletePartsResponse
// Summary: This is a type that defines a DeletePartsResponse.
// Service: Traceability
// Router: [DELETE] /parts
// Usage: output
type DeletePartsResponse struct {
	TraceID string `json:"traceId"`
}

// ToModel
// Summary: This is function to convert GetPartsResponse to []traceability.PartsModel.
// output: ([]traceability.PartsModel) converted to []traceability.PartsModel
// output: (error) error object
func (r GetPartsResponse) ToModel() ([]traceability.PartsModel, error) {
	ms := make([]traceability.PartsModel, len(r.Parts))
	for i, p := range r.Parts {
		m, err := p.ToModel()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return []traceability.PartsModel{}, err
		}
		ms[i] = m
	}
	return ms, nil
}

// FindByTraceID
// Summary: This is function to returns results find by traceId.
// input: traceID(uuid.UUID) ID of the trace
// output: (*GetPartsResponseParts) parts find by traceId
func (r GetPartsResponse) FindByTraceID(traceID uuid.UUID) *GetPartsResponseParts {
	for _, p := range r.Parts {
		if p.TraceID == traceID.String() {
			return &p
		}
	}
	return nil
}

// ToModel
// Summary: This is function to convert GetPartsResponseParts to []traceability.PartsModel.
// output: (traceability.PartsModel) converted to traceability.PartsModel
// output: (error) error object
func (r GetPartsResponseParts) ToModel() (traceability.PartsModel, error) {
	var err error

	m := traceability.PartsModel{
		PartsName:        r.PartsItem,
		SupportPartsName: r.SupportPartsItem,
		TerminatedFlag:   r.EndFlag,
		PartsLabelName:   r.PartsLabelName,
		PartsAddInfo1:    r.PartsAddInfo1,
		PartsAddInfo2:    r.PartsAddInfo2,
		PartsAddInfo3:    r.PartsAddInfo3,
	}

	var amountRequiredUnit traceability.AmountRequiredUnit
	if r.AmountUnitName != nil {
		amountRequiredUnit, err = traceability.NewAmountRequiredUnit(*r.AmountUnitName)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return traceability.PartsModel{}, err
		}
		m.AmountRequiredUnit = &amountRequiredUnit
	}

	traceID, err := uuid.Parse(r.TraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	m.TraceID = traceID

	operatorID, err := uuid.Parse(r.OperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	m.OperatorID = operatorID

	plantID, err := uuid.Parse(r.PlantID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	m.PlantID = &plantID

	return m, nil
}

// NextPrt
// Summary: This is function to convert GetPartsResponse to *string.
// output: (*string) error object
func (r GetPartsResponse) NextPrt() *string {
	if r.Next == "" {
		return nil
	}
	return &r.Next
}
