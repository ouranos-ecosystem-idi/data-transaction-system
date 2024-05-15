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
			input := traceability.GetPartsStructureModel{
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
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdの値が含まれない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdの値が不正な場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
				q.Set("traceId", "hoge")
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceId: Unexpected query parameter",
		},
		{
			name: "1-3. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
				q.Set("traceId", f.TraceId)
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-4. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", "partsStructure")
				q.Set("traceId", f.TraceId)
			},
			receive:     fmt.Errorf("Internal Server Error"),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
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
				if assert.Error(t, err) {
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

func TestProjectHandler_PutPartsStructure_Normal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"

	tests := []struct {
		name         string
		modifyInput  func(input *traceability.PutPartsStructureInput)
		expectStatus int
	}{
		{
			name: "2-1. 201: 正常系(親を新規作成、子を新規作成)",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = f.NewPutPartsStructureInput().ChildrenPartsInput
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-2. 201: 正常系(親を新規作成、子は空配列を指定)",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = &traceability.PutPartsInputs{}
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-3. 201: 正常系(親を指定、子を新規作成)",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = f.NewPutPartsStructureInput().ChildrenPartsInput
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-4. 201: 正常系(親を指定、子を指定)",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = f.NewPutPartsStructureInput().ChildrenPartsInput
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "2-5. 201: 正常系(親を指定、子は空配列を指定)",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = &traceability.PutPartsInputs{}
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			// t.Parallel()

			input := f.NewPutPartsStructureInput()

			test.modifyInput(&input)

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
			partsStructureUsecase.On("PutPartsStructure", c, partsStructureModel).Return(partsStructureModel, nil)

			err := partsHandler.PutPartsStructureModel(c)
			// エラーが発生しないことを確認
			if assert.NoError(t, err) {
				// ステータスコードが期待通りであることを確認
				assert.Equal(t, test.expectStatus, rec.Code)
				// モックの呼び出しが期待通りであることを確認
				partsStructureUsecase.AssertExpectations(t)
			}
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
// [x] 1-8.  400: バリデーションエラー：parentPartsModelのpartsNameの値が21文字以上の場合
// [x] 1-9.  400: バリデーションエラー：parentPartsModelのplantIdの値が未指定の場合
// [x] 1-10. 400: バリデーションエラー：parentPartsModelのplantIdの値がUUID以外の場合
// [x] 1-11. 400: バリデーションエラー：parentPartsModelのsupportPartsNameの値が11文字以上の場合
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
// [x] 1-22. 400: バリデーションエラー：childPartsModelのpartsNameの値が21文字以上の場合
// [x] 1-23. 400: バリデーションエラー：childPartsModelのplantIdの値が未指定の場合
// [x] 1-24. 400: バリデーションエラー：childPartsModelのplantIdの値がUUID以外の場合
// [x] 1-25. 400: バリデーションエラー：childPartsModelのsupportPartsNameの値が11文字以上の場合
// [x] 1-26. 400: バリデーションエラー：childPartsModelのterminatedFlagの値が未指定の場合
// [x] 1-27. 400: バリデーションエラー：childPartsModelのtraceIdの値がUUID以外の場合
// [x] 1-28. 400: バリデーションエラー：childPartsModelのamountRequiredの値が少数点以下6桁以上の場合
// [x] 1-29. 403: 認可エラー：parentPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合
// [x] 1-30. 403: 認可エラー：childPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合
// [x] 1-31. 500: システムエラー：更新処理エラー
// [x] 1-32. 500: システムエラー：更新処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutPartsStructure(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "partsStructure"
	tests := []struct {
		name         string
		modifyInput  func(input *traceability.PutPartsStructureInput)
		invalidInput any
		receive      error
		expectError  string
	}{
		{
			name: "1-1. 400: バリデーションエラー：parentPartsModelが未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = nil
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: cannot be blank.",
		},
		{
			name: "1-2. 400: バリデーションエラー：parentPartsModelのamountRequiredの値がnull以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.AmountRequired = common.Float64Ptr(1.0)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (amountRequired: must be blank.)",
		},
		{
			name: "1-3. 400: バリデーションエラー：parentPartsModelのamountRequiredに値が設定されている場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.AmountRequired = common.Float64Ptr(1.2)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (amountRequired: must be blank.)",
		},
		{
			name: "1-4. 400: バリデーションエラー：parentPartsModelのamountRequiredUnitの値が指定のEnum以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.AmountRequiredUnit = common.StringPtr(f.InvalidEnum)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (amountRequiredUnit: cannot be allowed 'invalid_enum'.)",
		},
		{
			name: "1-5. 400: バリデーションエラー：parentPartsModelのoperatorIdの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.OperatorID = ""
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (operatorId: cannot be blank.)",
		},
		{
			name: "1-6. 400: バリデーションエラー：parentPartsModelのoperatorIdの値がUUID以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.OperatorID = f.InvalidUUID
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (operatorId: invalid UUID.)",
		},
		{
			name: "1-7. 400: バリデーションエラー：parentPartsModelのpartsNameの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.PartsName = ""
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsName: cannot be blank.)",
		},
		{
			name: "1-8. 400: バリデーションエラー：parentPartsModelのpartsNameの値が21文字以上の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.PartsName = "123456789012345678901"
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (partsName: the length must be between 1 and 20.)",
		},
		{
			name: "1-9. 400: バリデーションエラー：parentPartsModelのplantIdの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.PlantID = ""
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (plantId: cannot be blank.)",
		},
		{
			name: "1-10. 400: バリデーションエラー：parentPartsModelのplantIdの値がUUID以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.PlantID = f.InvalidUUID
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (plantId: invalid UUID.)",
		},
		{
			name: "1-11. 400: バリデーションエラー：parentPartsModelのsupportPartsNameの値が11文字以上の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.SupportPartsName = common.StringPtr("12345678901")
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (supportPartsName: the length must be no more than 10.)",
		},
		{
			name: "1-12. 400: バリデーションエラー：parentPartsModelのterminatedFlagの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.TerminatedFlag = nil
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (terminatedFlag: cannot be blank.)",
		},
		{
			name: "1-13. 400: バリデーションエラー：parentPartsModelのtraceIdの値がUUID以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput.TraceID = common.StringPtr(f.InvalidUUID)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, parentPartsModel: (traceId: invalid UUID.)",
		},
		{
			name: "1-14. 400: バリデーションエラー：childPartsModelが未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ChildrenPartsInput = nil
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel: cannot be blank.",
		},
		{
			name: "1-15. 400: バリデーションエラー：childPartsModelのamountRequiredの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].AmountRequired = nil
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: cannot be blank.)",
		},
		{
			name: "1-16. 400: バリデーションエラー：childPartsModelのamountRequiredの値がマイナスの場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(-1.0)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: must be no less than 0.)",
		},
		{
			name: "1-17. 400: バリデーションエラー：childPartsModelのamountRequiredの値が桁数オーバーの場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(123456.0)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: must be no greater than 99999.99999.)",
		},
		{
			name: "1-19. 400: バリデーションエラー：childPartsModelのoperatorIdの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].OperatorID = ""
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (operatorId: cannot be blank.)",
		},
		{
			name: "1-20. 400: バリデーションエラー：childPartsModelのoperatorIdの値がUUID以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].OperatorID = f.InvalidUUID
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (operatorId: invalid UUID.)",
		},
		{
			name: "1-21. 400: バリデーションエラー：childPartsModelのpartsNameの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].PartsName = ""
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsName: cannot be blank.)",
		},
		{
			name: "1-22. 400: バリデーションエラー：childPartsModelのpartsNameの値が21文字以上の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].PartsName = "123456789012345678901"
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (partsName: the length must be between 1 and 20.)",
		},
		{
			name: "1-23. 400: バリデーションエラー：childPartsModelのplantIdの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].PlantID = ""
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (plantId: cannot be blank.)",
		},
		{
			name: "1-24. 400: バリデーションエラー：childPartsModelのplantIdの値がUUID以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].PlantID = f.InvalidUUID
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (plantId: invalid UUID.)",
		},
		{
			name: "1-25. 400: バリデーションエラー：childPartsModelのsupportPartsNameの値が11文字以上の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].SupportPartsName = common.StringPtr("12345678901")
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (supportPartsName: the length must be no more than 10.)",
		},
		{
			name: "1-26. 400: バリデーションエラー：childPartsModelのterminatedFlagの値が未指定の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].TerminatedFlag = nil
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (terminatedFlag: cannot be blank.)",
		},
		{
			name: "1-27. 400: バリデーションエラー：childPartsModelのtraceIdの値がUUID以外の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].TraceID = common.StringPtr(f.InvalidUUID)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (traceId: invalid UUID.)",
		},
		{
			name: "1-28. 400: バリデーションエラー：childPartsModelのamountRequiredの値が少数点以下6桁以上の場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				(*input.ChildrenPartsInput)[0].AmountRequired = common.Float64Ptr(1.123456)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, childrenPartsModel[0]: (amountRequired: must be a value up to the 5th decimal place.)",
		},
		{
			name: "1-29. 403: 認可エラー：parentPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &traceability.PutPartsInput{
					OperatorID:         "80762b76-cf76-4485-9a99-cbe609c677c8",
					TraceID:            &f.TraceId,
					PlantID:            f.PlantId,
					PartsName:          f.PartsName,
					SupportPartsName:   &f.SupportPartsName,
					TerminatedFlag:     &f.TerminatedFlag,
					AmountRequired:     nil,
					AmountRequiredUnit: &f.AmountRequiredUnit,
				}
				input.ChildrenPartsInput = f.NewPutPartsStructureInput().ChildrenPartsInput
			},
			expectError: "code=403, message={[dataspace] AccessDenied You do not have the necessary privileges",
		},
		{
			name: "1-30. 403: 認可エラー：childPartsModelのoperatorIDがjwtのoperatorIdと一致しない場合",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = &traceability.PutPartsInputs{
					traceability.PutPartsInput{
						OperatorID:         "80762b76-cf76-4485-9a99-cbe609c677c8",
						TraceID:            &f.TraceId,
						PlantID:            f.PlantId,
						PartsName:          f.PartsName,
						SupportPartsName:   &f.SupportPartsName,
						TerminatedFlag:     &f.TerminatedFlag,
						AmountRequired:     &f.AmountRequired,
						AmountRequiredUnit: &f.AmountRequiredUnit,
					},
				}
			},
			expectError: "code=403, message={[dataspace] AccessDenied You do not have the necessary privileges",
		},
		{
			name: "1-31. 500: システムエラー：更新処理エラー",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = f.NewPutPartsStructureInput().ChildrenPartsInput
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-32. 500: システムエラー：更新処理エラー",
			modifyInput: func(input *traceability.PutPartsStructureInput) {
				input.ParentPartsInput = &f.PutPartsInput
				input.ChildrenPartsInput = f.NewPutPartsStructureInput().ChildrenPartsInput
			},
			receive:     fmt.Errorf("Internal Server Error"),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var inputJSON []byte
			if test.invalidInput != nil {
				inputJSON, _ = json.Marshal(test.invalidInput)
			} else {
				input := f.NewPutPartsStructureInput()
				test.modifyInput(&input)
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
			partsStructureUsecase.On("PutPartsStructure", mock.Anything, mock.Anything).Return(traceability.PartsStructureModel{}, test.receive)
			partsHandler := handler.NewPartsStructureHandler(partsStructureUsecase)

			err := partsHandler.PutPartsStructureModel(c)
			// エラーが返されることを確認
			if assert.Error(t, err) {
				// エラーメッセージが期待通りであることを確認
				assert.ErrorContains(t, err, test.expectError)
			}
		})
	}
}
