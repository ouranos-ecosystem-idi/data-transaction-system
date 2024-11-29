package repository

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name TraceabilityRepository --output ../../test/mock --case underscore
type (
	TraceabilityRepository interface {

		// 部品情報検索API
		GetParts(c echo.Context, request traceabilityentity.GetPartsRequest, limit int) (traceabilityentity.GetPartsResponse, error)
		// 部品情報削除API
		DeleteParts(c echo.Context, request traceabilityentity.DeletePartsRequest) (traceabilityentity.DeletePartsResponse, common.ResponseHeaders, error)
		// 部品構成情報登録API
		GetPartsStructures(c echo.Context, request traceabilityentity.GetPartsStructuresRequest) (traceabilityentity.GetPartsStructuresResponse, error)
		PostPartsStructures(c echo.Context, request traceabilityentity.PostPartsStructuresRequest) (traceabilityentity.PostPartsStructuresResponse, common.ResponseHeaders, error)
		// 依頼情報検索API
		GetTradeRequests(c echo.Context, request traceabilityentity.GetTradeRequestsRequest) (traceabilityentity.GetTradeRequestsResponse, error)
		// 依頼情報登録API
		PostTradeRequests(c echo.Context, request traceabilityentity.PostTradeRequestsRequest) (traceabilityentity.PostTradeRequestsResponses, common.ResponseHeaders, error)
		// 依頼取消登録API
		PostTradeRequestsCancel(c echo.Context, request traceabilityentity.PostTradeRequestsCancelRequest) (traceabilityentity.PostTradeRequestsCancelResponse, common.ResponseHeaders, error)
		// 依頼差戻登録API
		PostTradeRequestsReject(c echo.Context, request traceabilityentity.PostTradeRequestsRejectRequest) (traceabilityentity.PostTradeRequestsRejectResponse, common.ResponseHeaders, error)
		// 受領依頼情報検索API
		// CFP情報取得API
		GetCfp(c echo.Context, request traceabilityentity.GetCfpRequest) (traceabilityentity.GetCfpResponses, error)
		GetTradeRequestsReceived(c echo.Context, request traceabilityentity.GetTradeRequestsReceivedRequest) (traceabilityentity.GetTradeRequestsReceivedResponse, error)
		// CFP情報登録API
		PostCfp(c echo.Context, requests traceabilityentity.PostCfpRequest) (traceabilityentity.PostCfpResponses, common.ResponseHeaders, error)
		// CFP証明書情報検索API
		GetCfpCertifications(c echo.Context, request traceabilityentity.GetCfpCertificationsRequest) (traceabilityentity.GetCfpCertificationsResponse, error)
		// 部品情報紐づけ登録API
		PostTrades(c echo.Context, request traceabilityentity.PostTradesRequest) (traceabilityentity.PostTradesResponse, common.ResponseHeaders, error)
	}
)
