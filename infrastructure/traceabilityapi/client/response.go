package client

import (
	"data-spaces-backend/domain/common"
	"net/http"
)

const HeaderXTrack = "X-Track"

type Response struct {
	Body    string
	Headers common.ResponseHeaders
}

type ResponseHeaders struct {
	XTrack string
}

func SetResponseHeaders(resp *http.Response) ResponseHeaders {
	return ResponseHeaders{
		XTrack: resp.Header.Get(HeaderXTrack),
	}
}
