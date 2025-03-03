package fixtures

import (
	"math/rand"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	// 変数名の昇順
	AmountRequired              = 1.0
	AmountRequiredUnit          = "liter"
	AmountRequiredUnit2         = traceability.AmountRequiredUnitKilogram
	CfpCertificationId          = "d9a38406-cae2-4679-b052-15a75f5531c5"
	CfpCertificationDescription = "サンプル証明書"
	CfpCertificateList          = []string{"https://www.example1.com/1"}
	CfpCertificateList2         = []string{"https://www.example1.com/2"}
	CfpId                       = "b8ba9414-0d24-ee00-81d4-f7e13343ebdd"
	CfpId2                      = "8cb39547-3d77-cbc5-2fde-507acb439f29"
	CfpType                     = traceability.CfpType("preProduction")
	Email                       = "testaccount_user122@example.com"
	GeR                         = 0.1
	GhgDeclaredUnit             = "kgCO2e/liter"
	GhgDeclaredUnit2            = "kgCO2e/kilogram"
	GhgEmission                 = 1.12345
	GlobalOperatorId            = "GlobalOperatorId"
	GlobalPlantId               = "GlobalPlantId"
	InvalidUUID                 = "invalid_uuid"
	InvalidEnum                 = "invalid_enum"
	OpenOperatorID              = "AAAA-BBBB"
	OpenPlantID                 = "AAAA-BBBB"
	OperatorAddress             = "東京都"
	OperatorId                  = "f99c9546-e76e-9f15-35b2-abb9c9b21698" // 独自
	OperatorID                  = "f99c9546-e76e-9f15-35b2-abb9c9b21698" // 独自
	OperatorID2                 = "02ad8c1e-3f64-4a92-a9cb-abb3c63f93c2"
	OperatorName                = "A株式会社"
	PartsName                   = "PartsA-002123"
	PlantAddress                = "東京都"
	PlantId                     = "eedf264e-cace-4414-8bd3-e10ce1c090e0"
	PlantID                     = "eedf264e-cace-4414-8bd3-e10ce1c090e0"
	PlantName                   = "A工場"
	RequestType                 = traceability.RequestType("CFP")
	// StatusId                    = "392eb4dc-fffc-23e1-8f63-080f653b0951"
	// StatusID                    = "392eb4dc-fffc-23e1-8f63-080f653b0951"
	StatusId         = "5185a435-c039-4196-bb34-0ee0c2395478"
	StatusID         = "5185a435-c039-4196-bb34-0ee0c2395478"
	SupportPartsName = "modelA"
	TeR              = 0.1
	TerminatedFlag   = false
	TiR              = 0.1
	// TraceId          = "d17833fe-22b7-4a4a-b097-bc3f2150c9a6"
	// TraceID          = "d17833fe-22b7-4a4a-b097-bc3f2150c9a6"
	TraceId      = "38bdd8a5-76a7-a53d-de12-725707b04a1b"
	TraceID      = "38bdd8a5-76a7-a53d-de12-725707b04a1b"
	TraceId2     = "06c9b015-4225-ba30-1ed3-6faf02cb3fe6"
	TraceID2     = "06c9b015-4225-ba30-1ed3-6faf02cb3fe6"
	TraceID3     = "2680ed32-19a3-435b-a094-23ff43aaa611"
	TraceID4     = "81259b24-e47e-449c-b68d-4f575f1fe7e6"
	TraceID5     = "49218d36-8903-4c16-8580-36dbb3c0bca6"
	TraceID6     = "7fa0df76-5efe-4769-96af-8b9aad7d1f66"
	TraceID7     = "b739b34b-c0ea-4fd6-a3f9-3c31ca0fec8f"
	TraceID8     = "c74b0977-5745-4843-84a1-08de6eb2d1e4"
	TraceIDChild = "1c2f37f5-25b9-dea5-346a-7b88035f2553"
	// TraceIDUpstream = "82223e60-0c65-142c-78b7-143c29272ecc" // 独自定義
	TraceIDDownstream = "087aaa4b-8974-4a0a-9c11-b2e66ed468c5"
	// TradeId                       = "97a72868-63e3-43fb-9997-488af61d3be7"
	// TradeID                       = "97a72868-63e3-43fb-9997-488af61d3be7"
	TradeId                       = "a84012cc-73fb-4f9b-9130-59ae546f7092"
	TradeID                       = "a84012cc-73fb-4f9b-9130-59ae546f7092"
	TradeID2                      = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	TradeIdUUID1                  = "00000000-0000-0000-0000-000000000001"
	TradeIDUUID1                  = "00000000-0000-0000-0000-000000000001"
	TradeIdUUID2                  = "00000000-0000-0000-0000-000000000002"
	TradeIDUUID2                  = "00000000-0000-0000-0000-000000000002"
	TradeRequestMessage           = "回答依頼時のメッセージが入ります"
	ResponseDueDate               = "2024-12-31"
	CompletedCount                = 0
	CompletedCountModifiedAt      = "2024-05-23T11:22:33Z"
	TradesCount                   = 0
	TradesCountModifiedAt         = "2024-05-24T22:33:44Z"
	PartsLabelName                = "PartsB"
	PartsAddInfo1                 = "Ver3.0"
	PartsAddInfo2                 = "2024-12-01-2024-12-31"
	PartsAddInfo3                 = "任意の情報が入ります"
	NotExistID                    = "00000000-0000-0000-0000-000000000000"
	UID                           = "uid"
	DummyTime, _                  = time.Parse("2006-01-02T15:04:05Z", "2024-05-01T00:00:00Z")
	AssertMessage                 = "比較対象の２つの値は定義順に関係なく、一致する必要があります。"
	UnmarshalMockFailureMessage   = "モック定義値のUnmarshalに失敗しました: %v"
	UnmarshalExpectFailureMessage = "正解値のUnmarshalに失敗しました: %v"

	PutTradeRequestInputWithNilTradeID = traceability.PutTradeRequestInput{
		Trade: traceability.PutTradeInput{
			TradeID:              nil,
			DownstreamOperatorID: OperatorId,
			UpstreamOperatorID:   OperatorId,
			DownstreamTraceID:    TraceId,
		},
		Status: traceability.PutStatusInput{
			StatusID:     &StatusId,
			TradeID:      nil,
			Message:      &TradeRequestMessage,
			RequestType:  RequestType,
			ReplyMessage: nil,
		},
	}
	PutTradeResponseInput = traceability.PutTradeResponseInput{
		OperatorID: uuid.MustParse("b1234567-1234-1234-1234-123456789012"),
		TradeID:    uuid.MustParse("a84012cc-73fb-4f9b-9130-59ae546f7092"),
		TraceID:    uuid.MustParse("38bdd8a5-76a7-a53d-de12-725707b04a1b"),
	}
	PutTradeResponseInput2 = traceability.PutTradeResponseInput{
		OperatorID: uuid.MustParse(OperatorID),
		TradeID:    uuid.MustParse(TradeID),
		TraceID:    uuid.MustParse(TraceID),
	}
)

func NewPutPartsInput() traceability.PutPartsInput {
	return traceability.PutPartsInput{
		OperatorID:         OperatorID,
		TraceID:            &TraceID,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   &SupportPartsName,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &AmountRequiredUnit,
		PartsLabelName:     &PartsLabelName,
		PartsAddInfo1:      &PartsAddInfo1,
		PartsAddInfo2:      &PartsAddInfo2,
		PartsAddInfo3:      &PartsAddInfo3,
	}
}

func NewPutPartsInput_RequiredOnly() traceability.PutPartsInput {
	return traceability.PutPartsInput{
		OperatorID:         OperatorID,
		TraceID:            &TraceID,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   &SupportPartsName,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &AmountRequiredUnit,
	}
}

func NewPutPartsInterface() interface{} {
	return map[string]interface{}{
		"operatorId":         OperatorId,
		"traceId":            TraceId,
		"plantId":            PlantId,
		"partsName":          PartsName,
		"supportPartsName":   &SupportPartsName,
		"terminatedFlag":     &TerminatedFlag,
		"amountRequired":     &AmountRequired,
		"amountRequiredUnit": &AmountRequiredUnit,
		"PartsLabelName":     &PartsLabelName,
		"PartsAddInfo1":      &PartsAddInfo1,
		"PartsAddInfo2":      &PartsAddInfo2,
		"PartsAddInfo3":      &PartsAddInfo3,
	}

}

func NewPutPartsStructureInterface() interface{} {
	return map[string]interface{}{
		"parentPartsModel": map[string]interface{}{
			"operatorId":         OperatorId,
			"traceId":            TraceId,
			"plantId":            PlantId,
			"partsName":          PartsName,
			"supportPartsName":   &SupportPartsName,
			"terminatedFlag":     &TerminatedFlag,
			"amountRequired":     nil,
			"amountRequiredUnit": &AmountRequiredUnit,
			"partsLabelName":     &PartsLabelName,
			"partsAddInfo1":      &PartsAddInfo1,
			"partsAddInfo2":      &PartsAddInfo2,
			"partsAddInfo3":      &PartsAddInfo3,
		},
		"childrenPartsModel": []interface{}{
			map[string]interface{}{
				"operatorId":         OperatorId,
				"traceId":            TraceId,
				"plantId":            PlantId,
				"partsName":          PartsName,
				"supportPartsName":   &SupportPartsName,
				"terminatedFlag":     &TerminatedFlag,
				"amountRequired":     &AmountRequired,
				"amountRequiredUnit": &AmountRequiredUnit,
				"partsLabelName":     &PartsLabelName,
				"partsAddInfo1":      &PartsAddInfo1,
				"partsAddInfo2":      &PartsAddInfo2,
				"partsAddInfo3":      &PartsAddInfo3,
			},
		},
	}
}

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
		PartsLabelName:     &PartsLabelName,
		PartsAddInfo1:      &PartsAddInfo1,
		PartsAddInfo2:      &PartsAddInfo2,
		PartsAddInfo3:      &PartsAddInfo3,
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
			PartsLabelName:     &PartsLabelName,
			PartsAddInfo1:      &PartsAddInfo1,
			PartsAddInfo2:      &PartsAddInfo2,
			PartsAddInfo3:      &PartsAddInfo3,
		},
	}
	return traceability.PutPartsStructureInput{
		ParentPartsInput:   &parentPartsModel,
		ChildrenPartsInput: &childrenPartsModel,
	}
}
func NewPutPartsStructureInput_Insert() traceability.PutPartsStructureInput {
	parentPartsModel := traceability.PutPartsInput{
		OperatorID:         OperatorId,
		TraceID:            nil,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   &SupportPartsName,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &AmountRequiredUnit,
		PartsLabelName:     &PartsLabelName,
		PartsAddInfo1:      &PartsAddInfo1,
		PartsAddInfo2:      &PartsAddInfo2,
		PartsAddInfo3:      &PartsAddInfo3,
	}
	childrenPartsModel := traceability.PutPartsInputs{
		traceability.PutPartsInput{
			OperatorID:         OperatorId,
			TraceID:            nil,
			PlantID:            PlantId,
			PartsName:          PartsName,
			SupportPartsName:   &SupportPartsName,
			TerminatedFlag:     &TerminatedFlag,
			AmountRequired:     &AmountRequired,
			AmountRequiredUnit: &AmountRequiredUnit,
			PartsLabelName:     &PartsLabelName,
			PartsAddInfo1:      &PartsAddInfo1,
			PartsAddInfo2:      &PartsAddInfo2,
			PartsAddInfo3:      &PartsAddInfo3,
		},
	}
	return traceability.PutPartsStructureInput{
		ParentPartsInput:   &parentPartsModel,
		ChildrenPartsInput: &childrenPartsModel,
	}
}
func NewPutPartsStructureInput_RequiredOnly() traceability.PutPartsStructureInput {
	parentPartsModel := traceability.PutPartsInput{
		OperatorID:         OperatorId,
		TraceID:            &TraceId,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   nil,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: nil,
		PartsLabelName:     nil,
		PartsAddInfo1:      nil,
		PartsAddInfo2:      nil,
		PartsAddInfo3:      nil,
	}
	childrenPartsModel := traceability.PutPartsInputs{
		traceability.PutPartsInput{
			OperatorID:         OperatorId,
			TraceID:            &TraceId,
			PlantID:            PlantId,
			PartsName:          PartsName,
			SupportPartsName:   nil,
			TerminatedFlag:     &TerminatedFlag,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			PartsLabelName:     nil,
			PartsAddInfo1:      nil,
			PartsAddInfo2:      nil,
			PartsAddInfo3:      nil,
		},
	}
	return traceability.PutPartsStructureInput{
		ParentPartsInput:   &parentPartsModel,
		ChildrenPartsInput: &childrenPartsModel,
	}
}
func NewPutPartsStructureInput_RequiredOnlyWithUndefined() traceability.PutPartsStructureInput {
	parentPartsModel := traceability.PutPartsInput{
		OperatorID:         OperatorId,
		TraceID:            &TraceId,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   nil,
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
			SupportPartsName:   nil,
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
func NewPutPartsStructureInput_NoComponent() traceability.PutPartsStructureInput {
	parentPartsModel := traceability.PutPartsInput{
		OperatorID:         OperatorId,
		TraceID:            &TraceId,
		PlantID:            PlantId,
		PartsName:          PartsName,
		SupportPartsName:   nil,
		TerminatedFlag:     &TerminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &AmountRequiredUnit,
		PartsLabelName:     &PartsLabelName,
		PartsAddInfo1:      &PartsAddInfo1,
		PartsAddInfo2:      &PartsAddInfo2,
		PartsAddInfo3:      &PartsAddInfo3,
	}
	childrenPartsModel := traceability.PutPartsInputs{}
	return traceability.PutPartsStructureInput{
		ParentPartsInput:   &parentPartsModel,
		ChildrenPartsInput: &childrenPartsModel,
	}
}

func NewPutCfpInterface() []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cfpId":           nil,
			"traceId":         TraceId,
			"ghgEmission":     0,
			"ghgDeclaredUnit": GhgDeclaredUnit,
			"cfpType":         traceability.CfpTypeMainProduction.ToString(),
			"dqrType":         traceability.DqrTypeMainProcessing.ToString(),
			"dqrValue": map[string]interface{}{
				"TeR": 0,
				"GeR": 0,
				"TiR": 0,
			},
		},
		map[string]interface{}{
			"cfpId":           nil,
			"traceId":         TraceId,
			"ghgEmission":     0,
			"ghgDeclaredUnit": GhgDeclaredUnit,
			"cfpType":         traceability.CfpTypePreProduction.ToString(),
			"dqrType":         traceability.DqrTypePreProcessing.ToString(),
			"dqrValue": map[string]interface{}{
				"TeR": 0,
				"GeR": 0,
				"TiR": 0,
			},
		},
		map[string]interface{}{
			"cfpId":           nil,
			"traceId":         TraceId,
			"ghgEmission":     10,
			"ghgDeclaredUnit": GhgDeclaredUnit,
			"cfpType":         traceability.CfpTypeMainComponent.ToString(),
			"dqrType":         traceability.DqrTypeMainProcessing.ToString(),
			"dqrValue": map[string]interface{}{
				"TeR": 0,
				"GeR": 0,
				"TiR": 0,
			},
		},
		map[string]interface{}{
			"cfpId":           nil,
			"traceId":         TraceId,
			"ghgEmission":     0,
			"ghgDeclaredUnit": GhgDeclaredUnit,
			"cfpType":         traceability.CfpTypePreComponent.ToString(),
			"dqrType":         traceability.DqrTypePreProcessing.ToString(),
			"dqrValue": map[string]interface{}{
				"TeR": 0,
				"GeR": 0,
				"TiR": 0,
			},
		},
	}
}
func NewPutCfpInputs2() traceability.PutCfpInputs {
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

func NewGetPartsStructureInput() traceability.GetPartsStructureInput {
	return traceability.GetPartsStructureInput{
		TraceID:    uuid.MustParse(TraceId),
		OperatorID: OperatorId,
	}
}

func NewGetCfpCertificationInput() traceability.GetCfpCertificationInput {
	return traceability.GetCfpCertificationInput{
		OperatorID: uuid.MustParse(OperatorId),
		TraceID:    uuid.MustParse(TraceId),
	}
}

func NewGetStatusInput() traceability.GetStatusInput {
	traceId := uuid.MustParse(TraceId)
	statusId := uuid.MustParse(StatusId)
	after := uuid.MustParse(StatusID)
	status := traceability.GetStatusInput{
		OperatorID:   uuid.MustParse(OperatorId),
		Limit:        1,
		After:        &after,
		StatusID:     &statusId,
		TraceID:      &traceId,
		StatusTarget: "",
	}
	return status
}

func NewGetCfpInput() traceability.GetCfpInput {
	return traceability.GetCfpInput{
		OperatorID: uuid.MustParse(OperatorId),
		TraceIDs:   []uuid.UUID{uuid.MustParse(TraceId)},
	}
}

func NewCfpModels() traceability.CfpModels {
	cfpUUID := uuid.MustParse(CfpId)
	return traceability.CfpModels{
		{
			CfpID:           &cfpUUID,
			TraceID:         uuid.MustParse(TraceID),
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
			CfpID:           &cfpUUID,
			TraceID:         uuid.MustParse(TraceID),
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
			CfpID:           &cfpUUID,
			TraceID:         uuid.MustParse(TraceID),
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
			CfpID:           &cfpUUID,
			TraceID:         uuid.MustParse(TraceID),
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

func NewGetPartsInput() traceability.GetPartsInput {
	parentFlag := true
	after := uuid.MustParse(TraceId)
	return traceability.GetPartsInput{
		TraceID:    &TraceId,
		OperatorID: OperatorId,
		PartsName:  &PartsName,
		PlantID:    &PlantId,
		ParentFlag: &parentFlag,
		Limit:      100,
		After:      &after,
	}
}

func NewGetPartsInput_RequiredOnly() traceability.GetPartsInput {
	parentFlag := true
	after := uuid.MustParse(TraceId)
	return traceability.GetPartsInput{
		TraceID:    &TraceID7,
		OperatorID: OperatorId,
		PartsName:  common.StringPtr("製品A7"),
		PlantID:    &PlantId,
		ParentFlag: &parentFlag,
		Limit:      100,
		After:      &after,
	}
}

func NewDeletePartsInput(traceId string) traceability.DeletePartsInput {
	return traceability.DeletePartsInput{
		TraceID: traceId,
	}
}

func NewPutPartsStructureEntityModel() traceability.PartsStructureEntity {
	plantId := uuid.MustParse(PlantID)
	supportPartsNameP := "A000001"
	supportPartsNameC := "B001"
	AmountRequiredUnit := traceability.AmountRequiredUnitKilogram
	partsModel := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			OperatorID:         uuid.MustParse(OperatorID),
			TraceID:            uuid.MustParse(TraceID3),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   &supportPartsNameP,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr(AmountRequiredUnit.ToString()),
			PartsLabelName:     &PartsLabelName,
			PartsAddInfo1:      &PartsAddInfo1,
			PartsAddInfo2:      &PartsAddInfo2,
			PartsAddInfo3:      &PartsAddInfo3,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				OperatorID:         uuid.MustParse(OperatorID),
				TraceID:            uuid.MustParse(TraceIDChild),
				PlantID:            plantId,
				PartsName:          "B01001",
				SupportPartsName:   &supportPartsNameC,
				TerminatedFlag:     false,
				AmountRequired:     common.Float64Ptr(2.1),
				AmountRequiredUnit: common.StringPtr(AmountRequiredUnit.ToString()),
				PartsLabelName:     &PartsLabelName,
				PartsAddInfo1:      &PartsAddInfo1,
				PartsAddInfo2:      &PartsAddInfo2,
				PartsAddInfo3:      &PartsAddInfo3,
				DeletedAt:          gorm.DeletedAt{Time: time.Now()},
				CreatedAt:          time.Now(),
				CreatedUserId:      "seed",
				UpdatedAt:          time.Now(),
				UpdatedUserId:      "seed",
			},
		},
	}
	return partsModel
}
func NewPutPartsStructureEntityModel_RequiredOnly() traceability.PartsStructureEntity {
	plantId := uuid.MustParse(PlantID)
	partsModel := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			OperatorID:         uuid.MustParse(OperatorID),
			TraceID:            uuid.MustParse(TraceID3),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: nil,
			PartsLabelName:     nil,
			PartsAddInfo1:      nil,
			PartsAddInfo2:      nil,
			PartsAddInfo3:      nil,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				OperatorID:         uuid.MustParse(OperatorID),
				TraceID:            uuid.MustParse(TraceIDChild),
				PlantID:            plantId,
				PartsName:          "B01001",
				SupportPartsName:   nil,
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: nil,
				PartsLabelName:     nil,
				PartsAddInfo1:      nil,
				PartsAddInfo2:      nil,
				PartsAddInfo3:      nil,
				DeletedAt:          gorm.DeletedAt{Time: time.Now()},
				CreatedAt:          time.Now(),
				CreatedUserId:      "seed",
				UpdatedAt:          time.Now(),
				UpdatedUserId:      "seed",
			},
		},
	}
	return partsModel
}
func NewPutPartsStructureEntityModel_NoComponent() traceability.PartsStructureEntity {
	plantId := uuid.MustParse(PlantID)
	AmountRequiredUnit := traceability.AmountRequiredUnitLiter
	partsModel := traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			OperatorID:         uuid.MustParse(OperatorID),
			TraceID:            uuid.MustParse(TraceID3),
			PlantID:            plantId,
			PartsName:          "B01",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr(AmountRequiredUnit.ToString()),
			PartsLabelName:     &PartsLabelName,
			PartsAddInfo1:      &PartsAddInfo1,
			PartsAddInfo2:      &PartsAddInfo2,
			PartsAddInfo3:      &PartsAddInfo3,
			DeletedAt:          gorm.DeletedAt{Time: time.Now()},
			CreatedAt:          time.Now(),
			CreatedUserId:      "seed",
			UpdatedAt:          time.Now(),
			UpdatedUserId:      "seed",
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{},
	}
	return partsModel
}

func NewPutTradeRequestModel() traceability.TradeRequestModel {
	tradeId := uuid.MustParse(TradeId)
	downstreamOperatorId := uuid.MustParse(OperatorId)
	upstreamOperatorId := uuid.MustParse(OperatorId)
	downstreamTraceId := uuid.MustParse(TraceId)
	statusId := uuid.MustParse(StatusId)
	cfpResponseStatusPending := traceability.CfpResponseStatusPending
	tradeTreeStatusUnterminated := traceability.TradeTreeStatusUnterminated
	partsModel := traceability.TradeRequestModel{
		TradeModel: traceability.TradeModel{
			TradeID:              &tradeId,
			DownstreamOperatorID: downstreamOperatorId,
			UpstreamOperatorID:   upstreamOperatorId,
			DownstreamTraceID:    downstreamTraceId,
			UpstreamTraceID:      nil,
		},
		StatusModel: traceability.StatusModel{
			StatusID: statusId,
			TradeID:  tradeId,
			RequestStatus: traceability.RequestStatus{
				CfpResponseStatus: &cfpResponseStatusPending,
				TradeTreeStatus:   &tradeTreeStatusUnterminated,
			},
			Message:         &TradeRequestMessage,
			ReplyMessage:    nil,
			RequestType:     traceability.RequestTypeCFP.ToString(),
			ResponseDueDate: &ResponseDueDate,
		},
	}
	return partsModel
}

func NewPutTradeResponseModel() traceability.TradeModel {
	tradeId := uuid.MustParse(TradeID)
	traceId := uuid.MustParse(TraceID)
	return traceability.TradeModel{
		TradeID:              &tradeId,
		DownstreamOperatorID: uuid.MustParse(OperatorId),
		UpstreamOperatorID:   uuid.MustParse(OperatorId),
		DownstreamTraceID:    uuid.MustParse(TraceIDDownstream),
		UpstreamTraceID:      &traceId,
	}
}

func GetPartsModelEntity(traceId string, terminatedFlag bool) traceability.PartsModelEntity {
	return traceability.PartsModelEntity{
		TraceID:            uuid.MustParse(traceId),
		OperatorID:         uuid.MustParse(OperatorID),
		PlantID:            uuid.MustParse(PlantID),
		PartsName:          "B01",
		SupportPartsName:   common.StringPtr("A000001"),
		TerminatedFlag:     terminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: common.StringPtr(traceability.AmountRequiredUnitKilogram.ToString()),
	}
}

func GetPartsStructureEntityModel() traceability.PartsStructureEntityModel {
	return traceability.PartsStructureEntityModel{
		TraceID:       uuid.MustParse(TraceIDChild),
		ParentTraceID: uuid.MustParse(TraceID),
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

func GetTradeEntityModel() traceability.TradeEntityModel {
	tradeId := uuid.MustParse(TradeID)
	operatorId := uuid.MustParse(OperatorID)
	traceId := uuid.MustParse(TraceID)
	res := traceability.TradeEntityModel{
		TradeID:              &tradeId,
		DownstreamOperatorID: operatorId,
		UpstreamOperatorID:   &operatorId,
		DownstreamTraceID:    traceId,
		UpstreamTraceID:      &traceId,
		TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
		DeletedAt:            gorm.DeletedAt{Time: time.Now()},
		CreatedAt:            time.Now(),
		CreatedUserID:        "seed",
		UpdatedAt:            time.Now(),
		UpdatedUserID:        "seed",
	}
	return res
}

func GetCfpEntityModels() traceability.CfpEntityModels {
	cfpId := uuid.MustParse(CfpId)
	return traceability.CfpEntityModels{
		{
			CfpID:              &cfpId,
			TraceID:            uuid.MustParse(TraceId),
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
			TraceID:            uuid.MustParse(TraceId),
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
			TraceID:            uuid.MustParse(TraceId),
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
			TraceID:            uuid.MustParse(TraceId),
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

func NewPutStatusInput() traceability.PutStatusInput {
	cfpResponseStatus := traceability.CfpResponseStatusCancel
	tradeTreeStatus := traceability.TradeTreeStatusTerminated
	return traceability.PutStatusInput{
		StatusID:     common.StringPtr(StatusId),
		TradeID:      common.StringPtr(TradeId),
		Message:      &TradeRequestMessage,
		ReplyMessage: nil,
		RequestType:  RequestType,
		PutRequestStatusInput: traceability.PutRequestStatusInput{
			CfpResponseStatus: &cfpResponseStatus,
			TradeTreeStatus:   &tradeTreeStatus,
		},
	}
}

func NewPutStatusInterface() interface{} {
	return map[string]interface{}{
		"statusId":     StatusId,
		"tradeId":      TradeId,
		"message":      TradeRequestMessage,
		"replyMessage": "",
		"requestType":  RequestType.ToString(),
		"putRequestStatusInput": map[string]interface{}{
			"cfpResponseStatus": traceability.CfpResponseStatusCancel.ToString(),
			"tradeTreeStatus":   traceability.TradeTreeStatusTerminated.ToString(),
		},
	}
}

func NewPutTradeRequestInput() traceability.PutTradeRequestInput {
	cfpResponseStatus := traceability.CfpResponseStatusPending
	tradeTreeStatus := traceability.TradeTreeStatusUnterminated

	return traceability.PutTradeRequestInput{
		Trade: traceability.PutTradeInput{
			TradeID:              &TradeId,
			DownstreamOperatorID: OperatorId,
			UpstreamOperatorID:   OperatorId,
			DownstreamTraceID:    TraceId,
		},
		Status: traceability.PutStatusInput{
			StatusID:        &StatusId,
			TradeID:         &TradeId,
			Message:         &TradeRequestMessage,
			RequestType:     RequestType,
			ReplyMessage:    nil,
			ResponseDueDate: ResponseDueDate,
			PutRequestStatusInput: traceability.PutRequestStatusInput{
				CfpResponseStatus: &cfpResponseStatus,
				TradeTreeStatus:   &tradeTreeStatus,
			},
		},
	}
}

func CraeteRamdomString(length int) string {
	rs1Letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, length)
	for i := range str {
		str[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(str)
}

func NewBatchCreateCFPInput() traceability.CfpEntityModels {
	return traceability.CfpEntityModels{
		&traceability.CfpEntityModel{
			CfpID:              common.UUIDPtr(uuid.MustParse(CfpId)),
			TraceID:            uuid.MustParse(TraceID3),
			GhgEmission:        common.Float64Ptr(0.1),
			GhgDeclaredUnit:    GhgDeclaredUnit,
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypePreProduction.ToString(),
			DqrType:            traceability.DqrTypePreProcessing.ToString(),
			TeR:                common.Float64Ptr(1.1),
			GeR:                common.Float64Ptr(2.1),
			TiR:                common.Float64Ptr(3.1),
			DeletedAt:          gorm.DeletedAt{},
			CreatedAt:          DummyTime,
			CreatedUserId:      "seed",
			UpdatedAt:          DummyTime,
			UpdatedUserId:      "seed",
		},
		&traceability.CfpEntityModel{
			CfpID:              common.UUIDPtr(uuid.MustParse(CfpId2)),
			TraceID:            uuid.MustParse(TraceID3),
			GhgEmission:        common.Float64Ptr(0.2),
			GhgDeclaredUnit:    GhgDeclaredUnit,
			CfpCertificateList: CfpCertificateList,
			CfpType:            traceability.CfpTypeMainProduction.ToString(),
			DqrType:            traceability.DqrTypePreProcessing.ToString(),
			TeR:                common.Float64Ptr(1.2),
			GeR:                common.Float64Ptr(2.2),
			TiR:                common.Float64Ptr(3.2),
			DeletedAt:          gorm.DeletedAt{},
			CreatedAt:          DummyTime,
			CreatedUserId:      "seed",
			UpdatedAt:          DummyTime,
			UpdatedUserId:      "seed",
		},
	}
}

func NewPutCFPInput() traceability.CfpEntityModel {
	return traceability.CfpEntityModel{
		CfpID:              common.UUIDPtr(uuid.MustParse(CfpId)),
		TraceID:            uuid.MustParse(TraceID3),
		GhgEmission:        common.Float64Ptr(0.1),
		GhgDeclaredUnit:    GhgDeclaredUnit,
		CfpCertificateList: CfpCertificateList,
		CfpType:            traceability.CfpTypePreProduction.ToString(),
		DqrType:            traceability.DqrTypePreProcessing.ToString(),
		TeR:                common.Float64Ptr(1.1),
		GeR:                common.Float64Ptr(2.1),
		TiR:                common.Float64Ptr(3.1),
		DeletedAt:          gorm.DeletedAt{},
		CreatedAt:          DummyTime,
		CreatedUserId:      "seed",
		UpdatedAt:          DummyTime,
		UpdatedUserId:      "seed",
	}
}

func NewPutTradeRequestModelInput() traceability.TradeRequestEntityModel {
	return traceability.TradeRequestEntityModel{
		TradeEntityModel: traceability.TradeEntityModel{
			TradeID:              common.UUIDPtr(uuid.MustParse(TradeId)),
			DownstreamOperatorID: uuid.MustParse(OperatorId),
			UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(OperatorId)),
			DownstreamTraceID:    uuid.MustParse(TraceId),
			UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(TraceID2)),
			TradeDate:            common.StringPtr("2023-09-25T14:30:00Z"),
			DeletedAt:            gorm.DeletedAt{},
			CreatedAt:            DummyTime,
			CreatedUserID:        "seed",
			UpdatedAt:            DummyTime,
			UpdatedUserID:        "seed",
		},
		StatusEntityModel: traceability.StatusEntityModel{
			StatusID:      uuid.MustParse(StatusId),
			TradeID:       uuid.MustParse(TradeID),
			Message:       &TradeRequestMessage,
			ReplyMessage:  common.StringPtr(""),
			RequestType:   "CFP",
			DeletedAt:     gorm.DeletedAt{},
			CreatedAt:     DummyTime,
			CreatedUserId: "seed",
			UpdatedAt:     DummyTime,
			UpdatedUserId: "seed",
		},
	}
}

func NewRequestStatus() traceability.RequestStatus {
	cfpResponseStatusCancel := traceability.CfpResponseStatusCancel
	tradeTreeStatusUnterminated := traceability.TradeTreeStatusUnterminated
	return traceability.RequestStatus{
		CfpResponseStatus: &cfpResponseStatusCancel,
		TradeTreeStatus:   &tradeTreeStatusUnterminated,
	}
}

func NewTradeEntityModel() traceability.TradeEntityModel {
	return traceability.TradeEntityModel{
		TradeID:              common.UUIDPtr(uuid.MustParse(TradeId)),
		DownstreamOperatorID: uuid.MustParse(OperatorID2),
		UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(OperatorID)),
		DownstreamTraceID:    uuid.MustParse(TraceID2),
		UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(TraceID)),
		TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
		DeletedAt:            gorm.DeletedAt{},
		CreatedAt:            DummyTime,
		CreatedUserID:        "seed",
		UpdatedAt:            DummyTime,
		UpdatedUserID:        "seed",
	}
}

func NewStatusModels() traceability.StatusEntityModels {
	return traceability.StatusEntityModels{
		NewStatusModel2(),
	}
}

func NewStatusModel2() traceability.StatusEntityModel {
	expectTime, _ := time.Parse("2006-01-02T15:04:05Z", "2024-05-01T00:00:00Z")
	return traceability.StatusEntityModel{
		StatusID:                 uuid.MustParse(StatusID),
		TradeID:                  uuid.MustParse(TradeID),
		CfpResponseStatus:        traceability.CfpResponseStatusComplete.ToString(),
		TradeTreeStatus:          traceability.TradeTreeStatusTerminated.ToString(),
		Message:                  common.StringPtr("来月中にご回答をお願いします。"),
		ReplyMessage:             nil,
		RequestType:              "CFP",
		ResponseDueDate:          "2024-05-01",
		CompletedCount:           common.IntPtr(1),
		CompletedCountModifiedAt: &expectTime,
		TradesCount:              common.IntPtr(1),
		TradesCountModifiedAt:    &expectTime,
		DeletedAt:                gorm.DeletedAt{},
		CreatedAt:                DummyTime,
		CreatedUserId:            "seed",
		UpdatedAt:                DummyTime,
		UpdatedUserId:            "seed",
	}
}

func NewPartsStructureModel() traceability.PartsStructureModel {

	return traceability.PartsStructureModel{
		ParentPartsModel: &traceability.PartsModel{
			TraceID:            uuid.MustParse(TraceID5),
			OperatorID:         uuid.MustParse(OperatorID),
			PlantID:            common.UUIDPtr(uuid.MustParse(PlantId)),
			PartsName:          "製品A3",
			SupportPartsName:   common.StringPtr("品番A3"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: &AmountRequiredUnit2,
			PartsLabelName:     common.StringPtr("PartsA3"),
			PartsAddInfo1:      common.StringPtr("Ver3.0"),
			PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
			PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
		},
		ChildrenPartsModel: []traceability.PartsModel{
			{
				TraceID:            uuid.MustParse(TraceID6),
				OperatorID:         uuid.MustParse(OperatorID),
				PlantID:            common.UUIDPtr(uuid.MustParse(PlantId)),
				PartsName:          "製品A5",
				SupportPartsName:   common.StringPtr("品番A5"),
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: &AmountRequiredUnit2,
				PartsLabelName:     common.StringPtr("PartsA5"),
				PartsAddInfo1:      common.StringPtr("Ver3.0"),
				PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
				PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
			},
		},
	}
}

func NewPartsStructureModel_RequiredOnly() traceability.PartsStructureModel {

	return traceability.PartsStructureModel{
		ParentPartsModel: &traceability.PartsModel{
			TraceID:            uuid.MustParse(TraceID7),
			OperatorID:         uuid.MustParse(OperatorID),
			PlantID:            common.UUIDPtr(uuid.MustParse(PlantId)),
			PartsName:          "製品A7",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: &AmountRequiredUnit2,
			PartsLabelName:     nil,
			PartsAddInfo1:      nil,
			PartsAddInfo2:      nil,
			PartsAddInfo3:      nil,
		},
		ChildrenPartsModel: []traceability.PartsModel{
			{
				TraceID:            uuid.MustParse(TraceID8),
				OperatorID:         uuid.MustParse(OperatorID),
				PlantID:            common.UUIDPtr(uuid.MustParse(PlantId)),
				PartsName:          "製品A8",
				SupportPartsName:   nil,
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: &AmountRequiredUnit2,
				PartsLabelName:     nil,
				PartsAddInfo1:      nil,
				PartsAddInfo2:      nil,
				PartsAddInfo3:      nil,
			},
		},
	}
}

func NewPartsStructureModel_RequiredOnlyWithUndefined() traceability.PartsStructureModel {

	return traceability.PartsStructureModel{
		ParentPartsModel: &traceability.PartsModel{
			TraceID:            uuid.MustParse(TraceID7),
			OperatorID:         uuid.MustParse(OperatorID),
			PlantID:            common.UUIDPtr(uuid.MustParse(PlantId)),
			PartsName:          "製品A7",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: &AmountRequiredUnit2,
		},
		ChildrenPartsModel: []traceability.PartsModel{
			{
				TraceID:            uuid.MustParse(TraceID8),
				OperatorID:         uuid.MustParse(OperatorID),
				PlantID:            common.UUIDPtr(uuid.MustParse(PlantId)),
				PartsName:          "製品A8",
				SupportPartsName:   nil,
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: &AmountRequiredUnit2,
			},
		},
	}
}

func NewPartsStructureEntity() traceability.PartsStructureEntity {
	return traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse(TraceID5),
			OperatorID:         uuid.MustParse(OperatorID),
			PlantID:            uuid.MustParse(PlantId),
			PartsName:          "製品A3",
			SupportPartsName:   common.StringPtr("品番A3"),
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr("kilogram"),
			PartsLabelName:     common.StringPtr("PartsA3"),
			PartsAddInfo1:      common.StringPtr("Ver3.0"),
			PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
			PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				TraceID:            uuid.MustParse(TraceID6),
				OperatorID:         uuid.MustParse(OperatorID),
				PlantID:            uuid.MustParse(PlantId),
				PartsName:          "製品A5",
				SupportPartsName:   common.StringPtr("品番A5"),
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: common.StringPtr("kilogram"),
				PartsLabelName:     common.StringPtr("PartsA5"),
				PartsAddInfo1:      common.StringPtr("Ver3.0"),
				PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
				PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
			},
		},
	}
}

func NewPartsStructureEntity_RequiredOnly() traceability.PartsStructureEntity {
	return traceability.PartsStructureEntity{
		ParentPartsEntity: &traceability.PartsModelEntity{
			TraceID:            uuid.MustParse(TraceID7),
			OperatorID:         uuid.MustParse(OperatorID),
			PlantID:            uuid.MustParse(PlantId),
			PartsName:          "製品A7",
			SupportPartsName:   nil,
			TerminatedFlag:     false,
			AmountRequired:     nil,
			AmountRequiredUnit: common.StringPtr("kilogram"),
			PartsLabelName:     nil,
			PartsAddInfo1:      nil,
			PartsAddInfo2:      nil,
			PartsAddInfo3:      nil,
		},
		ChildrenPartsEntity: traceability.PartsModelEntities{
			{
				TraceID:            uuid.MustParse(TraceID8),
				OperatorID:         uuid.MustParse(OperatorID),
				PlantID:            uuid.MustParse(PlantId),
				PartsName:          "製品A8",
				SupportPartsName:   nil,
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: common.StringPtr("kilogram"),
				PartsLabelName:     nil,
				PartsAddInfo1:      nil,
				PartsAddInfo2:      nil,
				PartsAddInfo3:      nil,
			},
		},
	}
}

func NewTraceabilityError(errorCode string, errorDescription string, relevantUUIDs []uuid.UUID) error {
	err := common.TraceabilityAPIErrorDetailDelete{
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
	}

	if len(relevantUUIDs) > 0 {
		relevantData := common.UUIDsToStrings(relevantUUIDs)
		err.RelevantData = &relevantData
	}

	traceabilityError := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			err,
		},
	}
	return traceabilityError.ToCustomError(400)
}
