package handler

import (
	"net/http"

	"data-spaces-backend/domain/common"

	"github.com/labstack/echo/v4"
)

// DeleteOuranos
// Summary: This is the function which call the handler depending on the dataTarget query parameter.
// input: c(echo.Context): echo context
// output: (error) error object
func (h *ouranosHandler) DeleteOuranos(c echo.Context) error {
	method := c.Request().Method
	operatorID := c.Get("operatorID").(string)

	dataTarget := c.QueryParam("dataTarget")

	switch dataTarget {
	case "parts":
		return h.partsHandler.DeletePartsModel(c)
	default:
		errDetails := common.UnexpectedQueryParameter("dataTarget")
		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceDataspace, common.Err400InvalidRequest, operatorID, dataTarget, method, errDetails))
	}
}
