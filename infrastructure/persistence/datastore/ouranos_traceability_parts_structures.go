package datastore

import (
	"fmt"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetPartsStructure
// Summary: This function get the partsStructure of a request and response.
// input: getPartsStructureInput(traceability.GetPartsStructureInput) target of the partsStructure
// output: (traceability.PartsStructureEntity) partsStructure entity
// output: (error) error object
func (r *ouranosRepository) GetPartsStructure(getPartsStructureInput traceability.GetPartsStructureInput) (traceability.PartsStructureEntity, error) {
	var (
		partsStructure traceability.PartsStructureEntity
		parentParts    traceability.PartsModelEntity
		childrenParts  []traceability.PartsModelEntity
	)

	err := r.db.Table("parts").
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
		Where(`
				parts.deleted_at IS NULL
				AND parts.trace_id = ?
				AND parts.operator_id = ?
			`, getPartsStructureInput.TraceID,
			getPartsStructureInput.OperatorID).
		Find(&parentParts).Error
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsStructureEntity{}, err
	}

	partsStructure.ParentPartsEntity = &parentParts

	err = r.db.Table("parts").
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
		Where(`
				parts.deleted_at IS NULL
				AND EXISTS (
					SELECT 1 FROM parts_structures
					WHERE parts_structures.parent_trace_id = ?
					AND parts_structures.trace_id = parts.trace_id 
				)
				AND parts.operator_id = ?
			`, getPartsStructureInput.TraceID,
			getPartsStructureInput.OperatorID).
		Order("parts.trace_id ASC").
		Find(&childrenParts).
		Error
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsStructureEntity{}, err
	}

	partsStructure.ChildrenPartsEntity = childrenParts

	return partsStructure, err
}

// PutPartsStructure
// Summary: This function put the partsStructure of a request and response.
// input: partsStructure(traceability.PartsStructureModel) target of the partsStructure
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (r *ouranosRepository) PutPartsStructure(
	partsStructure traceability.PartsStructureModel,
) (
	traceability.PartsStructureEntity, error,
) {

	response := traceability.PartsStructureEntity{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if partsStructure.ParentPartsModel.TraceID == uuid.Nil {
			partsStructure.ParentPartsModel.TraceID, _ = uuid.NewRandom()
		}

		var plantID uuid.UUID
		if partsStructure.ParentPartsModel.PlantID != nil {
			plantID = *partsStructure.ParentPartsModel.PlantID
		}
		var amountRequiredUnit *string
		if partsStructure.ParentPartsModel.AmountRequiredUnit != nil {
			a := partsStructure.ParentPartsModel.AmountRequiredUnit.ToString()
			amountRequiredUnit = &a
		}
		partsEntity := traceability.PartsModelEntity{
			TraceID:            partsStructure.ParentPartsModel.TraceID,
			OperatorID:         partsStructure.ParentPartsModel.OperatorID,
			PlantID:            plantID,
			PartsName:          partsStructure.ParentPartsModel.PartsName,
			SupportPartsName:   partsStructure.ParentPartsModel.SupportPartsName,
			TerminatedFlag:     partsStructure.ParentPartsModel.TerminatedFlag,
			AmountRequired:     partsStructure.ParentPartsModel.AmountRequired,
			AmountRequiredUnit: amountRequiredUnit,
			PartsLabelName:     partsStructure.ParentPartsModel.PartsLabelName,
			PartsAddInfo1:      partsStructure.ParentPartsModel.PartsAddInfo1,
			PartsAddInfo2:      partsStructure.ParentPartsModel.PartsAddInfo2,
			PartsAddInfo3:      partsStructure.ParentPartsModel.PartsAddInfo3,
		}
		res1 := tx.Table("parts").Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
					{Name: "trace_id"},
				},
				DoUpdates: clause.AssignmentColumns(
					[]string{
						"trace_id",
						"operator_id",
						"plant_id",
						"parts_name",
						"support_parts_name",
						"terminated_flag",
						"amount_required",
						"amount_required_unit",
						"parts_label_name",
						"parts_add_info1",
						"parts_add_info2",
						"parts_add_info3",
					}),
			},
		).Create(&partsEntity)
		if res1.Error != nil {
			logger.Set(nil).Errorf("DB Error: When Insert in parts : %v", res1.Error)
			return res1.Error
		}

		response.ParentPartsEntity = &partsEntity
		partsStructureEntity := traceability.PartsStructureEntityModel{
			TraceID: partsStructure.ParentPartsModel.TraceID,
		}

		res2 := tx.Table("parts_structures").Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
					{Name: "trace_id"}, {Name: "parent_trace_id"},
				},
				DoUpdates: clause.AssignmentColumns(
					[]string{
						"trace_id",
						"parent_trace_id",
					}),
			},
		).Create(&partsStructureEntity)
		if res2.Error != nil {
			logger.Set(nil).Errorf("DB Error: When Insert in parts : %v", res2.Error)
			return res2.Error
		}

		for i, v := range partsStructure.ChildrenPartsModel {

			if v.TraceID == uuid.Nil {
				v.TraceID, _ = uuid.NewRandom()
			}

			partsStructure.ChildrenPartsModel[i] = v

			var plantID uuid.UUID
			if v.PlantID != nil {
				plantID = *v.PlantID
			}
			var amountRequiredUnit string
			if v.AmountRequiredUnit != nil {
				amountRequiredUnit = v.AmountRequiredUnit.ToString()
			}
			childPartsEntity := traceability.PartsModelEntity{
				TraceID:            v.TraceID,
				OperatorID:         v.OperatorID,
				PlantID:            plantID,
				PartsName:          v.PartsName,
				SupportPartsName:   v.SupportPartsName,
				TerminatedFlag:     v.TerminatedFlag,
				AmountRequired:     v.AmountRequired,
				AmountRequiredUnit: &amountRequiredUnit,
				PartsLabelName:     v.PartsLabelName,
				PartsAddInfo1:      v.PartsAddInfo1,
				PartsAddInfo2:      v.PartsAddInfo2,
				PartsAddInfo3:      v.PartsAddInfo3,
			}

			response.ChildrenPartsEntity = append(response.ChildrenPartsEntity, childPartsEntity)
			res3 := tx.Table("parts").Clauses(
				clause.OnConflict{
					Columns: []clause.Column{
						{Name: "trace_id"},
					},
					DoUpdates: clause.AssignmentColumns(
						[]string{
							"trace_id",
							"operator_id",
							"plant_id",
							"parts_name",
							"support_parts_name",
							"terminated_flag",
							"amount_required",
							"amount_required_unit",
							"parts_label_name",
							"parts_add_info1",
							"parts_add_info2",
							"parts_add_info3",
						}),
				}).Create(&childPartsEntity)
			if res3.Error != nil {
				logger.Set(nil).Errorf("DB Error: When Insert in parts : %v", res3.Error)

				return res3.Error
			}

			chaildPartsStructureEntity := traceability.PartsStructureEntityModel{
				TraceID:       v.TraceID,
				ParentTraceID: partsStructure.ParentPartsModel.TraceID,
			}

			res4 := tx.Table("parts_structures").Clauses(
				clause.OnConflict{
					Columns: []clause.Column{
						{Name: "trace_id"}, {Name: "parent_trace_id"},
					},
					DoUpdates: clause.AssignmentColumns(
						[]string{
							"trace_id",
							"parent_trace_id",
						}),
				},
			).Create(&chaildPartsStructureEntity)
			if res4.Error != nil {
				logger.Set(nil).Errorf("DB Error: When Insert in parts : %v", res4.Error)
				return res4.Error
			}
		}

		return nil
	})

	return response, err
}

// DeletePartsStructure
// Summary: This function delete the partsStructure of a request and response.
// input: traceID(string) ID of the trace
// output: (error) error object
func (r *ouranosRepository) DeletePartsStructure(traceID string) error {
	result := r.db.Unscoped().Table("parts_structures").Where("trace_id = ?", traceID).Delete(nil)
	if result.Error != nil {
		return fmt.Errorf("failed to physically delete record from table parts_structures: %v", result.Error)
	}
	return nil
}

// GetPartsStructureByTraceId
// Summary: This function get the partsStructure by traceId of a request and response.
// input: traceID(string) ID of the trace
// output: (traceability.PartsStructureEntityModel) partsStructure model
// output: (error) error object
func (r *ouranosRepository) GetPartsStructureByTraceId(traceID string) (traceability.PartsStructureEntityModel, error) {
	var partsStructure traceability.PartsStructureEntityModel

	if err := r.db.Table("parts_structures").Where("trace_id = ?", traceID).Limit(1).First(&partsStructure).Error; err != nil {
		logger.Set(nil).Error(err.Error())

		return traceability.PartsStructureEntityModel{}, err
	}

	return partsStructure, nil
}

// ListParentPartsStructureByTraceId
// Summary: This is function which get PartsStructureEntityModels from partsStructures by using trace_id.
// input: traceID(string) ID of the trace
// output: (traceability.PartsStructureEntityModels) partsStructure model
// output: (error) error object
func (r *ouranosRepository) ListParentPartsStructureByTraceId(traceID string) (traceability.PartsStructureEntityModels, error) {
	var es traceability.PartsStructureEntityModels

	if err := r.db.Table("parts_structures").Where("trace_id = ? AND parent_trace_id <> ?", traceID, uuid.Nil.String()).Find(&es).Error; err != nil {
		logger.Set(nil).Error(err.Error())

		return traceability.PartsStructureEntityModels{}, err
	}

	return es, nil
}

// ListChildPartsStructureByTraceId
// Summary: This is function which get PartsStructureEntityModels from partsStructures by using trace_id.
// input: traceID(string) ID of the trace
// output: (traceability.PartsStructureEntityModels) partsStructure model
// output: (error) error object
func (r *ouranosRepository) ListChildPartsStructureByTraceId(traceID string) (traceability.PartsStructureEntityModels, error) {
	var es traceability.PartsStructureEntityModels

	if err := r.db.Table("parts_structures").Where("parent_trace_id = ?", traceID).Find(&es).Error; err != nil {
		logger.Set(nil).Error(err.Error())

		return traceability.PartsStructureEntityModels{}, err
	}

	return es, nil
}
