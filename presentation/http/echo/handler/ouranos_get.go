package handler

import (
	"net/http"

	"data-spaces-backend/domain/common"

	"github.com/labstack/echo/v4"
)

// GetOuranos
// Summary: This is the function which call the handler depending on the dataTarget query parameter.
// input: c(echo.Context) echo context
// output: (error) error object
func (h *ouranosHandler) GetOuranos(c echo.Context) error {
	method := c.Request().Method
	operatorID := c.Get("operatorID").(string)

	dataTarget := c.QueryParam("dataTarget")

	switch dataTarget {
	case "partsStructure":
		return h.partsStructureHandler.GetPartsStructureModel(c)
	case "parts":
		return h.partsHandler.GetPartsModel(c)
	case "tradeRequest":
		return h.tradeHandler.GetTradeRequest(c)
	case "tradeResponse":
		return h.tradeHandler.GetTradeResponse(c)
	case "cfp":
		return h.cfpHandler.GetCfp(c)
	case "cfpCertification":
		return h.cfpCertificationHandler.GetCfpCertification(c)
	case "status":
		return h.statusHandler.GetStatus(c)
	default:
		errDetails := common.UnexpectedQueryParameter("dataTarget")
		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
}
