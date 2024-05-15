package traceabilityentity

import "data-spaces-backend/domain/model/traceability"

// GetCfpCertificationsRequest
// Summary: This is struct which defines get cfp GetCfpCertificationsRequest.
// Service: Traceability
// Router: [GET] /cfpCertifications
// Usage: input
type GetCfpCertificationsRequest struct {
	OperatorID string `json:"operatorId"`
	TraceID    string `json:"traceId"`
}

// GetCfpCertificationsResponse
// Summary: This is a type that defines a list of GetCfpCertificationsResponseCfpCertification.
// Service: Traceability
// Router: [GET] /cfpCertifications
// Usage: output
type GetCfpCertificationsResponse []GetCfpCertificationsResponseCfpCertification

// GetCfpCertificationsResponseCfpCertification
// Summary: This is structure which defines GetCfpCertificationsResponseCfpCertification.
type GetCfpCertificationsResponseCfpCertification struct {
	CfpCertificationID          string                      `json:"cfpCertificationId"`
	TraceID                     string                      `json:"traceId"`
	CfpCertificationDescription string                      `json:"cfpCertificationDescription"`
	CfpCertificationFileInfo    *[]CfpCertificationFileInfo `json:"cfpCertificationFileInfo"`
	CreatedAt                   *string                     `json:"createdAt"`
}

// CfpCertificationFileInfo
// Summary: This is structure which defines CfpCertificationFileInfo.
type CfpCertificationFileInfo struct {
	OperatorID string `json:"operatorId"`
	FileID     string `json:"fileId"`
	FileName   string `json:"fileName"`
}

// ToModels
// Summary: This is function to convert GetCfpCertificationsResponse to CfpCertificationModels.
// output: (traceability.CfpCertificationModels) CfpCertificationModels object
func (r *GetCfpCertificationsResponse) ToModels() traceability.CfpCertificationModels {
	cfpCertificationModels := traceability.CfpCertificationModels{}
	for _, cfpCertification := range *r {
		cfpCertificationModel := cfpCertification.ToModel()
		cfpCertificationModels = append(cfpCertificationModels, cfpCertificationModel)
	}
	return cfpCertificationModels
}

// ToModel
// Summary: This is function to convert GetCfpCertificationsResponseCfpCertification to CfpCertificationModel.
// output: (traceability.CfpCertificationModel) CfpCertificationModel object
func (r GetCfpCertificationsResponseCfpCertification) ToModel() traceability.CfpCertificationModel {
	cfpCertificationModel := traceability.CfpCertificationModel{
		CfpCertificationID:          r.CfpCertificationID,
		TraceID:                     r.TraceID,
		CfpCertificationDescription: &r.CfpCertificationDescription,
	}
	if r.CfpCertificationFileInfo != nil {
		cfpCertificationFileInfos := make([]traceability.CfpCertificationFileInfo, len(*r.CfpCertificationFileInfo))
		for i, cfpCertificationFileInfo := range *r.CfpCertificationFileInfo {
			cfpCertificationFileInfos[i] = traceability.CfpCertificationFileInfo{
				OperatorID: cfpCertificationFileInfo.OperatorID,
				FileID:     cfpCertificationFileInfo.FileID,
				FileName:   cfpCertificationFileInfo.FileName,
			}
		}
		cfpCertificationModel.CfpCertificationFileInfo = &cfpCertificationFileInfos
	}
	return cfpCertificationModel
}
