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
// Get /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 200: 正常系(Queryパラメータ追加なし)
// [x] 2-2. 200: 正常系(limit指定)
// [x] 2-3. 200: 正常系(after指定)
// [x] 2-4. 200: 正常系(traceId指定)
// [x] 2-5. 200: 正常系(partsName指定)
// [x] 2-6. 200: 正常系(plantId指定)
// [x] 2-7. 200: 正常系(parentFlag指定)
// [x] 2-8. 200: 正常系(limit+after指定)
// [x] 2-9. 200: 正常系(limit+after+traceId指定)
// [x] 2-10. 200: 正常系(limit+after+traceId+partsName指定)
// [x] 2-11. 200: 正常系(limit+after+traceId+partsName+plantId指定)
// [x] 2-12. 200: 正常系(limit+after+traceId+partsName+plantId+parentFlag指定)
// [x] 2-13. 200: limitに値が設定されていない場合(limit=)
// [x] 1-4. 200: afterに値が設定されていない場合(after=)
// [x] 1-6. 200: traceIdに値が設定されていない場合(traceId=)
// [x] 1-8. 200: partsNameに値が設定されていない場合(partsName=)
// [x] 1-10. 200: partsNameが21文字以上の場合
// [x] 1-11. 200: plantIdに値が設定されていない場合(plantId=)
// [x] 1-13. 200: parentFlagに値が設定されていない場合(parentFlag=)
// [x] 1-15. 200: 1-6と1-8が同時に発生する場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetParts_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		after             *string
		expectStatus      int
	}{
		{
			name: "2-1. 200: 正常系(Queryパラメータ追加なし)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-2. 200: 正常系(limit指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "5")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-3. 200: 正常系(after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("after", f.TraceId)
			},
			after:        common.StringPtr(f.TraceId),
			expectStatus: http.StatusOK,
		},
		{
			name: "2-4. 正常系(traceId指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-5. 200: 正常系(partsName指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("partsName", f.PartsName)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-6. 200: 正常系(plantId指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("plantId", f.PlantId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-7. 200: 正常系(parentFlag指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("parentFlag", "true")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-8. 200: 正常系(limit+after指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "5")
				q.Set("after", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-9. 200: 正常系(limit+after+traceId指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "5")
				q.Set("after", f.TraceId)
				q.Set("traceId", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-10. 200: 正常系(limit+after+traceId+partsName指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "5")
				q.Set("after", f.TraceId)
				q.Set("traceId", f.TraceId)
				q.Set("partsName", f.PartsName)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-11. 200: 正常系(limit+after+traceId+partsName+plantId指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "5")
				q.Set("after", f.TraceId)
				q.Set("traceId", f.TraceId)
				q.Set("partsName", f.PartsName)
				q.Set("plantId", f.PlantId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-12. 200: 正常系(limit+after+traceId+partsName+plantId+parentFlag指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "5")
				q.Set("after", f.TraceId)
				q.Set("traceId", f.TraceId)
				q.Set("partsName", f.PartsName)
				q.Set("plantId", f.PlantId)
				q.Set("parentFlag", "true")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-13. 200: limitに値が設定されていない場合(limit=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-4. 200: afterに値が設定されていない場合(after=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("after", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-6. 200: traceIdに値が設定されていない場合(traceId=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-8. 200: partsNameに値が設定されていない場合(partsName=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("partsName", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-10. 200: partsNameが21文字以上の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("partsName", "123456789012345678901")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-11. 200: plantIdに値が設定されていない場合(plantId=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("plantId", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-13. 200: parentFlagに値が設定されていない場合(parentFlag=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("parentFlag", "")
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "1-15. 200: 1-6と1-8が同時に発生する場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", "")
				q.Set("partsName", "")
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				// t.Parallel()
				operatorID := f.OperatorId

				input := traceability.GetPartsInput{
					OperatorID: operatorID,
				}
				partModel := []traceability.PartsModel{}

				q := make(url.Values)
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				limit := q.Get("limit")
				after := q.Get("after")
				traceID := q.Get("traceId")
				partsName := q.Get("partsName")
				plantID := q.Get("plantId")
				parentFlag := q.Get("parentFlag")

				// クエリパラメータのチェック
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
				if traceID != "" {
					input.TraceID = &traceID
				}
				if partsName != "" {
					input.PartsName = &partsName
				}
				if plantID != "" {
					input.PlantID = &plantID
				}
				if after != "" {
					after, _ := uuid.Parse(f.TraceId)
					input.After = &after
				}
				if parentFlag != "" {
					var myBool bool = true
					input.ParentFlag = &myBool
				}

				partsUsecase := new(mocks.IPartsUsecase)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")
				partsUsecase.On("GetPartsList", c, input).Return(partModel, test.after, nil)

				err := partsHandler.GetPartsModel(c)
				// エラーが発生しないことを確認
				if assert.NoError(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// モックの呼び出しが期待通りであることを確認
					partsUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：limitがintでない場合
// [x] 1-2. 400: バリデーションエラー：limitが101以上の場合
// [x] 1-3. 400: バリデーションエラー：limitが0以下の場合
// [x] 1-4. 400: バリデーションエラー：afterに値が設定されていない場合(after=)
// [x] 1-5. 400: バリデーションエラー：afterがUUIDでない場合
// [x] 1-6. 400: バリデーションエラー：traceIdに値が設定されていない場合(traceId=)
// [x] 1-7. 400: バリデーションエラー：traceIdがUUIDでない場合
// [x] 1-8. 400: バリデーションエラー：partsNameに値が設定されていない場合(partsName=)
// [x] 1-9. 400: バリデーションエラー：partsNameがstringでない場合
// [x] 1-10. 400: バリデーションエラー：partsNameが21文字以上の場合
// [x] 1-11. 400: バリデーションエラー：plantIdに値が設定されていない場合(plantId=)
// [x] 1-12. 400: バリデーションエラー：plantIdがUUIDでない場合
// [x] 1-13. 400: バリデーションエラー：parentFlagに値が設定されていない場合(parentFlag=)
// [x] 1-14. 400: バリデーションエラー：parentFlagがboolean型でない場合
// [x] 1-15. 500: システムエラー：取得処理エラー
// [x] 1-16. 500: システムエラー：取得処理エラー
// [x] 1-17. 500: システムエラー：取得処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetParts(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：limitがintでない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "hoge")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：limitが101以上の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "101")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：limitが0以下の場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("limit", "0")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, limit: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：afterがUUIDでない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("after", f.InvalidUUID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, after: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-7. 400: バリデーションエラー：traceIdがUUIDでない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.InvalidUUID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-12. 400: バリデーションエラー：plantIdがUUIDでない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("plantId", f.InvalidUUID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, plantId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-14. 400: バリデーションエラー：parentFlagがboolean型でない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("parentFlag", f.InvalidUUID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, parentFlag: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-15. 400: バリデーションエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
			},
			receive:      common.NewCustomError(common.CustomErrorCode400, "Invalid request parameters", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-16. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-17. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
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
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsUsecase := new(mocks.IPartsUsecase)
				partsUsecase.On("GetPartsList", mock.Anything, mock.Anything).Return([]traceability.PartsModel{}, common.StringPtr(""), test.receive)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")

				err := partsHandler.GetPartsModel(c)
				e.HTTPErrorHandler(err, c)
				// エラーが返されることを確認
				if assert.Error(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// エラーメッセージが期待通りであることを確認
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/parts 正常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 201: 正常系(新規作成:amountRequiredが指定されている場合)
// [x] 1-2. 201: 正常系(新規作成: amountRequiredがnullの場合)
// [x] 1-3. 201: 正常系(更新: 全ての値に値が指定されている場合)
// // /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutParts_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name         string
		inputFunc    func() traceability.PutPartsInput
		expectStatus int
	}{
		{
			name: "1-2. 201: 正常系(新規作成: amountRequiredがnullの場合)",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.TraceID = nil
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "1-3. 201: 正常系(更新: 全ての値に値が指定されている場合)",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				return input
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
				inputJSON, _ := json.Marshal(input)

				putPartsStructureInput := traceability.PutPartsStructureInput{
					ParentPartsInput: &input,
				}

				partsModel, _ := input.ToModel()
				var partsStructureModel = traceability.PartsStructureModel{
					ParentPartsModel: &partsModel,
				}
				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(string(inputJSON)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsUsecase := new(mocks.IPartsUsecase)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")
				partsStructureUsecase.On("PutPartsStructure", c, putPartsStructureInput).Return(partsStructureModel, common.ResponseHeaders{}, nil)

				err := partsHandler.PutPartsModel(c)
				// エラーが発生しないことを確認
				if assert.NoError(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// モックの呼び出しが期待通りであることを確認
					partsStructureUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/parts 正常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 201: 正常系(入力値が最小値または最小桁数)
// [x] 2-2. 201: 正常系(入力値が最大値または最大桁数)
// [x] 2-3. 201: 正常系(nil許容項目がnil)
// [x] 2-4. 201: 正常系(任意項目が未定義)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutParts_Validation_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name         string
		inputFunc    func() traceability.PutPartsInput
		expectStatus int
	}{
		{
			name: "2-1. 201: 正常系(入力値が最小値または最小桁数)",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsName = "A"
				input.SupportPartsName = common.StringPtr("")
				input.PartsLabelName = common.StringPtr("")
				input.PartsAddInfo1 = common.StringPtr("")
				input.PartsAddInfo2 = common.StringPtr("")
				input.PartsAddInfo3 = common.StringPtr("")
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-2. 201: 正常系(入力値が最大値または最大桁数)",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsName = "12345678901234567890123456789012345678901234567890"
				input.SupportPartsName = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.PartsLabelName = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.PartsAddInfo1 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.PartsAddInfo2 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.PartsAddInfo3 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-3. 201: 正常系(nil許容項目がnil)",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsLabelName = nil
				input.PartsAddInfo1 = nil
				input.PartsAddInfo2 = nil
				input.PartsAddInfo3 = nil
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-4. 201: 正常系(任意項目が未定義)",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput_RequiredOnly()
				return input
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
				inputJSON, _ := json.Marshal(input)

				putPartsStructureInput := traceability.PutPartsStructureInput{
					ParentPartsInput: &input,
				}

				partsModel, _ := input.ToModel()
				var partsStructureModel = traceability.PartsStructureModel{
					ParentPartsModel: &partsModel,
				}
				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(string(inputJSON)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsUsecase := new(mocks.IPartsUsecase)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")
				partsStructureUsecase.On("PutPartsStructure", c, putPartsStructureInput).Return(partsStructureModel, common.ResponseHeaders{}, nil)

				err := partsHandler.PutPartsModel(c)
				// エラーが発生しないことを確認
				if assert.NoError(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// モックの呼び出しが期待通りであることを確認
					partsStructureUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：amountRequiredがnull以外の値の場合
// [x] 1-2. 400: バリデーションエラー：amountRequiredUnitが指定のEnumではない場合
// [x] 1-3. 400: バリデーションエラー：operatorIdの値が未指定の場合
// [x] 1-4. 400: バリデーションエラー：operatorIdの値がUUID以外の場合
// [x] 1-5. 400: バリデーションエラー：partsNameの値が未指定の場合
// [x] 1-6. 400: バリデーションエラー：partsNameの値が51文字以上の場合
// [x] 1-7. 400: バリデーションエラー：plantIdの値が未指定の場合
// [x] 1-8. 400: バリデーションエラー：plantIdの値がUUID以外の場合
// [x] 1-9. 400: バリデーションエラー：supportPartsNameの値が51文字以上の場合
// [x] 1-10. 400: バリデーションエラー：terminatedFlagの値が未指定の場合
// [x] 1-11. 400: バリデーションエラー：traceIdの値がUUID以外の場合
// [x] 1-12. 400: バリデーションエラー：operatorIdがstring形式でない
// [x] 1-13. 400: バリデーションエラー：operatorIDがjwtのoperatorIdと一致しない場合
// [x] 1-14. 500: システムエラー：更新処理エラー
// [x] 1-15. 500: システムエラー：更新処理エラー
// [x] 1-16. 400: バリデーションエラー：partsNameがstring形式でない
// [x] 1-17. 400: バリデーションエラー：supportPartsNameがstring形式でない
// [x] 1-18. 400: バリデーションエラー：terminatedFlagがboolean形式ではない
// [x] 1-19. 400: バリデーションエラー：amountRequiredがnumber形式ではない
// [x] 1-20. 400: バリデーションエラー：amountRequiredUnitがstring形式でない
// [x] 1-21. 400: バリデーションエラー：1-3と1-5が同時に発生する場合
// [x] 1-22. 400: バリデーションエラー：partsLabelNameの値が51文字以上の場合
// [x] 1-23. 400: バリデーションエラー：partsAddInfo1の値が51文字以上の場合
// [x] 1-24. 400: バリデーションエラー：partsAddInfo2の値が51文字以上の場合
// [x] 1-25. 400: バリデーションエラー：partsAddInfo3の値が51文字以上の場合
// [x] 1-26. 400: バリデーションエラー：partsLabelNameがstring形式でない
// [x] 1-27. 400: バリデーションエラー：partsAddInfo1がstring形式でない
// [x] 1-28. 400: バリデーションエラー：partsAddInfo2がstring形式でない
// [x] 1-29. 400: バリデーションエラー：partsAddInfo3がstring形式でない
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutParts(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name             string
		inputFunc        func() traceability.PutPartsInput
		invalidInputFunc func() interface{}
		receive          error
		expectError      string
		expectStatus     int
	}{
		{
			name: "1-1. 400: バリデーションエラー：amountRequiredの値がnull以外の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.AmountRequired = common.Float64Ptr(1.0)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, amountRequired: must be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：amountRequiredUnitが指定のEnumではない場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.AmountRequiredUnit = common.StringPtr("invalid_enum")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, amountRequiredUnit: cannot be allowed 'invalid_enum'",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：operatorIdの値が未指定の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.OperatorID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, operatorId: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：operatorIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.OperatorID = f.InvalidUUID
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, operatorId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：partsNameの値が未指定の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsName = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsName: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 400: バリデーションエラー：partsNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsName = "123456789012345678901234567890123456789012345678901"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsName: the length must be between 1 and 50.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-7. 400: バリデーションエラー：plantIdの値が未指定の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PlantID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, plantId: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-8. 400: バリデーションエラー：plantIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PlantID = f.InvalidUUID
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, plantId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-9. 400: バリデーションエラー：supportPartsNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.SupportPartsName = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, supportPartsName: the length must be no more than 50.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-10. 400: バリデーションエラー：terminatedFlagの値が未指定の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.TerminatedFlag = nil
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, terminatedFlag: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-11. 400: バリデーションエラー：traceIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.TraceID = common.StringPtr(f.InvalidUUID)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, traceId: invalid UUID.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-12. 400: バリデーションエラー：operatorIdがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["operatorId"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, operatorId: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-13. 400: バリデーションエラー：operatorIDがjwtのoperatorIdと一致しない場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.OperatorID = "80762b76-cf76-4485-9a99-cbe609c677c8"
				return input
			},
			expectError:  "code=403, message={[dataspace] AccessDenied You do not have the necessary privileges",
			expectStatus: http.StatusForbidden,
		},
		{
			name: "1-14. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				return input
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-15. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				return input
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-16. 400: バリデーションエラー：partsNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["partsName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-17. 400: バリデーションエラー：supportPartsNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["supportPartsName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, supportPartsName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-18. 400: バリデーションエラー：terminatedFlagがboolean形式ではない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["terminatedFlag"] = "value"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, terminatedFlag: Unmarshal type error: expected=bool, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-19. 400: バリデーションエラー：amountRequiredがnumber形式ではない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["amountRequired"] = "value"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, amountRequired: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-20. 400: バリデーションエラー：amountRequiredUnitがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["amountRequiredUnit"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, amountRequiredUnit: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-21. 400: バリデーションエラー：1-3と1-5が同時に発生する場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.OperatorID = ""
				input.PartsName = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, operatorId: cannot be blank; partsName: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-22. 400: バリデーションエラー：partsLabelNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsLabelName = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsLabelName: the length must be no more than 50.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-23. 400: バリデーションエラー：partsAddInfo1の値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsAddInfo1 = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsAddInfo1: the length must be no more than 50.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-24. 400: バリデーションエラー：partsAddInfo2の値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsAddInfo2 = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsAddInfo2: the length must be no more than 50.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-25. 400: バリデーションエラー：partsAddInfo3の値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsInput {
				input := f.NewPutPartsInput()
				input.PartsAddInfo3 = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsAddInfo3: the length must be no more than 50.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-26. 400: バリデーションエラー：partsLabelNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["partsLabelName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsLabelName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-27. 400: バリデーションエラー：partsAddInfo1がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["partsAddInfo1"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsAddInfo1: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-28. 400: バリデーションエラー：partsAddInfo2がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["partsAddInfo2"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsAddInfo2: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-29. 400: バリデーションエラー：partsAddInfo3がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsInterface()
				input.(map[string]interface{})["partsAddInfo3"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, partsAddInfo3: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				var inputJSON []byte
				if test.invalidInputFunc != nil {
					input := test.invalidInputFunc()
					inputJSON, _ = json.Marshal(input)
				} else {
					input := test.inputFunc()
					inputJSON, _ = json.Marshal(input)
				}
				q := make(url.Values)
				q.Set("dataTarget", dataTarget)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(string(inputJSON)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsUsecase := new(mocks.IPartsUsecase)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsStructureUsecase.On("PutPartsStructure", mock.Anything, mock.Anything).Return(traceability.PartsStructureModel{}, common.ResponseHeaders{}, test.receive)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")

				err := partsHandler.PutPartsModel(c)
				e.HTTPErrorHandler(err, c)
				// エラーが返されることを確認
				if assert.Error(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// エラーメッセージが期待通りであることを確認
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Delete /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 200: 正常系
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_DeleteParts_Normal(tt *testing.T) {
	var method = "DELETE"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		after             *string
		expectStatus      int
	}{
		{
			name: "2-1. 200: 正常系",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
			},
			expectStatus: http.StatusNoContent,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				// t.Parallel()

				input := traceability.DeletePartsInput{
					TraceID: f.TraceID,
				}

				q := make(url.Values)
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsUsecase := new(mocks.IPartsUsecase)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")
				partsUsecase.On("DeleteParts", c, input).Return(common.ResponseHeaders{}, nil)

				err := partsHandler.DeletePartsModel(c)
				// エラーが発生しないことを確認
				if assert.NoError(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// モックの呼び出しが期待通りであることを確認
					partsUsecase.AssertExpectations(t)
				}

				// レスポンスヘッダにX-Trackが含まれているかチェック
				_, ok := rec.Header()["X-Track"]
				assert.True(t, ok, "Header should have 'X-Track' key")
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Delete /api/v1/datatransport/parts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：traceIdに値が設定されていない場合(traceId=)
// [x] 1-2. 400: バリデーションエラー：traceIdがUUIDでない場合
// [x] 1-3. 500: システムエラー：取得処理エラー
// [x] 1-4. 500: システムエラー：取得処理エラー
// [x] 1-5. 500: システムエラー：取得処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_DeleteParts(tt *testing.T) {
	var method = "DELETE"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "parts"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdに値が設定されていない場合(traceId=)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", "")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdがUUIDでない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.InvalidUUID)
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
			},
			receive:      common.NewCustomError(common.CustomErrorCode400, "traceID has parents", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=400, message={[dataspace] BadRequest traceID has parents",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-5. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceID)
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
				test.modifyQueryParams(q)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				partsUsecase := new(mocks.IPartsUsecase)
				partsUsecase.On("DeleteParts", mock.Anything, mock.Anything).Return(common.ResponseHeaders{}, test.receive)
				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, "")

				err := partsHandler.DeletePartsModel(c)
				e.HTTPErrorHandler(err, c)
				// エラーが返されることを確認
				if assert.Error(t, err) {
					// ステータスコードが期待通りであることを確認
					assert.Equal(t, test.expectStatus, rec.Code)
					// エラーメッセージが期待通りであることを確認
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}
