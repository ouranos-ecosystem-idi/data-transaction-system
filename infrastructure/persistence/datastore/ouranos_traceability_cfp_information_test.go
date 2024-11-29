package datastore_test

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/infrastructure/persistence/datastore"
	f "data-spaces-backend/test/fixtures"
	testhelper "data-spaces-backend/test/test_helper"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// /////////////////////////////////////////////////////////////////////////////////
// CfpInformation GetCFPInformation テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_CFPInformation_GetCFPInformation(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.CfpEntityModel
	}{
		{
			name:  "1-1: 正常系 取得成功の場合",
			input: f.TraceID4,
			expect: traceability.CfpEntityModel{
				CfpID:              common.UUIDPtr(uuid.MustParse(f.CfpId)),
				TraceID:            uuid.MustParse(f.TraceID4),
				GhgEmission:        common.Float64Ptr(0.2),
				GhgDeclaredUnit:    f.GhgDeclaredUnit2,
				CfpCertificateList: nil,
				CfpType:            traceability.CfpTypePreProduction.ToString(),
				DqrType:            traceability.DqrTypePreProcessing.ToString(),
				TeR:                common.Float64Ptr(1.2),
				GeR:                common.Float64Ptr(2.2),
				TiR:                common.Float64Ptr(3.2),
				DeletedAt:          gorm.DeletedAt{},
				CreatedAt:          f.DummyTime,
				CreatedUserId:      "seed",
				UpdatedAt:          f.DummyTime,
				UpdatedUserId:      "seed",
			},
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
				actual, _ := r.GetCFPInformation(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}

}

// /////////////////////////////////////////////////////////////////////////////////
// CfpInformation GetCFPInformation テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_CFPInformation_GetCFPInformation_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：0件の場合",
			input:     f.NotExistID,
			dropQuery: "",
			expect:    fmt.Errorf("record not found"),
		},
		{
			name:      "2-2: 異常系：取得失敗の場合",
			input:     f.TraceID4,
			dropQuery: "DROP TABLE IF EXISTS cfp_infomation",
			expect:    fmt.Errorf("no such table: cfp_infomation"),
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
				if test.name == "2-2: 異常系：取得失敗の場合" {
					err = db.Exec(test.dropQuery).Error
					if err != nil {
						assert.Fail(t, "Errors occured by deleting DB")
					}
				}
				r := datastore.NewOuranosRepository(db)
				_, err = r.GetCFPInformation(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// CfpInformation DeleteCFPInformation テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：削除成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_CFPInformation_DeleteCFPInformation(tt *testing.T) {

	tests := []struct {
		name       string
		input      string
		checkQuery string
		before     int
		after      int
	}{
		{
			name:       "1-1: 正常系：削除成功の場合",
			input:      f.CfpId,
			checkQuery: "SELECT COUNT(*) FROM cfp_infomation WHERE cfp_id = ?",
			before:     1,
			after:      0,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				db, err := testhelper.NewMockDB()
				if err != nil {
					assert.Fail(t, "Errors occured by creating Mock DB")
				}
				r := datastore.NewOuranosRepository(db)
				var actualCount int
				db.Raw(test.checkQuery, test.input).Scan(&actualCount)
				assert.Equal(t, test.before, actualCount)
				err = r.DeleteCFPInformation(test.input)
				if assert.NoError(t, err) {
					db.Raw(test.checkQuery, test.input).Scan(&actualCount)
					assert.Equal(t, test.after, actualCount)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// CfpInformation DeleteCFPInformation テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：削除失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_CFPInformation_DeleteCFPInformation_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：削除失敗の場合",
			input:     f.NotExistID,
			dropQuery: "DROP TABLE IF EXISTS cfp_infomation",
			expect:    fmt.Errorf("failed to physically delete record from table cfp_infomation: no such table: cfp_infomation"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				db, err := testhelper.NewMockDB()
				if err != nil {
					assert.Fail(t, "Errors occured by creating Mock DB")
				}
				err = db.Exec(test.dropQuery).Error
				if err != nil {
					assert.Fail(t, "Errors occured by deleting DB")
				}
				r := datastore.NewOuranosRepository(db)
				err = r.DeleteCFPInformation(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
