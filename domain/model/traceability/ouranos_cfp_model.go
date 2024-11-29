package traceability

import (
	"fmt"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CfpModel
// Summary: This is structure which defines CfpModel.
// Service: Dataspace
// Router: [GET] /api/v1/datatransport?dataTarget=cfp
// Usage: output
type CfpModel struct {
	CfpID           *uuid.UUID      `json:"cfpId"`
	TraceID         uuid.UUID       `json:"traceId"`
	GhgEmission     *float64        `json:"ghgEmission"`
	GhgDeclaredUnit GhgDeclaredUnit `json:"ghgDeclaredUnit"`
	CfpType         string          `json:"cfpType"`
	DqrType         string          `json:"dqrType"`
	DqrValue        DqrValue        `json:"dqrValue"`
}

// DqrValue
// Summary: This is structure which defines DqrValue.
type DqrValue struct {
	TeR *float64 `json:"TeR"`
	GeR *float64 `json:"GeR"`
	TiR *float64 `json:"TiR"`
}

// CfpModels
// Summary: This is a type that defines a list of CfpModel.
type CfpModels []CfpModel

// CfpDeclaredUnit
// Summary: This is enum which defines GhgDeclaredUnit.
type GhgDeclaredUnit string

const (
	GhgDeclaredUnitKgCO2ePerLiter        GhgDeclaredUnit = "kgCO2e/liter"
	GhgDeclaredUnitKgCO2ePerKilogram     GhgDeclaredUnit = "kgCO2e/kilogram"
	GhgDeclaredUnitKgCO2ePerCubicMeter   GhgDeclaredUnit = "kgCO2e/cubic-meter"
	GhgDeclaredUnitKgCO2ePerKilowattHour GhgDeclaredUnit = "kgCO2e/kilowatt-hour"
	GhgDeclaredUnitKgCO2ePerMegajoule    GhgDeclaredUnit = "kgCO2e/megajoule"
	GhgDeclaredUnitKgCO2ePerTonKilometer GhgDeclaredUnit = "kgCO2e/ton-kilometer"
	GhgDeclaredUnitKgCO2ePerSquareMeter  GhgDeclaredUnit = "kgCO2e/square-meter"
	GhgDeclaredUnitKgCO2ePerUnit         GhgDeclaredUnit = "kgCO2e/unit"
)

// NewGhgDeclaredUnit
// Summary: This is the function to create new GhgDeclaredUnit.
// input: s(string) GhgDeclaredUnit string
// output: (GhgDeclaredUnit) GhgDeclaredUnit
// output: (error) error object
func NewGhgDeclaredUnit(s string) (GhgDeclaredUnit, error) {
	switch s {
	case GhgDeclaredUnitKgCO2ePerLiter.ToString():
		return GhgDeclaredUnitKgCO2ePerLiter, nil
	case GhgDeclaredUnitKgCO2ePerKilogram.ToString():
		return GhgDeclaredUnitKgCO2ePerKilogram, nil
	case GhgDeclaredUnitKgCO2ePerCubicMeter.ToString():
		return GhgDeclaredUnitKgCO2ePerCubicMeter, nil
	case GhgDeclaredUnitKgCO2ePerKilowattHour.ToString():
		return GhgDeclaredUnitKgCO2ePerKilowattHour, nil
	case GhgDeclaredUnitKgCO2ePerMegajoule.ToString():
		return GhgDeclaredUnitKgCO2ePerMegajoule, nil
	case GhgDeclaredUnitKgCO2ePerTonKilometer.ToString():
		return GhgDeclaredUnitKgCO2ePerTonKilometer, nil
	case GhgDeclaredUnitKgCO2ePerSquareMeter.ToString():
		return GhgDeclaredUnitKgCO2ePerSquareMeter, nil
	case GhgDeclaredUnitKgCO2ePerUnit.ToString():
		return GhgDeclaredUnitKgCO2ePerUnit, nil
	default:
		return "", fmt.Errorf(common.UnexpectedEnumError("GhgDeclaredUnit", s))
	}
}

// ToString
// Summary: This is the function to convert GhgDeclaredUnit to string.
// output: (string) converted to string
func (e GhgDeclaredUnit) ToString() string {
	return string(e)
}

// CfpEntityModel
// Summary: This is structure which defines CfpEntityModel.
// DBName: cfp_infomation
type CfpEntityModel struct {
	CfpID              *uuid.UUID     `json:"cfpId" gorm:"type:uuid"`
	TraceID            uuid.UUID      `json:"traceId" gorm:"type:uuid;not null"`
	GhgEmission        *float64       `json:"ghgEmission" gorm:"type:number"`
	GhgDeclaredUnit    string         `json:"ghgDeclaredUnit" gorm:"type:string"`
	CfpCertificateList []string       `json:"cfpCertificateList" gorm:"-"`
	CfpType            string         `json:"cfpType" gorm:"type:string"`
	DqrType            string         `json:"dqrType" gorm:"type:string"`
	TeR                *float64       `json:"TeR" gorm:"type:number"`
	GeR                *float64       `json:"GeR" gorm:"type:number"`
	TiR                *float64       `json:"TiR" gorm:"type:number"`
	DeletedAt          gorm.DeletedAt `json:"deletedAt"`
	CreatedAt          time.Time      `json:"createdAt" gorm:"<-:create "`
	CreatedUserId      string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	UpdatedUserId      string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
}

// CfpEntityModels
// Summary: This is a type that defines a list of CfpEntityModel.
type CfpEntityModels []*CfpEntityModel

// CfpEntityModelSet
// Summary: This is a type that defines a set of CfpEntityModel.
type CfpEntityModelSet []*CfpEntityModel

// CfpCertificateEntityModel
// Summary: This is structure which defines CfpCertificateEntityModel.
type CfpCertificateEntityModel struct {
	Id             int            `json:"id" gorm:"type:int"`
	CfpID          uuid.UUID      `json:"cfpId" gorm:"type:uuid"`
	CfpCertificate string         `json:"cfpCertificate"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"<-:create "`
	CreatedUserId  string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	UpdatedUserId  string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
}

// CfpType
// Summary: This is enum which defines CfpType.
type CfpType string

const (
	CfpTypePreProduction          CfpType = "preProduction"
	CfpTypeMainProduction         CfpType = "mainProduction"
	CfpTypePreComponent           CfpType = "preComponent"
	CfpTypeMainComponent          CfpType = "mainComponent"
	CfpTypePreProductionTotal     CfpType = "preProductionTotal"
	CfpTypeMainProductionTotal    CfpType = "mainProductionTotal"
	CfpTypePreComponentTotal      CfpType = "preComponentTotal"
	CfpTypeMainComponentTotal     CfpType = "mainComponentTotal"
	CfpTypePreProductionResponse  CfpType = "preProductionResponse"
	CfpTypeMainProductionResponse CfpType = "mainProductionResponse"
)

// ToString
// Summary: This is the function to convert CfpType to string.
// output: (string) converted to string
func (e CfpType) ToString() string {
	return string(e)
}

// IsTotal
// Summary: This is the function to check if CfpType is total.
// output: (bool) true: total, false: not total
func (e CfpType) IsTotal() bool {
	return e == CfpTypePreProductionTotal || e == CfpTypeMainProductionTotal || e == CfpTypePreComponentTotal || e == CfpTypeMainComponentTotal
}

// ToDqrType
// Summary: This is the function to convert CfpType to DqrType.
// output: (DqrType) dqr type
// output: (error) error object
func (e CfpType) ToDqrType() (DqrType, error) {
	switch e {
	case CfpTypePreProduction:
		return DqrTypePreProcessing, nil
	case CfpTypeMainProduction:
		return DqrTypeMainProcessing, nil
	case CfpTypePreComponent:
		return DqrTypePreProcessing, nil
	case CfpTypeMainComponent:
		return DqrTypeMainProcessing, nil
	case CfpTypePreProductionTotal:
		return DqrTypePreProcessingTotal, nil
	case CfpTypeMainProductionTotal:
		return DqrTypeMainProcessingTotal, nil
	case CfpTypePreComponentTotal:
		return DqrTypePreProcessingTotal, nil
	case CfpTypeMainComponentTotal:
		return DqrTypeMainProcessingTotal, nil
	case CfpTypePreProductionResponse:
		return DqrPreProcessingResponse, nil
	case CfpTypeMainProductionResponse:
		return DqrMainProcessingResponse, nil
	default:
		return "", fmt.Errorf(common.UnexpectedEnumError("CfpType", e.ToString()))
	}
}

// DqrType
// Summary: This is enum which defines DqrType.
type DqrType string

const (
	DqrTypePreProcessing       DqrType = "preProcessing"
	DqrTypeMainProcessing      DqrType = "mainProcessing"
	DqrTypePreProcessingTotal  DqrType = "preProcessingTotal"
	DqrTypeMainProcessingTotal DqrType = "mainProcessingTotal"
	DqrPreProcessingResponse   DqrType = "preProcessingResponse"
	DqrMainProcessingResponse  DqrType = "mainProcessingResponse"
)

// ToString
// Summary: This is the function to convert DqrType to string.
// output: (string) converted to string
func (e DqrType) ToString() string {
	return string(e)
}

// GetCfpInput
// Summary: This is structure which defines GetCfpInput.
type GetCfpInput struct {
	OperatorID uuid.UUID
	TraceIDs   []uuid.UUID
}

// PutCfpInput
// Summary: This is structure which defines PutCfpInput.
type PutCfpInput struct {
	CfpID           *string          `json:"cfpId"`
	TraceID         string           `json:"traceId"`
	GhgEmission     *float64         `json:"ghgEmission"`
	GhgDeclaredUnit string           `json:"ghgDeclaredUnit"`
	CfpType         CfpType          `json:"cfpType"`
	DqrType         DqrType          `json:"dqrType"`
	DqrValue        PutDqrValueInput `json:"dqrValue"`
}

// PutDqrValueInput
// Summary: This is structure which defines PutDqrValueInput.
type PutDqrValueInput struct {
	TeR *float64 `json:"TeR"`
	GeR *float64 `json:"GeR"`
	TiR *float64 `json:"TiR"`
}

// PutCfpInputs
// Summary: This is a type that defines a list of PutCfpInput.
// Service: Dataspace
// Router: [PUT] /api/v1/datatransport?dataTarget=cfp
// Usage: input
type PutCfpInputs []PutCfpInput

// GenerateCfpEntitisFromModels
// Summary: This is the function to generate CfpEntityModels from CfpModels.
// input: cfpModels(CfpModels) CfpModels object
// output: (CfpEntityModels) CfpEntityModels object
func GenerateCfpEntitisFromModels(cfpModels CfpModels) CfpEntityModels {
	es := make(CfpEntityModels, len(cfpModels))
	cfpId := uuid.New()
	for i, m := range cfpModels {
		e := NewCfpEntityModelWithID(
			cfpId,
			m.TraceID,
			m.GhgEmission,
			m.GhgDeclaredUnit.ToString(),
			m.CfpType,
			m.DqrType,
			m.DqrValue.TeR,
			m.DqrValue.GeR,
			m.DqrValue.TiR,
		)
		es[i] = &e
	}
	return es
}

// NewCfpEntityModel
// Summary: This is the function to create new CfpEntityModel.
// input: TraceID(uuid.UUID) ID of the trace
// input: GhgEmission(*float64) GHG emission value
// input: GhgDeclaredUnit(string) GHG declared unit
// input: CfpCertificateList([]string) list of cfp certificate
// input: CfpType(string) cfp type
// input: DqrType(string) dqr type
// input: TeR(*float64) TeR value
// input: GeR(*float64) GeR value
// input: TiR(*float64) TiR value
// output: (CfpEntityModel) CfpEntityModel object
func NewCfpEntityModel(
	TraceID uuid.UUID,
	GhgEmission *float64,
	GhgDeclaredUnit string,
	CfpCertificateList []string,
	CfpType string,
	DqrType string,
	TeR *float64,
	GeR *float64,
	TiR *float64,
) CfpEntityModel {
	t := time.Now()
	cfpID := uuid.New()
	return CfpEntityModel{
		CfpID:              &cfpID,
		TraceID:            TraceID,
		GhgEmission:        GhgEmission,
		GhgDeclaredUnit:    GhgDeclaredUnit,
		CfpCertificateList: CfpCertificateList,
		CfpType:            CfpType,
		DqrType:            DqrType,
		TeR:                TeR,
		GeR:                GeR,
		TiR:                TiR,
		CreatedAt:          t,
		DeletedAt:          gorm.DeletedAt{},
		CreatedUserId:      "tmp",
		UpdatedAt:          t,
		UpdatedUserId:      "tmp",
	}
}

// NewCfpEntityModelWithID
// Summary: This is the function to create new CfpEntityModel with ID.
// input: CfpID(uuid.UUID) ID of the cfp
// input: TraceID(uuid.UUID) ID of the trace
// input: GhgEmission(*float64) GHG emission value
// input: GhgDeclaredUnit(string) GHG declared unit
// input: CfpType(string) cfp type
// input: DqrType(string) dqr type
// input: TeR(*float64) TeR value
// input: GeR(*float64) GeR value
// input: TiR(*float64) TiR value
// output: (CfpEntityModel) CfpEntityModel object
func NewCfpEntityModelWithID(
	CfpID uuid.UUID,
	TraceID uuid.UUID,
	GhgEmission *float64,
	GhgDeclaredUnit string,
	CfpType string,
	DqrType string,
	TeR *float64,
	GeR *float64,
	TiR *float64,
) CfpEntityModel {
	t := time.Now()
	return CfpEntityModel{
		CfpID:           &CfpID,
		TraceID:         TraceID,
		GhgEmission:     GhgEmission,
		GhgDeclaredUnit: GhgDeclaredUnit,
		CfpType:         CfpType,
		DqrType:         DqrType,
		TeR:             TeR,
		GeR:             GeR,
		TiR:             TiR,
		CreatedAt:       t,
		DeletedAt:       gorm.DeletedAt{},
		CreatedUserId:   "tmp",
		UpdatedAt:       t,
		UpdatedUserId:   "tmp",
	}
}

// NewCfpEntityModelSetFromCfpEntityModels
// Summary: This is the function to create new CfpEntityModelSet from CfpEntityModels.
// input: cfpEntityModels(CfpEntityModels) CfpEntityModels object
// input: isParent(bool) true: parent, false: not parent
// output: (CfpEntityModelSet) CfpEntityModelSet object
// output: (error) error object
func NewCfpEntityModelSetFromCfpEntityModels(cfpEntityModels CfpEntityModels, isParent bool) (CfpEntityModelSet, error) {
	cfpID, err := cfpEntityModels.GetCommonCfpID()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return CfpEntityModelSet{}, err
	}
	traceID, err := cfpEntityModels.GetCommonTraceID()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return CfpEntityModelSet{}, err
	}
	ghgDeclaredUnit, err := cfpEntityModels.GetCommonGhgDeclaredUnit()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return CfpEntityModelSet{}, err
	}
	cfpCertificateList, err := cfpEntityModels.GetCommonCfpCertificateList()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return CfpEntityModelSet{}, err
	}

	var cfpSet CfpEntityModelSet
	if v := cfpEntityModels.GetPreProductionCfp(); v != nil {
		cfpSet = append(cfpSet, v)
	} else {
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "preProduction", DqrType: DqrTypePreProcessing.ToString()})
	}

	if v := cfpEntityModels.GetPreComponentCfp(); v != nil {
		cfpSet = append(cfpSet, v)
	} else {
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "preComponent", DqrType: DqrTypePreProcessing.ToString()})
	}

	if v := cfpEntityModels.GetMainProductionCfp(); v != nil {
		cfpSet = append(cfpSet, v)
	} else {
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "mainProduction", DqrType: DqrTypeMainProcessing.ToString()})
	}

	if v := cfpEntityModels.GetMainComponentCfp(); v != nil {
		cfpSet = append(cfpSet, v)
	} else {
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "mainComponent", DqrType: DqrTypeMainProcessing.ToString()})
	}
	if isParent {
		defaultGhgEmission := 0.0
		staticTeR, staticGeR := common.Float64Ptr(2.1), common.Float64Ptr(0)
		preProductionCfp := cfpSet.GetPreProductionCfp()
		mainProductionCfp := cfpSet.GetMainProductionCfp()
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "preProductionTotal", GhgEmission: preProductionCfp.GhgEmission, GhgDeclaredUnit: ghgDeclaredUnit, CfpCertificateList: cfpCertificateList, DqrType: DqrTypePreProcessingTotal.ToString(), TeR: staticTeR, GeR: staticGeR, TiR: nil})
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "preComponentTotal", GhgEmission: &defaultGhgEmission, GhgDeclaredUnit: ghgDeclaredUnit, CfpCertificateList: cfpCertificateList, DqrType: DqrTypePreProcessingTotal.ToString(), TeR: staticTeR, GeR: staticGeR, TiR: nil})
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "mainProductionTotal", GhgEmission: mainProductionCfp.GhgEmission, GhgDeclaredUnit: ghgDeclaredUnit, CfpCertificateList: cfpCertificateList, DqrType: DqrTypeMainProcessingTotal.ToString(), TeR: staticTeR, GeR: staticGeR, TiR: nil})
		cfpSet = append(cfpSet, &CfpEntityModel{CfpID: &cfpID, TraceID: traceID, CfpType: "mainComponentTotal", GhgEmission: &defaultGhgEmission, GhgDeclaredUnit: ghgDeclaredUnit, CfpCertificateList: cfpCertificateList, DqrType: DqrTypeMainProcessingTotal.ToString(), TeR: staticTeR, GeR: staticGeR, TiR: nil})
	}

	return cfpSet, nil
}

// NewCfpCertificateEntityModel
// Summary: This is the function to create new CfpCertificateEntityModel.
// input: id(int) ID of the cfp certificate
// input: cfpID(uuid.UUID) ID of the cfp
// input: cfpCertificate(string) cfp certificate
// output: (CfpCertificateEntityModel) CfpCertificateEntityModel object
func NewCfpCertificationEntityModel(
	id int,
	cfpID uuid.UUID,
	cfpCertificate string,
) CfpCertificateEntityModel {
	t := time.Now()
	return CfpCertificateEntityModel{
		Id:             id,
		CfpID:          cfpID,
		CfpCertificate: cfpCertificate,
		CreatedAt:      t,
		DeletedAt:      gorm.DeletedAt{},
		CreatedUserId:  "tmp",
		UpdatedAt:      t,
		UpdatedUserId:  "tmp",
	}
}

// NewEmptyCfpResponse
// Summary: This is the function to create new empty CfpEntityModels.
// input: traceID(uuid.UUID) ID of the trace
// output: (CfpEntityModels) CfpEntityModels object
func NewEmptyCfpResponse(traceID uuid.UUID) CfpEntityModels {
	pre := CfpEntityModel{TraceID: traceID, CfpType: "preProductionResponse", DqrType: DqrPreProcessingResponse.ToString()}
	main := CfpEntityModel{TraceID: traceID, CfpType: "mainProductionResponse", DqrType: DqrMainProcessingResponse.ToString()}
	cfps := CfpEntityModels{&pre, &main}
	return cfps
}

// ExtractByCfpType
// Summary: This is the function to extract CfpModel by CfpType.
// input: cfpType(CfpType) cfp type
// output: (CfpModel) CfpModel object
// output: (error) error object
func (ms CfpModels) ExtractByCfpType(cfpType CfpType) (CfpModel, error) {
	for _, m := range ms {
		if m.CfpType == cfpType.ToString() {
			return m, nil
		}
	}
	err := fmt.Errorf(common.CfpTypeNotFoundError(cfpType.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return CfpModel{}, err
}

// GetGhgEmission
// Summary: This is the function to get GHG emission value.
// input: cfpType(CfpType) cfp type
// output: (float64) GHG emission value
// output: (error) error object
func (ms CfpModels) GetGhgEmission(cfpType CfpType) (float64, error) {
	m, err := ms.ExtractByCfpType(cfpType)
	if err != nil {
		return 0, err
	}
	if m.GhgEmission == nil {
		return 0, err
	}
	return *m.GhgEmission, nil
}

// GetCommonCfpID
// Summary: This is the function to get common cfp ID.
// output: (*uuid.UUID) ID of the cfp
// output: (error) error object
func (ms CfpModels) GetCommonCfpID() (*uuid.UUID, error) {
	if !ms.isSameCfpID() {
		return nil, fmt.Errorf(common.InconsistentFieldError("cfpId"))
	}
	return ms[0].CfpID, nil
}

// GetCommonTraceID
// Summary: This is the function to get common trace ID.
// output: (uuid.UUID) ID of the trace
// output: (error) error object
func (ms CfpModels) GetCommonTraceID() (uuid.UUID, error) {
	if !ms.isSameTraceID() {
		return uuid.Nil, fmt.Errorf(common.InconsistentFieldError("traceId"))
	}
	return ms[0].TraceID, nil
}

// GetCommonGhgDeclaredUnit
// Summary: This is the function to get common GHG declared unit.
// output: (string) GHG declared unit
// output: (error) error object
func (ms CfpModels) GetCommonGhgDeclaredUnit() (string, error) {
	if !ms.isSameGhgDeclaredUnit() {
		return "", fmt.Errorf(common.InconsistentFieldError("ghgDeclaredUnits"))
	}
	return ms[0].GhgDeclaredUnit.ToString(), nil
}

// GetPreProcessingTeR
// Summary: This is the function to get pre-processing TeR value.
// output: (*float64) TeR value
// output: (error) error object
func (ms CfpModels) GetPreProcessingTeR() (*float64, error) {
	for _, m := range ms {
		if m.DqrType == DqrTypePreProcessing.ToString() {
			return m.DqrValue.TeR, nil
		}
	}
	err := fmt.Errorf(common.DqrTypeNotFoundError(DqrTypePreProcessing.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return nil, err
}

// GetPreProcessingGeR
// Summary: This is the function to get pre-processing GeR value.
// output: (*float64) GeR value
// output: (error) error object
func (ms CfpModels) GetPreProcessingGeR() (*float64, error) {
	for _, m := range ms {
		if m.DqrType == DqrTypePreProcessing.ToString() {
			return m.DqrValue.GeR, nil
		}
	}
	err := fmt.Errorf(common.DqrTypeNotFoundError(DqrTypePreProcessing.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return nil, err
}

// GetPreProcessingTiR
// Summary: This is the function to get pre-processing TiR value.
// output: (*float64) TiR value
// output: (error) error object
func (ms CfpModels) GetPreProcessingTiR() (*float64, error) {
	for _, m := range ms {
		if m.DqrType == DqrTypePreProcessing.ToString() {
			return m.DqrValue.TiR, nil
		}
	}
	err := fmt.Errorf(common.DqrTypeNotFoundError(DqrTypePreProcessing.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return nil, err
}

// GetMainProductionTeR
// Summary: This is the function to get main-processing TeR value.
// output: (*float64) TeR value
// output: (error) error object
func (ms CfpModels) GetMainProductionTeR() (*float64, error) {
	for _, m := range ms {
		if m.DqrType == DqrTypeMainProcessing.ToString() {
			return m.DqrValue.TeR, nil
		}
	}
	err := fmt.Errorf(common.DqrTypeNotFoundError(DqrTypeMainProcessing.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return nil, err
}

// GetMainProductionGeR
// Summary: This is the function to get main-processing GeR value.
// output: (*float64) GeR value
// output: (error) error object
func (ms CfpModels) GetMainProductionGeR() (*float64, error) {
	for _, m := range ms {
		if m.DqrType == DqrTypeMainProcessing.ToString() {
			return m.DqrValue.GeR, nil
		}
	}
	err := fmt.Errorf(common.DqrTypeNotFoundError(DqrTypeMainProcessing.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return nil, err
}

// GetMainProductionTiR
// Summary: This is the function to get main-processing TiR value.
// output: (*float64) TiR value
// output: (error) error object
func (ms CfpModels) GetMainProductionTiR() (*float64, error) {
	for _, m := range ms {
		if m.DqrType == DqrTypeMainProcessing.ToString() {
			return m.DqrValue.TiR, nil
		}
	}
	err := fmt.Errorf(common.DqrTypeNotFoundError(DqrTypeMainProcessing.ToString()))
	logger.Set(nil).Errorf(err.Error())
	return nil, err
}

// isSameCfpID
// Summary: This is the function to checks if all cfp IDs are the same.
// output: (bool) true: same, false: not same
func (ms CfpModels) isSameCfpID() bool {
	if len(ms) == 0 {
		return false
	}
	firstCfpID := ms[0].CfpID
	for _, e := range ms {
		if firstCfpID == nil && e.CfpID == nil {
			continue
		}

		if (firstCfpID == nil && e.CfpID != nil) || (firstCfpID != nil && e.CfpID == nil) {
			return false
		}

		if e.CfpID.String() != firstCfpID.String() {
			return false
		}
	}
	return true
}

// isSameTraceID
// Summary: This is the function to checks if all trace IDs are the same.
// output: (bool) true: same, false: not same
func (ms CfpModels) isSameTraceID() bool {
	if len(ms) == 0 {
		return false
	}
	firstTraceID := ms[0].TraceID
	for _, e := range ms {
		if e.TraceID != firstTraceID {
			return false
		}
	}
	return true
}

// isSameGhgDeclaredUnit
// Summary: This is the function to checks if all GHG declared units are the same.
// output: (bool) true: same, false: not same
func (ms CfpModels) isSameGhgDeclaredUnit() bool {
	if len(ms) == 0 {
		return false
	}
	firstGhgDeclaredUnit := ms[0].GhgDeclaredUnit
	for _, e := range ms {
		if e.GhgDeclaredUnit != firstGhgDeclaredUnit {
			return false
		}
	}
	return true
}

// SetCfpID
// Summary: This is the function to set cfp ID.
// input: cfpID(uuid.UUID) ID of the cfp
func (ms *CfpModels) SetCfpID(cfpID uuid.UUID) {
	for i := range *ms {
		(*ms)[i].CfpID = &cfpID
	}
}

// SortCfpModelsByTraceIDs
// Summary: This is the function to sort CfpModels by trace IDs.
// input: traceIDs([]uuid.UUID) list of trace IDs
// output: (CfpModels) sorted CfpModels
func (ms CfpModels) SortCfpModelsByTraceIDs(traceIDs []uuid.UUID) CfpModels {
	var sortedCfpModels CfpModels = CfpModels{}
	for _, traceID := range traceIDs {
		for _, m := range ms {
			if m.TraceID == traceID {
				sortedCfpModels = append(sortedCfpModels, m)
			}
		}
	}
	return sortedCfpModels
}

// Validate
// Summary: This is the function to validate PutCfpInputs.
// output: (error) error object
func (is PutCfpInputs) Validate() error {
	var errors []error
	for i, input := range is {
		if err := input.validate(); err != nil {
			newErr := fmt.Errorf("cfpModel[%d]: (%v)", i, err)
			errors = append(errors, newErr)
		}
	}

	if len(is) != 4 {
		errors = append(errors, fmt.Errorf(common.CfpElementsError()))
	}
	for _, cfpType := range []CfpType{CfpTypePreProduction, CfpTypeMainProduction, CfpTypePreComponent, CfpTypeMainComponent} {
		if !is.hasCfpType(cfpType) {
			errors = append(errors, fmt.Errorf(common.InsufficientCfpType(cfpType.ToString())))
		}
	}

	first := is[0]
	for _, i := range is {
		if !common.IsStrPtrValEqual(first.CfpID, i.CfpID) {
			errors = append(errors, fmt.Errorf(common.InconsistentFieldError("cfpId")))
		}
		if first.TraceID != i.TraceID {
			errors = append(errors, fmt.Errorf(common.InconsistentFieldError("traceId")))
		}
		if first.GhgDeclaredUnit != i.GhgDeclaredUnit {
			errors = append(errors, fmt.Errorf(common.InconsistentFieldError("ghgDeclaredUnit")))
		}
	}

	if errs := is.ValidateGhgEmission(); errs != nil {
		errors = append(errors, errs...)
	}

	if !is.isValidDqrValue() {
		errors = append(errors, fmt.Errorf(common.DqrValueInconsistentError()))
	}

	if len(errors) > 0 {
		return common.JoinErrors(errors)
	}

	return nil
}

// ValidateGhgEmission
// Summary: This is the function to validate GHG emission.
// output: ([]error) list of errors
func (is PutCfpInputs) ValidateGhgEmission() []error {
	var errors []error
	var preProcessingDqrValue *PutDqrValueInput
	var preProcessingGhgEmission *float64
	var mainProcessingDqrValue *PutDqrValueInput
	var mainProcessingGhgEmission *float64
	for _, i := range is {
		if i.DqrType == DqrTypePreProcessing {
			if i.DqrValue.hasValue() {
				preProcessingDqrValue = &i.DqrValue
			}
			if i.GhgEmission != nil && *i.GhgEmission > 0 {
				preProcessingGhgEmission = i.GhgEmission
			}
			continue
		}
		if i.DqrType == DqrTypeMainProcessing {
			if i.DqrValue.hasValue() {
				mainProcessingDqrValue = &i.DqrValue
			}
			if i.GhgEmission != nil && *i.GhgEmission > 0 {
				mainProcessingGhgEmission = i.GhgEmission
			}
			continue
		}
	}

	if preProcessingDqrValue != nil && preProcessingGhgEmission == nil {
		errors = append(errors, fmt.Errorf(common.InvalidGhgEmission(string(CfpTypePreProduction), string(CfpTypePreComponent))))
	}
	if mainProcessingDqrValue != nil && mainProcessingGhgEmission == nil {
		errors = append(errors, fmt.Errorf(common.InvalidGhgEmission(string(CfpTypeMainProduction), string(CfpTypeMainComponent))))
	}

	return errors
}

// isValidDqrValue
// Summary: This is the function that validates all dqr values.
// output: (bool) true: valid, false: invalid
func (is PutCfpInputs) isValidDqrValue() bool {
	var firstPreProcessing PutDqrValueInput
	var firstMainProcessing PutDqrValueInput
	for _, i := range is {
		if i.DqrType == DqrTypePreProcessing {
			if firstPreProcessing == (PutDqrValueInput{}) {
				firstPreProcessing = i.DqrValue
				continue
			}

			if !i.DqrValue.isSame(firstPreProcessing) {
				return false
			}
		}
		if i.DqrType == DqrTypeMainProcessing {
			if firstMainProcessing == (PutDqrValueInput{}) {
				firstMainProcessing = i.DqrValue
				continue
			}

			if !i.DqrValue.isSame(firstMainProcessing) {
				return false
			}
		}
	}
	return true
}

// isSame
// Summary: This is the function that checks if all dqr values are the same.
// input: other(PutDqrValueInput) PutDqrValueInput object
// output: (bool) true: same, false: not same
func (i PutDqrValueInput) isSame(other PutDqrValueInput) bool {
	if (i.TeR != nil && other.TeR == nil) || (i.TeR == nil && other.TeR != nil) || *i.TeR != *other.TeR {
		return false
	}
	if (i.GeR != nil && other.GeR == nil) || (i.GeR == nil && other.GeR != nil) || *i.GeR != *other.GeR {
		return false
	}
	if (i.TiR != nil && other.TiR == nil) || (i.TiR == nil && other.TiR != nil) || *i.TiR != *other.TiR {
		return false
	}
	return true
}

// hasValue
// Summary: This is the function that checks if dqr value has value.
// output: (bool) true: has value, false: no value
func (i PutDqrValueInput) hasValue() bool {
	if i.TeR != nil && *i.TeR > 0 {
		return true
	}
	if i.GeR != nil && *i.GeR > 0 {
		return true
	}
	if i.TiR != nil && *i.TiR > 0 {
		return true
	}
	return false
}

// Validate
// Summary: This is the function to validate PutCfpInput.
// output: (error) error object
func (i PutCfpInput) Validate() error {
	return i.validate()
}

// validate
// Summary: This is the function to validate PutCfpInput.
// output: (error) error object
func (i PutCfpInput) validate() error {
	errors := []error{}
	err := validation.ValidateStruct(&i,
		validation.Field(
			&i.CfpID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&i.TraceID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.GhgEmission,
			validation.Min(0.00000),
			validation.Max(99999.99999),
			validation.By(common.FloatPtr5thDecimal),
		),
		validation.Field(
			&i.GhgDeclaredUnit,
			validation.Required,
			validation.By(EnumGhgDeclaredUnitValid),
		),
		validation.Field(
			&i.CfpType,
			validation.Required,
			validation.By(EnumCfpTypeValid),
		),
		validation.Field(
			&i.DqrType,
			validation.Required,
			validation.By(EnumDqrTypeValidForPutCfp),
		),
	)
	if err != nil {
		errors = append(errors, err)
	}

	var cfpTypeAndDqrTypeErr error
	if !i.isValidCfpTypeAndDqrType() {
		cfpTypeAndDqrTypeErr = fmt.Errorf(common.InvalidCombination("cfpType", "dqrType"))
		errors = append(errors, cfpTypeAndDqrTypeErr)
	}

	var dqrValueErr error
	if err := i.DqrValue.validate(); err != nil {
		dqrValueErr = fmt.Errorf("dqrValue: (%v)", err)
		errors = append(errors, dqrValueErr)
	}

	if len(errors) > 0 {
		if dqrValueErr != nil || cfpTypeAndDqrTypeErr != nil {
			return common.JoinErrors(errors)
		} else {
			return err
		}
	}

	return nil
}

// isValidCfpTypeAndDqrType
// Summary: This is the function that checks if cfp type and dqr type are valid.
// output: (bool) true: valid, false: invalid
func (i PutCfpInput) isValidCfpTypeAndDqrType() bool {
	if i.CfpType == CfpTypePreProduction && i.DqrType != DqrTypePreProcessing {
		return false
	}

	if i.CfpType == CfpTypeMainProduction && i.DqrType != DqrTypeMainProcessing {
		return false
	}

	if i.CfpType == CfpTypePreComponent && i.DqrType != DqrTypePreProcessing {
		return false
	}

	if i.CfpType == CfpTypeMainComponent && i.DqrType != DqrTypeMainProcessing {
		return false
	}

	return true
}

// validate
// Summary: This is the function to validate PutDqrValueInput.
// output: (error) error object
func (i PutDqrValueInput) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.TeR,
			validation.Min(0.00000),
			validation.Max(99999.99999),
			validation.By(common.FloatPtr5thDecimal),
		),
		validation.Field(
			&i.GeR,
			validation.Min(0.00000),
			validation.Max(99999.99999),
			validation.By(common.FloatPtr5thDecimal),
		),
		validation.Field(
			&i.TiR,
			validation.Min(0.00000),
			validation.Max(99999.99999),
			validation.By(common.FloatPtr5thDecimal),
		),
	)
}

// ToModels
// Summary: This is the function to convert PutCfpInputs to CfpModels.
// output: (CfpModels) CfpModels object
// output: (error) error object
func (is PutCfpInputs) ToModels() (CfpModels, error) {
	ms := make(CfpModels, len(is))
	for i, input := range is {
		m, err := input.ToModel()
		if err != nil {
			return nil, err
		}
		ms[i] = m
	}
	return ms, nil
}

// ToModel
// Summary: This is the function to convert PutCfpInput to CfpModel.
// output: (CfpModel) CfpModel object
// output: (error) error object
func (i PutCfpInput) ToModel() (CfpModel, error) {
	var model CfpModel

	if i.CfpID != nil {
		cfpID, err := uuid.Parse(*i.CfpID)
		if err != nil {
			logger.Set(nil).Warnf(err.Error())

			return CfpModel{}, fmt.Errorf(common.InvalidUUIDError("cfpId"))
		}
		model.CfpID = &cfpID
	}

	traceID, err := uuid.Parse(i.TraceID)
	if err != nil {
		logger.Set(nil).Warnf(err.Error())

		return CfpModel{}, fmt.Errorf(common.InvalidUUIDError("traceId"))
	}
	model.TraceID = traceID

	ghgDeclaredUnit, err := NewGhgDeclaredUnit(i.GhgDeclaredUnit)
	if err != nil {
		logger.Set(nil).Warnf(err.Error())

		return CfpModel{}, err
	}
	model.GhgEmission = i.GhgEmission
	model.GhgDeclaredUnit = ghgDeclaredUnit
	model.CfpType = i.CfpType.ToString()
	model.DqrType = i.DqrType.ToString()
	model.DqrValue = DqrValue{
		TeR: i.DqrValue.TeR,
		GeR: i.DqrValue.GeR,
		TiR: i.DqrValue.TiR,
	}

	return model, nil
}

// hasCfpType
// Summary: This is the function to checks if cfp type exists in PutCfpInputs.
// input: cfpType(CfpType) cfp type
// output: (bool) true: exists, false: not exists
func (is PutCfpInputs) hasCfpType(cfpType CfpType) bool {
	for _, i := range is {
		if i.CfpType == cfpType {
			return true
		}
	}
	return false
}

// Update
// Summary: This is the function to update CfpEntityModel.
// input: ghgEmission(*float64) GHG emission value
// input: ghgDeclaredUnit(string) GHG declared unit
// input: dqrType(string) dqr type
// input: TeR(*float64) TeR value
// input: GeR(*float64) GeR value
// input: TiR(*float64) TiR value
func (e *CfpEntityModel) Update(
	ghgEmission *float64,
	ghgDeclaredUnit string,
	dqrType string,
	TeR *float64,
	GeR *float64,
	TiR *float64,
) {
	e.GhgEmission = ghgEmission
	e.GhgDeclaredUnit = ghgDeclaredUnit
	e.DqrType = dqrType
	e.TeR = TeR
	e.GeR = GeR
	e.TiR = TiR
}

// ToModel
// Summary: This is the function to convert CfpEntityModel to CfpModel.
// output: (CfpModel) CfpModel object
// output: (error) error object
func (e CfpEntityModel) ToModel() (CfpModel, error) {
	ghgDeclaredUnit, err := NewGhgDeclaredUnit(e.GhgDeclaredUnit)
	if err != nil {
		logger.Set(nil).Warnf(err.Error())

		return CfpModel{}, err
	}
	res := CfpModel{
		CfpID:           e.CfpID,
		TraceID:         e.TraceID,
		GhgEmission:     e.GhgEmission,
		GhgDeclaredUnit: ghgDeclaredUnit,
		CfpType:         e.CfpType,
		DqrType:         e.DqrType,
		DqrValue: DqrValue{
			TeR: e.TeR,
			GeR: e.GeR,
			TiR: e.TiR,
		},
	}
	return res, nil
}

// GetCommonCfpID
// Summary: This is the function to get common cfp ID.
// output: (uuid.UUID) ID of the cfp
// output: (error) error object
func (es CfpEntityModels) GetCommonCfpID() (uuid.UUID, error) {
	if !es.isSameCfpID() {
		return uuid.Nil, fmt.Errorf(common.CfpIDsInconsistentError())
	}
	return *(es[0].CfpID), nil
}

// GetCommonTraceID
// Summary: This is the function to get common trace ID.
// output: (uuid.UUID) ID of the trace
// output: (error) error object
func (es CfpEntityModels) GetCommonTraceID() (uuid.UUID, error) {
	if !es.isSameTraceID() {
		return uuid.Nil, fmt.Errorf(common.TraceIDsInconsistentError())
	}
	return es[0].TraceID, nil
}

// GetCommonGhgDeclaredUnit
// Summary: This is the function to get common GHG declared unit.
// output: (string) GHG declared unit
// output: (error) error object
func (es CfpEntityModels) GetCommonGhgDeclaredUnit() (string, error) {
	if !es.isSameGhgDeclaredUnit() {
		return "", fmt.Errorf(common.GhgDeclaredUnitsInconsistentError())
	}
	return es[0].GhgDeclaredUnit, nil
}

// GetCommonCfpCertificateList
// Summary: This is the function to get common cfp certificate list.
// output: ([]string) list of cfp certificate
// output: (error) error object
func (es CfpEntityModels) GetCommonCfpCertificateList() ([]string, error) {
	if !es.isSameCfpCertificateList() {
		return nil, fmt.Errorf(common.CfpCertificateListInconsistentError())
	}
	return es[0].CfpCertificateList, nil
}

// isSameCfpID
// Summary: This is the function to checks if all cfp IDs are the same.
// output: (bool) true: same, false: not same
func (es CfpEntityModels) isSameCfpID() bool {
	if len(es) == 0 {
		return false
	}
	firstCfpID := es[0].CfpID
	for _, e := range es {
		if e.CfpID.String() != firstCfpID.String() {
			return false
		}
	}
	return true
}

// isSameTraceID
// Summary: This is the function to checks if all trace IDs are the same.
// output: (bool) true: same, false: not same
func (es CfpEntityModels) isSameTraceID() bool {
	if len(es) == 0 {
		return false
	}
	firstTraceID := es[0].TraceID
	for _, e := range es {
		if e.TraceID != firstTraceID {
			return false
		}
	}
	return true
}

// isSameGhgDeclaredUnit
// Summary: This is the function to checks if all GHG declared units are the same.
// output: (bool) true: same, false: not same
func (es CfpEntityModels) isSameGhgDeclaredUnit() bool {
	if len(es) == 0 {
		return false
	}
	firstGhgDeclaredUnit := es[0].GhgDeclaredUnit
	for _, e := range es {
		if e.GhgDeclaredUnit != firstGhgDeclaredUnit {
			return false
		}
	}
	return true
}

// isSameCfpCertificateList
// Summary: This is the function to checks if all cfp certificate lists are the same.
// output: (bool) true: same, false: not same
func (es CfpEntityModels) isSameCfpCertificateList() bool {
	if len(es) == 0 {
		return false
	}
	cfpCertificateList := es[0].CfpCertificateList
	for _, e := range es {
		for i := range e.CfpCertificateList {
			if cfpCertificateList[i] != e.CfpCertificateList[i] {
				return false
			}
		}
	}
	return true
}

// ToModels
// Summary: This is the function to convert CfpEntityModels to CfpModels.
// output: ([]CfpModel) list of CfpModel
// output: (error) error object
func (es CfpEntityModels) ToModels() ([]CfpModel, error) {
	ms := make([]CfpModel, len(es))
	for i, e := range es {
		m, err := e.ToModel()
		if err != nil {
			logger.Set(nil).Warnf(err.Error())

			return nil, err
		}
		ms[i] = m
	}
	return ms, nil
}

// GetPreProductionCfp
// Summary: This is the function to get pre-production cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModels) GetPreProductionCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "preProduction" {
			return e
		}
	}
	return nil
}

// GetPreComponentCfp
// Summary: This is the function to get pre-component cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModels) GetPreComponentCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "preComponent" {
			return e
		}
	}
	return nil
}

// GetMainProductionCfp
// Summary: This is the function to get main-production cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModels) GetMainProductionCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "mainProduction" {
			return e
		}
	}
	return nil
}

// GetMainComponentCfp
// Summary: This is the function to get main-component cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModels) GetMainComponentCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "mainComponent" {
			return e
		}
	}
	return nil
}

// MakeCfpResponse
// Summary: This is the function to make cfp response.
// input: traceID(uuid.UUID) ID of the trace
// output: (CfpEntityModels) CfpEntityModels object
// output: (error) error object
func (es CfpEntityModels) MakeCfpResponse(traceID uuid.UUID) (CfpEntityModels, error) {
	var cfps CfpEntityModels

	staticTeR, staticGeR := common.Float64Ptr(2.1), common.Float64Ptr(0)
	ghgDeclaredUnit, err := es.GetCommonGhgDeclaredUnit()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return nil, err
	}

	if es.GetPreProductionCfp() != nil {
		pre := es.GetPreProductionCfp()
		pre.CfpID = nil
		pre.TraceID = traceID
		pre.CfpType = "preProductionResponse"
		pre.DqrType = DqrPreProcessingResponse.ToString()
		pre.TeR = staticTeR
		pre.GeR = staticGeR
		pre.TiR = nil
		pre.GhgDeclaredUnit = ghgDeclaredUnit
		cfps = append(cfps, pre)
	}

	if es.GetMainProductionCfp() != nil {
		main := es.GetMainProductionCfp()
		main.CfpID = nil
		main.TraceID = traceID
		main.CfpType = "mainProductionResponse"
		main.DqrType = DqrMainProcessingResponse.ToString()
		main.TeR = staticTeR
		main.GeR = staticGeR
		main.TiR = nil
		main.GhgDeclaredUnit = ghgDeclaredUnit
		cfps = append(cfps, main)
	}
	if len(es) == 0 {
		pre := CfpEntityModel{CfpType: "preProductionResponse", DqrType: DqrPreProcessingResponse.ToString(), TeR: staticTeR, GeR: staticGeR, TiR: nil}
		main := CfpEntityModel{CfpType: "mainProductionResponse", DqrType: DqrMainProcessingResponse.ToString(), TeR: staticTeR, GeR: staticGeR, TiR: nil}
		cfps = append(cfps, &pre, &main)
	}
	return cfps, nil
}

// ToModels
// Summary: This is the function to convert CfpEntityModelSet to CfpModels.
// output: ([]CfpModel) list of CfpModel
// output: (error) error object
func (es CfpEntityModelSet) ToModels() ([]CfpModel, error) {
	ms := make([]CfpModel, len(es))
	for i, e := range es {
		m, err := e.ToModel()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
		ms[i] = m
	}
	return ms, nil
}

// GetPreProductionCfp
// Summary: This is the function to get pre-production cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetPreProductionCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "preProduction" {
			return e
		}
	}
	return nil
}

// GetPreComponentCfp
// Summary: This is the function to get pre-component cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetPreComponentCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "preComponent" {
			return e
		}
	}
	return nil
}

// GetMainProductionCfp
// Summary: This is the function to get main-production cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetMainProductionCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "mainProduction" {
			return e
		}
	}
	return nil
}

// GetMainComponentCfp
// Summary: This is the function to get main-component cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetMainComponentCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "mainComponent" {
			return e
		}
	}
	return nil
}

// GetPreProductionTotalCfp
// Summary: This is the function to get pre-production total cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetPreProductionTotalCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "preProductionTotal" {
			return e
		}
	}
	return nil
}

// GetPreComponentTotalCfp
// Summary: This is the function to get pre-component total cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetPreComponentTotalCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "preComponentTotal" {
			return e
		}
	}
	return nil
}

// GetMainProductionTotalCfp
// Summary: This is the function to get main-production total cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetMainProductionTotalCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "mainProductionTotal" {
			return e
		}
	}
	return nil
}

// GetMainComponentTotalCfp
// Summary: This is the function to get main-component total cfp.
// output: (*CfpEntityModel) CfpEntityModel object
func (es CfpEntityModelSet) GetMainComponentTotalCfp() *CfpEntityModel {
	for _, e := range es {
		if e.CfpType == "mainComponentTotal" {
			return e
		}
	}
	return nil
}
