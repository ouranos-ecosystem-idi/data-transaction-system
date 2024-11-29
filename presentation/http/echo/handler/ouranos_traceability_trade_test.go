package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/presentation/http/echo/handler"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestProjectHandler_GetTradeRequest_Normal
// Summary: This is normal test class which confirm the operation of API #10 GetTradeRequest.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 2-1. 200: 正常系(Queryパラメータ追加なし)
// [x] 2-2. 200: 正常系(limit指定)
// [x] 2-3. 200: 正常系(after指定)
// [x] 2-4. 200: 正常系(traceIds指定、数は1)
// [x] 2-5. 200: 正常系(traceIds指定、数は2)
// [x] 2-6. 200: 正常系(traceIds指定、数は50)
// [x] 2-7. 200: 正常系(limit+after指定)
// [x] 2-8. 200: limitに値が設定されていない場合(limit=)
// [x] 2-9. 200: afterに値が設定されていない場合(after=)
// [x] 2-10. 200: traceIdsに値が設定されていない場合(traceIds=)
func TestProjectHandler_GetTradeRequest_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeRequest"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		after             string
		expectStatus      int
	}{
		{
			name:              "2-1. 200: 正常系(Queryパラメータ追加なし)",
			modifyQueryParams: func(q url.Values) {},
			expectStatus:      http.StatusOK,
		},
		{
			name: "2-2. 200: 正常系(limit指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "3")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-3. 200: 正常系(after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", f.TraceID)
			},
			after:        f.TraceID,
			expectStatus: http.StatusOK,
		},
		{
			name: "2-4. 200: 正常系(traceIds指定、数は1)",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceIds", f.TraceID)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-5. 200: 正常系(traceIds指定、数は2)",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceIds", common.GenerateUUIDString(2))
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-6. 200: 正常系(traceIds指定、数は50)",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceIds", common.GenerateUUIDString(50))
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-7. 200: 正常系(limit+after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "3")
				q.Set("after", f.TraceID)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-8. 200: limitに値が設定されていない場合(limit=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-9. 200: afterに値が設定されていない場合(after=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-10. 200: traceIdsに値が設定されていない場合(traceIds=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceIds", "")
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)
				test.modifyQueryParams(q)

				operatorUUID, _ := uuid.Parse(f.OperatorID)
				input := traceability.GetTradeRequestInput{
					OperatorID: operatorUUID,
				}
				limit := q.Get("limit")
				after := q.Get("after")

				// traceIDsを区切り文字で分割して配列に格納
				traceIDs := strings.Split(q.Get("traceIds"), ",")

				if limit == "" {
					input.Limit = 100
				} else {
					limit, err := strconv.Atoi(limit)
					if err != nil {
						fmt.Println("Conversion error:", err)
						return
					}
					input.Limit = limit
				}
				if after != "" {
					UUIDTraceID, _ := uuid.Parse(after)
					input.After = &UUIDTraceID
				}
				if len(traceIDs) > 0 {
					// UUIDの配列
					UUIDs := make([]uuid.UUID, len(traceIDs))

					// 文字列をUUIDに変換して配列に格納
					for i, str := range traceIDs {
						parsedUUID, err := uuid.Parse(str)
						if err != nil {
							fmt.Println("Error parsing UUID:", err)
							return
						}
						UUIDs[i] = parsedUUID
					}
					input.TraceIDs = UUIDs
				}
				var tradeRequestModel []traceability.TradeModel

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorID)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")
				tradeUsecase.On("GetTradeRequest", c, input).Return(tradeRequestModel, common.StringPtr(test.after), nil)

				err := tradeHandler.GetTradeRequest(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					tradeUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// TestProjectHandler_GetTradeRequest
// Summary: This is abnormal test class which confirm the operation of API #10 GetTradeRequest.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 1-1. 400: バリデーションエラー：limitの値が不正の場合
// [x] 1-2. 400: バリデーションエラー：limitが101以上の場合
// [x] 1-3. 400: バリデーションエラー：afterの値が不正の場合
// [x] 1-4. 400: バリデーションエラー：traceIdsの値が不正の場合
// [x] 1-5. 400: バリデーションエラー：limitが0の場合
// [x] 1-6. 400: バリデーションエラー：traceIdsが51件以上の場合
// [x] 1-7. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-8. 500: システムエラー：取得処理エラー
// [x] 1-9. 500: システムエラー：取得処理エラー
func TestProjectHandler_GetTradeRequest(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeRequest"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		modifyContexts    func(c echo.Context)
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：limitの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "three") // 数値変換できない
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：limitが101以上の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "101")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：afterの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", f.InvalidUUID) // UUID形式でない
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, after: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：traceIdsの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceIds", f.InvalidUUID) // UUID形式でない
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：limitが0の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "0")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 400: バリデーションエラー：traceIdsが51件以上の場合",
			modifyQueryParams: func(q url.Values) {
				traceIDs := make([]uuid.UUID, 51)
				for i := range traceIDs {
					traceIDs[i] = uuid.New()
				}
				traceIDsStr := common.JoinUUIDs(traceIDs, ",")
				q.Set("traceIds", traceIDsStr)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-7. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalid")
			},
			expectError:  "code=400, message={[auth] BadRequest Invalid or expired token",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-8. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-9. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-10. 400: バリデーションエラー：traceIdsにUUIDとそうでない値が混在する場合",
			modifyQueryParams: func(q url.Values) {
				traceIDs := make([]uuid.UUID, 10)
				for i := range traceIDs {
					traceIDs[i] = uuid.New()
				}
				traceIDsStr := common.JoinUUIDs(traceIDs, ",")
				traceIDsStr = traceIDsStr + f.InvalidUUID
				q.Set("traceIds", traceIDsStr)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				test.modifyContexts(c)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeUsecase.On("GetTradeRequest", mock.Anything, mock.Anything).Return([]traceability.TradeModel{}, common.StringPtr(""), test.receive)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")

				err := tradeHandler.GetTradeRequest(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// TestProjectHandler_PutTradeRequest_Normal
// Summary: This is normal test class which confirm the operation of API #7 PutTradeRequest.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 2-1.  201: 正常系(新規作成)
// [x] 2-2.  201: 正常系(更新)
func TestProjectHandler_PutTradeRequest_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeRequest"

	tests := []struct {
		name         string
		inputFunc    func() traceability.PutTradeRequestInput
		expectStatus int
	}{
		{
			name: "2-1. 201: 正常系(新規作成)",
			inputFunc: func() traceability.PutTradeRequestInput {
				input := f.NewPutTradeRequestInput()
				input.Trade.TradeID = nil
				input.Status.StatusID = nil
				input.Status.TradeID = nil
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-2. 201: 正常系(更新)",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				return i
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				input := test.inputFunc()
				tradeRequestModel := input.ToModel()

				inputJSON, _ := json.Marshal(input)
				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), bytes.NewBuffer(inputJSON))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorID)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")
				tradeUsecase.On("PutTradeRequest", c, input).Return(tradeRequestModel, common.ResponseHeaders{}, nil)

				err := tradeHandler.PutTradeRequest(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					tradeUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// TestProjectHandler_PutTradeRequest
// Summary: This is abnormal test class which confirm the operation of API #7 PutTradeRequest.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 1-1.  400: バリデーションエラー：tradeModelのtradeIdがUUID形式以外の場合
// [x] 1-2.  400: バリデーションエラー：tradeModelのdownstreamOperatorIdが未指定の場合
// [x] 1-3.  400: バリデーションエラー：tradeModelのdownstreamOperatorIdがUUID形式以外の場合
// [x] 1-4.  400: バリデーションエラー：tradeModelのupstreamOperatorIdが未指定の場合
// [x] 1-5.  400: バリデーションエラー：tradeModelのupstreamOperatorIdがUUID形式以外の場合
// [x] 1-6.  400: バリデーションエラー：tradeModelのdownstreamTraceIdが未指定の場合
// [x] 1-7.  400: バリデーションエラー：tradeModelのdownstreamTraceIdがUUID形式以外の場合
// [x] 1-8.  400: バリデーションエラー：statusModelのstatusIdがUUID形式以外の場合
// [x] 1-9.  400: バリデーションエラー：statusModelのtradeIdがUUID形式以外の場合
// [x] 1-10. 400: バリデーションエラー：statusModelのrequestTypeが未指定の場合
// [x] 1-11. 400: バリデーションエラー：statusModelのrequestTypeが指定された値以外の場合
// [x] 1-12. 400: バリデーションエラー：statusModelのmessageが101文字以上の場合
// [x] 1-13. 400: バリデーションエラー：tradeModelのtradeIdとstatusModelのtradeIdが一致しない場合
// [x] 1-14. 400: バリデーションエラー：tradeModelのtradeIdが設定されているかつstatusModelのtradeIdとstatusModelのstatusIdが設定されていない場合
// [x] 1-15. 400: バリデーションエラー：statusModelのtradeIdが設定されているかつtradeModelのtradeIdとstatusModelのstatusIdが設定されていない場合
// [x] 1-16. 400: バリデーションエラー：statusModelのstatusIdが設定されているかつtradeModelのtradeIdとstatusModelのtradeIdが設定されていない場合
// [x] 1-17. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-18. 500: システムエラー：更新処理エラー
// [x] 1-19. 500: システムエラー：更新処理エラー
// [x] 1-20. 400: バリデーションエラー：tradeModelの1-3と1-5が同時に発生する場合
// [x] 1-21. 400: バリデーションエラー：statusModelの1-8と1-9が同時に発生する場合
// [x] 1-22. 400: バリデーションエラー：tradeModelの1-3と1-5とstatusModelの1-8と1-9が同時に発生する場合
// [x] 1-23. 400: バリデーションエラー：1-3と1-13が同時に発生する場合
// [x] 1-24. 400: バリデーションエラー：1-3と1-14が同時に発生する場合
// [x] 1-25. 400: バリデーションエラー：1-13と1-14が同時に発生する場合
func TestProjectHandler_PutTradeRequest(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeRequest"

	tests := []struct {
		name           string
		inputFunc      func() traceability.PutTradeRequestInput
		modifyContexts func(c echo.Context)
		receive        error
		expectError    string
		expectStatus   int
	}{
		{
			name: "1-1. 400: バリデーションエラー：tradeModelのtradeIdがUUID形式以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.TradeID = common.StringPtr(f.InvalidUUID)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：tradeModelのdownstreamOperatorIdが未指定の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamOperatorID = ""
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamOperatorId: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：tradeModelのdownstreamOperatorIdがUUID形式以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamOperatorID = f.InvalidUUID
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamOperatorId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：tradeModelのupstreamOperatorIdが未指定の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.UpstreamOperatorID = ""
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, upstreamOperatorId: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：tradeModelのupstreamOperatorIdがUUID形式以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.UpstreamOperatorID = f.InvalidUUID
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, upstreamOperatorId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 400: バリデーションエラー：tradeModelのdownstreamTraceIdが未指定の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamTraceID = ""
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamTraceId: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-7. 400: バリデーションエラー：tradeModelのdownstreamTraceIdがUUID形式以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamTraceID = f.InvalidUUID
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamTraceId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-8. 400: バリデーションエラー：statusModelのstatusIdがUUID形式以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.StatusID = common.StringPtr(f.InvalidUUID)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, statusId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-9. 400: バリデーションエラー：statusModelのtradeIdがUUID形式以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.TradeID = common.StringPtr(f.InvalidUUID)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-10. 400: バリデーションエラー：statusModelのrequestTypeが未指定の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.RequestType = ""
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, requestType: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-11. 400: バリデーションエラー：statusModelのrequestTypeが指定された値以外の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.RequestType = "not request type"
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, requestType: cannot be allowed 'not request type'.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-12. 400: バリデーションエラー：statusModelのmessageが1001文字以上の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.Message = common.StringPtr(f.CraeteRamdomString(1001))
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, message: the length must be no more than 1000.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-13. 400: バリデーションエラー：tradeModelのtradeIdとstatusModelのtradeIdが一致しない場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.TradeID = common.StringPtr(f.TradeIDUUID1)
				i.Status.TradeID = common.StringPtr(f.TradeIDUUID2)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeModel.tradeId and statusModel.tradeId must be equal",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-14. 400: バリデーションエラー：tradeModelのtradeIdが設定されているかつstatusModelのtradeIdとstatusModelのstatusIdが設定されていない場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.TradeID = common.StringPtr(f.TradeIDUUID1)
				i.Status.TradeID = nil
				i.Status.StatusID = nil
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeModel.tradeId, statusModel.statusId, and statusModel.tradeId must all have values or all be null",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-15. 400: バリデーションエラー：statusModelのtradeIdが設定されているかつtradeModelのtradeIdとstatusModelのstatusIdが設定されていない場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.TradeID = nil
				i.Status.TradeID = common.StringPtr(f.TradeIDUUID1)
				i.Status.StatusID = nil
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeModel.tradeId, statusModel.statusId, and statusModel.tradeId must all have values or all be null",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-16. 400: バリデーションエラー：statusModelのstatusIdが設定されているかつtradeModelのtradeIdとstatusModelのtradeIdが設定されていない場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.TradeID = nil
				i.Status.TradeID = nil
				i.Status.StatusID = common.StringPtr(f.TradeIDUUID1)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeModel.tradeId, statusModel.statusId, and statusModel.tradeId must all have values or all be null",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-17. 400: バリデーションエラー：operatorIDがjwtのoperatorIdと一致しない場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "e03cc699-7234-31ed-86be-cc18c92208e6")
			},
			expectError:  "code=403, message={[dataspace] AccessDenied You do not have the necessary privileges",
			expectStatus: http.StatusForbidden,
		},
		{
			name: "1-18. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-19. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-20. 400: バリデーションエラー：tradeModelの1-3と1-5が同時に発生する場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamOperatorID = f.InvalidUUID
				i.Trade.UpstreamOperatorID = f.InvalidUUID
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamOperatorId: invalid UUID; upstreamOperatorId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-21. 400: バリデーションエラー：statusModelの1-8と1-9が同時に発生する場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.StatusID = common.StringPtr(f.InvalidUUID)
				i.Status.TradeID = common.StringPtr(f.InvalidUUID)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, statusId: invalid UUID; tradeId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-22. 400: バリデーションエラー：tradeModelの1-3と1-5とstatusModelの1-8と1-9が同時に発生する場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamOperatorID = f.InvalidUUID
				i.Trade.UpstreamOperatorID = f.InvalidUUID
				i.Status.StatusID = common.StringPtr(f.InvalidUUID)
				i.Status.TradeID = common.StringPtr(f.InvalidUUID)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamOperatorId: invalid UUID; upstreamOperatorId: invalid UUID.; statusId: invalid UUID; tradeId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-23. 400: バリデーションエラー：tradeModelの1-3と1-13が同時に発生する場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamOperatorID = f.InvalidUUID
				i.Trade.TradeID = common.StringPtr(f.TraceID)
				i.Status.TradeID = common.StringPtr(f.TraceID2)
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamOperatorId: invalid UUID.; tradeModel.tradeId and statusModel.tradeId must be equal.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-24. 400: バリデーションエラー：tradeModelの1-3と1-14が同時に発生する場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.DownstreamOperatorID = f.InvalidUUID
				i.Trade.TradeID = common.StringPtr(f.TradeIDUUID1)
				i.Status.TradeID = nil
				i.Status.StatusID = nil
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, downstreamOperatorId: invalid UUID.; tradeModel.tradeId, statusModel.statusId, and statusModel.tradeId must all have values or all be null",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-26. 400: バリデーションエラー：statusModelのresponseDueDateが未指定の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.ResponseDueDate = ""
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, responseDueDate: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-27. 400: バリデーションエラー：statusModelのresponseDueDateが不正な日付の場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Status.ResponseDueDate = "2024-02-30"
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, responseDueDate: must be a valid date.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-25. 400: バリデーションエラー：tradeModelの1-13と1-14が同時に発生する場合",
			inputFunc: func() traceability.PutTradeRequestInput {
				i := f.NewPutTradeRequestInput()
				i.Trade.TradeID = common.StringPtr(f.TradeIDUUID1)
				i.Status.TradeID = common.StringPtr(f.TraceID2)
				i.Status.StatusID = nil
				return i
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, tradeModel.tradeId, statusModel.statusId, and statusModel.tradeId must all have values or all be null; tradeModel.tradeId and statusModel.tradeId must be equal.",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				input := test.inputFunc()
				tradeRequestModel := input.ToModel()

				inputJSON, _ := json.Marshal(input)
				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), bytes.NewBuffer(inputJSON))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				test.modifyContexts(c)
				tradeUsecase := new(mocks.ITradeUsecase)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")
				tradeUsecase.On("PutTradeRequest", c, input).Return(tradeRequestModel, common.ResponseHeaders{}, test.receive)

				err := tradeHandler.PutTradeRequest(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// TestProjectHandler_GetTradeResponse_Normal
// Summary: This is abnormal test class which confirm the operation of API #12 GetTradeResponse.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 2-1. 200: 正常系(Queryパラメータ追加なし)
// [x] 2-2. 200: 正常系(limit指定)
// [x] 2-3. 200: 正常系(after指定)
// [x] 2-4. 200: 正常系(limit+after指定)
// [x] 2-5. 200: limitに値が設定されていない場合(limit=)
// [x] 2-6. 200: afterに値が設定されていない場合(after=)
func TestProjectHandler_GetTradeResponse_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeResponse"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		after             string
		expectStatus      int
	}{
		{
			name:              "2-1. 200: 正常系(Queryパラメータ追加なし)",
			modifyQueryParams: func(q url.Values) {},
			expectStatus:      http.StatusOK,
		},
		{
			name: "2-2. 200: 正常系(limit指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "3")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-3. 200: 正常系(after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", f.TraceID)
			},
			after:        f.TraceID,
			expectStatus: http.StatusOK,
		},
		{
			name: "2-4. 200: 正常系(limit+after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "3")
				q.Set("after", f.TraceID)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-5. 200: limitに値が設定されていない場合(limit=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-6. 200: afterに値が設定されていない場合(after=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", "")
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)
				test.modifyQueryParams(q)

				operatorUUID, _ := uuid.Parse(f.OperatorID)
				input := traceability.GetTradeResponseInput{
					OperatorID: operatorUUID,
				}
				var tradeResponseModel []traceability.TradeResponseModel
				limit := q.Get("limit")
				after := q.Get("after")

				if limit == "" {
					input.Limit = 100
				} else {
					limit, err := strconv.Atoi(limit)
					if err != nil {
						fmt.Println("Conversion error:", err)
						return
					}
					input.Limit = limit
				}
				if after != "" {
					uuidTraceID, _ := uuid.Parse(after)
					input.After = &uuidTraceID
				}

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorID)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")
				tradeUsecase.On("GetTradeResponse", c, input).Return(tradeResponseModel, common.StringPtr(test.after), nil)

				err := tradeHandler.GetTradeResponse(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					tradeUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// TestProjectHandler_GetTradeResponse
// Summary: This is abnormal test class which confirm the operation of API #12 GetTradeResponse.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 1-1. 400: バリデーションエラー：limitの値が不正の場合
// [x] 1-2. 400: バリデーションエラー：limitが101以上の場合
// [x] 1-3. 400: バリデーションエラー：limitの値が0の場合
// [x] 1-4. 400: バリデーションエラー：afterの値が不正の場合
// [x] 1-5. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-6. 500: システムエラー：取得処理エラー
// [x] 1-7. 500: システムエラー：取得処理エラー
func TestProjectHandler_GetTradeResponse(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeResponse"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		modifyContexts    func(c echo.Context)
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：limitの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "three") // 数値変換できない
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：limitが101以上の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "101")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：limitの値が0の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "0")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：afterの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", f.InvalidUUID) // UUID形式でない
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, after: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalid")
			},
			expectError:  "code=400, message={[auth] BadRequest Invalid or expired token",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-7. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				q.Set("dataTarget", dataTarget)
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				test.modifyContexts(c)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")
				tradeUsecase.On("GetTradeResponse", mock.Anything, mock.Anything).Return([]traceability.TradeResponseModel{}, common.StringPtr(""), test.receive)

				err := tradeHandler.GetTradeResponse(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					fmt.Println(err)
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// TestProjectHandler_PutTradeResponse_Normal
// Summary: This is normal test class which confirm the operation of API #13 PutTradeResponse.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 2-1. 201: 正常系(更新)
func TestProjectHandler_PutTradeResponse_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeResponse"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		expectStatus      int
	}{
		{
			name: "2-1. 201: 正常系(更新)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("tradeId", f.TradeID)
				q.Set("traceId", f.TraceID)
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				test.modifyQueryParams(q)
				input := f.NewPutTradeResponseInput()
				tradeModel := traceability.TradeModel{}

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorID)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")
				tradeUsecase.On("PutTradeResponse", c, input).Return(tradeModel, common.ResponseHeaders{}, nil)

				err := tradeHandler.PutTradeResponse(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					tradeUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// TestProjectHandler_PutTradeResponse
// Summary: This is abnormal test class which confirm the operation of API #13 PutTradeResponse.
// Target: ouranos_traceability_trade.go
// TestPattern:
// [x] 2-1. 201: 正常系(更新)
// [x] 1-1. 400: バリデーションエラー：traceIdが未指定の場合
// [x] 1-2. 400: バリデーションエラー：traceIdがUUID形式以外の場合
// [x] 1-3. 400: バリデーションエラー：tradeIdが未指定の場合
// [x] 1-4. 400: バリデーションエラー：tradeIdがUUID形式以外の場合
// [x] 1-5. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-6. 500: システムエラー：更新処理エラー
// [x] 1-7. 500: システムエラー：更新処理エラー
// [x] 1-8. 400: バリデーションエラー：traceIdに値が設定されていない場合(after=)
// [x] 1-9. 400: バリデーションエラー：tradeIdに値が設定されていない場合(after=)
func TestProjectHandler_PutTradeResponse(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "tradeResponse"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		modifyContexts    func(c echo.Context)
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdが未指定の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("tradeId", f.TradeID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdがUUID形式以外の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", "not uuid")
				q.Set("tradeId", f.TradeID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：tradeIdが未指定の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, tradeId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：tradeIdがUUID形式以外の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
				q.Set("tradeId", "not uuid")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, tradeId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalid")
			},
			expectError:  "code=400, message={[auth] BadRequest Invalid or expired token",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 500: システムエラー：更新処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("tradeId", f.TradeID)
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-7. 500: システムエラー：更新処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("tradeId", f.TradeID)
				q.Set("traceId", f.TraceID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-8. 400: バリデーションエラー：traceIdに値が設定されていない場合(after=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", "")
				q.Set("tradeId", f.TradeID)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-9. 400: バリデーションエラー：tradeIdに値が設定されていない場合(after=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
				q.Set("tradeId", "")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, tradeId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				test.modifyContexts(c)

				tradeUsecase := new(mocks.ITradeUsecase)
				tradeUsecase.On("PutTradeResponse", mock.Anything, mock.Anything).Return(traceability.TradeModel{}, common.ResponseHeaders{}, test.receive)
				tradeHandler := handler.NewTradeHandler(tradeUsecase, "")

				err := tradeHandler.PutTradeResponse(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}
