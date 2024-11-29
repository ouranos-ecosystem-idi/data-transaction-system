package usecase_test

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	f "data-spaces-backend/test/fixtures"
	mocks "data-spaces-backend/test/mock"
	"data-spaces-backend/usecase"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Get /api/v1/datatransport/cfpCertification テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常終了
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectUsecaseDatastore_GetCfpCertification(tt *testing.T) {

	var method = "GET"
	var endPoint = "/api/v1/datatransport"
	var dataTarget = "cfpCertification"

	dsResAll := traceability.CfpCertificationModels{
		{
			CfpCertificationID:          "d9a38406-cae2-4679-b052-15a75f5531c5",
			TraceID:                     "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
			CfpCertificationDescription: common.StringPtr("サンプル証明書"),
			CfpCertificationFileInfo: &[]traceability.CfpCertificationFileInfo{
				{
					OperatorID: "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					FileID:     "5c07e3e9-c0e5-4a1f-b6a5-78145f7d1855",
					FileName:   "B01_CFP.pdf",
				},
			},
		},
	}

	tests := []struct {
		name        string
		input       traceability.GetCfpCertificationInput
		receive     traceability.CfpCertificationModels
		expectData  traceability.CfpCertificationModels
		expectAfter *string
	}{
		{
			name:        "1-1. 200: 正常終了",
			input:       f.NewGetCfpCertificationInput(),
			receive:     dsResAll,
			expectData:  dsResAll,
			expectAfter: nil,
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

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				ouranosRepositoryMock := new(mocks.OuranosRepository)
				ouranosRepositoryMock.On("GetCFPCertifications", mock.Anything, mock.Anything).Return(test.receive, nil)

				usecase := usecase.NewCfpCertificationUsecase(ouranosRepositoryMock)

				actualRes, err := usecase.GetCfpCertification(c, test.input)
				// エラーが発生しないことを確認
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.ElementsMatch(t, test.expectData, actualRes, f.AssertMessage)
				}
			},
		)
	}
}
