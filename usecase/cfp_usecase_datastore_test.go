package usecase_test

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 終端部品
// [x] 1-2. 200: 親部品
// [x] 1-3. 200: 仕入部品
// [x] 1-4. 200: 子部品あり(子が終端)
// [x] 1-5. 200: 子部品あり(子が非終端)
// [x] 1-6. 200: 終端部品(CFPなし)
// [x] 1-7. 200: 終端部品(CFPなし)
// [x] 1-8. 200: 仕入部品(依頼回答なし)
// [x] 1-9. 200: 仕入部品(CFP回答なし)
// [x] 1-10. 200: 子部品あり(CFP回答なし)
// [x] 1-11. 200: 子部品あり(子が終端)(CFPなし)
// [x] 1-12. 200: 子部品あり(子が非終端)(依頼情報なし)
// [x] 1-13. 200: 子部品あり(子が非終端)(依頼回答なし)
// [x] 1-14. 200: 子部品あり(子が非終端)(CFP回答なし)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetCfp(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"
	cfpId := uuid.MustParse("892262ab-6795-4a97-bf25-d92c512ebb31")
	partsTerminated := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", true)
	cfpTerminated := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")

	partsParentOnly := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", false)
	partsStructureParentOnly := f.GetPartsStructureEntityModel("00000000-0000-0000-0000-000000000000", "2680ed32-19a3-435b-a094-23ff43aaa611")
	partsStructureEntityParentOnly := f.GetPartsStructureEntity("2680ed32-19a3-435b-a094-23ff43aaa611", []string{}, false)
	cfpParentOnly := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")

	partsImport := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", false)
	partsStructureImport := f.GetPartsStructureEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa612")
	partsStructureEntityImport := f.GetPartsStructureEntity("2680ed32-19a3-435b-a094-23ff43aaa611", []string{"2680ed32-19a3-435b-a094-23ff43aaa612"}, false)
	cfpImport := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa612")
	trade := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa611", true)
	tradeNoAnswer := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa611", false)

	partsWithChildParent := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", false)
	partsStructureWithChildParent := f.GetPartsStructureEntityModel("00000000-0000-0000-0000-000000000000", "2680ed32-19a3-435b-a094-23ff43aaa611")
	partsStructureEntityWithChild1 := f.GetPartsStructureEntity("2680ed32-19a3-435b-a094-23ff43aaa611", []string{"2680ed32-19a3-435b-a094-23ff43aaa612"}, true)
	partsStructureEntityWithChild2 := f.GetPartsStructureEntity("2680ed32-19a3-435b-a094-23ff43aaa611", []string{"2680ed32-19a3-435b-a094-23ff43aaa613"}, false)
	cfpWithChildParent := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")
	cfpWithChildChild1 := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa612")
	cfpWithChildChild2 := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa613")
	tradeChild1 := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa612", "2680ed32-19a3-435b-a094-23ff43aaa612", true)
	tradeChild2 := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa613", "2680ed32-19a3-435b-a094-23ff43aaa613", true)

	expectTerminated := []traceability.CfpModel{
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponent.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProduction.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponent.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProduction.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     common.Float64Ptr(0),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponentTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProductionTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     common.Float64Ptr(0),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponentTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProductionTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
	}

	expectUnTerminated := []traceability.CfpModel{
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponent.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProduction.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponent.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProduction.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     common.Float64Ptr(0),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponentTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProductionTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     common.Float64Ptr(0),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponentTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProductionTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
	}

	expectImport := []traceability.CfpModel{
		{
			CfpID:           nil,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProductionResponse.ToString(),
			DqrType:         traceability.DqrPreProcessingResponse.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           nil,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProductionResponse.ToString(),
			DqrType:         traceability.DqrMainProcessingResponse.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
	}

	expectWithChild := []traceability.CfpModel{
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponent.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProduction.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponent.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProduction.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponentTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProductionTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponentTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProductionTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
	}

	expectWithChildNoCfp := []traceability.CfpModel{
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponent.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProduction.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponent.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProduction.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: &f.TeR,
				GeR: &f.GeR,
				TiR: &f.TiR,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     common.Float64Ptr(0),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponentTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProductionTotal.ToString(),
			DqrType:         traceability.DqrTypePreProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     common.Float64Ptr(0),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponentTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
		{
			CfpID:           &cfpId,
			TraceID:         uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			GhgEmission:     &f.GhgEmission,
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProductionTotal.ToString(),
			DqrType:         traceability.DqrTypeMainProcessingTotal.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(2.1),
				GeR: common.Float64Ptr(0),
				TiR: nil,
			},
		},
	}

	tests := []struct {
		name                        string
		input                       traceability.GetCfpModel
		searchType                  int
		receiveParts                *traceability.PartsModelEntity
		receivePartsStructure       *traceability.PartsStructureEntityModel
		receivePartsStructureEntity *traceability.PartsStructureEntity
		receiveCfpParent            *traceability.CfpEntityModels
		receiveCfpParentError       error
		receiveCfpChild             *traceability.CfpEntityModels
		receiveCfpChildError        error
		receiveTrade                *traceability.TradeEntityModel
		receiveTradeError           error
		expect                      []traceability.CfpModel
	}{
		{
			name:                  "1-1. 200: 終端部品",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsTerminated,
			receivePartsStructure: &traceability.PartsStructureEntityModel{},
			receiveCfpParent:      &cfpTerminated,
			receiveTrade:          nil,
			expect:                expectTerminated,
		},
		{
			name:                        "1-2. 200: 親部品",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsParentOnly,
			receivePartsStructure:       &partsStructureParentOnly,
			receivePartsStructureEntity: &partsStructureEntityParentOnly,
			receiveCfpParent:            &cfpParentOnly,
			receiveTrade:                &trade,
			expect:                      expectUnTerminated,
		},
		{
			name:                        "1-3. 200: 仕入部品",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsImport,
			receivePartsStructure:       &partsStructureImport,
			receivePartsStructureEntity: &partsStructureEntityImport,
			receiveCfpParent:            &cfpImport,
			receiveTrade:                &trade,
			expect:                      expectImport,
		},
		{
			name:                        "1-4. 200: 子部品あり(子が終端)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild1,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild1,
			receiveTrade:                &trade,
			expect:                      expectWithChild,
		},
		{
			name:                        "1-5. 200: 子部品あり(子が非終端)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild2,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild2,
			receiveTrade:                &trade,
			expect:                      expectWithChild,
		},
		{
			name:                  "1-6. 200: 終端部品(CFPなし)",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsTerminated,
			receivePartsStructure: &traceability.PartsStructureEntityModel{},
			receiveCfpParent:      &traceability.CfpEntityModels{},
			receiveCfpParentError: gorm.ErrRecordNotFound,
			receiveTrade:          nil,
			expect:                []traceability.CfpModel{},
		},
		{
			name:                  "1-7. 200: 終端部品(CFPなし)",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsTerminated,
			receivePartsStructure: &traceability.PartsStructureEntityModel{},
			receiveCfpParent:      &traceability.CfpEntityModels{},
			receiveTrade:          nil,
			expect:                []traceability.CfpModel{},
		},
		{
			name:                        "1-8. 200: 仕入部品(依頼回答なし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsImport,
			receivePartsStructure:       &partsStructureImport,
			receivePartsStructureEntity: &partsStructureEntityImport,
			receiveCfpParent:            &cfpImport,
			receiveTrade:                &tradeNoAnswer,
			expect:                      []traceability.CfpModel{},
		},
		{
			name:                        "1-9. 200: 仕入部品(CFP回答なし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsImport,
			receivePartsStructure:       &partsStructureImport,
			receivePartsStructureEntity: &partsStructureEntityImport,
			receiveCfpParent:            &traceability.CfpEntityModels{},
			receiveTrade:                &trade,
			expect:                      []traceability.CfpModel{},
		},
		{
			name:                        "1-10. 200: 子部品あり(CFP回答なし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild1,
			receiveCfpParent:            &traceability.CfpEntityModels{},
			receiveCfpChild:             &cfpWithChildChild1,
			receiveTrade:                &trade,
			expect:                      []traceability.CfpModel{},
		},
		{
			name:                        "1-11. 200: 子部品あり(子が終端)(CFPなし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild1,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild1,
			receiveCfpChildError:        gorm.ErrRecordNotFound,
			receiveTrade:                &tradeChild1,
			expect:                      expectWithChildNoCfp,
		},
		{
			name:                        "1-12. 200: 子部品あり(子が非終端)(依頼情報なし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild2,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild2,
			receiveTrade:                &tradeChild2,
			receiveTradeError:           gorm.ErrRecordNotFound,
			expect:                      expectWithChildNoCfp,
		},
		{
			name:                        "1-13. 200: 子部品あり(子が非終端)(依頼回答なし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild2,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild2,
			receiveTrade:                &tradeNoAnswer,
			expect:                      expectWithChildNoCfp,
		},
		{
			name:                        "1-14. 200: 子部品あり(子が非終端)(CFP回答なし)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild2,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &traceability.CfpEntityModels{},
			receiveCfpChildError:        gorm.ErrRecordNotFound,
			receiveTrade:                &tradeChild2,
			expect:                      expectWithChildNoCfp,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				if test.receiveParts != nil {
					ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(*test.receiveParts, nil)
				}
				if test.receivePartsStructure != nil {
					ouranosRepositoryMock.On("GetPartsStructureByTraceId", mock.Anything).Return(*test.receivePartsStructure, nil)
				}
				if test.receivePartsStructureEntity != nil {
					ouranosRepositoryMock.On("GetPartsStructure", mock.Anything).Return(*test.receivePartsStructureEntity, nil)
				}
				if test.receiveCfpParent != nil {
					ouranosRepositoryMock.On("ListCFPsByTraceID", "2680ed32-19a3-435b-a094-23ff43aaa611").Return(*test.receiveCfpParent, test.receiveCfpParentError)
					if test.receiveCfpChild != nil {
						ouranosRepositoryMock.On("ListCFPsByTraceID", "2680ed32-19a3-435b-a094-23ff43aaa612").Return(*test.receiveCfpChild, test.receiveCfpChildError)
						ouranosRepositoryMock.On("ListCFPsByTraceID", "2680ed32-19a3-435b-a094-23ff43aaa613").Return(*test.receiveCfpChild, test.receiveCfpChildError)
					}
				}
				if test.receiveTrade != nil {
					ouranosRepositoryMock.On("GetTradeByDownstreamTraceID", mock.Anything).Return(*test.receiveTrade, test.receiveTradeError)
				}
				usecase := usecase.NewCfpUsecase(ouranosRepositoryMock)
				actualRes, err := usecase.GetCfp(c, test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, test.expect, actualRes, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 200: データ取得エラー(部品)
// [x] 2-2. 200: データ取得エラー(部品構成)
// [x] 2-3. 200: データ取得エラー(CFP)(終端)
// [x] 2-4. 200: データ取得エラー(単位不一致)(終端)
// [x] 2-5. 200: データ取得エラー(依頼)(仕入部品)
// [x] 2-6. 200: データ取得エラー(CFP)(仕入部品)
// [x] 2-7. 200: データ取得エラー(子部品)(子部品あり)
// [x] 2-8. 200: データ取得エラー(CFP親)(子部品あり)
// [x] 2-9. 200: データ取得エラー(CFP子)(子部品あり)
// [x] 2-10. 200: データ取得エラー(依頼子非終端)(子部品あり)
// [x] 2-11. 200: データ取得エラー(CFP子非終端)(子部品あり)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetCfp_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"
	accessError := fmt.Errorf("DB AccessError")
	formatError := fmt.Errorf("ghgDeclaredUnits must be same")
	partsTerminate := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", true)
	partsStructureTerminate := f.GetPartsStructureEntityModel("00000000-0000-0000-0000-000000000000", "2680ed32-19a3-435b-a094-23ff43aaa611")
	cfpTerminateNotSameUnit := f.GetCfpEntityModelsNotSame("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")

	partsImport := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", false)
	partsStructureImport := f.GetPartsStructureEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa612")
	cfpImport := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa612")
	tradeImport := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa611", true)

	partsWithChildParent := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", false)
	partsStructureWithChildParent := f.GetPartsStructureEntityModel("00000000-0000-0000-0000-000000000000", "2680ed32-19a3-435b-a094-23ff43aaa611")
	partsStructureEntityWithChild1 := f.GetPartsStructureEntity("2680ed32-19a3-435b-a094-23ff43aaa611", []string{"2680ed32-19a3-435b-a094-23ff43aaa612"}, true)
	partsStructureEntityWithChild2 := f.GetPartsStructureEntity("2680ed32-19a3-435b-a094-23ff43aaa611", []string{"2680ed32-19a3-435b-a094-23ff43aaa613"}, false)
	cfpWithChildParent := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")
	cfpWithChildChild1 := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa612")
	cfpWithChildChild2 := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa613")
	tradeChild2 := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa613", "2680ed32-19a3-435b-a094-23ff43aaa613", true)

	tests := []struct {
		name                             string
		input                            traceability.GetCfpModel
		searchType                       int
		receiveParts                     *traceability.PartsModelEntity
		receivePartsError                error
		receivePartsStructure            *traceability.PartsStructureEntityModel
		receivePartsStructureError       error
		receivePartsStructureEntity      *traceability.PartsStructureEntity
		receivePartsStructureEntityError error
		receiveCfpParent                 *traceability.CfpEntityModels
		receiveCfpParentError            error
		receiveCfpChild                  *traceability.CfpEntityModels
		receiveCfpChildError             error
		receiveTrade                     *traceability.TradeEntityModel
		receiveTradeError                error

		expect error
	}{
		{
			name:              "2-1. 400: データ取得エラー(部品)",
			input:             f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:      &traceability.PartsModelEntity{},
			receivePartsError: accessError,
			expect:            accessError,
		},
		{
			name:                       "2-2. 400: データ取得エラー(部品構成)",
			input:                      f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:               &partsTerminate,
			receivePartsStructure:      &traceability.PartsStructureEntityModel{},
			receivePartsStructureError: accessError,
			expect:                     accessError,
		},
		{
			name:                  "2-3. 400: データ取得エラー(CFP)(終端)",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsTerminate,
			receivePartsStructure: &traceability.PartsStructureEntityModel{},
			receiveCfpParent:      &traceability.CfpEntityModels{},
			receiveCfpParentError: accessError,
			expect:                accessError,
		},
		{
			name:                  "2-4. 400: データ取得エラー(単位不一致)(終端)",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsTerminate,
			receivePartsStructure: &partsStructureTerminate,
			receiveCfpParent:      &cfpTerminateNotSameUnit,
			expect:                formatError,
		},
		{
			name:                  "2-5. 400: データ取得エラー(依頼)(仕入部品)",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsImport,
			receivePartsStructure: &partsStructureImport,
			receiveCfpParent:      &cfpImport,
			receiveTrade:          &tradeImport,
			receiveTradeError:     accessError,
			expect:                accessError,
		},
		{
			name:                  "2-6. 400: データ取得エラー(CFP)(仕入部品)",
			input:                 f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:          &partsImport,
			receivePartsStructure: &partsStructureImport,
			receiveCfpParent:      &cfpImport,
			receiveCfpParentError: accessError,
			receiveTrade:          &tradeImport,
			expect:                accessError,
		},
		{
			name:                             "2-7. 400: データ取得エラー(子部品)(子部品あり)",
			input:                            f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                     &partsWithChildParent,
			receivePartsStructure:            &partsStructureWithChildParent,
			receivePartsStructureEntity:      &partsStructureEntityWithChild1,
			receivePartsStructureEntityError: accessError,
			receiveCfpParent:                 &cfpWithChildParent,
			expect:                           accessError,
		},
		{
			name:                        "2-8. 400: データ取得エラー(CFP親)(子部品あり)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild1,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpParentError:       accessError,
			receiveCfpChild:             &cfpWithChildChild1,
			expect:                      accessError,
		},
		{
			name:                        "2-9. 400: データ取得エラー(CFP子)(子部品あり)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild1,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild1,
			receiveCfpChildError:        accessError,
			expect:                      accessError,
		},
		{
			name:                        "2-10. 400: データ取得エラー(依頼子非終端)(子部品あり)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild2,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild2,
			receiveTrade:                &tradeChild2,
			receiveTradeError:           accessError,
			expect:                      accessError,
		},
		{
			name:                        "2-11. 400: データ取得エラー(CFP子非終端)(子部品あり)",
			input:                       f.NewGetCfpModel("2680ed32-19a3-435b-a094-23ff43aaa611"),
			receiveParts:                &partsWithChildParent,
			receivePartsStructure:       &partsStructureWithChildParent,
			receivePartsStructureEntity: &partsStructureEntityWithChild2,
			receiveCfpParent:            &cfpWithChildParent,
			receiveCfpChild:             &cfpWithChildChild2,
			receiveCfpChildError:        accessError,
			receiveTrade:                &tradeChild2,
			expect:                      accessError,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				if test.receiveParts != nil {
					ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(*test.receiveParts, test.receivePartsError)
				}
				if test.receivePartsStructure != nil {
					ouranosRepositoryMock.On("GetPartsStructureByTraceId", mock.Anything).Return(*test.receivePartsStructure, test.receivePartsStructureError)
				}
				if test.receivePartsStructureEntity != nil {
					ouranosRepositoryMock.On("GetPartsStructure", mock.Anything).Return(*test.receivePartsStructureEntity, test.receivePartsStructureEntityError)
				}
				if test.receiveCfpParent != nil {
					ouranosRepositoryMock.On("ListCFPsByTraceID", "2680ed32-19a3-435b-a094-23ff43aaa611").Return(*test.receiveCfpParent, test.receiveCfpParentError)
					if test.receiveCfpChild != nil {
						ouranosRepositoryMock.On("ListCFPsByTraceID", "2680ed32-19a3-435b-a094-23ff43aaa612").Return(*test.receiveCfpChild, test.receiveCfpChildError)
						ouranosRepositoryMock.On("ListCFPsByTraceID", "2680ed32-19a3-435b-a094-23ff43aaa613").Return(*test.receiveCfpChild, test.receiveCfpChildError)
					}

				}
				if test.receiveTrade != nil {
					ouranosRepositoryMock.On("GetTradeByDownstreamTraceID", mock.Anything).Return(*test.receiveTrade, test.receiveTradeError)
				}
				usecase := usecase.NewCfpUsecase(ouranosRepositoryMock)
				_, err := usecase.GetCfp(c, test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect, err, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了(新規)
// [x] 1-2. 200: 正常終了(更新)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutCfp(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"
	inputCfpsForCreate := f.NewCfpModels("2680ed32-19a3-435b-a094-23ff43aaa611")
	inputCfpsForUpdate := f.NewCfpModelsForUpdate("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")
	cfp := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")
	trade := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa611", true)
	parts := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", true)
	tests := []struct {
		name                   string
		input                  traceability.CfpModels
		isCreate               bool
		receiveDuplicateCfp    *traceability.CfpEntityModels
		receiveCfp             *traceability.CfpEntityModels
		receiveTrade           *traceability.TradeEntityModels
		receivePart            *traceability.PartsModelEntity
		receivePutTrade        *traceability.TradeEntityModel
		receiveCfpForUpdate    *traceability.CfpEntityModels
		receivePutCfpForUpdate *traceability.CfpEntityModels
		expect                 []traceability.CfpModel
	}{
		{
			name:                "1-1. 200: 正常終了(新規)",
			input:               inputCfpsForCreate,
			isCreate:            true,
			receiveDuplicateCfp: &traceability.CfpEntityModels{},
			receiveCfp:          &cfp,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receivePart:            &parts,
			receivePutTrade:        &trade,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 inputCfpsForUpdate,
		},
		{
			name:                   "1-2. 200: 正常終了(更新)",
			input:                  inputCfpsForUpdate,
			isCreate:               false,
			receiveDuplicateCfp:    nil,
			receiveCfp:             nil,
			receiveTrade:           nil,
			receivePart:            nil,
			receivePutTrade:        nil,
			receiveCfpForUpdate:    &cfp,
			receivePutCfpForUpdate: &cfp,
			expect:                 inputCfpsForUpdate,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				if test.isCreate {
					ouranosRepositoryMock.On("ListCFPsByTraceID", mock.Anything).Return(*test.receiveDuplicateCfp, nil)
					ouranosRepositoryMock.On("BatchCreateCFP", mock.Anything).Return(*test.receiveCfp, nil)
					ouranosRepositoryMock.On("ListTradeByUpstreamTraceID", mock.Anything).Return(*test.receiveTrade, nil)
					ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(*test.receivePart, nil)
					ouranosRepositoryMock.On("PutTradeResponse", mock.Anything, mock.Anything).Return(*test.receivePutTrade, nil)
				} else {
					for _, cfp := range *test.receiveCfpForUpdate {
						ouranosRepositoryMock.On("GetCFP", mock.Anything, cfp.CfpType).Return(*cfp, nil)
						ouranosRepositoryMock.On("PutCFP", *cfp).Return(*cfp, nil)
					}
				}

				usecase := usecase.NewCfpUsecase(ouranosRepositoryMock)
				actualRes, err := usecase.PutCfp(c, test.input, f.OperatorId)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, test.expect, actualRes, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 400: データ取得エラー(新規)
// [x] 2-2. 400: データ取得エラー(更新)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_PutCfp_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"
	inputCfpsForCreate := f.NewCfpModels("2680ed32-19a3-435b-a094-23ff43aaa611")
	inputCfpsForUpdate := f.NewCfpModelsForUpdate("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")
	cfp := f.GetCfpEntityModels("892262ab-6795-4a97-bf25-d92c512ebb31", "2680ed32-19a3-435b-a094-23ff43aaa611")
	trade := f.GetTradeEntityModel("2680ed32-19a3-435b-a094-23ff43aaa611", "2680ed32-19a3-435b-a094-23ff43aaa611", true)
	parts := f.GetPartsModelEntity("2680ed32-19a3-435b-a094-23ff43aaa611", true)
	accessError := fmt.Errorf("DB AccessError")
	duplicateError := fmt.Errorf("traceId %v already has cfps", "2680ed32-19a3-435b-a094-23ff43aaa611")
	tests := []struct {
		name                        string
		input                       traceability.CfpModels
		isCreate                    bool
		receiveDuplicateCfp         *traceability.CfpEntityModels
		receiveDuplicateCfpError    error
		receiveCfp                  *traceability.CfpEntityModels
		receiveCfpError             error
		receiveTrade                *traceability.TradeEntityModels
		receiveTradeError           error
		receivePart                 *traceability.PartsModelEntity
		receivePartError            error
		receivePutTrade             *traceability.TradeEntityModel
		receivePutTradeError        error
		receiveCfpForUpdate         *traceability.CfpEntityModels
		receiveCfpForUpdateError    error
		receivePutCfpForUpdate      *traceability.CfpEntityModels
		receivePutCfpForUpdateError error
		expect                      error
	}{
		{
			name:                "2-1. 400: データ取得エラー(CFP重複)(新規)",
			input:               inputCfpsForCreate,
			isCreate:            true,
			receiveDuplicateCfp: &cfp,
			receiveCfp:          &cfp,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receivePart:            &parts,
			receivePutTrade:        &trade,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 duplicateError,
		},
		{
			name:                     "2-2. 400: データ取得エラー(CFP重複)(新規)",
			input:                    inputCfpsForCreate,
			isCreate:                 true,
			receiveDuplicateCfp:      &traceability.CfpEntityModels{},
			receiveDuplicateCfpError: accessError,
			receiveCfp:               &cfp,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receivePart:            &parts,
			receivePutTrade:        &trade,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 accessError,
		},
		{
			name:                "2-3. 400: データ取得エラー(CFP登録)(新規)",
			input:               inputCfpsForCreate,
			isCreate:            true,
			receiveDuplicateCfp: &traceability.CfpEntityModels{},
			receiveCfp:          &cfp,
			receiveCfpError:     accessError,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receivePart:            &parts,
			receivePutTrade:        &trade,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 accessError,
		},
		{
			name:                "2-4. 400: データ取得エラー(依頼)(新規)",
			input:               inputCfpsForCreate,
			isCreate:            true,
			receiveDuplicateCfp: &traceability.CfpEntityModels{},
			receiveCfp:          &cfp,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receiveTradeError:      accessError,
			receivePart:            &parts,
			receivePutTrade:        &trade,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 accessError,
		},
		{
			name:                "2-5. 400: データ取得エラー(部品)(新規)",
			input:               inputCfpsForCreate,
			isCreate:            true,
			receiveDuplicateCfp: &traceability.CfpEntityModels{},
			receiveCfp:          &cfp,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receivePart:            &parts,
			receivePartError:       accessError,
			receivePutTrade:        &trade,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 accessError,
		},
		{
			name:                "2-6. 400: データ取得エラー(依頼更新)(新規)",
			input:               inputCfpsForCreate,
			isCreate:            true,
			receiveDuplicateCfp: &traceability.CfpEntityModels{},
			receiveCfp:          &cfp,
			receiveTrade: &traceability.TradeEntityModels{
				trade,
			},
			receivePart:            &parts,
			receivePutTrade:        &trade,
			receivePutTradeError:   accessError,
			receiveCfpForUpdate:    nil,
			receivePutCfpForUpdate: nil,
			expect:                 accessError,
		},
		{
			name:                     "2-7. 400: データ取得エラー(CFP取得)(更新)",
			input:                    inputCfpsForUpdate,
			isCreate:                 false,
			receiveDuplicateCfp:      nil,
			receiveCfp:               nil,
			receiveTrade:             nil,
			receivePart:              nil,
			receivePutTrade:          nil,
			receiveCfpForUpdate:      &cfp,
			receiveCfpForUpdateError: accessError,
			receivePutCfpForUpdate:   &cfp,
			expect:                   accessError,
		},
		{
			name:                        "2-8. 400: データ取得エラー(CFP更新)(更新)",
			input:                       inputCfpsForUpdate,
			isCreate:                    false,
			receiveDuplicateCfp:         nil,
			receiveCfp:                  nil,
			receiveTrade:                nil,
			receivePart:                 nil,
			receivePutTrade:             nil,
			receiveCfpForUpdate:         &cfp,
			receivePutCfpForUpdate:      &cfp,
			receivePutCfpForUpdateError: accessError,
			expect:                      accessError,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				if test.isCreate {
					ouranosRepositoryMock.On("ListCFPsByTraceID", mock.Anything).Return(*test.receiveDuplicateCfp, test.receiveDuplicateCfpError)
					ouranosRepositoryMock.On("BatchCreateCFP", mock.Anything).Return(*test.receiveCfp, test.receiveCfpError)
					ouranosRepositoryMock.On("ListTradeByUpstreamTraceID", mock.Anything).Return(*test.receiveTrade, test.receiveTradeError)
					ouranosRepositoryMock.On("GetPartByTraceID", mock.Anything).Return(*test.receivePart, test.receivePartError)
					ouranosRepositoryMock.On("PutTradeResponse", mock.Anything, mock.Anything).Return(*test.receivePutTrade, test.receivePutTradeError)
				} else {
					for _, cfp := range *test.receiveCfpForUpdate {
						ouranosRepositoryMock.On("GetCFP", mock.Anything, cfp.CfpType).Return(*cfp, test.receiveCfpForUpdateError)
						ouranosRepositoryMock.On("PutCFP", *cfp).Return(*cfp, test.receivePutCfpForUpdateError)
					}
				}

				usecase := usecase.NewCfpUsecase(ouranosRepositoryMock)
				_, err := usecase.PutCfp(c, test.input, f.OperatorId)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect, err, f.AssertMessage)
				}
			},
		)
	}
}
