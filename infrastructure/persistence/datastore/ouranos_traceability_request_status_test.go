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

	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus GetStatus テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_GetStatus(tt *testing.T) {

	tests := []struct {
		name              string
		inputOperatorID   string
		inputLimit        int
		inputStatusID     *string
		inputTraceID      *string
		inputStatusTarget string
		expect            traceability.StatusEntityModels
	}{
		{
			name:              "1-1: 正常系：1件以上の場合",
			inputOperatorID:   f.OperatorID2,
			inputLimit:        20,
			inputStatusID:     common.StringPtr(f.StatusID),
			inputTraceID:      common.StringPtr(f.TraceID2),
			inputStatusTarget: "REQUEST",
			expect:            f.NewStatusModels(),
		},
		{
			name:              "1-2: 正常系：0件の場合",
			inputOperatorID:   f.NotExistID,
			inputLimit:        20,
			inputStatusID:     common.StringPtr(f.StatusID),
			inputTraceID:      common.StringPtr(f.TraceID2),
			inputStatusTarget: "REQUEST",
			expect:            traceability.StatusEntityModels{},
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
				actual, err := r.GetStatus(test.inputOperatorID, test.inputLimit, test.inputStatusID, test.inputTraceID, test.inputStatusTarget)
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
// RequestStatus GetStatus テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_GetStatus_Abnormal(tt *testing.T) {

	tests := []struct {
		name              string
		inputOperatorID   string
		inputLimit        int
		inputStatusID     *string
		inputTraceID      *string
		inputStatusTarget string
		dropQuery         string
		expect            error
	}{
		{
			name:              "2-1: 異常系：取得失敗の場合",
			inputOperatorID:   f.OperatorID2,
			inputLimit:        20,
			inputStatusID:     common.StringPtr(f.StatusID),
			inputTraceID:      common.StringPtr(f.TraceID2),
			inputStatusTarget: "REQUEST",
			dropQuery:         "DROP TABLE IF EXISTS request_status",
			expect:            fmt.Errorf("no such table: request_status"),
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
				_, err = r.GetStatus(test.inputOperatorID, test.inputLimit, test.inputStatusID, test.inputTraceID, test.inputStatusTarget)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus GetStatusByTradeID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：取得成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_GetStatusByTradeID(tt *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect traceability.StatusEntityModel
	}{
		{
			name:   "1-1: 正常系 取得成功の場合",
			input:  f.TradeID,
			expect: f.NewStatusModel2(),
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
				actual, err := r.GetStatusByTradeID(test.input)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus GetStatusByTradeID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：0件の場合
// [x] 2-2. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_GetStatusByTradeID_Abnormal(tt *testing.T) {

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
			dropQuery: "DROP TABLE IF EXISTS request_status",
			expect:    fmt.Errorf("no such table: request_status"),
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
				_, err = r.GetStatusByTradeID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus CountStatus テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：1件以上の場合
// [x] 1-2. 正常系：0件の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_CountStatus(tt *testing.T) {

	tests := []struct {
		name              string
		inputOperatorID   string
		inputStatusID     *string
		inputTraceID      *string
		inputStatusTarget string
		expect            int
	}{
		{
			name:              "1-1: 正常系：1件以上の場合",
			inputOperatorID:   f.OperatorID2,
			inputStatusID:     common.StringPtr(f.StatusID),
			inputTraceID:      common.StringPtr(f.TraceID2),
			inputStatusTarget: "REQUEST",
			expect:            1,
		},
		{
			name:              "1-2: 正常系：0件の場合",
			inputOperatorID:   f.NotExistID,
			inputStatusID:     common.StringPtr(f.StatusID),
			inputTraceID:      common.StringPtr(f.TraceID2),
			inputStatusTarget: "REQUEST",
			expect:            0,
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
				actual, err := r.CountStatus(test.inputOperatorID, test.inputStatusID, test.inputTraceID, test.inputStatusTarget)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect, actual)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus CountStatus テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：取得失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_CountStatus_Abnormal(tt *testing.T) {

	tests := []struct {
		name              string
		inputOperatorID   string
		inputStatusID     *string
		inputTraceID      *string
		inputStatusTarget string
		dropQuery         string
		expect            error
	}{
		{
			name:              "2-1: 異常系：取得失敗の場合",
			inputOperatorID:   f.OperatorID2,
			inputStatusID:     common.StringPtr(f.StatusID),
			inputTraceID:      common.StringPtr(f.TraceID2),
			inputStatusTarget: "REQUEST",
			dropQuery:         "DROP TABLE IF EXISTS request_status",
			expect:            fmt.Errorf("no such table: request_status"),
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
				_, err = r.CountStatus(test.inputOperatorID, test.inputStatusID, test.inputTraceID, test.inputStatusTarget)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus DeleteRequestStatusByTradeID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：削除成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_DeleteRequestStatusByTradeID(tt *testing.T) {

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
			checkQuery: "SELECT COUNT(*) FROM request_status WHERE trade_id = ?",
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
				err = r.DeleteRequestStatusByTradeID(test.input)
				if assert.NoError(t, err) {
					db.Raw(test.checkQuery, test.input).Scan(&actualCount)
					assert.Equal(t, test.after, actualCount)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus DeleteRequestStatusByTradeID テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：削除失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_DeleteRequestStatusByTradeID_Abnormal(tt *testing.T) {

	tests := []struct {
		name      string
		input     string
		dropQuery string
		expect    error
	}{
		{
			name:      "2-1: 異常系：削除失敗の場合",
			input:     f.TradeID,
			dropQuery: "DROP TABLE IF EXISTS request_status",
			expect:    fmt.Errorf("failed to physically delete record from table request_status: no such table: request_status"),
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
				err = r.DeleteRequestStatusByTradeID(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus PutStatusCancel テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：更新成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_PutStatusCancel(tt *testing.T) {

	expectStatusModel := f.NewStatusModel2()
	expectStatusModel.CfpResponseStatus = traceability.CfpResponseStatusCancel.ToString()
	expectStatusModel.TradeTreeStatus = traceability.TradeTreeStatusUnterminated.ToString()

	tests := []struct {
		name            string
		inputStatusID   string
		inputOperatorID string
		expect          traceability.StatusEntityModel
	}{

		{
			name:            "1-1: 正常系 取得成功の場合",
			inputStatusID:   f.StatusID,
			inputOperatorID: f.OperatorID2,
			expect:          expectStatusModel,
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
				err = r.PutStatusCancel(test.inputStatusID, test.inputOperatorID)
				if assert.NoError(t, err) {
					test.expect.UpdatedAt = f.DummyTime
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus PutStatusCancel テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：更新失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_PutStatusCancel_Abnormal(tt *testing.T) {

	tests := []struct {
		name            string
		inputStatusID   string
		inputOperatorID string
		dropQuery       string
		expect          error
	}{
		{
			name:            "2-1: 異常系：更新失敗の場合",
			inputStatusID:   f.StatusID,
			inputOperatorID: f.OperatorID2,
			dropQuery:       "DROP TABLE IF EXISTS request_status",
			expect:          fmt.Errorf("no such table: request_status"),
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
				err = r.PutStatusCancel(test.inputStatusID, test.inputOperatorID)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// RequestStatus PutStatusReject テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 正常系：更新成功の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_PutStatusReject(tt *testing.T) {

	expectStatusModel := f.NewStatusModel2()
	expectStatusModel.CfpResponseStatus = traceability.CfpResponseStatusReject.ToString()
	expectStatusModel.TradeTreeStatus = traceability.TradeTreeStatusUnterminated.ToString()

	tests := []struct {
		name              string
		inputStatusID     string
		inputReplyMessage *string
		inputOperatorID   string
		expect            traceability.StatusEntityModel
	}{

		{
			name:              "1-1: 正常系 更新成功の場合",
			inputStatusID:     f.StatusID,
			inputReplyMessage: common.StringPtr(""),
			inputOperatorID:   f.OperatorID,
			expect:            expectStatusModel,
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
				actual, err := r.PutStatusReject(test.inputStatusID, test.expect.ReplyMessage, test.inputOperatorID)
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
// RequestStatus PutStatusReject テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1. 異常系：更新失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_RequestStatus_PutStatusReject_Abnormal(tt *testing.T) {

	tests := []struct {
		name              string
		inputStatusID     string
		inputReplyMessage *string
		inputOperatorID   string
		dropQuery         string
		expect            error
	}{
		{
			name:              "2-1: 異常系：更新失敗の場合",
			inputStatusID:     f.StatusID,
			inputReplyMessage: common.StringPtr(""),
			inputOperatorID:   f.OperatorID2,
			dropQuery:         "DROP TABLE IF EXISTS request_status",
			expect:            fmt.Errorf("no such table: request_status"),
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
				_, err = r.PutStatusReject(test.inputStatusID, test.inputReplyMessage, test.inputOperatorID)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
