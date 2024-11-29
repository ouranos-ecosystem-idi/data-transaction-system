package handler_test

import (
	"bytes"
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
			input := traceability.GetCfpInput{
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

			// レスポンスヘッダにX-Trackが含まれているかチェック
			_, ok := rec.Header()["X-Track"]
			assert.True(t, ok, "Header should have 'X-Track' key")
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
		expectStatus      int
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
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
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
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
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
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
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
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
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
			expectError:  "code=400, message={[dataspace] BadRequest Invalid request parameters, traceIds: Unexpected query parameter",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 400: バリデーションエラー：operatorIdがUUID形式ではない場合",
			modifyQueryParams: func(q url.Values) {
				q.Set("traceId", f.TraceId)
			},
			modifyContexts: func(c echo.Context) {
				c.Set("operatorID", "invalid")
			},
			expectError:  "code=400, message={[auth] BadRequest Invalid or expired token",
			expectStatus: http.StatusBadRequest,
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
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
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
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
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
			e.HTTPErrorHandler(err, c)
			if assert.Error(t, err) {
				assert.Equal(t, test.expectStatus, rec.Code)
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
		inputFunc    func() traceability.PutCfpInputs
		expectStatus int
	}{
		{
			name: "1-1. 201: 正常系：新規作成",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].CfpID = nil
				}
				return cfpInputs
			},
			expectStatus: http.StatusCreated,
		},
		{
			name: "1-2. 201: 正常系：更新",
			inputFunc: func() traceability.PutCfpInputs {
				return f.NewPutCfpInputs2()
			},
			expectStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(test.name, func(t *testing.T) {
			t.Parallel()

			inputs := test.inputFunc()
			inputJSON, _ := json.Marshal(inputs)

			q := make(url.Values)
			q.Set("dataTarget", dataTarget)

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), bytes.NewBuffer(inputJSON))
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
			cfpUsecase.On("PutCfp", c, inputs, f.OperatorId).Return(responseCfpModels, common.ResponseHeaders{}, nil)

			err := cfpHandler.PutCfp(c)
			if assert.NoError(t, err) {
				assert.Equal(t, test.expectStatus, rec.Code)
				cfpUsecase.AssertExpectations(t)
			}

			// レスポンスヘッダにX-Trackが含まれているかチェック
			_, ok := rec.Header()["X-Track"]
			assert.True(t, ok, "Header should have 'X-Track' key")
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
		name             string
		inputFunc        func() traceability.PutCfpInputs
		invalidInputFunc func() []interface{}
		receive          error
		expectError      string
		expectStatus     int
	}{
		{
			name: "1-1. 400: バリデーションエラー：traceIdが含まれない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].TraceID = ""
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (traceId: cannot be blank.); cfpModel[1]: (traceId: cannot be blank.); cfpModel[2]: (traceId: cannot be blank.); cfpModel[3]: (traceId: cannot be blank.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：traceIdがUUID形式ではない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].TraceID = "NOT_UUID_FORMAT"
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (traceId: invalid UUID.); cfpModel[1]: (traceId: invalid UUID.); cfpModel[2]: (traceId: invalid UUID.); cfpModel[3]: (traceId: invalid UUID.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 400: バリデーションエラー：cfpIdがUUID形式ではない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].CfpID = common.StringPtr("NOT_UUID_FORMAT")
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (cfpId: invalid UUID.); cfpModel[1]: (cfpId: invalid UUID.); cfpModel[2]: (cfpId: invalid UUID.); cfpModel[3]: (cfpId: invalid UUID.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-4. 400: バリデーションエラー：ghgEmissionがマイナスの場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputs[2].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(0)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfpModel[2]: (ghgEmission: must be no less than 0.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-5. 400: バリデーションエラー：ghgEmissionが上限値を超える場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(100000)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no greater than 99999.99999.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-6. 400: バリデーションエラー：ghgDeclaredUnitが含まれない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].GhgDeclaredUnit = ""
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgDeclaredUnit: cannot be blank.); cfpModel[1]: (ghgDeclaredUnit: cannot be blank.); cfpModel[2]: (ghgDeclaredUnit: cannot be blank.); cfpModel[3]: (ghgDeclaredUnit: cannot be blank.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-7. 400: バリデーションエラー：ghgDeclaredUnitが指定のEnum以外の場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].GhgDeclaredUnit = f.InvalidEnum
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.); cfpModel[1]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.); cfpModel[2]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.); cfpModel[3]: (ghgDeclaredUnit: cannot be allowed 'invalid_enum'.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-8. 400: バリデーションエラー：cfpTypeが含まれない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].CfpType = ""
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (cfpType: cannot be blank.); cfpModel[1]: (cfpType: cannot be blank.); cfpModel[2]: (cfpType: cannot be blank.); cfpModel[3]: (cfpType: cannot be blank.); cfpType preProduction is insufficient; cfpType mainProduction is insufficient; cfpType preComponent is insufficient; cfpType mainComponent is insufficient.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-9. 400: バリデーションエラー：cfpTypeが指定のEnumではない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].CfpType = "InvalidCfpType"
				}
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpModel[1]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpModel[2]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpModel[3]: (cfpType: cannot be allowed 'InvalidCfpType'.); cfpType preProduction is insufficient; cfpType mainProduction is insufficient; cfpType preComponent is insufficient; cfpType mainComponent is insufficient.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-10. 400: バリデーションエラー：cfpの入力が4つでない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInput := traceability.PutCfpInputs{cfpInputs[0]}
				return cfpInput
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfp models must be 4 elements",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-11. 400: バリデーションエラー：cfpTypeの４つが揃っていない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpType = traceability.CfpTypePreComponent
				cfpInputs[0].DqrType = traceability.DqrTypePreProcessing
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpType mainProduction is insufficient",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-12. 400: バリデーションエラー：cfpIdが一致していない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpID = &f.CfpId2
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same cfpId",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-13. 400: バリデーションエラー：traceIdが一致していない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].TraceID = f.TraceId2
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same traceId",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-14. 400: バリデーションエラー：ghgDeclaredUnitが一致していない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgDeclaredUnit = f.GhgDeclaredUnit2
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same ghgDeclaredUnit",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-16. 400: バリデーションエラー：dqrTypeが含まれない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrType = ""
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrType: cannot be blank.; ensure the combination of cfpType and dqrType is correct.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-17. 400: バリデーションエラー：dqrTypeが指定のEnumではない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrType = "InvalidDqrType"
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrType: cannot be allowed 'InvalidDqrType'.; ensure the combination of cfpType and dqrType is correct.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-18. 400: バリデーションエラー：dqrValue.TeRがマイナスの場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(-1.2)
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(-1.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TeR: must be no less than 0.).); cfpModel[2]: (dqrValue: (TeR: must be no less than 0.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-19. 400: バリデーションエラー：dqrValue.TeRが上限値を超える場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(100000)
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(100000)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TeR: must be no greater than 99999.99999.).); cfpModel[2]: (dqrValue: (TeR: must be no greater than 99999.99999.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-20. 400: バリデーションエラー：dqrValue.GeRがマイナスの場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(-1.2)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(-1.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (GeR: must be no less than 0.).); cfpModel[2]: (dqrValue: (GeR: must be no less than 0.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-21. 400: バリデーションエラー：dqrValue.GeRが上限値を超える場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(100000)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(100000)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (GeR: must be no greater than 99999.99999.).); cfpModel[2]: (dqrValue: (GeR: must be no greater than 99999.99999.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-22. 400: バリデーションエラー：dqrValue.TiRがマイナスの場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(-1.2)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(-1.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TiR: must be no less than 0.).); cfpModel[2]: (dqrValue: (TiR: must be no less than 0.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-23. 400: バリデーションエラー：dqrValue.TiRが上限値を超える場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(100000)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(100000)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TiR: must be no greater than 99999.99999.).); cfpModel[2]: (dqrValue: (TiR: must be no greater than 99999.99999.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-24. 400: バリデーションエラー：cfpTypeとdqrTypeが指定の組み合わせと一致していない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpType = traceability.CfpTypeMainProduction
				cfpInputs[0].DqrType = traceability.DqrTypePreProcessing
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ensure the combination of cfpType and dqrType is correct.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-25. 400: バリデーションエラー：dqrTypeに値が入っているが、対応するcfpTypeの値が1つも入っていない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(0)
				cfpInputs[2].GhgEmission = common.Float64Ptr(0)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, set ghgEmission to a value greater than 0 for cfpType of mainProduction or mainComponent.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-27. 400: バリデーションエラー：dqrTypeが同じdqrValueでDQRの値が一致しない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(0.2)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(0.2)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(0.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, different dqrValues are set for the same dqrType.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-28. 400: バリデーションエラー：複合メッセージの確認",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				// 1-3と1-5のエラーが混在
				for i := range cfpInputs {
					cfpInputs[i].TraceID = ""
				}
				cfpInputs[0].GhgEmission = common.Float64Ptr(100000)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no greater than 99999.99999; traceId: cannot be blank.); cfpModel[1]: (traceId: cannot be blank.); cfpModel[2]: (traceId: cannot be blank.); cfpModel[3]: (traceId: cannot be blank.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-29. 400: バリデーションエラー：ghgEmissionが小数点以下第5位までの値でない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(0.111111)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be a value up to the 5th decimal place.).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-30. 400: バリデーションエラー：dqrValue.TeRが小数点以下第5位までの値でない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0.111111)
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(0.111111)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TeR: must be a value up to the 5th decimal place.).); cfpModel[2]: (dqrValue: (TeR: must be a value up to the 5th decimal place.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-31. 400: バリデーションエラー：dqrValue.GeRが小数点以下第5位までの値でない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0.111111)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(0.111111)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (GeR: must be a value up to the 5th decimal place.).); cfpModel[2]: (dqrValue: (GeR: must be a value up to the 5th decimal place.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-32. 400: バリデーションエラー：dqrValue.TiRが小数点以下第5位までの値でない場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0.000001)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(0.000001)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (dqrValue: (TiR: must be a value up to the 5th decimal place.).); cfpModel[2]: (dqrValue: (TiR: must be a value up to the 5th decimal place.).).",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-33. 400: バリデーションエラー：ghgDeclaredUnitがstring形式でない",
			invalidInputFunc: func() []interface{} {
				inputs := f.NewPutCfpInterface()
				inputs[0].(map[string]interface{})["ghgDeclaredUnit"] = 1
				return inputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, ghgDeclaredUnit: Unmarshal type error: expected=string, got=number.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-34. 400: バリデーションエラー：ghgEmissionがdouble形式でない",
			invalidInputFunc: func() []interface{} {
				inputs := f.NewPutCfpInterface()
				inputs[0].(map[string]interface{})["ghgEmission"] = "hoge"
				return inputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, ghgEmission: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-35. 400: バリデーションエラー：dqrValue.TeRがdouble形式でない",
			invalidInputFunc: func() []interface{} {
				inputs := f.NewPutCfpInterface()
				inputs[0].(map[string]interface{})["dqrValue"].(map[string]interface{})["TeR"] = "hoge"
				return inputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, dqrValue.TeR: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-36. 400: バリデーションエラー：dqrValue.GeRがdouble形式でない",
			invalidInputFunc: func() []interface{} {
				inputs := f.NewPutCfpInterface()
				inputs[0].(map[string]interface{})["dqrValue"].(map[string]interface{})["GeR"] = "hoge"
				return inputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, dqrValue.GeR: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-37. 400: バリデーションエラー：dqrValue.TiRがdouble形式でない",
			invalidInputFunc: func() []interface{} {
				inputs := f.NewPutCfpInterface()
				inputs[0].(map[string]interface{})["dqrValue"].(map[string]interface{})["TiR"] = "hoge"
				return inputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, dqrValue.TiR: Unmarshal type error: expected=float64, got=string.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-38. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].CfpID = nil
				}
				return cfpInputs
			},
			receive:      common.NewCustomError(common.CustomErrorCode500, "Unexpected error occurred", common.StringPtr(""), common.HTTPErrorSourceDataspace),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-39. 500: システムエラー：更新処理エラー",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				for i := range cfpInputs {
					cfpInputs[i].CfpID = nil
				}
				return cfpInputs
			},
			receive:      fmt.Errorf("Internal Server Error"),
			expectError:  "code=500, message={[dataspace] InternalServerError Unexpected error occurred",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-40. 400: バリデーションエラー：1-4と1-10が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfp models must be 4 elements",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-41. 400: バリデーションエラー：1-4と1-11が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpType = traceability.CfpTypePreComponent
				cfpInputs[0].DqrType = traceability.DqrTypePreProcessing
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(0)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfpType mainProduction is insufficient",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-42. 400: バリデーションエラー：1-4と1-12が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[2].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputs[0].CfpID = &f.CfpId2
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); ensure all objects have the same cfpId",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-43. 400: バリデーションエラー：1-4と1-27が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); different dqrValues are set for the same dqrType.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-44. 400: バリデーションエラー：1-10と1-11が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpType = traceability.CfpTypePreComponent
				cfpInputs[0].DqrType = traceability.DqrTypePreProcessing
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0], cfpInputs[1], cfpInputs[2]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfp models must be 4 elements; cfpType mainProduction is insufficient",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-45. 400: バリデーションエラー：1-10と1-12が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpID = &f.CfpId2
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0], cfpInputs[1], cfpInputs[2]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfp models must be 4 elements; cfpType preComponent is insufficient; ensure all objects have the same cfpId",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-46. 400: バリデーションエラー：1-10と1-27が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0.2)
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0], cfpInputs[1], cfpInputs[2]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfp models must be 4 elements; cfpType preComponent is insufficient; different dqrValues are set for the same dqrType.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-47. 400: バリデーションエラー：1-11と1-12が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpType = traceability.CfpTypePreComponent
				cfpInputs[0].DqrType = traceability.DqrTypePreProcessing
				cfpInputs[0].CfpID = &f.CfpId2
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpType mainProduction is insufficient; ensure all objects have the same cfpId",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-48. 400: バリデーションエラー：1-11と1-27が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpType = traceability.CfpTypePreComponent
				cfpInputs[0].DqrType = traceability.DqrTypePreProcessing
				cfpInputs[0].GhgEmission = common.Float64Ptr(1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpType mainProduction is insufficient; different dqrValues are set for the same dqrType.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-49. 400: バリデーションエラー：1-12と1-27が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpID = &f.CfpId2
				cfpInputs[0].GhgEmission = common.Float64Ptr(1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0.2)
				return cfpInputs
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, ensure all objects have the same cfpId; different dqrValues are set for the same dqrType.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-50. 400: バリデーションエラー：1-4と1-10と1-11が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0], cfpInputs[1], cfpInputs[2]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfp models must be 4 elements; cfpType preComponent is insufficient;",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-51. 400: バリデーションエラー：1-4と1-10と1-11と1-12が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpID = &f.CfpId2
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0)
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0], cfpInputs[1], cfpInputs[2]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfp models must be 4 elements; cfpType preComponent is insufficient; ensure all objects have the same cfpId",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-52. 400: バリデーションエラー：1-4と1-10と1-11と1-12と1-27が同時に発生する場合",
			inputFunc: func() traceability.PutCfpInputs {
				cfpInputs := f.NewPutCfpInputs2()
				cfpInputs[0].CfpID = &f.CfpId2
				cfpInputs[0].GhgEmission = common.Float64Ptr(-1.2)
				cfpInputs[0].DqrValue.TeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.GeR = common.Float64Ptr(0.2)
				cfpInputs[0].DqrValue.TiR = common.Float64Ptr(0.2)
				cfpInputsPart := traceability.PutCfpInputs{cfpInputs[0], cfpInputs[1], cfpInputs[2]}
				return cfpInputsPart
			},
			expectError:  "code=400, message={[dataspace] BadRequest Validation failed, cfpModel[0]: (ghgEmission: must be no less than 0.); cfp models must be 4 elements; cfpType preComponent is insufficient; ensure all objects have the same cfpId; different dqrValues are set for the same dqrType.",
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

			cfpUsecase := new(mocks.ICfpUsecase)
			cfpUsecase.On("PutCfp", mock.Anything, mock.Anything, mock.Anything).Return([]traceability.CfpModel{}, common.ResponseHeaders{}, tc.receive)
			cfpHandler := handler.NewCfpHandler(cfpUsecase)

			err := cfpHandler.PutCfp(c)
			e.HTTPErrorHandler(err, c)
			if assert.Error(t, err) {
				assert.Equal(t, tc.expectStatus, rec.Code)
				assert.ErrorContains(t, err, tc.expectError)
			}
		})
	}
}
