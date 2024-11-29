package datastore_test

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/infrastructure/persistence/datastore"
	f "data-spaces-backend/test/fixtures"
	testhelper "data-spaces-backend/test/test_helper"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// /////////////////////////////////////////////////////////////////////////////////
// Cfp BatchCreateCFP テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：登録成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_BatchCreateCFP(tt *testing.T) {

	tests := []struct {
		name   string
		input  traceability.CfpEntityModels
		expect traceability.CfpEntityModels
	}{
		{
			name:   "1-1: 正常系 登録成功の場合",
			input:  f.NewBatchCreateCFPInput(),
			expect: f.NewBatchCreateCFPInput(),
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
				actual, err := r.BatchCreateCFP(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp BatchCreateCFP テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-2. 異常系：登録失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_BatchCreateCFP_Abnormal(tt *testing.T) {
	tests := []struct {
		name      string
		input     traceability.CfpEntityModels
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 0件の場合",
			input:     traceability.CfpEntityModels{},
			dropQuery: "",
			expect:    fmt.Errorf("cfp entities is empty"),
		},
		{
			name:      "2-2: 異常系：登録失敗の場合",
			input:     f.NewBatchCreateCFPInput(),
			dropQuery: "DROP TABLE IF EXISTS cfp_infomation",
			expect:    fmt.Errorf("failed to insert cfp_infomation record: no such table: cfp_infomation"),
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
				if test.name == "2-2: 異常系：登録失敗の場合" {
					err = db.Exec(test.dropQuery).Error
					if err != nil {
						assert.Fail(t, "Errors occured by deleting DB")
					}
				}
				r := datastore.NewOuranosRepository(db)
				_, err = r.BatchCreateCFP(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp GetCFP テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_GetCFP(tt *testing.T) {

	tests := []struct {
		name         string
		inputCfpID   string
		inputCfpType string
		expect       traceability.CfpEntityModel
	}{
		{
			name:         "1-1: 正常系 取得成功の場合",
			inputCfpID:   f.CfpId,
			inputCfpType: "preProduction",
			expect: traceability.CfpEntityModel{
				CfpID:              common.UUIDPtr(uuid.MustParse(f.CfpId)),
				TraceID:            uuid.MustParse(f.TraceID4),
				GhgEmission:        common.Float64Ptr(0.2),
				GhgDeclaredUnit:    f.GhgDeclaredUnit2,
				CfpCertificateList: f.CfpCertificateList,
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
				actual, err := r.GetCFP(test.inputCfpID, test.inputCfpType)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp GetCFP テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_GetCFP_Abnormal(tt *testing.T) {

	tests := []struct {
		name         string
		inputCfpID   string
		inputCfpType string
		dropQuery    string
		expect       error
	}{
		{
			name:         "2-2: 異常系：取得失敗の場合",
			inputCfpID:   f.CfpId,
			inputCfpType: "preProduction",
			dropQuery:    "DROP TABLE IF EXISTS cfp_infomation",
			expect:       fmt.Errorf("no such table: cfp_infomation"),
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
				_, err = r.GetCFP(test.inputCfpID, test.inputCfpType)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp ListCFPsByTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_ListCFPsByTraceID(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.CfpEntityModels
	}{
		{
			name:  "1-1: 正常系：1件以上の場合",
			input: f.TraceID4,
			expect: traceability.CfpEntityModels{
				&traceability.CfpEntityModel{
					CfpID:              common.UUIDPtr(uuid.MustParse(f.CfpId)),
					TraceID:            uuid.MustParse(f.TraceID4),
					GhgEmission:        common.Float64Ptr(0.2),
					GhgDeclaredUnit:    f.GhgDeclaredUnit2,
					CfpCertificateList: f.CfpCertificateList,
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
		},
		{
			name:   "1-2: 正常系：0件の場合",
			input:  f.NotExistID,
			expect: traceability.CfpEntityModels{},
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
				actual, err := r.ListCFPsByTraceID(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect), len(actual))
					for i, data := range test.expect {
						assert.Equal(t, data, actual[i])
					}
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp ListCFPsByTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_ListCFPsByTraceID_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：取得失敗の場合",
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
				err = db.Exec(test.dropQuery).Error
				if err != nil {
					assert.Fail(t, "Errors occured by deleting DB")
				}
				r := datastore.NewOuranosRepository(db)
				_, err = r.ListCFPsByTraceID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp PutCFP テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：更新成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_PutCFP(tt *testing.T) {

	tests := []struct {
		name   string
		input  traceability.CfpEntityModel
		expect traceability.CfpEntityModel
	}{

		{
			name:   "1-1: 正常系 取得成功の場合",
			input:  f.NewPutCFPInput(),
			expect: f.NewPutCFPInput(),
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
				actual, err := r.PutCFP(test.input)
				if assert.NoError(t, err) {
					assert.WithinDuration(t, time.Now(), actual.UpdatedAt, 3*time.Second)
					test.expect.UpdatedAt = f.DummyTime
					actual.UpdatedAt = f.DummyTime
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Cfp PutCFP テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：更新失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Cfp_PutCFP_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     traceability.CfpEntityModel
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：更新失敗の場合",
			input:     f.NewPutCFPInput(),
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
					assert.Fail(t, "Errors occured by creating Mock DB")
				}
				err = db.Exec(test.dropQuery).Error
				if err != nil {
					assert.Fail(t, "Errors occured by deleting DB")
				}
				r := datastore.NewOuranosRepository(db)
				_, err = r.PutCFP(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
