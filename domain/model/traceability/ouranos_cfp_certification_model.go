package traceability

import (
	"github.com/google/uuid"
)

// CfpCertificationModel
// Summary: This is structure which defines CfpCertificationModel.
type CfpCertificationModel struct {
	CfpCertificationID          string                      `json:"cfpCertificationId"`
	TraceID                     string                      `json:"traceId"`
	CfpCertificationDescription *string                     `json:"cfpCertificationDescription"`
	CfpCertificationFileInfo    *[]CfpCertificationFileInfo `json:"cfpCertificationFileInfo"`
}

// CfpCertificationModels
// Summary: This is a type that defines a list of CfpCertificationModel.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=cfpCertification
// Usage: output
type CfpCertificationModels []CfpCertificationModel

// CfpCertificationFileInfo
// Summary: This is structure which defines CfpCertificationFileInfo.
type CfpCertificationFileInfo struct {
	OperatorID string `json:"operatorId"`
	FileID     string `json:"fileId"`
	FileName   string `json:"fileName"`
}

// GetCfpCertificationInput
// Summary: This is structure which defines GetCfpCertificationInput.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=cfpCertification
// Usage: input
type GetCfpCertificationInput struct {
	OperatorID uuid.UUID `json:"operatorId"`
	TraceID    uuid.UUID `json:"traceId"`
}
