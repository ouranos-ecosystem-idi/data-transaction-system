package handler_test

import (
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

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 200: 正常系(Queryパラメータ追加なし)
// [x] 2-2. 200: 正常系(limit指定)
// [x] 2-3. 200: 正常系(after指定)
// [x] 2-4. 200: 正常系(statusTarget指定)
// [x] 2-5. 200: 正常系(statusId指定)
// [x] 2-7. 200: 正常系(limit+after指定)
// [x] 2-8. 200: 正常系(statusTarget=REQUEST+traceId指定)
// [x] 2-9. 200: 正常系(limit+after+statusId指定)
// [x] 2-10. 200: limitに値が設定されていない場合(limit=)
// [x] 2-11. 200: afterに値が設定されていない場合(after=)
// [x] 2-12. 200: statusIdに値が設定されていない場合(statusId=)
// [x] 2-13. 200: 2-9,2-10が同時に発生する場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetStatus_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "status"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		after             *string
		expectStatus      int
	}{
		{
			name: "2-1. 200: 正常系(Queryパラメータ追加なし)",
			modifyQueryParams: func(q url.Values) {
			},
			expectStatus: http.StatusOK,
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
				q.Set("after", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-4. 200: 正常系(statusTarget指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("statusTarget", "REQUEST")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-5. 200: 正常系(statusId指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("statusId", f.StatusId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-7. 200: 正常系(limit+after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "3")
				q.Set("after", f.TraceId)
			},
			after:        common.StringPtr(f.TraceId),
			expectStatus: http.StatusOK,
		},
		{
			name: "2-8. 200: 正常系(statusTarget=REQUEST+traceId指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("statusTarget", "REQUEST")
				q.Set("traceId", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-9. 200: 正常系(limit+after+statusTarget=REQUEST指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "3")
				q.Set("after", f.TraceId)
				q.Set("statusTarget", "REQUEST")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-10. 200: limitに値が設定されていない場合(limit=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-11. 00: afterに値が設定されていない場合(after=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("after", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-12. 200: statusIdに値が設定されていない場合(statusId=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("statusId", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-13. 200: 2-9,2-10が同時に発生する場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("limit", "")
				q.Set("after", "")
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			q := make(url.Values)
			test.modifyQueryParams(q)
			q.Set("dataTarget", dataTarget)
			operatorUUID, _ := uuid.Parse(f.OperatorId)
			input := traceability.GetStatusModel{
				OperatorID: operatorUUID,
			}
			statusModel := []traceability.StatusModel{}

			limit := q.Get("limit")
			after := q.Get("after")
			statusTarget := q.Get("statusTarget")
			statusID := q.Get("statusId")
			traceID := q.Get("traceId")

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
				after, _ := uuid.Parse(f.TraceId)
				input.After = &after
			}
			if statusTarget != "" {
				input.StatusTarget = "REQUEST"
			}
			if statusID != "" {
				statusId, _ := uuid.Parse(f.StatusId)
				input.StatusID = &statusId
			}
			if traceID != "" {
				traceId, _ := uuid.Parse(f.TraceId)
				input.TraceID = &traceId
			}

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			statusUsecase := new(mocks.IStatusUsecase)
			statusHandler := handler.NewStatusHandler(statusUsecase, "")
			statusUsecase.On("GetStatus", c, input).Return(statusModel, test.after, nil)

			err := statusHandler.GetStatus(c)
			// エラーが発生しないことを確認
			if assert.NoError(t, err) {
				// ステータスコードが期待通りであることを確認
				assert.Equal(t, test.expectStatus, rec.Code)
				// モックの呼び出しが期待通りであることを確認
				statusUsecase.AssertExpectations(t)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/status テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：limitの値が不正の場合
// [x] 1-2. 400: バリデーションエラー：limitが101以上の場合
// [x] 1-3. 400: バリデーションエラー：afterの値が不正の場合
// [x] 1-4. 400: バリデーションエラー：statusIdの値が不正の場合
// [x] 1-5. 400: バリデーションエラー：limitが0以下の場合
// [x] 1-6. 400: バリデーションエラー：traceIdの値が不正の場合
// [x] 1-7. 400: バリデーションエラー：statusTargetの値が不正の場合
// [x] 1-8. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-9. 500: システムエラー：取得処理エラー
// [x] 1-10. 500: システムエラー：取得処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetStatus_Abnormal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "status"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		modifyContexts    func(c echo.Context)
		receive           error
		expectError       string
	}{
		{
			name: "1-1. 400: バリデーションエラー：limitの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "three") // 数値変換できない文字列
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
		},
		{
			name: "1-2. 400: バリデーションエラー：limitが101以上の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "101")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
		},
		{
			name: "1-3. 400: バリデーションエラー：afterの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("after", f.InvalidUUID) // UUID変換できない文字列
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, after: Unexpected query parameter",
		},
		{
			name: "1-4. 400: バリデーションエラー：statusIdの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("statusId", f.InvalidUUID) // UUID変換できない文字列
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, statusId: Unexpected query parameter",
		},
		{
			name: "1-5. 400: バリデーションエラー：limitが0以下の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "0")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
		},
		{
			name: "1-6. 400: バリデーションエラー：traceIdの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.InvalidUUID) // UUID変換できない文字列
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
		},
		{
			name: "1-7. 400: バリデーションエラー：statusTargetの値が不正の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("statusTarget", f.InvalidEnum) // enumに存在しない文字列
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, statusTarget: Unexpected query parameter",
		},
		{
			name: "1-8. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("statusTarget", "REQUEST")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalidValue")
			},
			expectError: "code=400, message={[auth] BadRequest Invalid or expired token",
		},
		{
			name: "1-9. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("statusTarget", "REQUEST")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-10. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("statusTarget", "REQUEST")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:     fmt.Errorf("Internal Server Error"),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
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

			statusUsecase := new(mocks.IStatusUsecase)
			statusUsecase.On("GetStatus", mock.Anything, mock.Anything).Return([]traceability.StatusModel{}, common.StringPtr(""), test.receive)

			statusHandler := handler.NewStatusHandler(statusUsecase, "")

			err := statusHandler.GetStatus(c)
			// エラーが返ってくることを確認
			if assert.Error(t, err) {
				// エラーメッセージが期待通りであることを確認
				assert.ErrorContains(t, err, test.expectError)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/status 正常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常系：依頼取消
// [x] 1-2. 200: 正常系：依頼差戻：メッセージあり
// [x] 1-3. 200: 正常系：依頼差戻：メッセージなし
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutStatus_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "status"

	tests := []struct {
		name         string
		modifyInput  func(i *traceability.PutStatusInput) string
		expectStatus int
	}{
		{
			name: "1-1. 200: 正常系：依頼取消",
			modifyInput: func(i *traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusCancel
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "1-2. 200: 正常系：依頼差戻：メッセージあり",
			modifyInput: func(i *traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusReject
				i.ReplyMessage = "差戻メッセージ"
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "1-3. 200: 正常系：依頼差戻：メッセージなし",
			modifyInput: func(i *traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusReject
				i.ReplyMessage = ""
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			input := f.NewPutStatusInput()
			inputJSONStr := test.modifyInput(&input)

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(inputJSONStr))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			statusUsecase := new(mocks.IStatusUsecase)
			statusModel, _ := input.ToModel()
			cfpRequestStatus := statusModel.RequestStatus.CfpResponseStatus
			if cfpRequestStatus == traceability.CfpResponseStatusCancel {
				statusUsecase.On("PutStatusCancel", c, statusModel).Return(nil)
			} else {
				statusUsecase.On("PutStatusReject", c, statusModel).Return(nil)
			}
			statusHandler := handler.NewStatusHandler(statusUsecase, "")

			err := statusHandler.PutStatus(c)
			// エラーが発生しないことを確認
			if assert.NoError(t, err) {
				// ステータスコードが期待通りであることを確認
				assert.Equal(t, test.expectStatus, rec.Code)
				// モックの呼び出しが期待通りであることを確認
				statusUsecase.AssertExpectations(t)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/status 異常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：statusIdが含まれない場合
// [x] 1-2. 400: バリデーションエラー：statusIdの値が不正の場合
// [x] 1-3. 400: バリデーションエラー：tradeIdが含まれない場合
// [x] 1-4. 400: バリデーションエラー：tradeIdの値が不正の場合
// ケース削除：[x] 1-5. 400: バリデーションエラー：requestTypeが含まれない場合
// ケース削除：[x] 1-6. 400: バリデーションエラー：requestTypeの値が不正の場合
// [x] 1-7. 400: バリデーションエラー：replyMessageがstring型でない場合
// [x] 1-8. 400: バリデーションエラー：replyMessageが101文字以上の場合
// [x] 1-9. 400: バリデーションエラー：requestStatusのcfpResponseStatusが不正の場合
// ケース削除：[x] 1-10. 400: バリデーションエラー：requestStatusのtradeTreeStatusが不正の場合
// [x] 1-11. 400: バリデーションエラー：1-1と1-3が同時に発生した場合
// [x] 1-12. 400: バリデーションエラー：1-3と1-9が同時に発生した場合
// [x] 1-13. 500: システムエラー：更新処理エラー
// [x] 1-14. 500: システムエラー：更新処理エラー
// [x] 1-15. 500: システムエラー：更新処理エラー
// [x] 1-16. 500: システムエラー：更新処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutStatus_Abnormal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	dataTarget := "status"

	tests := []struct {
		name        string
		modifyInput func(i traceability.PutStatusInput) string
		receive     error
		expectError string
	}{
		{
			name: "1-1. 400: バリデーションエラー：statusIdが含まれない場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.StatusID = nil
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, statusId: cannot be blank.",
		},
		{
			name: "1-2. 400: バリデーションエラー：statusIdの値が不正の場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.StatusID = common.StringPtr(f.InvalidUUID)
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, statusId: invalid UUID.",
		},
		{
			name: "1-3. 400: バリデーションエラー：tradeIdが含まれない場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.TradeID = nil
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, tradeId: cannot be blank.",
		},
		{
			name: "1-4. 400: バリデーションエラー：tradeIdの値が不正の場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.TradeID = common.StringPtr(f.InvalidUUID)
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, tradeId: invalid UUID.",
		},
		{
			name: "1-7. 400: バリデーションエラー：replyMessageがstring型でない場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				inputJSON, _ := json.Marshal(i)
				inputJsonStr := string(inputJSON)

				var inputJsonMap map[string]interface{}
				// nolint
				json.Unmarshal([]byte(inputJsonStr), &inputJsonMap)
				inputJsonMap["replyMessage"] = 1
				// nolint
				inputJSON, _ = json.Marshal(inputJsonMap)

				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, replyMessage: Unmarshal type error: expected=string, got=number.",
		},
		{
			name: "1-8. 400: バリデーションエラー：replyMessageが101文字以上の場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.ReplyMessage = "Jv0ceJYX9Pa9zQTtlYPQLqLyUUhhNZ5EQCL2JDj9jLfrrgFK8MzV7zkaPvVj1wVtq5ESQGAbXrOhElxsVJzBjSxMBhwOa7hJwBrEc"
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, replyMessage: the length must be no more than 100.",
		},
		{
			name: "1-9. 400: バリデーションエラー：requestStatusのcfpResponseStatusが不正の場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatus(f.InvalidEnum)
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, requestStatus: (cfpResponseStatus: cannot be allowed 'invalid_enum').",
		},
		{
			name: "1-11. 400: バリデーションエラー：1-1と1-3が同時に発生した場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.StatusID = nil
				i.TradeID = nil
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, statusId: cannot be blank; tradeId: cannot be blank.",
		},
		{
			name: "1-12. 400: バリデーションエラー：1-3と1-9が同時に発生した場合",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.TradeID = nil
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatus(f.InvalidEnum)
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, tradeId: cannot be blank.; requestStatus: (cfpResponseStatus: cannot be allowed 'invalid_enum').",
		},
		{
			name: "1-13. 500: システムエラー：更新処理エラー",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusCancel
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-14. 500: システムエラー：更新処理エラー",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusCancel
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			receive:     fmt.Errorf("Internal Server Error"),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-15. 500: システムエラー：更新処理エラー",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusReject
				i.ReplyMessage = "差戻メッセージ"
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-16. 500: システムエラー：更新処理エラー",
			modifyInput: func(i traceability.PutStatusInput) string {
				i.PutRequestStatusInput.CfpResponseStatus = traceability.CfpResponseStatusReject
				i.ReplyMessage = "差戻メッセージ"
				inputJSON, _ := json.Marshal(i)
				return string(inputJSON)
			},
			receive:     fmt.Errorf("Internal Server Error"),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			input := f.NewPutStatusInput()
			inputJSONStr := test.modifyInput(input)

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(inputJSONStr))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			statusUsecase := new(mocks.IStatusUsecase)
			statusUsecase.On("PutStatusCancel", mock.Anything, mock.Anything).Return(test.receive)
			statusUsecase.On("PutStatusReject", mock.Anything, mock.Anything).Return(test.receive)
			statusHandler := handler.NewStatusHandler(statusUsecase, "")

			err := statusHandler.PutStatus(c)
			// エラーが返されることを確認
			if assert.Error(t, err) {
				// エラーメッセージが期待通りであることを確認
				assert.ErrorContains(t, err, test.expectError)
			}
		})
	}
}
