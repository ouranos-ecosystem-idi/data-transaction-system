package logger

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Set
// Summary: This is function to set logger.
// input: c(echo.Context) echo context
// output: (*zap.SugaredLogger) logger
func Set(c echo.Context) *zap.SugaredLogger {
	var operatorID string
	var trackID string
	if c != nil {
		i := c.Get("operatorID")
		if i != nil {
			operatorID = i.(string)
		}
		trackID = getTrackID(c.Request())
	}
	return zap.S().With("operator_id", operatorID, "track_id", trackID)
}

// getTrackID
// Summary: This is function to get request id.
// input: r(*http.Request) http request
// output: (string) ID of the track
func getTrackID(r *http.Request) string {
	traceHeader := r.Header.Get("X-Cloud-Trace-Context")
	traceParts := strings.Split(traceHeader, "/")
	trackID := ""
	if len(traceParts) > 0 {
		trackID = traceParts[0]
	}
	return trackID
}

// Log messages
const (
	AccessInfoLog      = "Access URL: %v"
	DataSpaceAPILog    = "DataSpaceAPI Access Path: %v, Header: %v, Request Body: %v, Response Body: %v, X-Track: %v"
	TraceabilityAPILog = "TraceabilityAPI Access URL: %v, Header: %v, Request Body: %v, Response Body: %v, X-Track: %v"
)
