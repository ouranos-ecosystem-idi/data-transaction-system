package datastore_test

import (
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/infrastructure/persistence/datastore"
	f "data-spaces-backend/test/fixtures"
	testhelper "data-spaces-backend/test/test_helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////////////////////////////////////////////////
// CfpCertification GetCFPCertifications テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_CFPCertification_GetCFPCertifications(tt *testing.T) {

	tests := []struct {
		name       string
		operatorID string
		traceID    string
		expect     traceability.CfpCertificationModels
	}{
		{
			name:       "1-1: 正常系：取得成功の場合",
			operatorID: f.OperatorID,
			traceID:    f.TraceID,
			expect:     f.CfpCertificationsData,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				db, err := testhelper.NewMockDB()
				if err != nil {
					assert.Fail(t, err.Error())
				}
				r := datastore.NewOuranosRepository(db)
				actual, err := r.GetCFPCertifications(test.operatorID, test.traceID)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}
