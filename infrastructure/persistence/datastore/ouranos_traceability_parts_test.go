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
)

// /////////////////////////////////////////////////////////////////////////////////
// Parts ListParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// [x] 1-3. 正常系：nil許容項目がnilの場合
// [x] 1-4. 正常系：任意項目が未定義の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_ListParts(tt *testing.T) {

	tests := []struct {
		name   string
		input  func() traceability.GetPartsInput
		expect traceability.PartsModelEntities
	}{
		{
			name: "1-1: 正常系：1件以上の場合",
			input: func() traceability.GetPartsInput {
				i := f.NewGetPartsInput()
				return i
			},
			expect: traceability.PartsModelEntities{
				traceability.PartsModelEntity{
					TraceID:            uuid.MustParse(f.TraceID),
					OperatorID:         uuid.MustParse(f.OperatorID),
					PlantID:            uuid.MustParse(f.PlantId),
					PartsName:          f.PartsName,
					SupportPartsName:   common.StringPtr("品番A1"),
					TerminatedFlag:     true,
					AmountRequired:     nil,
					AmountRequiredUnit: common.StringPtr("kilogram"),
					PartsLabelName:     common.StringPtr("PartsB"),
					PartsAddInfo1:      common.StringPtr("Ver2.0"),
					PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
					PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
				},
			},
		},
		{
			name: "1-2: 正常系：0件の場合",
			input: func() traceability.GetPartsInput {
				i := f.NewGetPartsInput()
				i.TraceID = &f.NotExistID
				return i
			},
			expect: traceability.PartsModelEntities{},
		},
		{
			name: "1-3: nil許容項目がnilの場合",
			input: func() traceability.GetPartsInput {
				i := f.NewGetPartsInput_RequiredOnly()
				return i
			},
			expect: traceability.PartsModelEntities{
				traceability.PartsModelEntity{
					TraceID:            uuid.MustParse(f.TraceID7),
					OperatorID:         uuid.MustParse(f.OperatorID),
					PlantID:            uuid.MustParse(f.PlantId),
					PartsName:          "製品A7",
					SupportPartsName:   nil,
					TerminatedFlag:     false,
					AmountRequired:     nil,
					AmountRequiredUnit: nil,
					PartsLabelName:     nil,
					PartsAddInfo1:      nil,
					PartsAddInfo2:      nil,
					PartsAddInfo3:      nil,
				},
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
				actual, err := r.ListParts(test.input())
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
// Parts ListParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_ListParts_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     traceability.GetPartsInput
		dropQuery string
		expect    error
	}{
		{
			name: "2-1: 異常系：取得失敗の場合",
			input: traceability.GetPartsInput{
				OperatorID: f.OperatorID,
				TraceID:    common.StringPtr(f.TraceID),
				PartsName:  common.StringPtr("製品A1"),
				PlantID:    common.StringPtr(f.PlantId),
				ParentFlag: common.BoolPtr(true),
				Limit:      1,
			},
			dropQuery: "DROP TABLE IF EXISTS parts",
			expect:    fmt.Errorf("no such table: parts"),
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
				_, err = r.ListParts(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Parts GetPartByTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_GetPartByTraceID(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.PartsModelEntity
	}{
		{
			name:  "1-1: 正常系 取得成功の場合",
			input: f.TraceID5,
			expect: traceability.PartsModelEntity{
				TraceID:            uuid.MustParse(f.TraceID5),
				OperatorID:         uuid.MustParse(f.OperatorID),
				PlantID:            uuid.MustParse(f.PlantId),
				PartsName:          "製品A3",
				SupportPartsName:   common.StringPtr("品番A3"),
				TerminatedFlag:     false,
				AmountRequired:     nil,
				AmountRequiredUnit: common.StringPtr("kilogram"),
				PartsLabelName:     common.StringPtr("PartsB"),
				PartsAddInfo1:      common.StringPtr("Ver2.0"),
				PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
				PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
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
				actual, err := r.GetPartByTraceID(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Parts GetPartByTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-2. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_GetPartByTraceID_Abnormal(tt *testing.T) {

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
			input:     f.TraceID5,
			dropQuery: "DROP TABLE IF EXISTS parts",
			expect:    fmt.Errorf("no such table: parts"),
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
				_, err = r.GetPartByTraceID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Parts CountPartsList テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_CountPartsList(tt *testing.T) {

	tests := []struct {
		name   string
		input  traceability.GetPartsInput
		expect int
	}{
		{
			name: "1-1: 正常系：1件以上の場合",
			input: traceability.GetPartsInput{
				OperatorID: f.OperatorID,
			},
			expect: 7,
		},
		{
			name: "1-2: 正常系：0件の場合",
			input: traceability.GetPartsInput{
				OperatorID: f.NotExistID,
			},
			expect: 0,
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
				actual, err := r.CountPartsList(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Parts CountPartsList テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_CountPartsList_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     traceability.GetPartsInput
		dropQuery string
		expect    error
	}{
		{
			name: "2-1: 異常系：取得失敗の場合",
			input: traceability.GetPartsInput{
				OperatorID: f.OperatorID,
			},
			dropQuery: "DROP TABLE IF EXISTS parts",
			expect:    fmt.Errorf("no such table: parts"),
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
				err = db.Exec(test.dropQuery).Error
				if err != nil {
					assert.Fail(t, "Errors occured by deleting DB")
				}
				_, err = r.CountPartsList(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Parts DeleteParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：削除成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_DeleteParts(tt *testing.T) {

	tests := []struct {
		name       string
		input      string
		checkQuery string
		before     int
		after      int
	}{
		{
			name:       "1-1: 正常系：削除成功の場合",
			input:      f.TraceID,
			checkQuery: "SELECT COUNT(*) FROM parts WHERE trace_id = ?",
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
				err = r.DeleteParts(test.input)
				if assert.NoError(t, err) {
					db.Raw(test.checkQuery, test.input).Scan(&actualCount)
					assert.Equal(t, test.after, actualCount)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Parts DeleteParts テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：削除失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_DeleteParts_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：削除失敗の場合",
			input:     f.TraceID,
			dropQuery: "DROP TABLE IF EXISTS parts",
			expect:    fmt.Errorf("failed to physically delete record from table parts: no such table: parts"),
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
				err = r.DeleteParts(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
