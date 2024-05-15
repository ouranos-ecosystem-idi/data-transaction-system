package datastore

import (
	"fmt"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
)

// ListParts
// Summary: This is a function to retrieve a list of part information.
// input: getPlantPartsModel(traceability.GetPartsModel) parts model
// output: (traceability.PartsModels) parts models
// output: (error) Error object
func (r *ouranosRepository) ListParts(getPlantPartsModel traceability.GetPartsModel) (traceability.PartsModelEntities, error) {
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
			parts.amount_required_unit
		`).
		Where(`parts.deleted_at IS NULL AND parts.operator_id = ?`, getPlantPartsModel.OperatorID)

	if getPlantPartsModel.TraceID != nil {
		query = query.Where("parts.trace_id = ?", *getPlantPartsModel.TraceID)
	}
	if getPlantPartsModel.PartsName != nil {
		query = query.Where("parts.parts_name = ?", *getPlantPartsModel.PartsName)
	}
	if getPlantPartsModel.PlantID != nil {
		query = query.Where("parts.plant_id = ?", *getPlantPartsModel.PlantID)
	}
	if getPlantPartsModel.ParentFlag != nil {
		if *getPlantPartsModel.ParentFlag {
			query = query.
				Joins("INNER JOIN parts_structures ON parts_structures.trace_id = parts.trace_id").
				Where("parts_structures.parent_trace_id = ?", uuid.Nil.String())
		}
	}

	err = query.
		Limit(getPlantPartsModel.Limit).
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
// input: getPlantPartsModel(traceability.GetPartsModel) parts model
// output: (int) parts count
// output: (error) Error object
func (r *ouranosRepository) CountPartsList(getPlantPartsModel traceability.GetPartsModel) (int, error) {
	var count int64
	if err := r.db.Table("parts").Where("deleted_at IS NULL AND operator_id = ?", getPlantPartsModel.OperatorID).Count(&count).Error; err != nil {
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
