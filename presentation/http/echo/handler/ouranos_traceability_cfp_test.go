package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 200: 正常系(traceIds指定、数は1)
// [x] 2-2. 200: 正常系(traceIds指定、数は2)
// [x] 2-3. 200: 正常系(traceIds指定、数は100)
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetCfp_Normal(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		expectStatus      int
	}{
		{
			name: "2-1. 200: 正常系(traceIds指定、数は1)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", f.TraceId)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-2. 200: 正常系(traceIds指定、数は2)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", common.GenerateUUIDString(2))
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "2-3. 200: 正常系(traceIds指定、数は100)",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", common.GenerateUUIDString(50))
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			q := make(url.Values)
			test.modifyQueryParams(q)
			operatorUUID, _ := uuid.Parse(f.OperatorId)
			input := traceability.GetCfpModel{
				OperatorID: operatorUUID,
			}
			cfpmodel := []traceability.CfpModel{}
			// traceIdsを区切り文字で分割して配列に格納
			traceIds := strings.Split(q.Get("traceIds"), ",")
			if len(traceIds) > 0 {
				// UUIDの配列
				uuids := make([]uuid.UUID, len(traceIds))

				// 文字列をUUIDに変換して配列に格納
				for i, str := range traceIds {
					parsedUUID, err := uuid.Parse(str)
					if err != nil {
						fmt.Println("Error parsing UUID:", err)
						return
					}
					uuids[i] = parsedUUID
				}
				input.TraceIDs = uuids
			}

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			cfpUsecase := new(mocks.ICfpUsecase)
			cfpHandler := handler.NewCfpHandler(cfpUsecase)
			cfpUsecase.On("GetCfp", c, input).Return(cfpmodel, nil, nil)

			err := cfpHandler.GetCfp(c)
			if assert.NoError(t, err) {
				assert.Equal(t, test.expectStatus, rec.Code)
				cfpUsecase.AssertExpectations(t)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：traceIdが含まれない場合
// [x] 1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合
// [x] 1-3. 400: バリデーションエラー：traceIdsにUUIDとそうでない値が混在する場合
// [x] 1-4. 400: バリデーションエラー：traceIdsの複数指定の方法が誤っている場合
// [x] 1-5. 400: バリデーションエラー：traceIdが上限数の50以上指定されている場合
// [x] 1-6. 400: バリデーションエラー：operatorIdがUUID形式ではない場合
// [x] 1-7. 500: システムエラー：取得処理エラー
// [x] 1-8. 500: システムエラー：取得処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_GetCfp(tt *testing.T) {
	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	tests := []struct {
		name              string
		modifyQueryParams func(q url.Values)
		modifyContexts    func(c echo.Context)
		receive           error
		expectError       string
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdが含まれない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceId", uuid.New().String())
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", "NOT_UUID_FORMAT")
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
		},
		{
			name: "1-3. 400: バリデーションエラー：traceIdsにUUIDとそうでない値が混在する場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				traceIds := common.GenerateUUIDString(2)
				traceIds += "INVALIDUUID"
				q.Set("traceIds", traceIds)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
		},
		{
			name: "1-4. 400: バリデーションエラー：traceIdsの複数指定の方法が誤っている場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				traceIds := uuid.New().String()
				traceIds += "&"
				traceIds += uuid.New().String()
				q.Set("traceIds", traceIds)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
		},
		{
			name: "1-5. 400: バリデーションエラー：traceIdが最大の50項目以上指定されている場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", common.GenerateUUIDString(51))
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			expectError: "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
		},
		{
			name: "1-6. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceId)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalid")
			},
			expectError: "code=400, message={[auth] BadRequest Invalid or expired token",
		},
		{
			name: "1-7. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", common.GenerateUUIDString(1))
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", f.OperatorId)
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-8. 500: システムエラー：取得処理エラー",
			modifyQueryParams: func(q url.Values) {
				q.Set("dataTarget", dataTarget)
				q.Set("traceIds", common.GenerateUUIDString(1))
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

			cfpUsecase := new(mocks.ICfpUsecase)
			cfpHandler := handler.NewCfpHandler(cfpUsecase)
			cfpUsecase.On("GetCfp", mock.Anything, mock.Anything).Return([]traceability.CfpModel{}, test.receive)

			err := cfpHandler.GetCfp(c)
			if assert.Error(t, err) {
				assert.ErrorContains(t, err, test.expectError)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PUT /api/v1/datatransport/cfp 正常系テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 201: 正常系：新規作成
// [x] 1-2. 201: 正常系：更新
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutCfp_Nomal(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	tests := []struct {
		name         string
		makeInput    func(inputs traceability.PutCfpInputs) string
		expectStatus int
	}{
		{
			name: "1-1. 201: 正常系：新規作成",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpID = nil
				inputs[1].CfpID = nil
				inputs[2].CfpID = nil
				inputs[3].CfpID = nil
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "1-2. 201: 正常系：更新",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			inputs := f.NewPutCfpInputs()
			inputJSON := test.makeInput(inputs)

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(inputJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			cfpUsecase := new(mocks.ICfpUsecase)
			cfpHandler := handler.NewCfpHandler(cfpUsecase)
			cfpModels, _ := inputs.ToModels()
			responseCfpModels := []traceability.CfpModel{
				cfpModels[0],
				cfpModels[1],
				cfpModels[2],
				cfpModels[3],
			}
			cfpUsecase.On("PutCfp", c, cfpModels, f.OperatorId).Return(responseCfpModels, nil)

			err := cfpHandler.PutCfp(c)
			if assert.NoError(t, err) {
				assert.Equal(t, test.expectStatus, rec.Code)
				cfpUsecase.AssertExpectations(t)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v1/datatransport/cfp テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー：traceIdが含まれない場合
// [x] 1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合
// [x] 1-3. 400: バリデーションエラー：cfpIdがUUID形式ではない場合
// [x] 1-4. 400: バリデーションエラー：ghgEmissionがマイナスの場合
// [x] 1-5. 400: バリデーションエラー：ghgEmissionが上限値を超える場合
// [x] 1-6. 400: バリデーションエラー：ghgDeclaredUnitが含まれない場合
// [x] 1-7. 400: バリデーションエラー：ghgDeclaredUnitが指定のEnum以外の場合
// [x] 1-8. 400: バリデーションエラー：cfpTypeが含まれない場合
// [x] 1-9. 400: バリデーションエラー：cfpTypeが指定のEnumではない場合
// [x] 1-10. 400: バリデーションエラー：cfpの入力が4つでない場合
// [x] 1-11. 400: バリデーションエラー：cfpTypeの４つが揃っていない場合
// [x] 1-12. 400: バリデーションエラー：cfpIdが一致していない場合
// [x] 1-13. 400: バリデーションエラー：traceIdが一致していない場合
// [x] 1-14. 400: バリデーションエラー：ghgDeclaredUnitが一致していない場合
// [x] 1-15. 400: バリデーションエラー：※ケース除去 cfpCertificateListが一致していない場合
// [x] 1-16. 400: バリデーションエラー：dqrTypeが含まれない場合
// [x] 1-17. 400: バリデーションエラー：dqrTypeが指定のEnumではない場合
// [x] 1-18. 400: バリデーションエラー：dqrValue.TeRがマイナスの場合
// [x] 1-19. 400: バリデーションエラー：dqrValue.TeRが上限値を超える場合
// [x] 1-20. 400: バリデーションエラー：dqrValue.GeRがマイナスの場合
// [x] 1-21. 400: バリデーションエラー：dqrValue.GeRが上限値を超える場合
// [x] 1-22. 400: バリデーションエラー：dqrValue.TiRがマイナスの場合
// [x] 1-23. 400: バリデーションエラー：dqrValue.TiRが上限値を超える場合
// [x] 1-24. 400: バリデーションエラー：cfpTypeとdqrTypeが指定の組み合わせと一致していない場合
// [x] 1-25. 400: バリデーションエラー：dqrTypeに値が入っているが、対応するcfpTypeの値が1つも入っていない場合
// [x] 1-26. 400: バリデーションエラー：※ケース除去 cfpTypeに値が入っているが、対応するdqrTypeの値が1つも入っていない場合
// [x] 1-27. 400: バリデーションエラー：dqrTypeが同じdqrValueでDQRの値が一致しない場合
// [x] 1-28. 400: バリデーションエラー：複合メッセージの確認
// [x] 1-29. 400: バリデーションエラー：ghgEmissionが小数点以下第5位までの値でない場合
// [x] 1-30. 400: バリデーションエラー：dqrValue.TeRが小数点以下第5位までの値でない場合
// [x] 1-31. 400: バリデーションエラー：dqrValue.GeRが小数点以下第5位までの値でない場合
// [x] 1-32. 400: バリデーションエラー：dqrValue.TiRが小数点以下第5位までの値でない場合
// [x] 1-33. 400: バリデーションエラー：ghgDeclaredUnitがstring形式でない
// [x] 1-34. 400: バリデーションエラー：ghgEmissionがdouble形式でない
// [x] 1-35. 400: バリデーションエラー：dqrValue.TeRがdouble形式でない
// [x] 1-36. 400: バリデーションエラー：dqrValue.GeRがdouble形式でない
// [x] 1-37. 400: バリデーションエラー：dqrValue.TiRがdouble形式でない
// [x] 1-38. 500: システムエラー：更新処理エラー
// [x] 1-39. 500: システムエラー：更新処理エラー
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutCfp(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfp"

	tests := []struct {
		name        string
		makeInput   func(inputs traceability.PutCfpInputs) string
		receive     error
		expectError string
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdが含まれない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].TraceID = ""
				inputs[1].TraceID = ""
				inputs[2].TraceID = ""
				inputs[3].TraceID = ""
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (traceId: cannot be blank.); cfpModel[1]: (traceId: cannot be blank.); cfpModel[2]: (traceId: cannot be blank.); cfpModel[3]: (traceId: cannot be blank.).",
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].TraceID = "NOT_UUID_FORMAT"
				inputs[1].TraceID = "NOT_UUID_FORMAT"
				inputs[2].TraceID = "NOT_UUID_FORMAT"
				inputs[3].TraceID = "NOT_UUID_FORMAT"
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (traceId: invalid UUID.); cfpModel[1]: (traceId: invalid UUID.); cfpModel[2]: (traceId: invalid UUID.); cfpModel[3]: (traceId: invalid UUID.).",
		},
		{
			name: "1-3. 400: バリデーションエラー：cfpIdがUUID形式ではない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpID = common.StringPtr("NOT_UUID_FORMAT")
				inputs[1].CfpID = common.StringPtr("NOT_UUID_FORMAT")
				inputs[2].CfpID = common.StringPtr("NOT_UUID_FORMAT")
				inputs[3].CfpID = common.StringPtr("NOT_UUID_FORMAT")
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (cfpId: invalid UUID.); cfpModel[1]: (cfpId: invalid UUID.); cfpModel[2]: (cfpId: invalid UUID.); cfpModel[3]: (cfpId: invalid UUID.).",
		},
		{
			name: "1-4. 400: バリデーションエラー：ghgEmissionがマイナスの場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgEmission = common.Float64Ptr(-1.2)
				inputs[0].DqrValue.TeR = common.Float64Ptr(0)
				inputs[0].DqrValue.GeR = common.Float64Ptr(0)
				inputs[0].DqrValue.TiR = common.Float64Ptr(0)
				inputs[2].GhgEmission = common.Float64Ptr(-1.2)
				inputs[2].DqrValue.TeR = common.Float64Ptr(0)
				inputs[2].DqrValue.GeR = common.Float64Ptr(0)
				inputs[2].DqrValue.TiR = common.Float64Ptr(0)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfpModel[2]: (ghgEmission: must be no less than 0.).",
		},
		{
			name: "1-5. 400: バリデーションエラー：ghgEmissionが上限値を超える場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgEmission = common.Float64Ptr(100000)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no greater than 99999.99999.).",
		},
		{
			name: "1-6. 400: バリデーションエラー：ghgDeclaredUnitが含まれない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgDeclaredUnit = ""
				inputs[1].GhgDeclaredUnit = ""
				inputs[2].GhgDeclaredUnit = ""
				inputs[3].GhgDeclaredUnit = ""
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgDeclaredUnit: cannot be blank.); cfpModel[1]: (ghgDeclaredUnit: cannot be blank.); cfpModel[2]: (ghgDeclaredUnit: cannot be blank.); cfpModel[3]: (ghgDeclaredUnit: cannot be blank.).",
		},
		{
			name: "1-7. 400: バリデーションエラー：ghgDeclaredUnitが指定のEnum以外の場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgDeclaredUnit = f.InvalidEnum
				inputs[1].GhgDeclaredUnit = f.InvalidEnum
				inputs[2].GhgDeclaredUnit = f.InvalidEnum
				inputs[3].GhgDeclaredUnit = f.InvalidEnum
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.); cfpModel[1]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.); cfpModel[2]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.); cfpModel[3]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.).",
		},
		{
			name: "1-8. 400: バリデーションエラー：cfpTypeが含まれない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpType = ""
				inputs[1].CfpType = ""
				inputs[2].CfpType = ""
				inputs[3].CfpType = ""
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (cfpType: cannot be blank.); cfpModel[1]: (cfpType: cannot be blank.); cfpModel[2]: (cfpType: cannot be blank.); cfpModel[3]: (cfpType: cannot be blank.); cfpType preProduction is insufficient; cfpType mainProduction is insufficient; cfpType preComponent is insufficient; cfpType mainComponent is insufficient.",
		},
		{
			name: "1-9. 400: バリデーションエラー：cfpTypeが指定のEnumではない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpType = "InvalidCfpType"
				inputs[1].CfpType = "InvalidCfpType"
				inputs[2].CfpType = "InvalidCfpType"
				inputs[3].CfpType = "InvalidCfpType"
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpModel[1]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpModel[2]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpModel[3]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpType preProduction is insufficient; cfpType mainProduction is insufficient; cfpType preComponent is insufficient; cfpType mainComponent is insufficient.",
		},
		{
			name: "1-10. 400: バリデーションエラー：cfpの入力が4つでない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs = traceability.PutCfpInputs{inputs[0]}
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfp models must be 4 elements",
		},
		{
			name: "1-11. 400: バリデーションエラー：cfpTypeの４つが揃っていない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpType = traceability.CfpTypePreComponent
				inputs[0].DqrType = traceability.DqrTypePreProcessing
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpType mainProduction is insufficient",
		},
		{
			name: "1-12. 400: バリデーションエラー：cfpIdが一致していない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpID = &f.CfpId2
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same cfpId",
		},
		{
			name: "1-13. 400: バリデーションエラー：traceIdが一致していない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].TraceID = f.TraceId2
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same traceId",
		},
		{
			name: "1-14. 400: バリデーションエラー：ghgDeclaredUnitが一致していない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgDeclaredUnit = f.GhgDeclaredUnit2
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same ghgDeclaredUnit",
		},
		{
			name: "1-16. 400: バリデーションエラー：dqrTypeが含まれない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrType = ""
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrType: cannot be blank.; ensure the combination of cfpType and dqrType is correct.).",
		},
		{
			name: "1-17. 400: バリデーションエラー：dqrTypeが指定のEnumではない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrType = "InvalidDqrType"
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrType: cannot be allowed 'InvalidDqrType'.; ensure the combination of cfpType and dqrType is correct.).",
		},
		{
			name: "1-18. 400: バリデーションエラー：dqrValue.TeRがマイナスの場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TeR = common.Float64Ptr(-1.2)
				inputs[2].DqrValue.TeR = common.Float64Ptr(-1.2)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TeR: must be no less than 0.).); cfpModel[2]: (dqrValue: (TeR: must be no less than 0.).).",
		},
		{
			name: "1-19. 400: バリデーションエラー：dqrValue.TeRが上限値を超える場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TeR = common.Float64Ptr(100000)
				inputs[2].DqrValue.TeR = common.Float64Ptr(100000)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TeR: must be no greater than 99999.99999.).); cfpModel[2]: (dqrValue: (TeR: must be no greater than 99999.99999.).).",
		},
		{
			name: "1-20. 400: バリデーションエラー：dqrValue.GeRがマイナスの場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.GeR = common.Float64Ptr(-1.2)
				inputs[2].DqrValue.GeR = common.Float64Ptr(-1.2)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (GeR: must be no less than 0.).); cfpModel[2]: (dqrValue: (GeR: must be no less than 0.).).",
		},
		{
			name: "1-21. 400: バリデーションエラー：dqrValue.GeRが上限値を超える場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.GeR = common.Float64Ptr(100000)
				inputs[2].DqrValue.GeR = common.Float64Ptr(100000)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (GeR: must be no greater than 99999.99999.).); cfpModel[2]: (dqrValue: (GeR: must be no greater than 99999.99999.).).",
		},
		{
			name: "1-22. 400: バリデーションエラー：dqrValue.TiRがマイナスの場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TiR = common.Float64Ptr(-1.2)
				inputs[2].DqrValue.TiR = common.Float64Ptr(-1.2)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TiR: must be no less than 0.).); cfpModel[2]: (dqrValue: (TiR: must be no less than 0.).).",
		},
		{
			name: "1-23. 400: バリデーションエラー：dqrValue.TiRが上限値を超える場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TiR = common.Float64Ptr(100000)
				inputs[2].DqrValue.TiR = common.Float64Ptr(100000)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TiR: must be no greater than 99999.99999.).); cfpModel[2]: (dqrValue: (TiR: must be no greater than 99999.99999.).).",
		},
		{
			name: "1-24. 400: バリデーションエラー：cfpTypeとdqrTypeが指定の組み合わせと一致していない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpType = traceability.CfpTypeMainProduction
				inputs[0].DqrType = traceability.DqrTypePreProcessing
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ensure the combination of cfpType and dqrType is correct.).",
		},
		{
			name: "1-25. 400: バリデーションエラー：dqrTypeに値が入っているが、対応するcfpTypeの値が1つも入っていない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgEmission = common.Float64Ptr(0)
				inputs[2].GhgEmission = common.Float64Ptr(0)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, set ghgEmission to a value greater than 0 for cfpType of mainProduction or mainComponent.",
		},
		{
			name: "1-27. 400: バリデーションエラー：dqrTypeが同じdqrValueでDQRの値が一致しない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TeR = common.Float64Ptr(0.2)
				inputs[0].DqrValue.GeR = common.Float64Ptr(0.2)
				inputs[0].DqrValue.TiR = common.Float64Ptr(0.2)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, different dqrValues are set for the same dqrType.",
		},
		{
			name: "1-28. 400: バリデーションエラー：複合メッセージの確認",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				// 1-3と1-5のエラーが混在
				inputs[0].TraceID = ""
				inputs[1].TraceID = ""
				inputs[2].TraceID = ""
				inputs[3].TraceID = ""
				inputs[0].GhgEmission = common.Float64Ptr(100000)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no greater than 99999.99999; traceId: cannot be blank.); cfpModel[1]: (traceId: cannot be blank.); cfpModel[2]: (traceId: cannot be blank.); cfpModel[3]: (traceId: cannot be blank.).",
		},
		{
			name: "1-29. 400: バリデーションエラー：ghgEmissionが小数点以下第5位までの値でない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].GhgEmission = common.Float64Ptr(0.111111)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be a value up to the 5th decimal place.).",
		},
		{
			name: "1-30. 400: バリデーションエラー：dqrValue.TeRが小数点以下第5位までの値でない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TeR = common.Float64Ptr(0.000001)
				inputs[2].DqrValue.TeR = common.Float64Ptr(0.000001)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TeR: must be a value up to the 5th decimal place.).); cfpModel[2]: (dqrValue: (TeR: must be a value up to the 5th decimal place.).).",
		},
		{
			name: "1-31. 400: バリデーションエラー：dqrValue.GeRが小数点以下第5位までの値でない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.GeR = common.Float64Ptr(0.000001)
				inputs[2].DqrValue.GeR = common.Float64Ptr(0.000001)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (GeR: must be a value up to the 5th decimal place.).); cfpModel[2]: (dqrValue: (GeR: must be a value up to the 5th decimal place.).).",
		},
		{
			name: "1-32. 400: バリデーションエラー：dqrValue.TiRが小数点以下第5位までの値でない場合",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].DqrValue.TiR = common.Float64Ptr(0.000001)
				inputs[2].DqrValue.TiR = common.Float64Ptr(0.000001)
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TiR: must be a value up to the 5th decimal place.).); cfpModel[2]: (dqrValue: (TiR: must be a value up to the 5th decimal place.).).",
		},
		{
			name: "1-33. 400: バリデーションエラー：ghgDeclaredUnitがstring形式でない",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputJSON, _ := json.Marshal(inputs)
				inputJsonStr := string(inputJSON)

				var inputJsonMapArr []map[string]interface{}
				// nolint
				json.Unmarshal([]byte(inputJsonStr), &inputJsonMapArr)

				inputJsonMapArr[0]["ghgDeclaredUnit"] = 1
				testInputJson, _ := json.Marshal(inputJsonMapArr)
				return string(testInputJson)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, ghgDeclaredUnit: Unmarshal type error: expected=string, got=number.",
		},
		{
			name: "1-34. 400: バリデーションエラー：ghgEmissionがdouble形式でない",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputJSON, _ := json.Marshal(inputs)
				inputJsonStr := string(inputJSON)

				var inputJsonMapArr []map[string]interface{}
				// nolint
				json.Unmarshal([]byte(inputJsonStr), &inputJsonMapArr)
				inputJsonMapArr[0]["ghgEmission"] = "hoge"
				testInputJson, _ := json.Marshal(inputJsonMapArr)
				return string(testInputJson)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, ghgEmission: Unmarshal type error: expected=float64, got=string.",
		},
		{
			name: "1-35. 400: バリデーションエラー：dqrValue.TeRがdouble形式でない",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputJSON, _ := json.Marshal(inputs)
				inputJsonStr := string(inputJSON)

				var inputJsonMapArr []map[string]interface{}
				// nolint
				json.Unmarshal([]byte(inputJsonStr), &inputJsonMapArr)
				inputJsonMapArr[0]["dqrValue"].(map[string]interface{})["TeR"] = "hoge"
				testInputJson, _ := json.Marshal(inputJsonMapArr)
				return string(testInputJson)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, dqrValue.TeR: Unmarshal type error: expected=float64, got=string.",
		},
		{
			name: "1-36. 400: バリデーションエラー：dqrValue.GeRがdouble形式でない",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputJSON, _ := json.Marshal(inputs)
				inputJsonStr := string(inputJSON)

				var inputJsonMapArr []map[string]interface{}
				// nolint
				json.Unmarshal([]byte(inputJsonStr), &inputJsonMapArr)
				inputJsonMapArr[0]["dqrValue"].(map[string]interface{})["GeR"] = "hoge"
				testInputJson, _ := json.Marshal(inputJsonMapArr)
				return string(testInputJson)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, dqrValue.GeR: Unmarshal type error: expected=float64, got=string.",
		},
		{
			name: "1-37. 400: バリデーションエラー：dqrValue.TiRがdouble形式でない",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputJSON, _ := json.Marshal(inputs)
				inputJsonStr := string(inputJSON)

				var inputJsonMapArr []map[string]interface{}
				// nolint
				json.Unmarshal([]byte(inputJsonStr), &inputJsonMapArr)
				inputJsonMapArr[0]["dqrValue"].(map[string]interface{})["TiR"] = "hoge"
				testInputJson, _ := json.Marshal(inputJsonMapArr)
				return string(testInputJson)
			},
			expectError: "code=400, message={[dataspace] BadRequest Validation failed, dqrValue.TiR: Unmarshal type error: expected=float64, got=string.",
		},
		{
			name: "1-38. 500: システムエラー：更新処理エラー",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpID = nil
				inputs[1].CfpID = nil
				inputs[2].CfpID = nil
				inputs[3].CfpID = nil
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			receive:     common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
		{
			name: "1-39. 500: システムエラー：更新処理エラー",
			makeInput: func(inputs traceability.PutCfpInputs) string {
				inputs[0].CfpID = nil
				inputs[1].CfpID = nil
				inputs[2].CfpID = nil
				inputs[3].CfpID = nil
				inputJSON, _ := json.Marshal(inputs)
				return string(inputJSON)
			},
			receive:     fmt.Errorf("Internal Server Error"),
			expectError: "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
		},
	}

	for _, tc := range tests {
		tc := tc
		tt.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			inputsBase := f.NewPutCfpInputs()        // 元のinput構造体を準備
			inputJSONStr := tc.makeInput(inputsBase) // テストケースに応じてinputを変更

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(string(inputJSONStr)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)
			c.SetPath(endPoint)
			c.Set("operatorID", f.OperatorId)

			cfpUsecase := new(mocks.ICfpUsecase)
			cfpUsecase.On("PutCfp", mock.Anything, mock.Anything, mock.Anything).Return([]traceability.CfpModel{}, tc.receive)
			cfpHandler := handler.NewCfpHandler(cfpUsecase)

			err := cfpHandler.PutCfp(c)
			if assert.Error(t, err) {
				assert.ErrorContains(t, err, tc.expectError)
			}
		})
	}
}
