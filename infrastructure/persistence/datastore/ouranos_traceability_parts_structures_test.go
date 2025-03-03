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
// PartsStructure GetPartsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// [x] 1-2. 正常系：nil許容項目がnilの場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_GetPartsStructure(tt *testing.T) {

	tests := []struct {
		name   string
		input  traceability.GetPartsStructureInput
		expect traceability.PartsStructureEntity
	}{
		{
			name: "1-1: 正常系 取得成功の場合",
			input: traceability.GetPartsStructureInput{
				TraceID:    uuid.MustParse(f.TraceID5),
				OperatorID: f.OperatorID,
			},
			expect: traceability.PartsStructureEntity{
				ParentPartsEntity: &traceability.PartsModelEntity{
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
				},
				ChildrenPartsEntity: traceability.PartsModelEntities{
					{
						TraceID:            uuid.MustParse(f.TraceID6),
						OperatorID:         uuid.MustParse(f.OperatorID),
						PlantID:            uuid.MustParse(f.PlantId),
						PartsName:          "製品A5",
						SupportPartsName:   common.StringPtr("品番A5"),
						TerminatedFlag:     false,
						AmountRequired:     nil,
						AmountRequiredUnit: common.StringPtr("kilogram"),
						PartsLabelName:     common.StringPtr("PartsB"),
						PartsAddInfo1:      common.StringPtr("Ver2.0"),
						PartsAddInfo2:      common.StringPtr("2024-12-01-2024-12-31"),
						PartsAddInfo3:      common.StringPtr("任意の情報が入ります"),
					},
				},
			},
		},
		{
			name: "1-2: 正常系 nil許容項目がnilの場合",
			input: traceability.GetPartsStructureInput{
				TraceID:    uuid.MustParse(f.TraceID7),
				OperatorID: f.OperatorID,
			},
			expect: traceability.PartsStructureEntity{
				ParentPartsEntity: &traceability.PartsModelEntity{
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
				ChildrenPartsEntity: traceability.PartsModelEntities{
					{
						TraceID:            uuid.MustParse(f.TraceID8),
						OperatorID:         uuid.MustParse(f.OperatorID),
						PlantID:            uuid.MustParse(f.PlantId),
						PartsName:          "製品A8",
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
				actual, err := r.GetPartsStructure(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect.ParentPartsEntity, actual.ParentPartsEntity)
					for i, m := range actual.ChildrenPartsEntity {
						assert.Equal(t, test.expect.ChildrenPartsEntity[i], m)
					}
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure GetPartsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_GetPartsStructure_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     traceability.GetPartsStructureInput
		dropQuery string
		expect    error
	}{
		{
			name: "2-1: 異常系：取得失敗の場合",
			input: traceability.GetPartsStructureInput{
				TraceID:    uuid.MustParse(f.TraceID5),
				OperatorID: f.OperatorID,
			},
			dropQuery: "DROP TABLE IF EXISTS parts_structures",
			expect:    fmt.Errorf("no such table: parts_structures"),
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
				_, err = r.GetPartsStructure(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure PutPartsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：更新成功の場合
// [x] 1-2. 正常系：nil許容項目がnilの場合
// [x] 1-3. 正常系：任意項目が未定義の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_PutPartsStructure(tt *testing.T) {

	tests := []struct {
		name   string
		input  traceability.PartsStructureModel
		expect traceability.PartsStructureEntity
	}{
		{
			name:   "1-1: 正常系 更新成功の場合",
			input:  f.NewPartsStructureModel(),
			expect: f.NewPartsStructureEntity(),
		},
		{
			name:   "1-2: 正常系 nil許容項目がnilの場合",
			input:  f.NewPartsStructureModel_RequiredOnly(),
			expect: f.NewPartsStructureEntity_RequiredOnly(),
		},
		{
			name:   "1-3: 正常系 任意項目が未定義の場合",
			input:  f.NewPartsStructureModel_RequiredOnlyWithUndefined(),
			expect: f.NewPartsStructureEntity_RequiredOnly(),
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
				actual, err := r.PutPartsStructure(test.input)
				if assert.NoError(t, err) {
					test.expect.ParentPartsEntity.CreatedAt = f.DummyTime
					test.expect.ParentPartsEntity.UpdatedAt = f.DummyTime
					actual.ParentPartsEntity.CreatedAt = f.DummyTime
					actual.ParentPartsEntity.UpdatedAt = f.DummyTime
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure PutPartsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：更新失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_PutPartsStructure_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     traceability.PartsStructureModel
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：更新失敗の場合",
			input:     f.NewPartsStructureModel(),
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
				_, err = r.PutPartsStructure(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure DeletePartsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：削除成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_DeletePartsStructure(tt *testing.T) {

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
			checkQuery: "SELECT COUNT(*) FROM parts_structures WHERE trace_id = ?",
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
				err = r.DeletePartsStructure(test.input)
				if assert.NoError(t, err) {
					db.Raw(test.checkQuery, test.input).Scan(&actualCount)
					assert.Equal(t, test.after, actualCount)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure DeletePartsStructure テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：削除失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_DeletePartsStructure_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：削除失敗の場合",
			input:     f.NotExistID,
			dropQuery: "DROP TABLE IF EXISTS parts_structures",
			expect:    fmt.Errorf("failed to physically delete record from table parts_structures: no such table: parts_structures"),
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
				err = r.DeletePartsStructure(test.input)
				//fmt.Println(err.Error())
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure GetPartsStructureByTraceId テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_GetPartsStructureByTraceId(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.PartsStructureEntityModel
	}{
		{
			name:  "1-1: 正常系 取得成功の場合",
			input: f.TraceID6,
			expect: traceability.PartsStructureEntityModel{
				TraceID:       uuid.MustParse(f.TraceID6),
				ParentTraceID: uuid.MustParse(f.TraceID5),
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
				actual, err := r.GetPartsStructureByTraceId(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect.TraceID, actual.TraceID)
					assert.Equal(t, test.expect.ParentTraceID, actual.ParentTraceID)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// PartsStructure GetPartsStructureByTraceId テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-2. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Parts_GetPartsStructureByTraceId_Abnormal(tt *testing.T) {

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
			input:     f.TraceID6,
			dropQuery: "DROP TABLE IF EXISTS parts_structures",
			expect:    fmt.Errorf("no such table: parts_structures"),
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
				_, err = r.GetPartsStructureByTraceId(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
