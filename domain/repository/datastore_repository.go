package repository

import (
	"data-spaces-backend/domain/model/traceability"
)

//go:generate mockery --name OuranosRepository --output ../../test/mock --case underscore
type (
	OuranosRepository interface {
		// Parts
		ListParts(getPlantPartsModel traceability.GetPartsInput) (traceability.PartsModelEntities, error)
		GetPartByTraceID(traceID string) (traceability.PartsModelEntity, error)
		CountPartsList(getPlantPartsModel traceability.GetPartsInput) (int, error)
		DeleteParts(traceID string) error
		DeletePartsWithCFP(traceID string) error

		// PartsStructure
		GetPartsStructure(getPartsStructureInput traceability.GetPartsStructureInput) (traceability.PartsStructureEntity, error)
		GetPartsStructureByTraceId(traceID string) (traceability.PartsStructureEntityModel, error)
		ListParentPartsStructureByTraceId(traceID string) (traceability.PartsStructureEntityModels, error)
		ListChildPartsStructureByTraceId(traceID string) (traceability.PartsStructureEntityModels, error)

		PutPartsStructure(partsStructure traceability.PartsStructureModel) (traceability.PartsStructureEntity, error)
		DeletePartsStructure(traceID string) error

		// Trade
		GetTradeRequest(downstreamOperatorID string, limit int, traceIDs []string) (traceability.TradeEntityModels, error)
		GetTradeResponse(upstreamOperatorID string, limit int) (traceability.TradeEntityModels, error)
		GetTradeByDownstreamTraceID(donwstreamTraceID string) (traceability.TradeEntityModel, error)
		GetTrade(tradeID string) (traceability.TradeEntityModel, error)
		ListTradeByUpstreamTraceID(upstreamTraceID string) (traceability.TradeEntityModels, error)
		ListTradeByDownstreamTraceID(downstreamTraceID string) (traceability.TradeEntityModels, error)
		CountTradeRequest(downstreamOperatorID string) (int, error)
		CountTradeResponse(upstreamOperatorID string) (int, error)
		PutTradeRequest(tradeRequestEntityModel traceability.TradeRequestEntityModel) (traceability.TradeRequestEntityModel, error)
		PutTradeResponse(putTradeResponseInput traceability.PutTradeResponseInput, requestStatusValue traceability.RequestStatus) (traceability.TradeEntityModel, error)
		ListTradesByOperatorID(operatorID string) (traceability.TradeEntityModels, error)
		DeleteTrade(tradeID string) error

		// RequestStatus
		GetStatusByTradeID(tradeID string) (traceability.StatusEntityModel, error)
		GetStatus(operatorID string, limit int, statusID *string, traceID *string, statusTarget string) (traceability.StatusEntityModels, error)
		CountStatus(operatorID string, statusID *string, traceID *string, statusTarget string) (int, error)
		PutStatusCancel(statusID string, operatorID string) error
		PutStatusReject(statusID string, replyMessage *string, operatorID string) (traceability.StatusEntityModel, error)
		DeleteRequestStatusByTradeID(tradeID string) error

		// CFP
		BatchCreateCFP(es traceability.CfpEntityModels) (traceability.CfpEntityModels, error)
		GetCFP(cfpID string, cfpType string) (traceability.CfpEntityModel, error)
		ListCFPsByTraceID(traceID string) (traceability.CfpEntityModels, error)
		PutCFP(e traceability.CfpEntityModel) (traceability.CfpEntityModel, error)

		// CFPInfomation
		GetCFPInformation(traceID string) (traceability.CfpEntityModel, error)
		DeleteCFPInformation(cfpID string) error

		// CFPCertification
		GetCFPCertifications(operatorID string, traceID string) (traceability.CfpCertificationModels, error)
	}
)
