package handler

import "github.com/labstack/echo/v4"

type (
	OuranosHandler interface {
		GetOuranos(c echo.Context) error
		PutOuranos(c echo.Context) error
		DeleteOuranos(c echo.Context) error
	}

	ouranosHandler struct {
		cfpHandler              ICfpHandler
		cfpCertificationHandler ICfpCertificationHandler
		partsHandler            IPartsHandler
		partsStructureHandler   IPartsStructureHandler
		tradeHandler            ITradeHandler
		statusHandler           IStatusHandler
	}
)

// NewOuranosHandler
// Summary: This is function which creates new OuranosHandler.
// input: cfpHandler(ICfpHandler) CfpHandler
// input: cfpCertificationHandler(ICfpCertificationHandler) CfpCertificationHandler
// input: partsHandler(IPartsHandler) PartsHandler
// input: partsStructureHandler(IPartsStructureHandler) PartsStructureHandler
// input: tradeHandler(ITradeHandler) TradeHandler
// input: statusHandler(IStatusHandler) StatusHandler
// output: (OuranosHandler) OuranosHandler object
func NewOuranosHandler(
	cfpHandler ICfpHandler,
	cfpCertificationHandler ICfpCertificationHandler,
	partsHandler IPartsHandler,
	partsStructureHandler IPartsStructureHandler,
	tradeHandler ITradeHandler,
	statusHandler IStatusHandler,
) OuranosHandler {
	return &ouranosHandler{
		cfpHandler,
		cfpCertificationHandler,
		partsHandler,
		partsStructureHandler,
		tradeHandler,
		statusHandler,
	}
}
