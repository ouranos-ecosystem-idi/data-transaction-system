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
// Trades GetTradeRequest テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTradeRequest(tt *testing.T) {

	tests := []struct {
		name                      string
		inputDownstreamOperatorID string
		inputLimit                int
		inputDownstreamTraceIds   []string
		expect                    traceability.TradeEntityModels
	}{
		{
			name:                      "1-1: 正常系：1件以上の場合",
			inputDownstreamOperatorID: f.OperatorID,
			inputLimit:                20,
			inputDownstreamTraceIds:   []string{f.TraceID3, f.TraceID2},
			expect: traceability.TradeEntityModels{
				{
					TradeID:              common.UUIDPtr(uuid.MustParse(f.TradeID2)),
					DownstreamOperatorID: uuid.MustParse(f.OperatorID),
					UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(f.OperatorID2)),
					DownstreamTraceID:    uuid.MustParse(f.TraceID3),
					UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(f.TraceID4)),
					TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
					DeletedAt:            gorm.DeletedAt{},
					CreatedAt:            f.DummyTime,
					CreatedUserID:        "seed",
					UpdatedAt:            f.DummyTime,
					UpdatedUserID:        "seed",
				},
			},
		},
		{
			name:                      "1-2: 正常系：0件の場合",
			inputDownstreamOperatorID: f.OperatorID,
			inputLimit:                20,
			inputDownstreamTraceIds:   []string{f.NotExistID},
			expect:                    traceability.TradeEntityModels{},
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
				actual, err := r.GetTradeRequest(test.inputDownstreamOperatorID, test.inputLimit, test.inputDownstreamTraceIds)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect), len(actual))
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTradeRequest テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTradeRequest_Abnormal(tt *testing.T) {

	tests := []struct {
		name                      string
		inputDownstreamOperatorID string
		inputLimit                int
		inputDownstreamTraceIds   []string
		dropQuery                 string
		expect                    error
	}{
		{
			name:                      "2-1: 異常系：取得失敗の場合",
			inputDownstreamOperatorID: f.OperatorID,
			inputLimit:                20,
			inputDownstreamTraceIds:   []string{f.TraceID2, f.TraceID3},
			dropQuery:                 "DROP TABLE IF EXISTS trades",
			expect:                    fmt.Errorf("no such table: trades"),
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
				_, err = r.GetTradeRequest(test.inputDownstreamOperatorID, test.inputLimit, test.inputDownstreamTraceIds)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTradeResponse テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTradeResponse(tt *testing.T) {

	tests := []struct {
		name                    string
		inputUpstreamOperatorID string
		inputLimit              int
		expect                  traceability.TradeEntityModels
	}{
		{
			name:                    "1-1: 正常系：1件以上の場合",
			inputUpstreamOperatorID: f.OperatorID,
			inputLimit:              1,
			expect: traceability.TradeEntityModels{
				{
					TradeID:              common.UUIDPtr(uuid.MustParse(f.TradeID)),
					DownstreamOperatorID: uuid.MustParse(f.OperatorID2),
					UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(f.OperatorID)),
					DownstreamTraceID:    uuid.MustParse(f.TraceID2),
					UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(f.TraceID)),
					TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
					DeletedAt:            gorm.DeletedAt{},
					CreatedAt:            f.DummyTime,
					CreatedUserID:        "seed",
					UpdatedAt:            f.DummyTime,
					UpdatedUserID:        "seed",
				},
			},
		},
		{
			name:                    "1-2: 正常系：0件の場合",
			inputUpstreamOperatorID: f.NotExistID,
			inputLimit:              20,
			expect:                  traceability.TradeEntityModels{},
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
				actual, _ := r.GetTradeResponse(test.inputUpstreamOperatorID, test.inputLimit)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect), len(actual))
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTradeResponse テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTradeResponse_Abnormal(tt *testing.T) {

	tests := []struct {
		name                    string
		inputUpstreamOperatorID string
		inputLimit              int
		dropQuery               string
		expect                  error
	}{
		{
			name:                    "2-1: 異常系：取得失敗の場合",
			inputUpstreamOperatorID: f.OperatorID,
			inputLimit:              20,
			dropQuery:               "DROP TABLE IF EXISTS trades",
			expect:                  fmt.Errorf("no such table: trades"),
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
				_, err = r.GetTradeResponse(test.inputUpstreamOperatorID, test.inputLimit)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTradeByDownstreamTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTradeByDownstreamTraceID(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.TradeEntityModel
	}{
		{
			name:  "1-1: 正常系 取得成功の場合",
			input: f.TraceID2,
			expect: traceability.TradeEntityModel{
				TradeID:              common.UUIDPtr(uuid.MustParse(f.TradeID)),
				DownstreamOperatorID: uuid.MustParse(f.OperatorID2),
				UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(f.OperatorID)),
				DownstreamTraceID:    uuid.MustParse(f.TraceID2),
				UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(f.TraceID)),
				TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
				DeletedAt:            gorm.DeletedAt{},
				CreatedAt:            f.DummyTime,
				CreatedUserID:        "seed",
				UpdatedAt:            f.DummyTime,
				UpdatedUserID:        "seed",
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
				actual, _ := r.GetTradeByDownstreamTraceID(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTradeByDownstreamTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-2. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTradeByDownstreamTraceID_Abnormal(tt *testing.T) {

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
			input:     f.TraceID2,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.GetTradeByDownstreamTraceID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTrade テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTrade(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.TradeEntityModel
	}{
		{
			name:  "1-1: 正常系 取得成功の場合",
			input: f.TradeID,
			expect: traceability.TradeEntityModel{
				TradeID:              common.UUIDPtr(uuid.MustParse(f.TradeID)),
				DownstreamOperatorID: uuid.MustParse(f.OperatorID2),
				UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(f.OperatorID)),
				DownstreamTraceID:    uuid.MustParse(f.TraceID2),
				UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(f.TraceID)),
				TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
				DeletedAt:            gorm.DeletedAt{},
				CreatedAt:            f.DummyTime,
				CreatedUserID:        "seed",
				UpdatedAt:            f.DummyTime,
				UpdatedUserID:        "seed",
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
				actual, _ := r.GetTrade(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades GetTrade テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-2. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_GetTrade_Abnormal(tt *testing.T) {

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
			input:     f.TradeID,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.GetTradeByDownstreamTraceID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades ListTradeByUpstreamTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_ListTradeByUpstreamTraceID(tt *testing.T) {

	tests := []struct {
		name       string
		input      string
		inputLimit int
		expect     traceability.TradeEntityModels
	}{
		{
			name:  "1-1: 正常系：1件以上の場合",
			input: f.TraceID,
			expect: traceability.TradeEntityModels{
				{
					TradeID:              common.UUIDPtr(uuid.MustParse(f.TradeID)),
					DownstreamOperatorID: uuid.MustParse(f.OperatorID2),
					UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(f.OperatorID)),
					DownstreamTraceID:    uuid.MustParse(f.TraceID2),
					UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(f.TraceID)),
					TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
					DeletedAt:            gorm.DeletedAt{},
					CreatedAt:            f.DummyTime,
					CreatedUserID:        "seed",
					UpdatedAt:            f.DummyTime,
					UpdatedUserID:        "seed",
				},
			},
		},
		{
			name:   "1-2: 正常系：0件の場合",
			input:  f.NotExistID,
			expect: traceability.TradeEntityModels{},
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
				actual, _ := r.ListTradeByUpstreamTraceID(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, len(test.expect), len(actual))
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades ListTradeByUpstreamTraceID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_ListTradeByUpstreamTraceID_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：取得失敗の場合",
			input:     f.TraceID,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.ListTradeByUpstreamTraceID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades ListTradesByOperatorID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_ListTradesByOperatorID(tt *testing.T) {

	tests := []struct {
		name        string
		input       string
		expectCount int
		expect      traceability.TradeEntityModel
	}{
		{
			name:        "1-1: 正常系：1件以上の場合",
			input:       f.OperatorID2,
			expectCount: 4,
			expect: traceability.TradeEntityModel{
				TradeID:              common.UUIDPtr(uuid.MustParse(f.TradeID)),
				DownstreamOperatorID: uuid.MustParse(f.OperatorID2),
				UpstreamOperatorID:   common.UUIDPtr(uuid.MustParse(f.OperatorID)),
				DownstreamTraceID:    uuid.MustParse(f.TraceID2),
				UpstreamTraceID:      common.UUIDPtr(uuid.MustParse(f.TraceID)),
				TradeDate:            common.StringPtr("2024-05-01T00:00:00Z"),
				DeletedAt:            gorm.DeletedAt{},
				CreatedAt:            f.DummyTime,
				CreatedUserID:        "seed",
				UpdatedAt:            f.DummyTime,
				UpdatedUserID:        "seed",
			},
		},
		{
			name:        "1-2: 正常系：0件の場合",
			input:       f.NotExistID,
			expectCount: 0,
			expect:      traceability.TradeEntityModel{},
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
				actual, _ := r.ListTradesByOperatorID(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectCount, len(actual))
					if test.name == "1-1: 正常系：1件以上の場合" {
						assert.Contains(t, actual, test.expect)
					}
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades ListTradesByOperatorID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_ListTradesByOperatorID_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：取得失敗の場合",
			input:     f.OperatorID2,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.ListTradesByOperatorID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades CountTradeRequest テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_CountTradeRequest(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "1-1: 正常系：1件以上の場合",
			input:  f.OperatorID,
			expect: 2,
		},
		{
			name:   "1-2: 正常系：0件の場合",
			input:  f.NotExistID,
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
				actual, _ := r.CountTradeRequest(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades CountTradeRequest テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_CountTradeRequest_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：取得失敗の場合",
			input:     f.OperatorID,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.CountTradeRequest(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades CountTradeResponse テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_CountTradeResponse(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "1-1: 正常系：1件以上の場合",
			input:  f.OperatorID2,
			expect: 2,
		},
		{
			name:   "1-2: 正常系：0件の場合",
			input:  f.NotExistID,
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
				actual, _ := r.CountTradeResponse(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades CountTradeResponse テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_CountTradeResponse_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：取得失敗の場合",
			input:     f.OperatorID2,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.CountTradeResponse(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades DeleteTrade テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：削除成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_DeleteTrade(tt *testing.T) {

	tests := []struct {
		name       string
		input      string
		checkQuery string
		before     int
		after      int
	}{
		{
			name:       "1-1: 正常系：削除成功の場合",
			input:      f.TradeID,
			checkQuery: "SELECT COUNT(*) FROM trades WHERE trade_id = ?",
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
				err = r.DeleteTrade(test.input)
				if assert.NoError(t, err) {
					db.Raw(test.checkQuery, test.input).Scan(&actualCount)
					assert.Equal(t, test.after, actualCount)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades DeleteTrade テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：削除失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_DeleteTrade_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：取得失敗の場合",
			input:     f.TradeID,
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("failed to physically delete record from table trades : no such table: trades"),
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
				err = r.DeleteTrade(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades PutTradeRequest テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：更新成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_PutTradeRequest(tt *testing.T) {

	tests := []struct {
		name   string
		input  traceability.TradeRequestEntityModel
		expect traceability.TradeRequestEntityModel
	}{

		{
			name:   "1-1: 正常系 更新成功の場合",
			input:  f.NewPutTradeRequestModelInput(),
			expect: f.NewPutTradeRequestModelInput(),
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
				actual, err := r.PutTradeRequest(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades PutTradeRequest テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：更新失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_PutTradeRequest_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     traceability.TradeRequestEntityModel
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：更新失敗の場合",
			input:     f.NewPutTradeRequestModelInput(),
			dropQuery: "DROP TABLE IF EXISTS trades",
			expect:    fmt.Errorf("no such table: trades"),
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
				_, err = r.PutTradeRequest(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Trades PutTradeResponse テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：更新成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_PutTradeResponse(tt *testing.T) {

	tests := []struct {
		name                    string
		inputTradeResponseInput traceability.PutTradeResponseInput
		inputRequestStatus      traceability.RequestStatus
		expect                  traceability.TradeEntityModel
	}{

		{
			name:                    "1-1: 正常系 更新成功の場合",
			inputTradeResponseInput: f.PutTradeResponseInput2,
			inputRequestStatus:      f.NewRequestStatus(),
			expect:                  f.NewTradeEntityModel(),
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
				actual, err := r.PutTradeResponse(test.inputTradeResponseInput, test.inputRequestStatus)
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
// Trades PutTradeResponse テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：更新失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Trade_PutTradeResponse_Abnormal(tt *testing.T) {

	tests := []struct {
		name                    string
		inputTradeResponseInput traceability.PutTradeResponseInput
		inputRequestStatus      traceability.RequestStatus
		dropQuery               string
		expect                  error
	}{
		{
			name:                    "2-1: 異常系：更新失敗の場合",
			inputTradeResponseInput: f.PutTradeResponseInput2,
			inputRequestStatus:      f.NewRequestStatus(),
			dropQuery:               "DROP TABLE IF EXISTS trades",
			expect:                  fmt.Errorf("no such table: trades"),
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
				_, err = r.PutTradeResponse(test.inputTradeResponseInput, test.inputRequestStatus)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
