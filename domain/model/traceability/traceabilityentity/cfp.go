package traceabilityentity

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
)

// GetCfpRequest
// Summary: This is structure which defines GetCfpRequest.
// Service: Traceability
// Router: [GET] /cfp
// Usage: input
type GetCfpRequest struct {
	OperatorID string `json:"operatorId"`
	TraceID    string `json:"traceId"`
}

// GetCfpResponses
// Summary: This is a type that defines a list of GetCfpResponseCfp.
// Service: Traceability
// Router: [GET] /cfp
// Usage: output
type GetCfpResponses []GetCfpResponse

// GetCfpResponse
// Summary: This is structure which defines GetCfpResponseCfp.
type GetCfpResponse struct {
	Cfp      GetCfpResponseCfp      `json:"cfp"`
	TotalCfp GetCfpResponseTotalCfp `json:"totalCfp"`
}

// GetCfpResponseCfp
// Summary: This is structure which defines GetCfpResponseCfpInfo.
type GetCfpResponseCfp struct {
	CfpID                           string  `json:"cfpId"`
	TraceID                         string  `json:"traceId"`
	PreProcessingOwnEmissions       float64 `json:"preProcessingOwnOriginatedEmissions"`
	MainProductionOwnEmissions      float64 `json:"mainProductionOwnOriginatedEmissions"`
	PreProcessingSupplierEmissions  float64 `json:"preProcessingSupplierOriginatedEmissions"`
	MainProductionSupplierEmissions float64 `json:"mainProductionSupplierOriginatedEmissions"`
	EmissionsUnitName               string  `json:"emissionsUnitName"`
	CfpComment                      string  `json:"cfpComment"`
	Dqr                             Dqr     `json:"dqr"`
	ParentFlag                      bool    `json:"parentFlag"`
}

// GetCfpResponseTotalCfp
// Summary: This is structure which defines GetCfpResponseTotalCfpInfo.
type GetCfpResponseTotalCfp struct {
	TotalPreProcessingOwnOriginatedEmissions       *float64 `json:"totalPreProcessingOwnOriginatedEmissions"`
	TotalMainProductionOwnOriginatedEmissions      *float64 `json:"totalMainProductionOwnOriginatedEmissions"`
	TotalPreProcessingSupplierOriginatedEmissions  *float64 `json:"totalPreProcessingSupplierOriginatedEmissions"`
	TotalMainProductionSupplierOriginatedEmissions *float64 `json:"totalMainProductionSupplierOriginatedEmissions"`
	TotalEmissionsUnitName                         *string  `json:"emissionsUnitName"`
	TotalDqr                                       Dqr      `json:"totalDqr"`
}

// Dqr
// Summary: This is structure which defines Dqr.
type Dqr struct {
	PreProcessingTeR  *float64 `json:"preProcessingTeR"`
	PreProcessingGeR  *float64 `json:"preProcessingGeR"`
	PreProcessingTiR  *float64 `json:"preProcessingTiR"`
	MainProductionTeR *float64 `json:"mainProductionTeR"`
	MainProductionGeR *float64 `json:"mainProductionGeR"`
	MainProductionTiR *float64 `json:"mainProductionTiR"`
}

// createCfpModel
// Summary: This is function which creates CfpModel.
// input: cfpID(uuid.UUID) ID of the cfp
// input: traceID(uuid.UUID) ID of the trace
// input: ghgEmission(float64) GHG emission value
// input: ghgDeclaredUnitStr(string) GHG declared unit
// input: cfpType(string) CFP type
// input: TeR(*float64) TeR value
// input: GeR(*float64) GeR value
// input: TiR(*float64) TiR value
// output: (traceability.CfpModel) CfpModel object
// output: (error) error object
func createCfpModel(cfpID, traceID uuid.UUID, ghgEmission *float64, ghgDeclaredUnitStr string, cfpType string, TeR *float64, GeR *float64, TiR *float64) (traceability.CfpModel, error) {
	ghgDeclaredUnit, err := traceability.NewGhgDeclaredUnit(ghgDeclaredUnitStr)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.CfpModel{}, err
	}

	dqrType, err := traceability.CfpType(cfpType).ToDqrType()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.CfpModel{}, err
	}

	return traceability.CfpModel{
		CfpID:           &cfpID,
		TraceID:         traceID,
		GhgEmission:     ghgEmission,
		GhgDeclaredUnit: ghgDeclaredUnit,
		CfpType:         cfpType,
		DqrType:         dqrType.ToString(),
		DqrValue: traceability.DqrValue{
			TeR: TeR,
			GeR: GeR,
			TiR: TiR,
		},
	}, nil
}

// ToModels
// Summary: This is function to convert GetCfpResponse to CfpModels.
// input: parts(GetPartsResponse) GetPartsResponse object
// output: (traceability.CfpModels) CfpModels object
// output: (error) error object
func (r GetCfpResponses) ToModels() (traceability.CfpModels, error) {

	ms := []traceability.CfpModel{}

	for _, cfp := range r {
		cfpID, err := uuid.Parse(cfp.Cfp.CfpID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
		traceID, err := uuid.Parse(cfp.Cfp.TraceID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}

		kv := map[string]*float64{
			traceability.CfpTypePreProduction.ToString():       &cfp.Cfp.PreProcessingOwnEmissions,
			traceability.CfpTypeMainProduction.ToString():      &cfp.Cfp.MainProductionOwnEmissions,
			traceability.CfpTypePreComponent.ToString():        &cfp.Cfp.PreProcessingSupplierEmissions,
			traceability.CfpTypeMainComponent.ToString():       &cfp.Cfp.MainProductionSupplierEmissions,
			traceability.CfpTypePreProductionTotal.ToString():  cfp.TotalCfp.TotalPreProcessingOwnOriginatedEmissions,
			traceability.CfpTypeMainProductionTotal.ToString(): cfp.TotalCfp.TotalMainProductionOwnOriginatedEmissions,
			traceability.CfpTypePreComponentTotal.ToString():   cfp.TotalCfp.TotalPreProcessingSupplierOriginatedEmissions,
			traceability.CfpTypeMainComponentTotal.ToString():  cfp.TotalCfp.TotalMainProductionSupplierOriginatedEmissions,
		}

		kvRepresentativeness := map[string][]*float64{
			traceability.CfpTypePreProduction.ToString():       {cfp.Cfp.Dqr.PreProcessingTeR, cfp.Cfp.Dqr.PreProcessingGeR, cfp.Cfp.Dqr.PreProcessingTiR},
			traceability.CfpTypeMainProduction.ToString():      {cfp.Cfp.Dqr.MainProductionTeR, cfp.Cfp.Dqr.MainProductionGeR, cfp.Cfp.Dqr.MainProductionTiR},
			traceability.CfpTypePreComponent.ToString():        {cfp.Cfp.Dqr.PreProcessingTeR, cfp.Cfp.Dqr.PreProcessingGeR, cfp.Cfp.Dqr.PreProcessingTiR},
			traceability.CfpTypeMainComponent.ToString():       {cfp.Cfp.Dqr.MainProductionTeR, cfp.Cfp.Dqr.MainProductionGeR, cfp.Cfp.Dqr.MainProductionTiR},
			traceability.CfpTypePreProductionTotal.ToString():  {cfp.TotalCfp.TotalDqr.PreProcessingTeR, cfp.TotalCfp.TotalDqr.PreProcessingGeR, cfp.TotalCfp.TotalDqr.PreProcessingTiR},
			traceability.CfpTypeMainProductionTotal.ToString(): {cfp.TotalCfp.TotalDqr.MainProductionTeR, cfp.TotalCfp.TotalDqr.MainProductionGeR, cfp.TotalCfp.TotalDqr.MainProductionTiR},
			traceability.CfpTypePreComponentTotal.ToString():   {cfp.TotalCfp.TotalDqr.PreProcessingTeR, cfp.TotalCfp.TotalDqr.PreProcessingGeR, cfp.TotalCfp.TotalDqr.PreProcessingTiR},
			traceability.CfpTypeMainComponentTotal.ToString():  {cfp.TotalCfp.TotalDqr.MainProductionTeR, cfp.TotalCfp.TotalDqr.MainProductionGeR, cfp.TotalCfp.TotalDqr.MainProductionTiR},
		}

		for cfpType, emission := range kv {
			if !cfp.Cfp.ParentFlag && traceability.CfpType(cfpType).IsTotal() {
				continue
			}
			representativenessValues := kvRepresentativeness[cfpType]
			m, err := createCfpModel(cfpID, traceID, emission, cfp.Cfp.EmissionsUnitName, cfpType, representativenessValues[0], representativenessValues[1], representativenessValues[2])
			if err != nil {
				logger.Set(nil).Errorf(err.Error())

				return nil, err
			}
			ms = append(ms, m)
		}
	}

	return ms, nil
}

// PostCfpRequest
// Summary: This is structure which defines PostCfpRequest.
// Service: Traceability
// Router: [POST] /cfp
// Usage: input
type PostCfpRequest struct {
	OperatorID string          `json:"operatorId"`
	Cfp        PostCfpRequests `json:"cfp"`
}

// PostCfpRequestCfp
// Summary: This is structure which defines PostCfpRequestCfp.
type PostCfpRequestCfp struct {
	CfpID                                     *string               `json:"cfpId"`
	TraceID                                   string                `json:"traceId"`
	PreProcessingOwnOriginatedEmissions       *float64              `json:"preProcessingOwnOriginatedEmissions"`
	MainProductionOwnOriginatedEmissions      *float64              `json:"mainProductionOwnOriginatedEmissions"`
	PreProcessingSupplierOriginatedEmissions  *float64              `json:"preProcessingSupplierOriginatedEmissions"`
	MainProductionSupplierOriginatedEmissions *float64              `json:"mainProductionSupplierOriginatedEmissions"`
	EmissionsUnitName                         string                `json:"emissionsUnitName"`
	CfpComment                                *string               `json:"cfpComment"`
	Dqr                                       *PostCfpRequestCfpDqr `json:"dqr"`
}

// PostCfpRequestCfpDqr
// Summary: This is structure which defines PostCfpRequestCfpDqr.
type PostCfpRequestCfpDqr struct {
	PreProcessingTeR  *float64 `json:"preProcessingTeR"`
	PreProcessingGeR  *float64 `json:"preProcessingGeR"`
	PreProcessingTiR  *float64 `json:"preProcessingTiR"`
	MainProductionTeR *float64 `json:"mainProductionTeR"`
	MainProductionGeR *float64 `json:"mainProductionGeR"`
	MainProductionTiR *float64 `json:"mainProductionTiR"`
}

// PostCfpRequests
// Summary: This is a type that defines a list of PostCfpRequestCfp.
type PostCfpRequests []PostCfpRequestCfp

// NewPostCfpRequestFromModel
// Summary: This is function to create new PostCfpRequest from CfpModels.
// input: cfpModels(traceability.CfpModels) CfpModels object
// input: operatorID(string) ID of the operator
// output: (PostCfpRequest) PostCfpRequest object
// output: (error) error object
func NewPostCfpRequestFromModel(cfpModels traceability.CfpModels, operatorID string) (PostCfpRequest, error) {
	cfpID, err := cfpModels.GetCommonCfpID()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}
	traceID, err := cfpModels.GetCommonTraceID()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}
	preProcessingOwnOriginatedEmissions, err := cfpModels.GetGhgEmission(traceability.CfpTypePreProduction)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}
	mainProductionOwnOriginatedEmissions, err := cfpModels.GetGhgEmission(traceability.CfpTypeMainProduction)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}
	preProcessingSupplierOriginatedEmissions, err := cfpModels.GetGhgEmission(traceability.CfpTypePreComponent)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}
	mainProductionSupplierOriginatedEmissions, err := cfpModels.GetGhgEmission(traceability.CfpTypeMainComponent)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}
	emissionsUnitName, err := cfpModels.GetCommonGhgDeclaredUnit()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	preProcessingTeR, err := cfpModels.GetPreProcessingTeR()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	preProcessingGeR, err := cfpModels.GetPreProcessingGeR()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	preProcessingTiR, err := cfpModels.GetPreProcessingTiR()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	mainProductionTeR, err := cfpModels.GetMainProductionTeR()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	mainProductionGeR, err := cfpModels.GetMainProductionGeR()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	mainProductionTiR, err := cfpModels.GetMainProductionTiR()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return PostCfpRequest{}, err
	}

	dqr := PostCfpRequestCfpDqr{
		PreProcessingTeR:  preProcessingTeR,
		PreProcessingGeR:  preProcessingGeR,
		PreProcessingTiR:  preProcessingTiR,
		MainProductionTeR: mainProductionTeR,
		MainProductionGeR: mainProductionGeR,
		MainProductionTiR: mainProductionTiR,
	}

	c := PostCfpRequestCfp{
		CfpID:                                     common.UUIDPtrToStringPtr(cfpID),
		TraceID:                                   traceID.String(),
		PreProcessingOwnOriginatedEmissions:       &preProcessingOwnOriginatedEmissions,
		MainProductionOwnOriginatedEmissions:      &mainProductionOwnOriginatedEmissions,
		PreProcessingSupplierOriginatedEmissions:  &preProcessingSupplierOriginatedEmissions,
		MainProductionSupplierOriginatedEmissions: &mainProductionSupplierOriginatedEmissions,
		EmissionsUnitName:                         emissionsUnitName,
		CfpComment:                                nil,
		Dqr:                                       &dqr,
	}

	cs := PostCfpRequests{c}

	r := PostCfpRequest{
		OperatorID: operatorID,
		Cfp:        cs,
	}
	return r, nil
}

// PostCfpResponse
// Summary: This is structure which defines PostCfpResponse.
// Service: Traceability
// Router: [POST] /cfp
// Usage: output
type PostCfpResponse struct {
	TraceID string `json:"traceId"`
	CfpID   string `json:"cfpId"`
}

// PostCfpResponses
// Summary: This is a type that defines a list of PostCfpResponse.
type PostCfpResponses []PostCfpResponse

// GetCfpID
// Summary: This is function which gets cfp ID.
// output: (string) ID of the cfp
func (r PostCfpResponses) GetCfpID() string {
	if len(r) > 0 {
		return r[0].CfpID
	}
	return ""
}
