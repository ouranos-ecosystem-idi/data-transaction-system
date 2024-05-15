package fixtures

import (
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	// 変数名の昇順
	AmountRequired                = 1.0
	AmountRequiredUnit            = "liter"
	CfpCertificationId            = "d9a38406-cae2-4679-b052-15a75f5531c5"
	CfpCertificationDescription   = "サンプル証明書"
	CfpCertificateList            = []string{"https://www.example1.com/1"}
	CfpCertificateList2           = []string{"https://www.example1.com/2"}
	CfpId                         = "d9a38406-cae2-4679-b052-15a75f5531f5"
	CfpId2                        = "8cb39547-3d77-cbc5-2fde-507acb439f29"
	CfpType                       = traceability.CfpType("preProduction")
	Email                         = "testaccount_user122@example.com"
	GeR                           = 0.1
	GhgDeclaredUnit               = "kgCO2e/liter"
	GhgDeclaredUnit2              = "kgCO2e/kilogram"
	GhgEmission                   = 1.12345
	GlobalOperatorId              = "GlobalOperatorId"
	GlobalPlantId                 = "GlobalPlantId"
	InvalidUUID                   = "invalid_uuid"
	InvalidEnum                   = "invalid_enum"
	OpenOperatorID                = "AAAA-BBBB"
	OpenPlantID                   = "AAAA-BBBB"
	OperatorAddress               = "東京都"
	OperatorId                    = "e03cc699-7234-31ed-86be-cc18c92208e5"
	OperatorID                    = "e03cc699-7234-31ed-86be-cc18c92208e5"
	OperatorName                  = "A株式会社"
	PartsName                     = "PartsA-002123"
	PlantAddress                  = "東京都"
	PlantId                       = "eedf264e-cace-4414-8bd3-e10ce1c090e0"
	PlantName                     = "A工場"
	RequestType                   = traceability.RequestType("CFP")
	StatusId                      = "d9a38406-cae2-4679-b052-15a75f5531f6"
	StatusID                      = "d9a38406-cae2-4679-b052-15a75f5531f6"
	SupportPartsName              = "modelA"
	TeR                           = 0.1
	TerminatedFlag                = false
	TiR                           = 0.1
	TraceId                       = "d17833fe-22b7-4a4a-b097-bc3f2150c9a6"
	TraceID                       = "d17833fe-22b7-4a4a-b097-bc3f2150c9a6"
	TraceId2                      = "06c9b015-4225-ba30-1ed3-6faf02cb3fe6"
	TraceID2                      = "06c9b015-4225-ba30-1ed3-6faf02cb3fe6"
	TradeId                       = "97a72868-63e3-43fb-9997-488af61d3be7"
	TradeID                       = "97a72868-63e3-43fb-9997-488af61d3be7"
	TradeIdUUID1                  = "00000000-0000-0000-0000-000000000001"
	TradeIDUUID1                  = "00000000-0000-0000-0000-000000000001"
	TradeIdUUID2                  = "00000000-0000-0000-0000-000000000002"
	TradeIDUUID2                  = "00000000-0000-0000-0000-000000000002"
	TradeRequestMessage           = "回答依頼時のメッセージが入ります"
	UID                           = "uid"
	AssertMessage                 = "比較対象の２つの値は定義順に関係なく、一致する必要があります。"
	UnmarshalMockFailureMessage   = "モック定義値のUnmarshalに失敗しました: %v"
	UnmarshalExpectFailureMessage = "正解値のUnmarshalに失敗しました: %v"

	// Input
	PutCfpInput = traceability.PutCfpInput{
		CfpID:           &CfpId,
		TraceID:         TraceId,
		GhgEmission:     &GhgEmission,
		GhgDeclaredUnit: GhgDeclaredUnit,
		CfpType:         CfpType,
	}
	PutPartsInput = traceability.PutPartsInput{
		OperatorID:         OperatorId,
		TraceID:            &TraceId,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   &SupportPartsName,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &AmountRequiredUnit,
	}
	PutTradeRequestInput = traceability.PutTradeRequestInput{
		Trade: traceability.PutTradeInput{
			TradeID:              &TradeId,
			DownstreamOperatorID: OperatorId,
			UpstreamOperatorID:   OperatorId,
			DownstreamTraceID:    TraceId,
		},
		Status: traceability.PutStatusInput{
			StatusID:    &StatusId,
			TradeID:     &TradeId,
			Message:     TradeRequestMessage,
			RequestType: RequestType,
		},
	}
	PutTradeResponseInput = traceability.PutTradeResponseInput{
		OperatorID: uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		TradeID:    uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
		TraceID:    uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
	}
)

func NewPutPartsStructureInput() traceability.PutPartsStructureInput {
	parentPartsModel := traceability.PutPartsInput{
		OperatorID:         OperatorId,
		TraceID:            &TraceId,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   &SupportPartsName,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &AmountRequiredUnit,
	}
	childrenPartsModel := traceability.PutPartsInputs{
		traceability.PutPartsInput{
			OperatorID:         OperatorId,
			TraceID:            &TraceId,
			PlantID:            PlantId,
			PartsName:          PartsName,
			SupportPartsName:   &SupportPartsName,
			TerminatedFlag:     &TerminatedFlag,
			AmountRequired:     &AmountRequired,
			AmountRequiredUnit: &AmountRequiredUnit,
		},
	}
	return traceability.PutPartsStructureInput{
		ParentPartsInput:   &parentPartsModel,
		ChildrenPartsInput: &childrenPartsModel,
	}
}

func NewPutCfpInputs() traceability.PutCfpInputs {
	mainProd := traceability.PutCfpInput{
		CfpID:           &CfpId,
		TraceID:         TraceId,
		GhgEmission:     &GhgEmission,
		GhgDeclaredUnit: GhgDeclaredUnit,
		CfpType:         traceability.CfpTypeMainProduction,
		DqrType:         traceability.DqrTypeMainProcessing,
		DqrValue: traceability.PutDqrValueInput{
			TeR: &TeR,
			GeR: &GeR,
			TiR: &TiR,
		},
	}
	preProd := traceability.PutCfpInput{
		CfpID:           &CfpId,
		TraceID:         TraceId,
		GhgEmission:     &GhgEmission,
		GhgDeclaredUnit: GhgDeclaredUnit,
		CfpType:         traceability.CfpTypePreProduction,
		DqrType:         traceability.DqrTypePreProcessing,
		DqrValue: traceability.PutDqrValueInput{
			TeR: &TeR,
			GeR: &GeR,
			TiR: &TiR,
		},
	}
	mainComp := traceability.PutCfpInput{
		CfpID:           &CfpId,
		TraceID:         TraceId,
		GhgEmission:     &GhgEmission,
		GhgDeclaredUnit: GhgDeclaredUnit,
		CfpType:         traceability.CfpTypeMainComponent,
		DqrType:         traceability.DqrTypeMainProcessing,
		DqrValue: traceability.PutDqrValueInput{
			TeR: &TeR,
			GeR: &GeR,
			TiR: &TiR,
		},
	}
	preComp := traceability.PutCfpInput{
		CfpID:           &CfpId,
		TraceID:         TraceId,
		GhgEmission:     &GhgEmission,
		GhgDeclaredUnit: GhgDeclaredUnit,
		CfpType:         traceability.CfpTypePreComponent,
		DqrType:         traceability.DqrTypePreProcessing,
		DqrValue: traceability.PutDqrValueInput{
			TeR: &TeR,
			GeR: &GeR,
			TiR: &TiR,
		},
	}

	return traceability.PutCfpInputs{
		mainProd,
		preProd,
		mainComp,
		preComp,
	}
}

func NewGetPartsStructureModel() traceability.GetPartsStructureModel {
	return traceability.GetPartsStructureModel{
		TraceID:    uuid.MustParse(TraceId),
		OperatorID: OperatorId,
	}
}

func NewGetCfpCertificationModel() traceability.GetCfpCertificationModel {
	return traceability.GetCfpCertificationModel{
		OperatorID: uuid.MustParse(OperatorId),
		TraceID:    uuid.MustParse(TraceId),
	}
}

func NewGetStatusModel(statusTarget int) traceability.GetStatusModel {
	traceId := uuid.MustParse(TraceId)
	statusId := uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478")
	after := uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478")
	status := traceability.GetStatusModel{
		OperatorID: uuid.MustParse(OperatorId),
		Limit:      1,
		After:      &after,
		StatusID:   &statusId,
		TraceID:    &traceId,
	}
	if statusTarget == 1 {
		status.StatusTarget = traceability.Request
	} else if statusTarget == 2 {
		status.StatusTarget = traceability.Response
	}
	return status
}

func NewStatusModel() traceability.StatusModel {
	return traceability.StatusModel{
		StatusID: uuid.MustParse("5185a435-c039-4196-bb34-0ee0c2395478"),
		TradeID:  uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
		RequestStatus: traceability.RequestStatus{
			CfpResponseStatus: traceability.CfpResponseStatusCancel,
			TradeTreeStatus:   traceability.TradeTreeStatusUnterminated,
		},
		Message:      common.StringPtr("A01のCFP値を回答ください"),
		ReplyMessage: common.StringPtr("A01のCFP値を回答しました"),
		RequestType:  RequestType.ToString(),
	}
}

func NewGetCfpModel(traceId string) traceability.GetCfpModel {
	return traceability.GetCfpModel{
		OperatorID: uuid.MustParse(OperatorId),
		TraceIDs:   []uuid.UUID{uuid.MustParse(traceId)},
	}
}

func NewCfpModels(traceId string) traceability.CfpModels {
	return traceability.CfpModels{
		{
			CfpID:           nil,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProduction.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
		{
			CfpID:           nil,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponent.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
		{
			CfpID:           nil,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProduction.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
		{
			CfpID:           nil,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponent.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
	}
}

func NewCfpModelsForUpdate(cfpId string, traceId string) traceability.CfpModels {
	cfpIdUid := uuid.MustParse(cfpId)
	return traceability.CfpModels{
		{
			CfpID:           &cfpIdUid,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreProduction.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
		{
			CfpID:           &cfpIdUid,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypePreComponent.ToString(),
			DqrType:         traceability.DqrTypePreProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
		{
			CfpID:           &cfpIdUid,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainProduction.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
		{
			CfpID:           &cfpIdUid,
			TraceID:         uuid.MustParse(traceId),
			GhgEmission:     common.Float64Ptr(GhgEmission),
			GhgDeclaredUnit: traceability.GhgDeclaredUnitKgCO2ePerKilogram,
			CfpType:         traceability.CfpTypeMainComponent.ToString(),
			DqrType:         traceability.DqrTypeMainProcessing.ToString(),
			DqrValue: traceability.DqrValue{
				TeR: common.Float64Ptr(TeR),
				GeR: common.Float64Ptr(GeR),
				TiR: common.Float64Ptr(TiR),
			},
		},
	}
}

func NewGetTradeRequestInput() traceability.GetTradeRequestInput {
	operatorId := uuid.MustParse(OperatorId)
	traceId := uuid.MustParse(TraceId)
	after := uuid.MustParse(TradeId)
	return traceability.GetTradeRequestInput{
		OperatorID: operatorId,
		TraceIDs:   []uuid.UUID{traceId},
		Limit:      1,
		After:      &after,
	}
}

func NewGetTradeResponseInput() traceability.GetTradeResponseInput {
	operatorId := uuid.MustParse(OperatorId)
	after := uuid.MustParse(TradeId)
	return traceability.GetTradeResponseInput{
		OperatorID: operatorId,
		Limit:      1,
		After:      &after,
	}
}

func NewPutTradeResponseInput() traceability.PutTradeResponseInput {
	operatorId := uuid.MustParse(OperatorId)
	tradeId := uuid.MustParse(TradeId)
	traceId := uuid.MustParse(TraceId)
	return traceability.PutTradeResponseInput{
		OperatorID: operatorId,
		TradeID:    tradeId,
		TraceID:    traceId,
	}
}

func NewGetPartsModel() traceability.GetPartsModel {
	parentFlag := true
	after := uuid.MustParse(TraceId)
	return traceability.GetPartsModel{
		TraceID:    &TraceId,
		OperatorID: OperatorId,
		PartsName:  &PartsName,
		PlantID:    &PlantId,
		ParentFlag: &parentFlag,
		Limit:      100,
		After:      &after,
	}
}

func NewPutPartsStructureModel() traceability.PartsStructureModel {
	plantId := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	supportPartsNameP := "A000001"
	supportPartsNameC := "B001"
	AmountRequiredUnit := traceability.AmountRequiredUnitKilogram
	partsModel := traceability.PartsStructureModel{
		ParentPartsModel: &traceability.PartsModel{
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			PlantID:            &plantId,
			PartsName:          "B01",
			SupportPartsName:   &supportPartsNameP,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: &AmountRequiredUnit,
		},
		ChildrenPartsModel: []traceability.PartsModel{
			{
				OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
				TraceID:            uuid.MustParse("4d987ed4-f1b0-4bf1-8795-1fdb25300e34"),
				PlantID:            &plantId,
				PartsName:          "B01001",
				SupportPartsName:   &supportPartsNameC,
				TerminatedFlag:     false,
				AmountRequired:     common.Float64Ptr(2.1),
				AmountRequiredUnit: &AmountRequiredUnit,
			},
		},
	}
	return partsModel
}

func NewPutPartsStructureEntityModel() traceability.PartsStructureEntity {
	plantId := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	supportPartsNameP := "A000001"
	supportPartsNameC := "B001"
	AmountRequiredUnit := traceability.AmountRequiredUnitKilogram
	partsModel := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
			TraceID:            uuid.MustParse("2680ed32-19a3-435b-a094-23ff43aaa611"),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   &supportPartsNameP,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr(AmountRequiredUnit.ToString()),
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
				TraceID:            uuid.MustParse("4d987ed4-f1b0-4bf1-8795-1fdb25300e34"),
				PlantID:            plantId,
				PartsName:          "B01001",
				SupportPartsName:   &supportPartsNameC,
				TerminatedFlag:     false,
				AmountRequired:     common.Float64Ptr(2.1),
				AmountRequiredUnit: common.StringPtr(AmountRequiredUnit.ToString()),
			},
		},
	}
	return partsModel
}

func NewPutTradeRequestModel(create bool) traceability.TradeRequestModel {
	tradeId := uuid.Nil
	if !create {
		tradeId = uuid.MustParse(TradeId)
	}
	downstreamOperatorId := uuid.MustParse(OperatorId)
	upstreamOperatorId := uuid.MustParse(OperatorId)
	downstreamTraceId := uuid.MustParse(TraceId)
	upstreamTraceId := uuid.MustParse(TraceId2)
	statusId := uuid.MustParse(StatusId)
	partsModel := traceability.TradeRequestModel{
		TradeModel: traceability.TradeModel{
			TradeID:              &tradeId,
			DownstreamOperatorID: downstreamOperatorId,
			UpstreamOperatorID:   upstreamOperatorId,
			DownstreamTraceID:    downstreamTraceId,
			UpstreamTraceID:      &upstreamTraceId,
		},
		StatusModel: traceability.StatusModel{
			StatusID: statusId,
			TradeID:  tradeId,
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus: traceability.CfpResponseStatusPending,
				TradeTreeStatus:   traceability.TradeTreeStatusUnterminated,
			},
			Message:      &TradeRequestMessage,
			ReplyMessage: &TradeRequestMessage,
			RequestType:  traceability.RequestTypeCFP.ToString(),
		},
	}
	return partsModel
}

func NewPutTradeResponseModel() traceability.TradeModel {
	tradeId := uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092")
	traceId2 := uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b")
	return traceability.TradeModel{
		TradeID:              &tradeId,
		DownstreamOperatorID: uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		UpstreamOperatorID:   uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		DownstreamTraceID:    uuid.MustParse("087aaa4b-8974-4a0a-9c11-b2e66ed468c5"),
		UpstreamTraceID:      &traceId2,
	}
}

func GetPartsModel(traceId string, terminatedFlag bool) traceability.PartsModel {
	plantId := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	amountRequiredUnit := traceability.AmountRequiredUnitKilogram
	return traceability.PartsModel{
		TraceID:            uuid.MustParse(traceId),
		OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PlantID:            &plantId,
		PartsName:          "B01",
		SupportPartsName:   common.StringPtr("A000001"),
		TerminatedFlag:     terminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &amountRequiredUnit,
	}
}

func GetPartsModelEntity(traceId string, terminatedFlag bool) traceability.PartsModelEntity {
	return traceability.PartsModelEntity{
		TraceID:            uuid.MustParse(traceId),
		OperatorID:         uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PlantID:            uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		PartsName:          "B01",
		SupportPartsName:   common.StringPtr("A000001"),
		TerminatedFlag:     terminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: common.StringPtr(traceability.AmountRequiredUnitKilogram.ToString()),
	}
}

func GetPartsStructureEntityModel(parentTraceId string, childTraceId string) traceability.PartsStructureEntityModel {
	return traceability.PartsStructureEntityModel{
		TraceID:       uuid.MustParse(childTraceId),
		ParentTraceID: uuid.MustParse(parentTraceId),
		DeletedAt:     gorm.DeletedAt{Time: time.Now()},
		CreatedAt:     time.Now(),
		CreatedUserID: "seed",
		UpdatedAt:     time.Now(),
		UpdatedUserID: "seed",
	}
}

func GetPartsStructureEntity(parentTraceID string, childTraceIDs []string, terminatedFlag bool) traceability.PartsStructureEntity {
	parentPartsEntity := GetPartsModelEntity(parentTraceID, false)
	res := traceability.PartsStructureEntity{
		ParentPartsEntity: &parentPartsEntity,
	}
	if len(childTraceIDs) == 0 {
		return res
	}
	childrenPartsEntity := make(traceability.PartsModelEntities, len(childTraceIDs))
	for idx, childTraceID := range childTraceIDs {
		childrenPartsEntity[idx] = GetPartsModelEntity(childTraceID, terminatedFlag)
	}
	res.ChildrenPartsEntity = childrenPartsEntity
	return res
}

func GetTradeEntityModel(tradeIdstr string, traceIdstr string, answered bool) traceability.TradeEntityModel {
	tradeId := uuid.MustParse(tradeIdstr)
	operatorId := uuid.MustParse("b1234567-1234-1234-1234-123456789012")
	traceId := uuid.MustParse(traceIdstr)
	res := traceability.TradeEntityModel{
		TradeID:              &tradeId,
		DownstreamOperatorID: operatorId,
		UpstreamOperatorID:   &operatorId,
		DownstreamTraceID:    traceId,
		TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
		DeletedAt:            gorm.DeletedAt{Time: time.Now()},
		CreatedAt:            time.Now(),
		CreatedUserID:        "seed",
		UpdatedAt:            time.Now(),
		UpdatedUserID:        "seed",
	}
	if answered {
		res.UpstreamTraceID = &traceId
	}
	return res
}

func GetCfpEntityModels(cfpIdstr string, traceId string) traceability.CfpEntityModels {
	cfpId := uuid.MustParse(cfpIdstr)
	return traceability.CfpEntityModels{
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerKilogram.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypePreComponent.ToString(),
			DqrType:            traceability.DqrTypePreProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerKilogram.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypePreProduction.ToString(),
			DqrType:            traceability.DqrTypePreProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerKilogram.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypeMainComponent.ToString(),
			DqrType:            traceability.DqrTypeMainProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerKilogram.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypeMainProduction.ToString(),
			DqrType:            traceability.DqrTypeMainProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
	}
}

func GetCfpEntityModelsNotSame(cfpIdstr string, traceId string) traceability.CfpEntityModels {
	cfpId := uuid.MustParse(cfpIdstr)
	return traceability.CfpEntityModels{
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerKilogram.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypePreComponent.ToString(),
			DqrType:            traceability.DqrTypePreProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerLiter.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypePreProduction.ToString(),
			DqrType:            traceability.DqrTypePreProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerKilogram.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypeMainComponent.ToString(),
			DqrType:            traceability.DqrTypeMainProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(traceId),
			GhgEmission:        &GhgEmission,
			GhgDeclaredUnit:    traceability.GhgDeclaredUnitKgCO2ePerLiter.ToString(),
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypeMainProduction.ToString(),
			DqrType:            traceability.DqrTypeMainProcessing.ToString(),
			TeR:                &TeR,
			GeR:                &GeR,
			TiR:                &TiR,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
	}
}

func NewPutStatusInput() traceability.PutStatusInput {
	return traceability.PutStatusInput{
		StatusID:    &StatusId,
		TradeID:     &TradeId,
		Message:     TradeRequestMessage,
		RequestType: RequestType,
		PutRequestStatusInput: traceability.PutRequestStatusInput{
			CfpResponseStatus: traceability.CfpResponseStatusCancel,
			TradeTreeStatus:   traceability.TradeTreeStatusTerminated,
		},
	}
}
