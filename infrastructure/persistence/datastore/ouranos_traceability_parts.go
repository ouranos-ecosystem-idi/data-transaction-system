package datastore

import (
	"fmt"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ListParts
// Summary: This is a function to retrieve a list of part information.
// input: getPartsInput(traceability.GetPartsInput) parts model
// output: (traceability.PartsModels) parts models
// output: (error) Error object
func (r *ouranosRepository) ListParts(getPartsInput traceability.GetPartsInput) (traceability.PartsModelEntities, error) {
	var (
		partsList traceability.PartsModelEntities
		err       error
	)

	query := r.db.Table("parts").
		Select(`
			parts.trace_id,
			parts.operator_id,
			parts.plant_id,
			parts.parts_name,
			parts.support_parts_name,
			parts.terminated_flag,
			parts.amount_required,
			parts.amount_required_unit,
			parts.parts_label_name,
			parts.parts_add_info1,
			parts.parts_add_info2,
			parts.parts_add_info3
		`).
		Where(`parts.deleted_at IS NULL AND parts.operator_id = ?`, getPartsInput.OperatorID)

	if getPartsInput.TraceID != nil {
		query = query.Where("parts.trace_id = ?", *getPartsInput.TraceID)
	}
	if getPartsInput.PartsName != nil {
		query = query.Where("parts.parts_name = ?", *getPartsInput.PartsName)
	}
	if getPartsInput.PlantID != nil {
		query = query.Where("parts.plant_id = ?", *getPartsInput.PlantID)
	}
	if getPartsInput.ParentFlag != nil {
		if *getPartsInput.ParentFlag {
			query = query.
				Joins("INNER JOIN parts_structures ON parts_structures.trace_id = parts.trace_id").
				Where("parts_structures.parent_trace_id = ?", uuid.Nil.String())
		}
	}

	err = query.
		Limit(getPartsInput.Limit).
		Order(`parts_name ASC`).
		Order(`support_parts_name ASC`).
		Find(&partsList).
		Error

	return partsList, err
}

// GetPartByTraceID
// Summary: This function is used to retrieve the results of filtering the part information by traceId.
// input: traceID(string) ID of the trace
// output: (traceability.PartsModelEntity) parts entity
// output: (error) Error object
func (r *ouranosRepository) GetPartByTraceID(traceID string) (traceability.PartsModelEntity, error) {
	var part traceability.PartsModelEntity

	if err := r.db.Table("parts").Where("trace_id = ?", traceID).Limit(1).First(&part).Error; err != nil {
		logger.Set(nil).Error(err.Error())
		return traceability.PartsModelEntity{}, err
	}
	return part, nil
}

// CountPartsList
// Summary: This function is used to retrieve the results of filtering the part information by traceId.
// input: getPartsInput(traceability.GetPartsInput) parts model
// output: (int) parts count
// output: (error) Error object
func (r *ouranosRepository) CountPartsList(getPartsInput traceability.GetPartsInput) (int, error) {
	var count int64
	if err := r.db.Table("parts").Where("deleted_at IS NULL AND operator_id = ?", getPartsInput.OperatorID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// DeleteParts
// Summary: This function deletes the part information.
// input: traceID(string) ID of the trace
// output: (error) Error object
func (r *ouranosRepository) DeleteParts(traceID string) error {
	result := r.db.Unscoped().Table("parts").Where("trace_id = ?", traceID).Delete(nil)
	if result.Error != nil {
		return fmt.Errorf("failed to physically delete record from table parts: %v", result.Error)
	}
	return nil
}

// DeletePartsWithCFP
// Summary: This function deletes the part and CFP information.
// input: traceID(string) ID of the trace
// output: (error) Error object
func (r *ouranosRepository) DeletePartsWithCFP(traceID string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Table("parts").Where("trace_id = ?", traceID).Delete(nil).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return fmt.Errorf("failed to physically delete record from table parts: %v", err)
	}
	return nil
}
