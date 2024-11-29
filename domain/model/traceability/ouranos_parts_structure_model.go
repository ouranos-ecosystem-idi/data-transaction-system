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

// PartsStructureModel
// Summary: This is structure which defines partsStructure model.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=PartsStructure
// Usage: output
type PartsStructureModel struct {
	ParentPartsModel   *PartsModel  `json:"parentPartsModel"`
	ChildrenPartsModel []PartsModel `json:"childrenPartsModel"`
}

// PartsStructureEntityModel
// Summary: This is structure which defines PartsStructureEntityModel.
// DBName: partsStructure
type PartsStructureEntityModel struct {
	TraceID       uuid.UUID      `json:"traceId" gorm:"type:uuid"`
	ParentTraceID uuid.UUID      `json:"parentTraceId" gorm:"type:uuid"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
	CreatedAt     time.Time      `json:"createdAt" gorm:"<-:create"`
	CreatedUserID string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	UpdatedUserID string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
}

// PartsStructureEntity
// Summary: Used for processing put.
// DBName: partsStructure
type PartsStructureEntity struct {
	ParentPartsEntity   *PartsModelEntity  `json:"parentPartsModel"`
	ChildrenPartsEntity PartsModelEntities `json:"childrenPartsModel"`
}

// PartsStructureEntities
// Summary: This is structure which defines list of partsStructure model.
type PartsStructureEntities []*PartsStructureEntity

// PartsStructureEntities
// Summary: This is structure which defines list of partsStructure model.
type PartsStructureEntityModels []PartsStructureEntityModel

// GetPartsStructureInput
// Summary: This is structure which defines partsStructure model.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=PartsStructure
// Usage: input
type GetPartsStructureInput struct {
	TraceID    uuid.UUID `json:"traceId" gorm:"type:uuid"`
	OperatorID string    `json:"operatorId" gorm:"type:string;not null"`
}

// PutPartsStructureInput
// Summary: This is structure which defines partsStructure model.
// Service: Dataspace
// Router: [PUT] /api/v1/datatransport?dataTarget=PartsStructure
// Usage: input
type PutPartsStructureInput struct {
	ParentPartsInput   *PutPartsInput  `json:"parentPartsModel"`
	ChildrenPartsInput *PutPartsInputs `json:"childrenPartsModel"`
}

// validate
// Summary: This is the function to validate GetPartsStructureInput.
// output: (error) Error object
func (m GetPartsStructureInput) Validate() error {
	if err := m.validate(); err != nil {
		logger.Set(nil).Warn(err.Error())
		return err
	}

	return nil
}

// validate
// Summary: This is the function to validate GetPartsStructureInput.
// output: (error) Error object
func (m GetPartsStructureInput) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(
			&m.TraceID,
			validation.By(common.UUIDNotNil),
		),
		validation.Field(
			&m.OperatorID,
			validation.Required,
			validation.RuneLength(36, 36),
		),
	)
}

// validate
// Summary: This is the function to validate PutPartsStructureInput.
// output: (error) Error object
func (i PutPartsStructureInput) Validate() error {
	return i.validate()
}

// validate
// Summary: This is the function to validate PutPartsStructureInput.
// output: (error) Error object
func (i PutPartsStructureInput) validate() error {
	errors := []error{}

	if i.ParentPartsInput == nil {
		newErr := fmt.Errorf("parentPartsModel: %v", common.ErrorMessageCannotBeBlank)
		logger.Set(nil).Warn(newErr.Error())

		errors = append(errors, newErr)
	} else if err := i.ParentPartsInput.validate(); err != nil {
		newErr := fmt.Errorf("parentPartsModel: (%v)", err)
		logger.Set(nil).Warn(newErr.Error())

		errors = append(errors, newErr)
	}

	if i.ChildrenPartsInput == nil {
		newErr := fmt.Errorf("childrenPartsModel: %v", common.ErrorMessageCannotBeBlank)
		logger.Set(nil).Warn(newErr.Error())

		errors = append(errors, newErr)
	} else {
		for i, v := range *i.ChildrenPartsInput {
			if err := v.validateForChild(); err != nil {
				newErr := fmt.Errorf("childrenPartsModel[%d]: (%v)", i, err)
				logger.Set(nil).Warn(newErr.Error())

				errors = append(errors, newErr)
			}
		}
	}

	if len(errors) > 0 {
		return common.JoinErrors(errors)
	}

	return nil
}

// HasChild
// Summary: This is the function to check child parts.
// output: (bool) true if child parts exist
func (i PutPartsStructureInput) HasChild() bool {
	return i.ChildrenPartsInput != nil && len(*i.ChildrenPartsInput) > 0
}

// validate
// Summary: This is the function to validate PutPartsStructureInput.
// output: (error) Error object
func (e PartsStructureEntity) ToModel() (PartsStructureModel, error) {
	var parent *PartsModel
	if e.ParentPartsEntity != nil {
		model, err := e.ParentPartsEntity.ToModel()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return PartsStructureModel{}, err
		}
		parent = &model
	}
	childrent, err := e.ChildrenPartsEntity.ToModels()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PartsStructureModel{}, err
	}

	return PartsStructureModel{
		ParentPartsModel:   parent,
		ChildrenPartsModel: childrent,
	}, nil
}

// IsParent
// Summary: This is the function to check parent parts.
// output: (bool) check ParentTraceID
func (e PartsStructureEntityModel) IsParent() bool {
	return e.ParentTraceID == uuid.Nil
}
