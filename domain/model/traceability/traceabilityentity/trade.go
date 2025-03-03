package traceabilityentity

import (
	"time"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetTradeRequestsRequest
// Summary: This is structure which defines GetTradeRequestsRequest.
// Service: Traceability
// Router: [GET] /tradeRequests
// Usage: input
type GetTradeRequestsRequest struct {
	OperatorID string  `json:"operatorId"`
	TraceID    *string `json:"traceId"`
	After      *string `json:"after"`
}

// GetTradeRequestsResponse
// Summary: This is structure which defines GetTradeRequestsResponse.
// Service: Traceability
// Router: [GET] /tradeRequests
// Usage: output
type GetTradeRequestsResponse struct {
	TradeRequests []GetTradeRequestsResponseTradeRequest `json:"tradeRequests"`
	Next          string                                 `json:"next"`
}

// GetTradeRequestsResponseTradeRequest
// Summary: This is structure which defines GetTradeRequestsResponseTradeRequest.
type GetTradeRequestsResponseTradeRequest struct {
	Request  GetTradeRequestsResponseRequest   `json:"request"`
	Trade    GetTradeRequestsResponseTrade     `json:"trade"`
	Response *GetTradeRequestsResponseResponse `json:"response"`
}

// GetTradeRequestsResponseRequest
// Summary: This is structure which defines GetTradeRequestsResponseRequest.
type GetTradeRequestsResponseRequest struct {
	RequestID                string  `json:"requestId"`
	RequestType              string  `json:"requestType"`
	RequestStatus            string  `json:"requestStatus"`
	RequestedToOperatorID    string  `json:"requestedToOperatorId"`
	RequestedAt              string  `json:"requestedAt"`
	RequestMessage           string  `json:"requestMessage"`
	ReplyMessage             *string `json:"replyMessage"`
	ResponseDueDate          *string `json:"responseDueDate"`
	CompletedCount           *int    `json:"completedCount"`
	CompletedCountModifiedAt *string `json:"completedCountModifiedAt"`
}

// GetTradeRequestsResponseTrade
// Summary: This is structure which defines GetTradeRequestsResponseTrade.
type GetTradeRequestsResponseTrade struct {
	TradeID               string                                  `json:"tradeId"`
	TreeStatus            string                                  `json:"treeStatus"`
	Downstream            GetTradeRequestsResponseTradeDownstream `json:"downstream"`
	TradeRelation         GetTradeRequestsResponseTradeRelation   `json:"tradeRelation"`
	TradesCount           *int                                    `json:"tradesCount"`
	TradesCountModifiedAt *string                                 `json:"tradesCountModifiedAt"`
}

// GetTradeRequestsResponseTradeDownstream
// Summary: This is structure which defines GetTradeRequestsResponseTradeDownstream.
type GetTradeRequestsResponseTradeDownstream struct {
	DownstreamAmountUnitName string `json:"downstreamAmountUnitName"`
}

// GetTradeRequestsResponseTradeRelation
// Summary: This is structure which defines GetTradeRequestsResponseTradeRelation.
type GetTradeRequestsResponseTradeRelation struct {
	UpstreamOperatorID string  `json:"upstreamOperatorId"`
	DownstreamTraceID  string  `json:"downstreamTraceId"`
	UpstreamTraceID    *string `json:"upstreamTraceId"`
}

// GetTradeRequestsResponseResponse
// Summary: This is structure which defines GetTradeRequestsResponseResponse.
type GetTradeRequestsResponseResponse struct {
	ResponseID                      string                                             `json:"responseId"`
	ResponseType                    string                                             `json:"responseType"`
	ResponsedAt                     string                                             `json:"responsedAt"`
	ResponsePreProcessingEmissions  *float64                                           `json:"responsePreProcessingEmissions"`
	ResponseMainProductionEmissions *float64                                           `json:"responseMainProductionEmissions"`
	EmissionsUnitName               string                                             `json:"emissionsUnitName"`
	CFPCertificationFileInfo        []GetTradeRequestsResponseCFPCertificationFileInfo `json:"cfpCertificationFileInfo"`
	ResponseDqr                     GetTradeRequestsResponseResponseDqr                `json:"responseDqr"`
}

// GetTradeRequestsResponseCFPCertificationFileInfo
// Summary: This is structure which defines GetTradeRequestsResponseCFPCertificationFileInfo.
type GetTradeRequestsResponseCFPCertificationFileInfo struct {
	FileID   string `json:"fileId"`
	FileName string `json:"fileName"`
}

// GetTradeRequestsResponseResponseDqr
// Summary: This is structure which defines GetTradeRequestsResponseResponseDqr.
type GetTradeRequestsResponseResponseDqr struct {
	PreProcessingTeR  *float64 `json:"preProcessingTeR"`
	PreProcessingGeR  *float64 `json:"preProcessingGeR"`
	PreProcessingTiR  *float64 `json:"preProcessingTiR"`
	MainProductionTeR *float64 `json:"mainProductionTeR"`
	MainProductionGeR *float64 `json:"mainProductionGeR"`
	MainProductionTiR *float64 `json:"mainProductionTiR"`
}

// GetNextPtr
// Summary: This is function which get next ID for paging.
// output: (*string) pointer next ID
func (r GetTradeRequestsResponse) GetNextPtr() *string {
	if r.Next == "" {
		return nil
	}
	return &r.Next
}

// ToTradeModels
// Summary: This is function which convert GetTradeRequestsResponse to TradeModels.
// input: c(echo.Context) echo Context
// output: (TradeModels) TradeModels object
// output: (error) error object
func (r GetTradeRequestsResponse) ToTradeModels(c echo.Context) (traceability.TradeModels, error) {
	var err error
	tradeModels := make(traceability.TradeModels, len(r.TradeRequests))
	for i, res := range r.TradeRequests {
		tradeModels[i], err = res.ToTradeModel(c)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}
	return tradeModels, nil
}

// ToCfpCertificationModels
// Summary: This is function which convert GetTradeRequestsResponse to TradeModels.
// input: c(echo.Context) echo Context
// output: (TradeModels) TradeModels object
// output: (error) error object
func (r GetTradeRequestsResponse) ToCfpCertificationModels(c echo.Context) (traceability.TradeModels, error) {
	var err error
	tradeModels := make(traceability.TradeModels, len(r.TradeRequests))
	for i, res := range r.TradeRequests {
		tradeModels[i], err = res.ToTradeModel(c)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}
	return tradeModels, nil
}

// ToTradeModel
// Summary: This is function which convert GetTradeRequestsResponseTradeRequest to TradeModel.
// input: c(echo.Context) echo Context
// output: (TradeModel) TradeModel object
// output: (error) error object
func (r GetTradeRequestsResponseTradeRequest) ToTradeModel(c echo.Context) (traceability.TradeModel, error) {
	tradeID, err := uuid.Parse(r.Trade.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	operatorID := c.Get("operatorID").(string)
	downstreamOperatorID, err := uuid.Parse(operatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	upstreamOperatorID, err := uuid.Parse(r.Request.RequestedToOperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	downstreamTraceID, err := uuid.Parse(r.Trade.TradeRelation.DownstreamTraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	var varupstreamTraceID *uuid.UUID
	if r.Trade.TradeRelation.UpstreamTraceID != nil && *r.Trade.TradeRelation.UpstreamTraceID != "" {
		upstreamTraceID, err := uuid.Parse(*r.Trade.TradeRelation.UpstreamTraceID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return traceability.TradeModel{}, err
		}
		varupstreamTraceID = &upstreamTraceID
	}

	m := traceability.TradeModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: downstreamOperatorID,
		UpstreamOperatorID:   upstreamOperatorID,
		DownstreamTraceID:    downstreamTraceID,
		UpstreamTraceID:      varupstreamTraceID,
	}

	return m, nil
}

// ToStatusModels
// Summary: This is function which convert GetTradeRequestsResponseTradeRequest to StatusModels.
// output: (StatusModels) StatusModels object
// output: (error) error object
func (r GetTradeRequestsResponse) ToStatusModels() (traceability.StatusModels, error) {
	var err error
	statusModels := make(traceability.StatusModels, len(r.TradeRequests))
	for i, res := range r.TradeRequests {
		statusModels[i], err = res.ToStatusModel()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}
	return statusModels, nil
}

// ToStatusModel
// Summary: This is function which convert GetTradeRequestsResponseTradeRequest to StatusModel.
// output: (StatusModel) StatusModel object
// output: (error) error object
func (r GetTradeRequestsResponseTradeRequest) ToStatusModel() (traceability.StatusModel, error) {
	statusID, err := uuid.Parse(r.Request.RequestID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, err
	}

	tradeID, err := uuid.Parse(r.Trade.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, err
	}

	cfpResponseStatus := traceability.CfpResponseStatus(r.Request.RequestStatus)
	tradeTreeStatus := traceability.TradeTreeStatus(r.Trade.TreeStatus)
	m := traceability.StatusModel{
		StatusID: statusID,
		TradeID:  tradeID,
		RequestStatus: traceability.RequestStatus{
			CfpResponseStatus:        &cfpResponseStatus,
			TradeTreeStatus:          &tradeTreeStatus,
			CompletedCount:           r.Request.CompletedCount,
			CompletedCountModifiedAt: r.Request.CompletedCountModifiedAt,
			TradesCount:              r.Trade.TradesCount,
			TradesCountModifiedAt:    r.Trade.TradesCountModifiedAt,
		},
		Message:         &r.Request.RequestMessage,
		ReplyMessage:    r.Request.ReplyMessage,
		RequestType:     r.Request.RequestType,
		ResponseDueDate: r.Request.ResponseDueDate,
	}

	return m, nil
}

// ToStatusModelsForSort
// Summary: This is function which convert GetTradeRequestsResponse to array of StatusModelForSort.
// output: (StatusModelForSort) array of StatusModelForSort
// output: (error) error object
func (r GetTradeRequestsResponse) ToStatusModelsForSort() ([]traceability.StatusModelForSort, error) {
	var err error
	statusModels := make([]traceability.StatusModelForSort, len(r.TradeRequests))
	for i, res := range r.TradeRequests {
		statusModels[i], err = res.ToStatusModelForSort()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}
	return statusModels, nil
}

// ToStatusModelForSort
// Summary: This is function which convert GetTradeRequestsResponseTradeRequest to StatusModelForSort.
// output: (StatusModelForSort) StatusModelForSort object
// output: (error) error object
func (r GetTradeRequestsResponseTradeRequest) ToStatusModelForSort() (traceability.StatusModelForSort, error) {
	statusID, err := uuid.Parse(r.Request.RequestID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}

	tradeID, err := uuid.Parse(r.Trade.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}

	cfpResponseStatus := traceability.CfpResponseStatus(r.Request.RequestStatus)
	tradeTreeStatus := traceability.TradeTreeStatus(r.Trade.TreeStatus)
	m := traceability.StatusModel{
		StatusID: statusID,
		TradeID:  tradeID,
		RequestStatus: traceability.RequestStatus{
			CfpResponseStatus:        &cfpResponseStatus,
			TradeTreeStatus:          &tradeTreeStatus,
			CompletedCount:           r.Request.CompletedCount,
			CompletedCountModifiedAt: r.Request.CompletedCountModifiedAt,
			TradesCount:              r.Trade.TradesCount,
			TradesCountModifiedAt:    r.Trade.TradesCountModifiedAt,
		},
		Message:         &r.Request.RequestMessage,
		ReplyMessage:    r.Request.ReplyMessage,
		RequestType:     r.Request.RequestType,
		ResponseDueDate: r.Request.ResponseDueDate,
	}
	// convert r.Request.RequestedAt to time.Time
	requestedAt, err := time.Parse(time.RFC3339, r.Request.RequestedAt)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}
	mForSort := traceability.StatusModelForSort{
		StatusModel: m,
		RequestedAt: requestedAt,
	}
	return mForSort, nil
}

// ToCfpModels
// Summary: This is function which convert GetTradeRequestsResponse to CfpModels.
// output: (CfpModels) CfpModels object
// output: (error) error object
func (r GetTradeRequestsResponse) ToCfpModels() (traceability.CfpModels, error) {
	allCfpModels := []traceability.CfpModel{}
	for _, res := range r.TradeRequests {

		// Skip if CFP information is not registered
		if res.Response == nil {
			continue
		}

		cfpModels, err := res.ToCfpModels()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
		allCfpModels = append(allCfpModels, cfpModels...)
	}
	return allCfpModels, nil
}

// ToCfpModels
// Summary: This is function which convert GetTradeRequestsResponseTradeRequest to array of CfpModels.
// output: (CfpModels) CfpModels object
// output: (error) error object
func (r GetTradeRequestsResponseTradeRequest) ToCfpModels() (traceability.CfpModels, error) {

	traceID, err := uuid.Parse(r.Trade.TradeRelation.DownstreamTraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return nil, err
	}

	kv := map[string]*float64{
		traceability.CfpTypePreProductionResponse.ToString():  r.Response.ResponsePreProcessingEmissions,
		traceability.CfpTypeMainProductionResponse.ToString(): r.Response.ResponseMainProductionEmissions,
	}

	cfpModels := traceability.CfpModels{}
	for cfpType, emission := range kv {
		ghgDeclaredUnit, err := traceability.NewGhgDeclaredUnit(r.Response.EmissionsUnitName)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
		cfpTypeEnum := traceability.CfpType(cfpType)
		dqrTypeEnum, err := cfpTypeEnum.ToDqrType()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}

		var dqr traceability.DqrValue
		if cfpTypeEnum == traceability.CfpTypePreProductionResponse {
			dqr = traceability.DqrValue{
				TeR: r.Response.ResponseDqr.PreProcessingTeR,
				GeR: r.Response.ResponseDqr.PreProcessingGeR,
				TiR: r.Response.ResponseDqr.PreProcessingTiR,
			}
		} else {
			dqr = traceability.DqrValue{
				TeR: r.Response.ResponseDqr.MainProductionTeR,
				GeR: r.Response.ResponseDqr.MainProductionGeR,
				TiR: r.Response.ResponseDqr.MainProductionTiR,
			}
		}
		m := traceability.CfpModel{
			CfpID:           nil,
			TraceID:         traceID,
			GhgEmission:     emission,
			GhgDeclaredUnit: ghgDeclaredUnit,
			CfpType:         cfpType,
			DqrType:         dqrTypeEnum.ToString(),
			DqrValue:        dqr,
		}
		cfpModels = append(cfpModels, m)
	}
	return cfpModels, nil
}

// GetTradeRequestsReceivedRequest
// Summary: This is structure which defines GetTradeRequestsReceivedRequest.
// Service: Traceability
// Router: [GET] /tradeRequestsReceived
// Usage: input
type GetTradeRequestsReceivedRequest struct {
	OperatorID        string     `json:"operatorId"`
	RequestID         *string    `json:"requestId"`
	RequestedDateFrom *time.Time `json:"requestedDateFrom" format:"2006-01-02"`
	RequestedDateTo   *time.Time `json:"requestedDateTo" format:"2006-01-02"`
	After             *string    `json:"after"`
}

// GetTradeRequestsReceivedResponse
// Summary: This is structure which defines GetTradeRequestsReceivedResponse.
// Service: Traceability
// Router: [GET] /tradeRequestsReceived
// Usage: output
type GetTradeRequestsReceivedResponse struct {
	TradeRequests []GetTradeRequestsReceivedResponseTradeRequest `json:"tradeRequests"`
	Next          string                                         `json:"next"`
}

// GetTradeRequestsReceivedResponseTradeRequest
// Summary: This is structure which defines GetTradeRequestsReceivedResponseTradeRequest.
type GetTradeRequestsReceivedResponseTradeRequest struct {
	Request GetTradeRequestsReceivedResponseRequest `json:"request"`
	Trade   GetTradeRequestsReceivedResponseTrade   `json:"trade"`
}

// GetTradeRequestsReceivedResponseTradeRequest
// Summary: This is structure which defines GetTradeRequestsReceivedResponseTradeRequest.
type GetTradeRequestsReceivedResponseRequest struct {
	RequestID                string  `json:"requestId"`
	RequestType              string  `json:"requestType"`
	RequestStatus            string  `json:"requestStatus"`
	RequestedFromOperatorID  string  `json:"requestedFromOperatorId"`
	RequestedAt              string  `json:"requestedAt"`
	RequestMessage           string  `json:"requestMessage"`
	ReplyMessage             *string `json:"replyMessage"`
	ResponseDueDate          *string `json:"responseDueDate"`
	CompletedCount           *int    `json:"completedCount"`
	CompletedCountModifiedAt *string `json:"completedCountModifiedAt"`
}

// GetTradeRequestsReceivedResponseTrade
// Summary: This is structure which defines GetTradeRequestsReceivedResponseTrade.
type GetTradeRequestsReceivedResponseTrade struct {
	TradeID               string                                          `json:"tradeId"`
	TradeRelation         GetTradeRequestsReceivedResponseTradeRelation   `json:"tradeRelation"`
	TreeStatus            string                                          `json:"treeStatus"`
	Downstream            GetTradeRequestsReceivedResponseTradeDownstream `json:"downstream"`
	TradesCount           *int                                            `json:"tradesCount"`
	TradesCountModifiedAt *string                                         `json:"tradesCountModifiedAt"`
}

// GetTradeRequestsReceivedResponseTradeRelation
// Summary: This is structure which defines GetTradeRequestsReceivedResponseTradeRelation.
type GetTradeRequestsReceivedResponseTradeRelation struct {
	DownstreamOperatorID string  `json:"downstreamOperatorId"`
	DownstreamTraceID    string  `json:"downstreamTraceId"`
	UpstreamTraceID      *string `json:"upstreamTraceId"`
}

// GetTradeRequestsReceivedResponseTradeDownstream
// Summary: This is structure which defines GetTradeRequestsReceivedResponseTradeDownstream.
type GetTradeRequestsReceivedResponseTradeDownstream struct {
	DownstreamPartsItem        string  `json:"downstreamPartsItem"`
	DownstreamSupportPartsItem string  `json:"downstreamSupportPartsItem"`
	DownstreamPlantID          string  `json:"downstreamPlantId"`
	DownstreamAmountUnitName   string  `json:"downstreamAmountUnitName"`
	DownstreamPartsLabelName   *string `json:"downstreamPartsLabelName"`
	DownstreamPartsAddInfo1    *string `json:"downstreamPartsAddInfo1"`
	DownstreamPartsAddInfo2    *string `json:"downstreamPartsAddInfo2"`
	DownstreamPartsAddInfo3    *string `json:"downstreamPartsAddInfo3"`
}

// GetNextPtr
// Summary: This is function which get next ID for paging.
// output: (*string) pointer next ID
func (r GetTradeRequestsReceivedResponse) GetNextPtr() *string {
	if r.Next == "" {
		return nil
	}
	return &r.Next
}

// ToStatusModel
// Summary: This is function which convert GetTradeRequestsReceivedResponse to StatusModels.
// output: (StatusModel) StatusModels object
// output: (error) error object
func (r GetTradeRequestsReceivedResponse) ToStatusModels() (traceability.StatusModels, error) {
	var err error
	res := make([]traceability.StatusModel, len(r.TradeRequests))
	for i, tr := range r.TradeRequests {
		res[i], err = tr.ToStatusModel()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}

	return res, nil
}

// ToStatusModelsForSort
// Summary: This is function which convert GetTradeRequestsReceivedResponse to array of StatusModelForSort.
// output: (StatusModelForSort) array of StatusModelForSort
// output: (error) error object
func (r GetTradeRequestsReceivedResponse) ToStatusModelsForSort() ([]traceability.StatusModelForSort, error) {
	var err error
	res := make([]traceability.StatusModelForSort, len(r.TradeRequests))
	for i, tr := range r.TradeRequests {
		res[i], err = tr.ToStatusModelForSort()
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}

	return res, nil
}

// ExtractModelByTradeID
// Summary: This is function which extract TradeModel from GetTradeRequestsReceivedResponse by tradeID.
// input: tradeID(uuid.UUID) value of tradeID
// output: (TradeModel) TradeModel object
// output: (error) error object
func (r GetTradeRequestsReceivedResponse) ExtractModelByTradeID(tradeID uuid.UUID) (traceability.TradeModel, error) {
	for _, tr := range r.TradeRequests {
		if tr.Trade.TradeID == tradeID.String() {
			return tr.ToModelWithUpstreamOperatorIDNil()
		}
	}

	return traceability.TradeModel{}, nil
}

// ToModelWithUpstreamOperatorIDNil
// Summary: This is function which convert GetTradeRequestsReceivedResponseTradeRequest to TradeModel.
// output: (TradeModel) TradeModel object
// output: (error) error object
func (r GetTradeRequestsReceivedResponseTradeRequest) ToModelWithUpstreamOperatorIDNil() (traceability.TradeModel, error) {
	var err error
	tradeID, err := uuid.Parse(r.Trade.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	downstreamOperatorID, err := uuid.Parse(r.Trade.TradeRelation.DownstreamOperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	downstreamTraceID, err := uuid.Parse(r.Trade.TradeRelation.DownstreamTraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	var varupstreamTraceID *uuid.UUID
	if r.Trade.TradeRelation.UpstreamTraceID != nil && *r.Trade.TradeRelation.UpstreamTraceID != "" {
		upstreamTraceID, err := uuid.Parse(*r.Trade.TradeRelation.UpstreamTraceID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return traceability.TradeModel{}, err
		}
		varupstreamTraceID = &upstreamTraceID
	}

	m := traceability.TradeModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: downstreamOperatorID,
		UpstreamOperatorID:   uuid.Nil,
		DownstreamTraceID:    downstreamTraceID,
		UpstreamTraceID:      varupstreamTraceID,
	}

	return m, nil
}

// ToTradeResponseModels
// Summary: This is function which convert GetTradeRequestsReceivedResponseTradeRequest to array of TradeResponseModel.
// input: upstreamOperatorID(uuid.UUID) value of upstreamOperatorID
// output: ([]traceability.TradeResponseModel) array of TradeResponseModel
// output: (error) error object
func (r GetTradeRequestsReceivedResponse) ToTradeResponseModels(upstreamOperatorID uuid.UUID) ([]traceability.TradeResponseModel, error) {
	var err error
	res := make([]traceability.TradeResponseModel, len(r.TradeRequests))
	for i, tr := range r.TradeRequests {
		res[i], err = tr.ToTradeResponseModel(upstreamOperatorID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}
	}

	return res, nil
}

// ToTradeResponseModel
// Summary: This is function which convert GetTradeRequestsReceivedResponseTradeRequest to TradeResponseModel.
// input: upstreamOperatorID(uuid.UUID) value of upstreamOperatorID
// output: (traceability.TradeResponseModel) TradeResponseModel object
// output: (error) error object
func (r GetTradeRequestsReceivedResponseTradeRequest) ToTradeResponseModel(upstreamOperatorID uuid.UUID) (traceability.TradeResponseModel, error) {
	trade, err := r.Trade.ToModel(upstreamOperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeResponseModel{}, err
	}
	status, err := r.ToStatusModel()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeResponseModel{}, err
	}
	parts, err := r.ToPartsModel()
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeResponseModel{}, err
	}
	m := traceability.TradeResponseModel{
		TradeModel:  trade,
		StatusModel: status,
		PartsModel:  parts,
	}
	return m, nil
}

// ToStatusModel
// Summary: This is function which convert GetTradeRequestsReceivedResponseTradeRequest to StatusModel.
// output: (StatusModel) StatusModel object
// output: (error) error object
func (r GetTradeRequestsReceivedResponseTradeRequest) ToStatusModel() (traceability.StatusModel, error) {
	statusID, err := uuid.Parse(r.Request.RequestID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, nil
	}
	tradeID, err := uuid.Parse(r.Trade.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModel{}, nil
	}

	cfpResponseStatus := traceability.CfpResponseStatus(r.Request.RequestStatus)
	tradeTreeStatus := traceability.TradeTreeStatus(r.Trade.TreeStatus)

	requestStatus := traceability.RequestStatus{
		CfpResponseStatus:        &cfpResponseStatus,
		TradeTreeStatus:          &tradeTreeStatus,
		CompletedCount:           r.Request.CompletedCount,
		CompletedCountModifiedAt: r.Request.CompletedCountModifiedAt,
		TradesCount:              r.Trade.TradesCount,
		TradesCountModifiedAt:    r.Trade.TradesCountModifiedAt,
	}
	requestMessage := r.Request.RequestMessage
	replyMessage := r.Request.ReplyMessage
	requestType := r.Request.RequestType
	responseDueDate := r.Request.ResponseDueDate

	// requestStatus
	m := traceability.StatusModel{
		StatusID:        statusID,
		TradeID:         tradeID,
		RequestStatus:   requestStatus,
		Message:         &requestMessage,
		ReplyMessage:    replyMessage,
		RequestType:     requestType,
		ResponseDueDate: responseDueDate,
	}
	return m, nil
}

// ToStatusModelForSort
// Summary: This is function which convert GetTradeRequestsReceivedResponseTradeRequest to StatusModel.
// output: (StatusModelForSort) StatusModelForSort object
// output: (error) error object
func (r GetTradeRequestsReceivedResponseTradeRequest) ToStatusModelForSort() (traceability.StatusModelForSort, error) {
	var err error
	statusID, err := uuid.Parse(r.Request.RequestID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}
	tradeID, err := uuid.Parse(r.Trade.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}
	cfpResponseStatus, err := traceability.NewCfpResponseStatus(r.Request.RequestStatus)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}
	tradeTreeStatus, err := traceability.NewTradeTreeStatus(r.Trade.TreeStatus)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}

	m := traceability.StatusModel{
		StatusID: statusID,
		TradeID:  tradeID,
		RequestStatus: traceability.RequestStatus{
			CfpResponseStatus:        &cfpResponseStatus,
			TradeTreeStatus:          &tradeTreeStatus,
			CompletedCount:           r.Request.CompletedCount,
			CompletedCountModifiedAt: r.Request.CompletedCountModifiedAt,
			TradesCount:              r.Trade.TradesCount,
			TradesCountModifiedAt:    r.Trade.TradesCountModifiedAt,
		},
		Message:         &r.Request.RequestMessage,
		ReplyMessage:    r.Request.ReplyMessage,
		RequestType:     r.Request.RequestType,
		ResponseDueDate: r.Request.ResponseDueDate,
	}

	requestedAt, err := time.Parse(time.RFC3339, r.Request.RequestedAt)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusModelForSort{}, err
	}

	mForSort := traceability.StatusModelForSort{
		StatusModel: m,
		RequestedAt: requestedAt,
	}
	return mForSort, nil
}

// ToPartsModel
// Summary: This is function which convert GetTradeRequestsReceivedResponseTradeRequest to PartsModel.
// output: (PartsModel) PartsModel object
// output: (error) error object
func (r GetTradeRequestsReceivedResponseTradeRequest) ToPartsModel() (traceability.PartsModel, error) {
	traceID, err := uuid.Parse(r.Trade.TradeRelation.DownstreamTraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	operatorID, err := uuid.Parse(r.Trade.TradeRelation.DownstreamOperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	plantID, err := uuid.Parse(r.Trade.Downstream.DownstreamPlantID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	partsName := r.Trade.Downstream.DownstreamPartsItem
	supportPartsName := r.Trade.Downstream.DownstreamSupportPartsItem
	// Non-terminal because there is always an upstream component
	terminatedFlag := false
	amountRequiredUnit, err := traceability.NewAmountRequiredUnit(r.Trade.Downstream.DownstreamAmountUnitName)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.PartsModel{}, err
	}
	partsLabelName := r.Trade.Downstream.DownstreamPartsLabelName
	partsAddInfo1 := r.Trade.Downstream.DownstreamPartsAddInfo1
	partsAddInfo2 := r.Trade.Downstream.DownstreamPartsAddInfo2
	partsAddInfo3 := r.Trade.Downstream.DownstreamPartsAddInfo3

	m := traceability.PartsModel{
		TraceID:            traceID,
		OperatorID:         operatorID,
		PlantID:            &plantID,
		PartsName:          partsName,
		SupportPartsName:   &supportPartsName,
		TerminatedFlag:     terminatedFlag,
		AmountRequired:     nil,
		AmountRequiredUnit: &amountRequiredUnit,
		PartsLabelName:     partsLabelName,
		PartsAddInfo1:      partsAddInfo1,
		PartsAddInfo2:      partsAddInfo2,
		PartsAddInfo3:      partsAddInfo3,
	}
	return m, nil

}

// ToModel
// Summary: This is function which convert GetTradeRequestsReceivedResponseTrade to PartsModel.
// input: upstreamOperatorID(uuid.UUID) value of upstreamOperatorID
// output: (TradeModel) TradeModel object
// output: (error) error object
func (r GetTradeRequestsReceivedResponseTrade) ToModel(upstreamOperatorID uuid.UUID) (traceability.TradeModel, error) {
	tradeID, err := uuid.Parse(r.TradeID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}
	downstreamOperatorID, err := uuid.Parse(r.TradeRelation.DownstreamOperatorID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}
	downstreamTraceID, err := uuid.Parse(r.TradeRelation.DownstreamTraceID)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeModel{}, err
	}

	var varupstreamTraceID *uuid.UUID
	if r.TradeRelation.UpstreamTraceID != nil && *r.TradeRelation.UpstreamTraceID != "" {
		upstreamTraceID, err := uuid.Parse(*r.TradeRelation.UpstreamTraceID)
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return traceability.TradeModel{}, err
		}
		varupstreamTraceID = &upstreamTraceID
	}

	m := traceability.TradeModel{
		TradeID:              &tradeID,
		DownstreamOperatorID: downstreamOperatorID,
		UpstreamOperatorID:   upstreamOperatorID,
		DownstreamTraceID:    downstreamTraceID,
		UpstreamTraceID:      varupstreamTraceID,
	}
	return m, nil
}

// PostTradeRequestsRequest
// Summary: This is structure which defines PostTradeRequestsRequest.
// Service: Traceability
// Router: [POST] /tradeRequests
// Usage: input
type PostTradeRequestsRequest struct {
	OperatorID    string                                 `json:"operatorId"`
	TradeRequests []PostTradeRequestsRequestTradeRequest `json:"tradeRequests"`
}

// PostTradeRequestsRequestTradeRequest
// Summary: This is structure which defines PostTradeRequestsRequestTradeRequest.
type PostTradeRequestsRequestTradeRequest struct {
	DownstreamTraceID  string  `json:"downstreamTraceId"`
	UpstreamOperatorID string  `json:"upstreamOperatorId"`
	RequestType        string  `json:"requestType"`
	RequestMessage     *string `json:"requestMessage"`
	ResponseDueDate    *string `json:"responseDueDate"`
}

// PostTradeRequestsResponse
// Summary: This is structure which defines PostTradeRequestsResponse.
type PostTradeRequestsResponse struct {
	TradeID           string `json:"tradeId"`
	RequestID         string `json:"requestId"`
	DownstreamTraceID string `json:"downstreamTraceId"`
}

// PostTradeRequestsResponses
// Summary: This is structure which defines PostTradeRequestsResponses.
// Service: Traceability
// Router: [POST] /tradeRequests
// Usage: output
type PostTradeRequestsResponses []PostTradeRequestsResponse

// NewPostTradeRequestRequestFromModel
// Summary: This is function which convert TradeRequestModel to PostTradeRequestsRequest.
// input: m(TradeRequestModel) TradeRequestModel object
// output: (PostTradeRequestsRequest) PostTradeRequestsRequest object
func NewPostTradeRequestRequestFromModel(m traceability.TradeRequestModel) PostTradeRequestsRequest {
	tradeRequests := []PostTradeRequestsRequestTradeRequest{
		{
			DownstreamTraceID:  m.TradeModel.DownstreamTraceID.String(),
			UpstreamOperatorID: m.TradeModel.UpstreamOperatorID.String(),
			RequestType:        m.StatusModel.RequestType,
			RequestMessage:     m.StatusModel.Message,
			ResponseDueDate:    m.StatusModel.ResponseDueDate,
		},
	}

	req := PostTradeRequestsRequest{
		OperatorID:    m.TradeModel.DownstreamOperatorID.String(),
		TradeRequests: tradeRequests,
	}

	return req
}

// PostTradesRequest
// Summary: This is structure which defines PostTradesRequest.
// Service: Traceability
// Router: [POST] /trades
// Usage: input
type PostTradesRequest struct {
	OperatorID string `json:"operatorId"`
	TradeID    string `json:"tradeId"`
	TraceID    string `json:"traceId"`
}

// PostTradesResponse
// Summary: This is structure which defines PostTradesResponse.
// Service: Traceability
// Router: [POST] /trades
// Usage: output
type PostTradesResponse struct {
	TradeID string `json:"tradeId"`
}

// ToCertificationModels
// Summary: This is function which convert GetTradeRequestsResponse to CfpCertificationModels.
// output: (CfpCertificationModels) CfpCertificationModels object
// output: (error) error object
func (r GetTradeRequestsResponse) ToCertificationModels() (traceability.CfpCertificationModels, error) {
	allCertificationModels := traceability.CfpCertificationModels{}
	for _, res := range r.TradeRequests {

		// Skip if not answered
		if res.Response == nil {
			continue
		}

		cfpCertificationModel := res.ToCertificationModel()
		allCertificationModels = append(allCertificationModels, cfpCertificationModel)
	}
	return allCertificationModels, nil
}

// ToCertificationModel
// Summary: This is function which convert GetTradeRequestsResponseTradeRequest to CfpCertificationModel.
// output: (CfpCertificationModel) CfpCertificationModel object
func (r GetTradeRequestsResponseTradeRequest) ToCertificationModel() traceability.CfpCertificationModel {
	return NewCfpCertificationModel(
		r.Trade.TradeID,
		r.Trade.TradeRelation.UpstreamTraceID,
		r.Trade.TradeRelation.UpstreamOperatorID,
		nil,
		r.Response.CFPCertificationFileInfo,
	)
}

// NewCfpCertificationModel
// Summary: This is function which create CfpCertificationModel.
// input: CfpCertificationID(string) value of *string
// input: TraceID(*string) pointer of TraceID
// input: OperatorID(string) value of OperatorID
// input: CfpCertificationDescription(*string) pointer of CfpCertificationDescription
// input: GetTradeRequestsResponseCFPCertificationFileInfo([]GetTradeRequestsResponseCFPCertificationFileInfo) array of GetTradeRequestsResponseCFPCertificationFileInfo
// output: (CfpCertificationModel) CfpCertificationModel object
func NewCfpCertificationModel(
	CfpCertificationID string,
	TraceID *string,
	OperatorID string,
	CfpCertificationDescription *string,
	GetTradeRequestsResponseCFPCertificationFileInfo []GetTradeRequestsResponseCFPCertificationFileInfo,
) traceability.CfpCertificationModel {
	es := []traceability.CfpCertificationFileInfo{}
	if len(GetTradeRequestsResponseCFPCertificationFileInfo) > 0 {
		for _, m := range GetTradeRequestsResponseCFPCertificationFileInfo {
			e := traceability.CfpCertificationFileInfo{
				OperatorID: OperatorID,
				FileID:     m.FileID,
				FileName:   m.FileName,
			}
			es = append(es, e)
		}
	}
	res := traceability.CfpCertificationModel{
		CfpCertificationID:          CfpCertificationID,
		CfpCertificationDescription: CfpCertificationDescription,
	}
	if len(es) > 0 {
		res.CfpCertificationFileInfo = &es
	}
	if TraceID != nil {
		res.TraceID = *TraceID
	}
	return res
}
