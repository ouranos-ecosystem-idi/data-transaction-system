package traceabilityentity

import (
	"fmt"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
)

// GetPartsStructuresRequest
// Summary: This is structure which defines GetPartsStructuresRequest.
// Service: Traceability
// Router: [GET] /partsStructure
// Usage: input
type GetPartsStructuresRequest struct {
	OperatorID    string `json:"operatorId"`
	ParentTraceID string `json:"parentTraceId"`
}

// GetPartsStructuresResponse
// Summary: This is structure which defines GetPartsStructuresResponse.
// Service: Traceability
// Router: [GET] /partsStructure
// Usage: output
type GetPartsStructuresResponse struct {
	Parent   *GetPartsStructuresResponseParent    `json:"parent"`
	Children []GetPartsStructuresResponseChildren `json:"children"`
}

// GetPartsStructuresResponseParent
// Summary: Define the main items of the component parts.
type GetPartsStructuresResponseParent struct {
	TraceID          string  `json:"traceId"`
	PartsItem        string  `json:"partsItem"`
	SupportPartsItem *string `json:"supportPartsItem"`
	PlantID          string  `json:"plantId"`
	OperatorID       string  `json:"operatorId"`
	AmountUnitName   *string `json:"amountUnitName"`
	EndFlag          bool    `json:"endFlag"`
	PartsLabelName   *string `json:"partsLabelName"`
	PartsAddInfo1    *string `json:"partsAddInfo1"`
	PartsAddInfo2    *string `json:"partsAddInfo2"`
	PartsAddInfo3    *string `json:"partsAddInfo3"`
}

// GetPartsStructuresResponseChildren
// Summary: Define sub-items of the component parts.
type GetPartsStructuresResponseChildren struct {
	PartsStructureID string   `json:"partsStructureId"`
	TraceID          string   `json:"traceId"`
	PartsItem        string   `json:"partsItem"`
	SupportPartsItem *string  `json:"supportPartsItem"`
	PlantID          string   `json:"plantId"`
	OperatorID       string   `json:"operatorId"`
	AmountUnitName   *string  `json:"amountUnitName"`
	EndFlag          bool     `json:"endFlag"`
	Amount           *float64 `json:"amount"`
	Revision         int      `json:"revision"`
	PartsLabelName   *string  `json:"partsLabelName"`
	PartsAddInfo1    *string  `json:"partsAddInfo1"`
	PartsAddInfo2    *string  `json:"partsAddInfo2"`
	PartsAddInfo3    *string  `json:"partsAddInfo3"`
}

// ToModel
// Summary: This is function to convert GetPartsStructuresResponse to traceability.PartsStructureModel.
// output: m(traceability.PartsStructureModel) partsStructure model
// output: err(error) error object
func (r GetPartsStructuresResponse) ToModel() (m traceability.PartsStructureModel, err error) {
	if r.Parent == nil {
		m.ChildrenPartsModel = []traceability.PartsModel{}
		return m, nil
	}

	parentPartsModel, err := r.Parent.ToModel()
	if err != nil {
		return m, err
	}
	m.ParentPartsModel = &parentPartsModel

	m.ChildrenPartsModel = []traceability.PartsModel{}
	for _, c := range r.Children {
		child, err := c.ToModel()
		if err != nil {
			return m, err
		}
		m.ChildrenPartsModel = append(m.ChildrenPartsModel, child)
	}

	return m, nil
}

// ToModel
// Summary: This is function to convert GetPartsStructuresResponseParent to traceability.PartsModel.
// output: (traceability.PartsModel) parts model
// output: (error) error object
func (r GetPartsStructuresResponseParent) ToModel() (traceability.PartsModel, error) {
	traceID, err := uuid.Parse(r.TraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	plantID, err := uuid.Parse(r.PlantID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	operatorID, err := uuid.Parse(r.OperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}

	model := traceability.PartsModel{
		TraceID:          traceID,
		PartsName:        r.PartsItem,
		SupportPartsName: r.SupportPartsItem,
		PlantID:          &plantID,
		OperatorID:       operatorID,
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
		model.AmountRequiredUnit = &amountRequiredUnit
	}

	return model, nil
}

// ToModel
// Summary: This is function to convert GetPartsStructuresResponseChildren to traceability.PartsModel.
// output: (traceability.PartsModel) parts model
// output: (error) error object
func (r GetPartsStructuresResponseChildren) ToModel() (traceability.PartsModel, error) {
	traceID, err := uuid.Parse(r.TraceID)
	if err != nil {
		return traceability.PartsModel{}, err
	}

	plantID, err := uuid.Parse(r.PlantID)
	if err != nil {
		return traceability.PartsModel{}, err
	}

	operatorID, err := uuid.Parse(r.OperatorID)
	if err != nil {
		return traceability.PartsModel{}, err
	}

	var amountUnitName traceability.AmountRequiredUnit
	if r.AmountUnitName != nil {
		amountUnitName, err = traceability.NewAmountRequiredUnit(*r.AmountUnitName)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return traceability.PartsModel{}, err
		}
	}

	return traceability.PartsModel{
		TraceID:            traceID,
		OperatorID:         operatorID,
		PlantID:            &plantID,
		PartsName:          r.PartsItem,
		SupportPartsName:   r.SupportPartsItem,
		TerminatedFlag:     r.EndFlag,
		AmountRequired:     r.Amount,
		AmountRequiredUnit: &amountUnitName,
		PartsLabelName:     r.PartsLabelName,
		PartsAddInfo1:      r.PartsAddInfo1,
		PartsAddInfo2:      r.PartsAddInfo2,
		PartsAddInfo3:      r.PartsAddInfo3,
	}, nil
}

// PostPartsStructuresRequest
// Summary: This is structure which defines PostPartsStructuresRequest.
// Service: Traceability
// Router: [POST] /partsStructure
// Usage: input
type PostPartsStructuresRequest struct {
	OperatorID string                            `json:"operatorId"`
	Parent     PostPartsStructuresRequestParent  `json:"parent"`
	Children   []PostPartsStructuresRequestChild `json:"children"`
}

// PostPartsStructuresRequestParent
// Summary: Define the main items of the component parts.
type PostPartsStructuresRequestParent struct {
	PartsItem        string  `json:"partsItem"`
	SupportPartsItem *string `json:"supportPartsItem"`
	PlantID          string  `json:"plantId"`
	OperatorID       string  `json:"operatorId"`
	AmountUnitName   *string `json:"amountUnitName"`
	EndFlag          bool    `json:"endFlag"`
	PartsLabelName   *string `json:"partsLabelName"`
	PartsAddInfo1    *string `json:"partsAddInfo1"`
	PartsAddInfo2    *string `json:"partsAddInfo2"`
	PartsAddInfo3    *string `json:"partsAddInfo3"`
}

// PostPartsStructuresRequestChild
// Summary: Define sub-items of the component parts.
type PostPartsStructuresRequestChild struct {
	PartsItem        string   `json:"partsItem"`
	SupportPartsItem *string  `json:"supportPartsItem"`
	PlantID          string   `json:"plantId"`
	OperatorID       string   `json:"operatorId"`
	AmountUnitName   *string  `json:"amountUnitName"`
	EndFlag          bool     `json:"endFlag"`
	Amount           *float64 `json:"amount"`
	PartsLabelName   *string  `json:"partsLabelName"`
	PartsAddInfo1    *string  `json:"partsAddInfo1"`
	PartsAddInfo2    *string  `json:"partsAddInfo2"`
	PartsAddInfo3    *string  `json:"partsAddInfo3"`
}

// PostPartsStructuresResponse
// Summary: This is structure which defines PostPartsStructuresResponse.
// Service: Traceability
// Router: [POST] /partsStructure
// Usage: output
type PostPartsStructuresResponse struct {
	Parent   PostPartsStructuresResponseParent  `json:"parent"`
	Children []PostPartsStructuresResponseChild `json:"children"`
}

// PostPartsStructuresResponseParent
// Summary: Define the main items of the component parts.
type PostPartsStructuresResponseParent struct {
	TraceID          string  `json:"traceId"`
	PartsItem        string  `json:"partsItem"`
	SupportPartsItem *string `json:"supportPartsItem"`
	PartsLabelName   *string `json:"partsLabelName"`
	PartsAddInfo1    *string `json:"partsAddInfo1"`
	PartsAddInfo2    *string `json:"partsAddInfo2"`
	PartsAddInfo3    *string `json:"partsAddInfo3"`
}

// PostPartsStructuresResponseChild
// Summary: Define sub-items of the component parts.
type PostPartsStructuresResponseChild struct {
	PartsStructureID string  `json:"partsStructureId"`
	TraceID          string  `json:"traceId"`
	PartsItem        string  `json:"partsItem"`
	PlantID          string  `json:"plantId"`
	SupportPartsItem *string `json:"supportPartsItem"`
	PartsLabelName   *string `json:"partsLabelName"`
	PartsAddInfo1    *string `json:"partsAddInfo1"`
	PartsAddInfo2    *string `json:"partsAddInfo2"`
	PartsAddInfo3    *string `json:"partsAddInfo3"`
}

// ParentTraceID
// Summary: This is function to convert PostPartsStructuresResponse to uuid.UUID.
// output: (uuid.UUID) parentTraceId
// output: (error) error object
func (r PostPartsStructuresResponse) ParentTraceID() (uuid.UUID, error) {
	traceIDStr := r.Parent.TraceID
	traceID, err := uuid.Parse(traceIDStr)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return uuid.Nil, err
	}
	return traceID, nil
}

// ChildTraceID
// Summary: This is function to convert PostPartsStructuresResponse to uuid.UUID.
// input: partsItem(string) partsItem
// input: supportPartsItem(*string) supportPartsItem
// output: (uuid.UUID) parentTraceId
// output: (error) error object
func (r PostPartsStructuresResponse) ChildTraceID(partsItem string, supportPartsItem *string, plantID string) (uuid.UUID, error) {
	var isFound bool = false
	for _, resChild := range r.Children {
		if plantID == resChild.PlantID {
			if partsItem == resChild.PartsItem {
				if supportPartsItem == nil && resChild.SupportPartsItem == nil {
					isFound = true
				}
				if supportPartsItem != nil && resChild.SupportPartsItem != nil && *supportPartsItem == *resChild.SupportPartsItem {
					isFound = true
				}
				if isFound {
					childTraceIDStr := resChild.TraceID
					childTraceID, err := uuid.Parse(childTraceIDStr)
					if err != nil {
						logger.Set(nil).Errorf(err.Error())

						return uuid.Nil, err
					}
					return childTraceID, nil
				}
			}
		}
	}
	logger.Set(nil).Errorf(common.NotFoundInResponseError(partsItem, supportPartsItem))

	return uuid.Nil, fmt.Errorf(common.NotFoundInResponseError(partsItem, supportPartsItem))
}

// NewPostPartsStructureRequestFromModel
// Summary: This is function to convert traceability.PartsStructureModel to PostPartsStructuresRequest.
// input: m(traceability.PartsStructureModel) partsStructure model
// output: (PostPartsStructuresRequest)) api request
func NewPostPartsStructureRequestFromModel(m traceability.PartsStructureModel) PostPartsStructuresRequest {
	reqParent := PostPartsStructuresRequestParent{
		PartsItem:        m.ParentPartsModel.PartsName,
		SupportPartsItem: m.ParentPartsModel.SupportPartsName,
		PlantID:          m.ParentPartsModel.PlantID.String(),
		OperatorID:       m.ParentPartsModel.OperatorID.String(),
		EndFlag:          m.ParentPartsModel.TerminatedFlag,
		PartsLabelName:   m.ParentPartsModel.PartsLabelName,
		PartsAddInfo1:    m.ParentPartsModel.PartsAddInfo1,
		PartsAddInfo2:    m.ParentPartsModel.PartsAddInfo2,
		PartsAddInfo3:    m.ParentPartsModel.PartsAddInfo3,
	}

	var amountUnitName string
	if m.ParentPartsModel.AmountRequiredUnit != nil {
		amountUnitName = m.ParentPartsModel.AmountRequiredUnit.ToString()
		reqParent.AmountUnitName = &amountUnitName
	}

	reqChildren := make([]PostPartsStructuresRequestChild, len(m.ChildrenPartsModel))
	for i, child := range m.ChildrenPartsModel {
		var amountUnitNameChild string
		if child.AmountRequiredUnit != nil {
			amountUnitNameChild = child.AmountRequiredUnit.ToString()
		}

		reqChild := PostPartsStructuresRequestChild{
			PartsItem:        child.PartsName,
			SupportPartsItem: child.SupportPartsName,
			OperatorID:       child.OperatorID.String(),
			PlantID:          child.PlantID.String(),
			AmountUnitName:   &amountUnitNameChild,
			EndFlag:          child.TerminatedFlag,
			Amount:           child.AmountRequired,
			PartsLabelName:   child.PartsLabelName,
			PartsAddInfo1:    child.PartsAddInfo1,
			PartsAddInfo2:    child.PartsAddInfo2,
			PartsAddInfo3:    child.PartsAddInfo3,
		}
		reqChildren[i] = reqChild
	}

	req := PostPartsStructuresRequest{
		OperatorID: m.ParentPartsModel.OperatorID.String(),
		Parent:     reqParent,
		Children:   reqChildren,
	}

	return req
}
