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

// PartsModel
// Summary: This is structure which defines parts model.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=parts
// Usage: input
type PartsModel struct {
	TraceID            uuid.UUID           `json:"traceId"`
	OperatorID         uuid.UUID           `json:"operatorId"`
	PlantID            *uuid.UUID          `json:"plantId"`
	PartsName          string              `json:"partsName"`
	SupportPartsName   *string             `json:"supportPartsName"`
	TerminatedFlag     bool                `json:"terminatedFlag"`
	AmountRequired     *float64            `json:"amountRequired"`
	AmountRequiredUnit *AmountRequiredUnit `json:"amountRequiredUnit"`
	PartsLabelName     *string             `json:"partsLabelName"`
	PartsAddInfo1      *string             `json:"partsAddInfo1"`
	PartsAddInfo2      *string             `json:"partsAddInfo2"`
	PartsAddInfo3      *string             `json:"partsAddInfo3"`
}

// PartsModels
// Summary: This is structure which defines list of parts model.
type PartsModels []PartsModel

// AmountRequiredUnit
// Summary: This is structure which defines amount required unit.
type AmountRequiredUnit string

const (
	AmountRequiredUnitLiter        AmountRequiredUnit = "liter"
	AmountRequiredUnitKilogram     AmountRequiredUnit = "kilogram"
	AmountRequiredUnitCubicMeter   AmountRequiredUnit = "cubic-meter"
	AmountRequiredUnitKilowattHour AmountRequiredUnit = "kilowatt-hour"
	AmountRequiredUnitMegajoule    AmountRequiredUnit = "megajoule"
	AmountRequiredUnitTonKilometer AmountRequiredUnit = "ton-kilometer"
	AmountRequiredUnitSquareMeter  AmountRequiredUnit = "square-meter"
	AmountRequiredUnitUnit         AmountRequiredUnit = "unit"
	AmountRequiredUnitEmpty        AmountRequiredUnit = ""
)

// NewAmountRequiredUnit
// Summary: This is the function to convert string to AmountRequiredUnit.
// input: s(string) amount unit name
// output: (AmountRequiredUnit) amount required unit
// output: (error) Error object
func NewAmountRequiredUnit(s string) (AmountRequiredUnit, error) {
	switch s {
	case AmountRequiredUnitLiter.ToString(),
		AmountRequiredUnitKilogram.ToString(),
		AmountRequiredUnitCubicMeter.ToString(),
		AmountRequiredUnitKilowattHour.ToString(),
		AmountRequiredUnitMegajoule.ToString(),
		AmountRequiredUnitTonKilometer.ToString(),
		AmountRequiredUnitSquareMeter.ToString(),
		AmountRequiredUnitUnit.ToString(),
		AmountRequiredUnitEmpty.ToString():
		return AmountRequiredUnit(s), nil
	default:
		return "", fmt.Errorf(common.UnexpectedEnumError("AmountRequiredUnit", s))
	}
}

// ToString
// Summary: This is the function to convert AmountRequiredUnit to string.
// output: (string) converted to string
func (e AmountRequiredUnit) ToString() string {
	return string(e)
}

// MaskAmountRequired
// Summary: The function converts the received activity amount to nil.
// output: (PartsModels) converted to AmountRequired = nil
func (ms PartsModelEntities) MaskAmountRequired() PartsModelEntities {
	for i, m := range ms {
		ms[i] = m.MaskAmountRequired()
	}
	return ms
}

// MaskAmountRequired
// Summary: The function converts the received activity amount to nil.
// output: (PartsModel)  converted to AmountRequired = nil
func (m PartsModelEntity) MaskAmountRequired() PartsModelEntity {
	m.AmountRequired = nil
	return m
}

// PartsModelEntity
// Summary: This is structure which defines PartsModelEntity.
// DBName: parts
type PartsModelEntity struct {
	TraceID            uuid.UUID      `json:"traceId" gorm:"type:uuid"`
	OperatorID         uuid.UUID      `json:"operatorId" gorm:"type:uuid;not null"`
	PlantID            uuid.UUID      `json:"plantId" gorm:"type:uuid;not null"`
	PartsName          string         `json:"partsName" gorm:"type:string;not null"`
	SupportPartsName   *string        `json:"supportPartsName" gorm:"type:string"`
	TerminatedFlag     bool           `json:"terminatedFlag" gorm:"type:boolean;not null"`
	AmountRequired     *float64       `json:"amountRequired" gorm:"type:float(64)"`
	AmountRequiredUnit *string        `json:"amountRequiredUnit" gorm:"type:string"`
	DeletedAt          gorm.DeletedAt `json:"deletedAt"`
	CreatedAt          time.Time      `json:"createdAt" gorm:"<-:create"`
	CreatedUserId      string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	UpdatedUserId      string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
	PartsLabelName     *string        `json:"partsLabelName" gorm:"type:string"`
	PartsAddInfo1      *string        `json:"partsAddInfo1" gorm:"type:string"`
	PartsAddInfo2      *string        `json:"partsAddInfo2" gorm:"type:string"`
	PartsAddInfo3      *string        `json:"partsAddInfo3" gorm:"type:string"`
}

// PartsModelEntities
// Summary: This is structure which defines parts model.
type PartsModelEntities []PartsModelEntity

// GetPartsInput
// Summary: Defines the type of request used for usecase of parts.
type GetPartsInput struct {
	OperatorID string
	TraceID    *string `json:"traceId"`
	PartsName  *string `json:"partsName"`
	PlantID    *string `json:"plantId"`
	ParentFlag *bool   `json:"parentFlag"`
	Limit      int     `json:"limit"`
	After      *uuid.UUID
}

// PutPartsInput
// Summary: Defines the type of request used in the parts information update.
// Service: Dataspace
// Router: [PUT] /api/v1/datatransport/parts
// Usage: input
type PutPartsInput struct {
	OperatorID         string   `json:"operatorId"`
	TraceID            *string  `json:"traceId"`
	PlantID            string   `json:"plantId"`
	PartsName          string   `json:"partsName"`
	SupportPartsName   *string  `json:"supportPartsName"`
	TerminatedFlag     *bool    `json:"terminatedFlag"`
	AmountRequired     *float64 `json:"amountRequired"`
	AmountRequiredUnit *string  `json:"amountRequiredUnit"`
	PartsLabelName     *string  `json:"partsLabelName"`
	PartsAddInfo1      *string  `json:"partsAddInfo1"`
	PartsAddInfo2      *string  `json:"partsAddInfo2"`
	PartsAddInfo3      *string  `json:"partsAddInfo3"`
}

// PutPartsInputs
// Summary: This is structure which defines list of parts model.
type PutPartsInputs []PutPartsInput

// DeletePartsInput
// Summary: Defines the type of request used for usecase of parts.
type DeletePartsInput struct {
	TraceID string `json:"traceId"`
}

// validate
// Summary: This is the function to validate PutPartsInput.
// output: (error) Error object
func (i PutPartsInput) Validate() error {
	return i.validate()
}

func (i PutPartsInput) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.OperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.TraceID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.PlantID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.PartsName,
			validation.Required,
			validation.RuneLength(1, 50),
		),
		validation.Field(
			&i.SupportPartsName,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.TerminatedFlag,
			validation.By(common.BoolPtrNotNil),
		),
		validation.Field(
			&i.AmountRequired,
			validation.Nil,
		),
		validation.Field(
			&i.AmountRequiredUnit,
			validation.By(EnumAmountRequiredUnitValidOrNil),
		),
		validation.Field(
			&i.PartsLabelName,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.PartsAddInfo1,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.PartsAddInfo2,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.PartsAddInfo3,
			validation.RuneLength(0, 50),
		),
	)
}

// validateForChild
// Summary: This is the function to validate PutPartsInput.
// output: (error) Error object
func (i PutPartsInput) validateForChild() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.OperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.TraceID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.PlantID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.PartsName,
			validation.Required,
			validation.RuneLength(1, 50),
		),
		validation.Field(
			&i.SupportPartsName,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.TerminatedFlag,
			validation.By(common.BoolPtrNotNil),
		),
		validation.Field(
			&i.AmountRequired,
			validation.NotNil,
			validation.Min(0.00000),
			validation.Max(99999.99999),
			validation.By(common.FloatPtr5thDecimal),
		),
		validation.Field(
			&i.AmountRequiredUnit,
			validation.NotNil,
			validation.By(EnumAmountRequiredUnitValidOrNil),
		),
		validation.Field(
			&i.PartsLabelName,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.PartsAddInfo1,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.PartsAddInfo2,
			validation.RuneLength(0, 50),
		),
		validation.Field(
			&i.PartsAddInfo3,
			validation.RuneLength(0, 50),
		),
	)
}

// ToModel
// Summary: This is the function to convert PartsModelEntity to PartsModel.
// output: (PartsModel) converted to PartsModel
// output: (error) Error object
func (e PartsModelEntity) ToModel() (PartsModel, error) {
	var err error

	model := PartsModel{
		TraceID:          e.TraceID,
		OperatorID:       e.OperatorID,
		PlantID:          &e.PlantID,
		PartsName:        e.PartsName,
		SupportPartsName: e.SupportPartsName,
		TerminatedFlag:   e.TerminatedFlag,
		AmountRequired:   e.AmountRequired,
		PartsLabelName:   e.PartsLabelName,
		PartsAddInfo1:    e.PartsAddInfo1,
		PartsAddInfo2:    e.PartsAddInfo2,
		PartsAddInfo3:    e.PartsAddInfo3,
	}

	var amountRequiredUnit AmountRequiredUnit
	if e.AmountRequiredUnit != nil {
		amountRequiredUnit, err = NewAmountRequiredUnit(*e.AmountRequiredUnit)
		if err != nil {
			return PartsModel{}, err
		}

		model.AmountRequiredUnit = &amountRequiredUnit
	}

	return model, nil
}

// ToModels
// Summary: This is the function to convert PartsModelEntities to []PartsModel.
// output: ([]PartsModel) converted to []PartsModel
// output: (error) Error object
func (es PartsModelEntities) ToModels() ([]PartsModel, error) {
	ms := make([]PartsModel, len(es))
	for i, e := range es {
		m, err := e.ToModel()
		if err != nil {
			return nil, err
		}
		ms[i] = m
	}
	return ms, nil

}

// ToModelWithMasking
// Summary: The function converts the received activity amount to nil.
// output: (PartsModel) converted to PartsModel
// output: (error) Error object
func (e PartsModelEntity) ToModelWithMasking() (PartsModel, error) {
	m, err := e.ToModel()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PartsModel{}, err
	}
	m.AmountRequired = nil
	return m, nil
}

// ToModel
// Summary: This is the function to convert PutPartsInput to PartsModel.
// output: (PartsModel) converted to PartsModel
// output: (error) Error object
func (i PutPartsInput) ToModel() (PartsModel, error) {
	var model PartsModel

	operatorID, err := uuid.Parse(i.OperatorID)
	if err != nil {
		logger.Set(nil).Error(err.Error())

		return PartsModel{}, fmt.Errorf(common.InvalidUUIDError("operatorId"))
	}
	model.OperatorID = operatorID

	if i.TraceID != nil {
		traceID, err := uuid.Parse(*i.TraceID)
		if err != nil {
			logger.Set(nil).Error(err.Error())

			return PartsModel{}, fmt.Errorf(common.InvalidUUIDError("traceId"))
		}
		model.TraceID = traceID
	}

	plantID, err := uuid.Parse(i.PlantID)
	if err != nil {
		logger.Set(nil).Error(err.Error())

		return PartsModel{}, fmt.Errorf(common.InvalidUUIDError("plantId"))
	}
	model.PlantID = &plantID

	var amountRequiredUnit AmountRequiredUnit
	if i.AmountRequiredUnit != nil {
		amountRequiredUnit, err = NewAmountRequiredUnit(*i.AmountRequiredUnit)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return PartsModel{}, err
		}
		model.AmountRequiredUnit = &amountRequiredUnit
	}

	model.PartsName = i.PartsName
	model.SupportPartsName = i.SupportPartsName
	model.TerminatedFlag = *i.TerminatedFlag
	model.AmountRequired = i.AmountRequired
	model.PartsLabelName = i.PartsLabelName
	model.PartsAddInfo1 = i.PartsAddInfo1
	model.PartsAddInfo2 = i.PartsAddInfo2
	model.PartsAddInfo3 = i.PartsAddInfo3

	return model, nil
}

// ToModels
// Summary: This is the function to convert PutPartsInputs to []PartsModel.
// output: ([]PartsModel) converted to []PartsModel
// output: (error) Error object
func (is PutPartsInputs) ToModels() ([]PartsModel, error) {
	models := make([]PartsModel, 0, len(is))

	for _, i := range is {
		m, err := i.ToModel()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return []PartsModel{}, nil
		}
		models = append(models, m)
	}

	return models, nil
}
