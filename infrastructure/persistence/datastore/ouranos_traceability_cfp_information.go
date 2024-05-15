package datastore

import (
	"fmt"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"
)

// GetCFPInformation
// Summary: This is a function to get cfp entity model.
// input: traceID(string) ID of the trace
// output: (traceability.CfpEntityModel) cfp entity model
// output: (error) error object
func (r *ouranosRepository) GetCFPInformation(traceID string) (traceability.CfpEntityModel, error) {
	var result traceability.CfpEntityModel

	if err := r.db.Table("cfp_infomation").Where("trace_id = ?", traceID).First(&result).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.CfpEntityModel{}, err
	}
	return result, nil

}

// DeleteCFPInformation
// Summary: This is a function to delete cfp entity model.
// input: cfpID(string) ID of the cfp
// output: (error) error object
func (r *ouranosRepository) DeleteCFPInformation(cfpID string) error {
	result := r.db.Unscoped().Table("cfp_infomation").Where("cfp_id = ?", cfpID).Delete(nil)
	if result.Error != nil {
		return fmt.Errorf("failed to physically delete record from table cfp_infomation: %v", result.Error)
	}
	return nil
}
