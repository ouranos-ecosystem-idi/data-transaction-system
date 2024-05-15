package testhelper

import (
	"data-spaces-backend/presentation/http/echo/handler"
	mocks "data-spaces-backend/test/mock"
)

func NewMockHandler(host string) handler.OuranosHandler {
	cfpUsecase := new(mocks.ICfpUsecase)
	cfpHandler := handler.NewCfpHandler(cfpUsecase)
	cfpCertificationUsecase := new(mocks.ICfpCertificationUsecase)
	cfpCertificationHandler := handler.NewCfpCertificationHandler(cfpCertificationUsecase)
	partsUsecase := new(mocks.IPartsUsecase)
	partsStructureUsecase := new(mocks.IPartsStructureUsecase)
	partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, host)
	partsStructureHandler := handler.NewPartsStructureHandler(partsStructureUsecase)
	tradeUsecase := new(mocks.ITradeUsecase)
	tradeHandler := handler.NewTradeHandler(tradeUsecase, host)
	statusUsecase := new(mocks.IStatusUsecase)
	statusHandler := handler.NewStatusHandler(statusUsecase, host)
	h := handler.NewOuranosHandler(cfpHandler, cfpCertificationHandler, partsHandler, partsStructureHandler, tradeHandler, statusHandler)

	return h
}
