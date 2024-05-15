package datastore

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
)

// GetCFPCertifications
// Summary: This is function which get cfp certification.
// input: operatorID(string) ID of the operator
// input: traceID(string) ID of the trace
// output: (traceability.CfpCertificationModels) CfpCertificationModels object
// output: (error) error object
func (r *ouranosRepository) GetCFPCertifications(operatorID string, traceID string) (traceability.CfpCertificationModels, error) {
	return traceability.CfpCertificationModels{
		{
			CfpCertificationID:          "d9a38406-cae2-4679-b052-15a75f5531c5",
			TraceID:                     traceID,
			CfpCertificationDescription: common.StringPtr("証明書の説明"),
			CfpCertificationFileInfo: &[]traceability.CfpCertificationFileInfo{
				{
					OperatorID: operatorID,
					FileID:     "5c07e3e9-c0e5-4a1f-b6a5-78145f7d1855",
					FileName:   "ダミーファイル.pdf",
				},
			},
		},
	}, nil
}
