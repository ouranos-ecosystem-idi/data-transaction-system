package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
// Get /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 200: 正常系(トレース識別子指定)
// /////////////////////////////////////////////////////////////////////////////////

func TestProjectHandler_GetPartsStructure_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		expectStatus      int
	}{
		{
			name: "2-1. 200: 正常系(トレース識別子指定)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			partsStructureModel := traceability.PartsStructureModel{}

			q := make(url.Values)
			test.modifyQueryParams(q)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			traceId, _ := uuid.Parse(q.Get("traceId"))
			operatorID := f.OperatorId
			input := traceability.GetPartsStructureInput{
				TraceID:    traceId,
				OperatorID: operatorID,
			}

			partsStructureUsecase := new(mocks.IPartsStructureUsecase)
			partsHandler := handler.NewPartsStructureHandler(partsStructureUsecase)
			partsStructureUsecase.On("GetPartsStructure", c, input).Return(partsStructureModel, nil)

			err := partsHandler.GetPartsStructureModel(c)
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
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：traceIdの値が含まれない場合
// [x] 1-2. 400: バリデーションエラー：traceIdの値が不正の場合
// [x] 1-3. 500: システムエラー：取得処理エラー
// [x] 1-4. 500: システムエラー：取得処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetPartsStructure(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		invalidInput      any
		receive           error
		expectError       string
		expectStatus      int
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdの値が含まれない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdの値が不正な場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
				q.Set("traceId", "hoge")
			},
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
				q.Set("traceId", f.TraceId)
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-4. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
				q.Set("traceId", f.TraceId)
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

				partsStructureUsecase := new(mocks.IPartsStructureUsecase)
				partsStructureUsecase.On("GetPartsStructure", mock.Anything, mock.Anything).Return(traceability.PartsStructureModel{}, test.receive)
				partsHandler := handler.NewPartsStructureHandler(partsStructureUsecase)

				err := partsHandler.GetPartsStructureModel(c)
				e.HTTPErrorHandler(err, c)
				if assert.Error(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 201: 正常系(親を新規作成、子を新規作成)
// [x] 2-2. 201: 正常系(親を新規作成、子は空配列を指定)
// [x] 2-3. 201: 正常系(親を指定、子を新規作成)
// [x] 2-4. 201: 正常系(親を指定、子を指定)
// [x] 2-5. 201: 正常系(親を指定、子は空配列を指定)
// /////////////////////////////////////////////////////////////////////////////////

func TestProjectHandler_PutPartsStructure_ParentChild_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	tests := []struct {
		name         string
		inputFunc    func() traceability.PutPartsStructureInput
		expectStatus int
	}{
		{
			name: "2-1. 201: 正常系(親を新規作成、子を新規作成)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.TraceID = nil
				(*input.ChildrenPartsInput)[0].TraceID = nil
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-2. 201: 正常系(親を新規作成、子は空配列を指定)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.TraceID = nil
				input.ChildrenPartsInput = &traceability.PutPartsInputs{}
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-3. 201: 正常系(親を指定、子を新規作成)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].TraceID = nil
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-4. 201: 正常系(親を指定、子を指定)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-5. 201: 正常系(親を指定、子は空配列を指定)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ChildrenPartsInput = &traceability.PutPartsInputs{}
				return input
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			input := test.inputFunc()
			inputJSON, _ := json.Marshal(input)

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), bytes.NewBuffer(inputJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			parentPartsModel, _ := input.ParentPartsInput.ToModel()

			childrenPartsModel, _ := input.ChildrenPartsInput.ToModels()

			partsStructureModel := traceability.PartsStructureModel{
				ParentPartsModel:   &parentPartsModel,
				ChildrenPartsModel: childrenPartsModel,
			}

			partsStructureUsecase := new(mocks.IPartsStructureUsecase)
			partsHandler := handler.NewPartsStructureHandler(partsStructureUsecase)
			partsStructureUsecase.On("PutPartsStructure", c, test.inputFunc()).Return(partsStructureModel, common.ResponseHeaders{}, nil)

			err := partsHandler.PutPartsStructureModel(c)
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
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 3-1. 201: 正常系(入力値が最小値または最小桁数)
// [x] 3-2. 201: 正常系(入力値が最大値または最大桁数)
// [x] 3-3. 201: 正常系(nil許容項目がnil)
// [x] 3-4. 201: 正常系(任意項目が未定義)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutPartsStructure_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	tests := []struct {
		name         string
		inputFunc    func() traceability.PutPartsStructureInput
		expectStatus int
	}{
		{
			name: "3-1. 201: 正常系(入力値が最小値または最小桁数)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsName = "A"
				input.ParentPartsInput.SupportPartsName = common.StringPtr("")
				input.ParentPartsInput.PartsLabelName = common.StringPtr("")
				input.ParentPartsInput.PartsAddInfo1 = common.StringPtr("")
				input.ParentPartsInput.PartsAddInfo2 = common.StringPtr("")
				input.ParentPartsInput.PartsAddInfo3 = common.StringPtr("")
				(*input.ChildrenPartsInput)[0].PartsName = "A"
				(*input.ChildrenPartsInput)[0].SupportPartsName = common.StringPtr("")
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(0)
				(*input.ChildrenPartsInput)[0].PartsLabelName = common.StringPtr("")
				(*input.ChildrenPartsInput)[0].PartsAddInfo1 = common.StringPtr("")
				(*input.ChildrenPartsInput)[0].PartsAddInfo2 = common.StringPtr("")
				(*input.ChildrenPartsInput)[0].PartsAddInfo3 = common.StringPtr("")
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "3-2. 201: 正常系(入力値が最大値または最大桁数)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsName = "12345678901234567890123456789012345678901234567890"
				input.ParentPartsInput.SupportPartsName = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.ParentPartsInput.PartsLabelName = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.ParentPartsInput.PartsAddInfo1 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.ParentPartsInput.PartsAddInfo2 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				input.ParentPartsInput.PartsAddInfo3 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				(*input.ChildrenPartsInput)[0].PartsName = "12345678901234567890123456789012345678901234567890"
				(*input.ChildrenPartsInput)[0].SupportPartsName = common.StringPtr("12345678901234567890123456789012345678901234567890")
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(99999.99999)
				(*input.ChildrenPartsInput)[0].PartsLabelName = common.StringPtr("12345678901234567890123456789012345678901234567890")
				(*input.ChildrenPartsInput)[0].PartsAddInfo1 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				(*input.ChildrenPartsInput)[0].PartsAddInfo2 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				(*input.ChildrenPartsInput)[0].PartsAddInfo3 = common.StringPtr("12345678901234567890123456789012345678901234567890")
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "3-3. 201: 正常系(nil許容項目がnil)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsLabelName = nil
				input.ParentPartsInput.PartsAddInfo1 = nil
				input.ParentPartsInput.PartsAddInfo2 = nil
				input.ParentPartsInput.PartsAddInfo3 = nil
				(*input.ChildrenPartsInput)[0].PartsLabelName = nil
				(*input.ChildrenPartsInput)[0].PartsAddInfo1 = nil
				(*input.ChildrenPartsInput)[0].PartsAddInfo2 = nil
				(*input.ChildrenPartsInput)[0].PartsAddInfo3 = nil
				return input
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "3-4. 201: 正常系(任意項目が未定義)",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput_RequiredOnlyWithUndefined()
				return input
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			input := test.inputFunc()
			inputJSON, _ := json.Marshal(input)

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), bytes.NewBuffer(inputJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			parentPartsModel, _ := input.ParentPartsInput.ToModel()

			childrenPartsModel, _ := input.ChildrenPartsInput.ToModels()

			partsStructureModel := traceability.PartsStructureModel{
				ParentPartsModel:   &parentPartsModel,
				ChildrenPartsModel: childrenPartsModel,
			}

			partsStructureUsecase := new(mocks.IPartsStructureUsecase)
			partsHandler := handler.NewPartsStructureHandler(partsStructureUsecase)
			partsStructureUsecase.On("PutPartsStructure", c, test.inputFunc()).Return(partsStructureModel, common.ResponseHeaders{}, nil)

			err := partsHandler.PutPartsStructureModel(c)
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
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/partsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1.  400: バリデーションエラー：parentPartsModelが未指定の場合
// [x] 1-2.  400: バリデーションエラー：parentPartsModelのamountRequiredの値がマイナスの場合
// [x] 1-3.  400: バリデーションエラー：parentPartsModelのamountRequiredの値が値が桁数オーバーの場合
// [x] 1-4.  400: バリデーションエラー：parentPartsModelのamountRequiredUnitの値が指定のEnum以外の場合
// [x] 1-5.  400: バリデーションエラー：parentPartsModelのoperatorIdの値が未指定の場合
// [x] 1-6.  400: バリデーションエラー：parentPartsModelのoperatorIdの値がUUID以外の場合
// [x] 1-7.  400: バリデーションエラー：parentPartsModelのpartsNameの値が未指定の場合
// [x] 1-8.  400: バリデーションエラー：parentPartsModelのpartsNameの値が51文字以上の場合
// [x] 1-9.  400: バリデーションエラー：parentPartsModelのplantIdの値が未指定の場合
// [x] 1-10. 400: バリデーションエラー：parentPartsModelのplantIdの値がUUID以外の場合
// [x] 1-11. 400: バリデーションエラー：parentPartsModelのsupportPartsNameの値が51文字以上の場合
// [x] 1-12. 400: バリデーションエラー：parentPartsModelのterminatedFlagの値が未指定の場合
// [x] 1-13. 400: バリデーションエラー：parentPartsModelのtraceIdの値がUUID以外の場合
// [x] 1-14. 400: バリデーションエラー：childPartsModelが未指定の場合
// [x] 1-15. 400: バリデーションエラー：childPartsModelのamountRequiredの値が未指定の場合
// [x] 1-16. 400: バリデーションエラー：childPartsModelのamountRequiredの値がマイナスの場合
// [x] 1-17. 400: バリデーションエラー：childPartsModelのamountRequiredの値が値が桁数オーバーの場合
// [x] 1-18. 400: バリデーションエラー：childPartsModelのamountRequiredUnitの値が21文字以上の場合
// [x] 1-19. 400: バリデーションエラー：childPartsModelのoperatorIdの値が未指定の場合
// [x] 1-20. 400: バリデーションエラー：childPartsModelのoperatorIdの値がUUID以外の場合
// [x] 1-21. 400: バリデーションエラー：childPartsModelのpartsNameの値が未指定の場合
// [x] 1-22. 400: バリデーションエラー：childPartsModelのpartsNameの値が51文字以上の場合
// [x] 1-23. 400: バリデーションエラー：childPartsModelのplantIdの値が未指定の場合
// [x] 1-24. 400: バリデーションエラー：childPartsModelのplantIdの値がUUID以外の場合
// [x] 1-25. 400: バリデーションエラー：childPartsModelのsupportPartsNameの値が51文字以上の場合
// [x] 1-26. 400: バリデーションエラー：childPartsModelのterminatedFlagの値が未指定の場合
// [x] 1-27. 400: バリデーションエラー：childPartsModelのtraceIdの値がUUID以外の場合
// [x] 1-28. 400: バリデーションエラー：childPartsModelのamountRequiredの値が少数点以下6桁以上の場合
// [x] 1-29. 403: 認可エラー：parentPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合
// [x] 1-30. 403: 認可エラー：childPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合
// [x] 1-31. 500: システムエラー：更新処理エラー
// [x] 1-32. 500: システムエラー：更新処理エラー
// [x] 1-33. 400: バリデーションエラー：parentPartsModelのpartsNameがstring形式でない
// [x] 1-34. 400: バリデーションエラー：parentPartsModelのsupportPartsNameがstring形式でない
// [x] 1-35. 400: バリデーションエラー：parentPartsModelのterminatedFlagがboolean形式ではない
// [x] 1-36. 400: バリデーションエラー：parentPartsModelのamountRequiredがnumber形式ではない
// [x] 1-37. 400: バリデーションエラー：parentPartsModelのamountRequiredUnitがstring形式でない
// [x] 1-38. 400: バリデーションエラー：childPartsModelのpartsNameがstring形式でない
// [x] 1-39. 400: バリデーションエラー：childPartsModelのsupportPartsNameがstring形式でない
// [x] 1-40. 400: バリデーションエラー：childPartsModelのterminatedFlagがboolean形式ではない
// [x] 1-41. 400: バリデーションエラー：childPartsModelのamountRequiredがnumber形式ではない
// [x] 1-42. 400: バリデーションエラー：childPartsModelのamountRequiredUnitがstring形式でない
// [x] 1-43. 400: バリデーションエラー：parentPartsModelの1-3と1-5が同時に発生する場合
// [x] 1-44. 400: バリデーションエラー：childPartsModelの1-19と1-21が同時に発生する場合
// [x] 1-45. 400: バリデーションエラー：1-5と1-9と1-19と1-21が同時に発生する場合
// [x] 1-46. 400: バリデーションエラー：parentPartsModelのpartsLabelNameの値が51文字以上の場合
// [x] 1-47. 400: バリデーションエラー：parentPartsModelのpartsAddInfo1の値が51文字以上の場合
// [x] 1-48. 400: バリデーションエラー：parentPartsModelのpartsAddInfo2の値が51文字以上の場合
// [x] 1-49. 400: バリデーションエラー：parentPartsModelのpartsAddInfo3の値が51文字以上の場合
// [x] 1-50. 400: バリデーションエラー：childPartsModelのpartsLabelNameの値が51文字以上の場合
// [x] 1-51. 400: バリデーションエラー：childPartsModelのpartsAddInfo1の値が51文字以上の場合
// [x] 1-52. 400: バリデーションエラー：childPartsModelのpartsAddInfo2の値が51文字以上の場合
// [x] 1-53. 400: バリデーションエラー：childPartsModelのpartsAddInfo3の値が51文字以上の場合
// [x] 1-54. 400: バリデーションエラー：parentPartsModelのpartsLabelNameがstring形式でない
// [x] 1-55. 400: バリデーションエラー：parentPartsModelのpartsAddInfo1がstring形式でない
// [x] 1-56. 400: バリデーションエラー：parentPartsModelのpartsAddInfo2がstring形式でない
// [x] 1-57. 400: バリデーションエラー：parentPartsModelのpartsAddInfo3がstring形式でない
// [x] 1-58. 400: バリデーションエラー：childPartsModelのpartsLabelNameがstring形式でない
// [x] 1-59. 400: バリデーションエラー：childPartsModelのpartsAddInfo1がstring形式でない
// [x] 1-60. 400: バリデーションエラー：childPartsModelのpartsAddInfo2がstring形式でない
// [x] 1-61. 400: バリデーションエラー：childPartsModelのpartsAddInfo3がstring形式でない
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutPartsStructure(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"
	tests := []struct {
		name             string
		inputFunc        func() traceability.PutPartsStructureInput
		invalidInputFunc func() interface{}
		receive          error
		expectError      string
		expectStatus     int
	}{
		{
			name: "1-1. 400: バリデーションエラー：parentPartsModelが未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput = nil
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：parentPartsModelのamountRequiredの値がnull以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.AmountRequired = common.Float64Ptr(1.0)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (amountRequired: must be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：parentPartsModelのamountRequiredに値が設定されている場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.AmountRequired = common.Float64Ptr(1.2)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (amountRequired: must be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：parentPartsModelのamountRequiredUnitの値が指定のEnum以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.AmountRequiredUnit = common.StringPtr(f.InvalidEnum)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (amountRequiredUnit: cannot be allowed 'invalid_enum'.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：parentPartsModelのoperatorIdの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.OperatorID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (operatorId: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 400: バリデーションエラー：parentPartsModelのoperatorIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.OperatorID = f.InvalidUUID
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (operatorId: invalid UUID.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-7. 400: バリデーションエラー：parentPartsModelのpartsNameの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsName = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsName: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-8. 400: バリデーションエラー：parentPartsModelのpartsNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsName = "123456789012345678901234567890123456789012345678901"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsName: the length must be between 1 and 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-9. 400: バリデーションエラー：parentPartsModelのplantIdの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PlantID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (plantId: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-10. 400: バリデーションエラー：parentPartsModelのplantIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PlantID = f.InvalidUUID
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (plantId: invalid UUID.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-11. 400: バリデーションエラー：parentPartsModelのsupportPartsNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.SupportPartsName = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (supportPartsName: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-12. 400: バリデーションエラー：parentPartsModelのterminatedFlagの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.TerminatedFlag = nil
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (terminatedFlag: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-13. 400: バリデーションエラー：parentPartsModelのtraceIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.TraceID = common.StringPtr(f.InvalidUUID)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (traceId: invalid UUID.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-14. 400: バリデーションエラー：childPartsModelが未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ChildrenPartsInput = nil
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-15. 400: バリデーションエラー：childPartsModelのamountRequiredの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].AmountRequired = nil
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: is required.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-16. 400: バリデーションエラー：childPartsModelのamountRequiredの値がマイナスの場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(-1.0)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: must be no less than 0.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-17. 400: バリデーションエラー：childPartsModelのamountRequiredの値が桁数オーバーの場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(123456.0)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: must be no greater than 99999.99999.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-19. 400: バリデーションエラー：childPartsModelのoperatorIdの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].OperatorID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (operatorId: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-20. 400: バリデーションエラー：childPartsModelのoperatorIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].OperatorID = f.InvalidUUID
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (operatorId: invalid UUID.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-21. 400: バリデーションエラー：childPartsModelのpartsNameの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsName = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsName: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-22. 400: バリデーションエラー：childPartsModelのpartsNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsName = "123456789012345678901234567890123456789012345678901"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsName: the length must be between 1 and 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-23. 400: バリデーションエラー：childPartsModelのplantIdの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PlantID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (plantId: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-24. 400: バリデーションエラー：childPartsModelのplantIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PlantID = f.InvalidUUID
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (plantId: invalid UUID.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-25. 400: バリデーションエラー：childPartsModelのsupportPartsNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].SupportPartsName = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (supportPartsName: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-26. 400: バリデーションエラー：childPartsModelのterminatedFlagの値が未指定の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].TerminatedFlag = nil
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (terminatedFlag: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-27. 400: バリデーションエラー：childPartsModelのtraceIdの値がUUID以外の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].TraceID = common.StringPtr(f.InvalidUUID)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (traceId: invalid UUID.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-28. 400: バリデーションエラー：childPartsModelのamountRequiredの値が少数点以下6桁以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(1.123456)
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: must be a value up to the 5th decimal place.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-29. 403: 認可エラー：parentPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.OperatorID = "80762b76-cf76-4485-9a99-cbe609c677c8"
				return input
			},
			expectError:  "code=403, message={[dataspace] AccessDenied You do not have the necessary privileges",
			expectStatus: http.StatusForbidden,
		},
		{
			name: "1-30. 403: 認可エラー：childPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].OperatorID = "80762b76-cf76-4485-9a99-cbe609c677c8"
				return input
			},
			expectError:  "code=403, message={[dataspace] AccessDenied You do not have the necessary privileges",
			expectStatus: http.StatusForbidden,
		},
		{
			name: "1-31. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				return input
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-32. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				return input
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-33. 400: バリデーションエラー：parentPartsModelのpartsNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.partsName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-34. 400: バリデーションエラー：parentPartsModelのsupportPartsNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["supportPartsName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.supportPartsName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-35. 400: バリデーションエラー：parentPartsModelのterminatedFlagがboolean形式ではない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["terminatedFlag"] = "value"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.terminatedFlag: Unmarshal type error: expected=bool, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-36. 400: バリデーションエラー：parentPartsModelのamountRequiredがnumber形式ではない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["amountRequired"] = "value"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.amountRequired: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-37. 400: バリデーションエラー：parentPartsModelのamountRequiredUnitがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["amountRequiredUnit"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.amountRequiredUnit: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-38. 400: バリデーションエラー：childrenPartsModelのpartsNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["partsName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.partsName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-39. 400: バリデーションエラー：childrenPartsModelのsupportPartsNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["supportPartsName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.supportPartsName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-40. 400: バリデーションエラー：childrenPartsModelのterminatedFlagがboolean形式ではない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["terminatedFlag"] = "value"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.terminatedFlag: Unmarshal type error: expected=bool, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-41. 400: バリデーションエラー：childrenPartsModelのamountRequiredがnumber形式ではない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["amountRequired"] = "value"
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.amountRequired: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-42. 400: バリデーションエラー：childrenPartsModelのamountRequiredUnitがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["amountRequiredUnit"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.amountRequiredUnit: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-43. 400: バリデーションエラー：parentPartsModelの1-3と1-5が同時に発生する場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsName = ""
				input.ParentPartsInput.OperatorID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (operatorId: cannot be blank; partsName: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-44. 400: バリデーションエラー：childPartsModelの1-19と1-21が同時に発生する場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsName = ""
				(*input.ChildrenPartsInput)[0].OperatorID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (operatorId: cannot be blank; partsName: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-45. 400: バリデーションエラー：1-5と1-9と1-19と1-21が同時に発生する場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				input.ParentPartsInput.PartsName = ""
				input.ParentPartsInput.OperatorID = ""
				(*input.ChildrenPartsInput)[0].PartsName = ""
				(*input.ChildrenPartsInput)[0].OperatorID = ""
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (operatorId: cannot be blank; partsName: cannot be blank.); childrenPartsModel[0]: (operatorId: cannot be blank; partsName: cannot be blank.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-46. 400: バリデーションエラー：parentPartsModelのpartsLabelNameの値が51文字以上の場合",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsLabelName"] = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsLabelName: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-47. 400: バリデーションエラー：parentPartsModelのpartsAddInfo1の値が51文字以上の場合",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsAddInfo1"] = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsAddInfo1: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-48. 400: バリデーションエラー：parentPartsModelのpartsAddInfo2の値が51文字以上の場合",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsAddInfo2"] = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsAddInfo2: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-49. 400: バリデーションエラー：parentPartsModelのpartsAddInfo3の値が51文字以上の場合",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsAddInfo3"] = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsAddInfo3: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-50. 400: バリデーションエラー：childPartsModelのpartsLabelNameの値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsLabelName = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsLabelName: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-51. 400: バリデーションエラー：childPartsModelのpartsAddInfo1の値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsAddInfo1 = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsAddInfo1: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-52. 400: バリデーションエラー：childPartsModelのpartsAddInfo2の値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsAddInfo2 = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsAddInfo2: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-53. 400: バリデーションエラー：childPartsModelのpartsAddInfo3の値が51文字以上の場合",
			inputFunc: func() traceability.PutPartsStructureInput {
				input := f.NewPutPartsStructureInput()
				(*input.ChildrenPartsInput)[0].PartsAddInfo3 = common.StringPtr("123456789012345678901234567890123456789012345678901")
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsAddInfo3: the length must be no more than 50.)",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-54. 400: バリデーションエラー：parentPartsModelのpartsLabelNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsLabelName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.partsLabelName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-55. 400: バリデーションエラー：parentPartsModelのpartsAddInfo1がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsAddInfo1"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.partsAddInfo1: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-56. 400: バリデーションエラー：parentPartsModelのpartsAddInfo2がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsAddInfo2"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.partsAddInfo2: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-57. 400: バリデーションエラー：parentPartsModelのpartsAddInfo3がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["parentPartsModel"].(map[string]interface{})["partsAddInfo3"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel.partsAddInfo3: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-58. 400: バリデーションエラー：childPartsModelのpartsLabelNameがstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["partsLabelName"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.partsLabelName: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-59. 400: バリデーションエラー：childPartsModelのpartsAddInfo1がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["partsAddInfo1"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.partsAddInfo1: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-60. 400: バリデーションエラー：childPartsModelのpartsAddInfo2がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["partsAddInfo2"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.partsAddInfo2: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-61. 400: バリデーションエラー：childPartsModelのpartsAddInfo3がstring形式でない",
			invalidInputFunc: func() interface{} {
				input := f.NewPutPartsStructureInterface()
				input.(map[string]interface{})["childrenPartsModel"].([]interface{})[0].(map[string]interface{})["partsAddInfo3"] = 1
				return input
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel.partsAddInfo3: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		tc := tc
		tt.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var inputJSON []byte
			if tc.invalidInputFunc != nil {
				input := tc.invalidInputFunc()
				inputJSON, _ = json.Marshal(input)
			} else {
				input := tc.inputFunc()
				inputJSON, _ = json.Marshal(input)
			}

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), bytes.NewBuffer(inputJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			partsStructureUsecase := new(mocks.IPartsStructureUsecase)
			partsStructureUsecase.On("PutPartsStructure", mock.Anything, mock.Anything).Return(traceability.PartsStructureModel{}, common.ResponseHeaders{}, tc.receive)
			partsHandler := handler.NewPartsStructureHandler(partsStructureUsecase)

			err := partsHandler.PutPartsStructureModel(c)
			e.HTTPErrorHandler(err, c)
			// エラーが返されることを確認
			if assert.Error(t, err) {
				// ステータスコードが期待通りであることを確認
				assert.Equal(t, tc.expectStatus, rec.Code)
				// エラーメッセージが期待通りであることを確認
				assert.ErrorContains(t, err, tc.expectError)
			}
		})
	}
}
