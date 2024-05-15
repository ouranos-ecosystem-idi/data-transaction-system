package datastore

import (
	"fmt"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"
)

// BatchCreateCFP
// Summary: This is a function to batch create cfp entity models.
// input: es(traceability.CfpEntityModels) list of cfp entity models
// output: (traceability.CfpEntityModels) list of cfp entity models
// output: (error) error object
func (r *ouranosRepository) BatchCreateCFP(es traceability.CfpEntityModels) (traceability.CfpEntityModels, error) {
	if len(es) == 0 {
		logger.Set(nil).Errorf("cfp entities is empty")

		return nil, fmt.Errorf("cfp entities is empty")
	}

	for _, e := range es {
		if res := r.db.Table("cfp_infomation").Create(&e); res.Error != nil {
			logger.Set(nil).Errorf("failed to insert cfp_infomation record: %v", res.Error)

			return nil, fmt.Errorf("failed to insert cfp_infomation record: %v", res.Error)
		}
	}

	return es, nil
}

// GetCFP
// Summary: This is a function to get cfp entity model.
// input: cfpID(string) ID of the cfp
// input: cfpType(string) cfp type
// output: (traceability.CfpEntityModel) cfp entity model
// output: (error) error object
func (r *ouranosRepository) GetCFP(cfpID string, cfpType string) (traceability.CfpEntityModel, error) {
	var cfp traceability.CfpEntityModel
	var cfpCertificates []traceability.CfpCertificateEntityModel

	if err := r.db.Table("cfp_infomation").Where("cfp_id = ? AND cfp_type = ?", cfpID, cfpType).Find(&cfp).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.CfpEntityModel{}, err
	}

	if err := r.db.Table("cfp_certificates").Where("cfp_id = ?", cfpID).Order("created_at DESC").Find(&cfpCertificates).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.CfpEntityModel{}, err
	}

	for _, cfpCertificate := range cfpCertificates {
		cfp.CfpCertificateList = append(cfp.CfpCertificateList, cfpCertificate.CfpCertificate)
	}

	return cfp, nil

}

// ListCFPsByTraceID
// Summary: This is a function to list cfp entity models by trace ID.
// input: traceID(string) ID of the trace
// output: (traceability.CfpEntityModels) list of cfp entity models
// output: (error) error object
func (r *ouranosRepository) ListCFPsByTraceID(traceID string) (traceability.CfpEntityModels, error) {
	var cfps traceability.CfpEntityModels
	var cfpCertificates []traceability.CfpCertificateEntityModel

	if err := r.db.Table("cfp_infomation").Where("trace_id = ?", traceID).Order("updated_at DESC").Find(&cfps).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.CfpEntityModels{}, err
	}
	for _, cfp := range cfps {
		if err := r.db.Table("cfp_certificates").Where("cfp_id = ?", cfp.CfpID).Order("created_at DESC").Find(&cfpCertificates).Error; err != nil {
			logger.Set(nil).Errorf(err.Error())

			return traceability.CfpEntityModels{}, err
		}

		for _, cfpCertificate := range cfpCertificates {
			cfp.CfpCertificateList = append(cfp.CfpCertificateList, cfpCertificate.CfpCertificate)
		}
	}
	return cfps, nil
}

// PutCFP
// Summary: This is a function to put cfp entity model.
// input: e(traceability.CfpEntityModel) cfp entity model
// output: (traceability.CfpEntityModel) cfp entity model
// output: (error) error object
func (r *ouranosRepository) PutCFP(e traceability.CfpEntityModel) (traceability.CfpEntityModel, error) {
	if err := r.db.Table("cfp_infomation").Where("cfp_id = ? AND cfp_type = ?", e.CfpID, e.CfpType).Updates(&e).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.CfpEntityModel{}, err
	}

	result := r.db.Unscoped().Table("cfp_certificates").Where("cfp_id = ?", e.CfpID).Delete(nil)
	if result.Error != nil {
		logger.Set(nil).Errorf("failed to physically delete record from table cfp_certificates: %v", result.Error)

		return traceability.CfpEntityModel{}, fmt.Errorf("failed to physically delete record from table cfp_certificates: %v", result.Error)
	}

	for i, cfpCertificate := range e.CfpCertificateList {
		if e.CfpID != nil {
			certificationEntity := traceability.NewCfpCertificationEntityModel(i+1, *e.CfpID, cfpCertificate)
			if result := r.db.Table("cfp_certificates").Create(&certificationEntity); result.Error != nil {
				logger.Set(nil).Errorf("failed to insert cfp_certificates record: %v", result.Error)

				return traceability.CfpEntityModel{}, fmt.Errorf("failed to insert cfp_certificates record: %v", result.Error)
			}
		}
	}
	return e, nil
}
