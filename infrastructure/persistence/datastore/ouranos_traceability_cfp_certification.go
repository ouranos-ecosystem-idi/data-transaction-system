package datastore

import (
	"data-spaces-backend/domain/model/traceability"
	f "data-spaces-backend/test/fixtures"
)

// GetCFPCertifications
// Summary: This is function which get cfp certification.
// input: operatorID(string) ID of the operator
// input: traceID(string) ID of the trace
// output: (traceability.CfpCertificationModels) CfpCertificationModels object
// output: (error) error object
func (r *ouranosRepository) GetCFPCertifications(operatorID string, traceID string) (traceability.CfpCertificationModels, error) {
	return f.CfpCertificationsData, nil
}
