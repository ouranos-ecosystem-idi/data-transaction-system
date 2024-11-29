package datastore

import (
	"fmt"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetTradeRequest
// Summary: This is function which get TradeEntityModels from trades by using downstream_operator_id and downstream_trace_id.
// input: downstreamOperatorID(string) value of downstreamOperatorID
// input: limit(int) value of limit
// input: downstreamTraceIDs([]string) list of downstreamTraceIDs
// output: (TradeEntityModels) TradeEntityModels object
// output: (error) error object
func (r *ouranosRepository) GetTradeRequest(downstreamOperatorID string, limit int, downstreamTraceIDs []string) (traceability.TradeEntityModels, error) {
	var es traceability.TradeEntityModels

	query := r.db.Table("trades").
		Joins("INNER JOIN request_status ON trades.trade_id = request_status.trade_id").
		Where("trades.downstream_operator_id = ?", downstreamOperatorID)

	if len(downstreamTraceIDs) > 0 {
		query = query.Where("trades.downstream_trace_id IN (?)", downstreamTraceIDs)
	}

	if err := query.
		Order("request_status.created_at DESC").
		Limit(limit).
		Find(&es).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.TradeEntityModels{}, err
	}

	return es, nil
}

// GetTradeResponse
// Summary: This is function which get TradeEntityModels from trades by using upstream_operator_id.
// input: upstreamOperatorID(string) value of upstreamOperatorID
// input: limit(int) value of limit
// output: (TradeEntityModels) TradeEntityModels object
// output: (error) error object
func (r *ouranosRepository) GetTradeResponse(upstreamOperatorID string, limit int) (traceability.TradeEntityModels, error) {
	var es traceability.TradeEntityModels

	err := r.db.Table("trades").
		Joins("INNER JOIN request_status ON trades.trade_id = request_status.trade_id").
		Where("trades.upstream_operator_id = ?", upstreamOperatorID).
		Order("request_status.created_at DESC").
		Limit(limit).
		Find(&es).
		Error
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeEntityModels{}, err
	}
	return es, nil
}

// GetTradeByDownstreamTraceID
// Summary: This is function which get TradeEntityModel from trades by using downstream_trace_id.
// input: donwstreamTraceID(string) value of donwstreamTraceID
// output: (TradeEntityModels) TradeEntityModels object
// output: (error) error object
func (r *ouranosRepository) GetTradeByDownstreamTraceID(donwstreamTraceID string) (traceability.TradeEntityModel, error) {
	var e traceability.TradeEntityModel

	if err := r.db.Table("trades").Where("downstream_trace_id = ?", donwstreamTraceID).First(&e).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeEntityModel{}, err
	}
	return e, nil
}

// GetTrade
// Summary: This is function which get TradeEntityModel from trades by using trade_id.
// input: tradeID(string) value of tradeID
// output: (TradeEntityModel) TradeEntityModel object
// output: (error) error object
func (r *ouranosRepository) GetTrade(tradeID string) (traceability.TradeEntityModel, error) {
	var e traceability.TradeEntityModel

	if err := r.db.Table("trades").Where("trade_id = ?", tradeID).Limit(1).First(&e).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeEntityModel{}, err
	}
	return e, nil
}

// ListTradeByUpstreamTraceID
// Summary: This is function which get TradeEntityModels from trades by using upstream_trace_id.
// input: upstreamTraceID(string) value of upstreamTraceID
// output: (TradeEntityModels) TradeEntityModels object
// output: (error) error object
func (r *ouranosRepository) ListTradeByUpstreamTraceID(upstreamTraceID string) (traceability.TradeEntityModels, error) {
	var es traceability.TradeEntityModels

	if err := r.db.Table("trades").Where("upstream_trace_id = ?", upstreamTraceID).Find(&es).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return nil, err
	}
	return es, nil
}

// ListTradeByDownstreamTraceID
// Summary: This is function which get TradeEntityModels from trades by using downstream_trace_id.
// input: downstreamTraceID(string) value of downstreamTraceID
// output: (TradeEntityModels) TradeEntityModels object
// output: (error) error object
func (r *ouranosRepository) ListTradeByDownstreamTraceID(downstreamTraceID string) (traceability.TradeEntityModels, error) {
	var es traceability.TradeEntityModels

	if err := r.db.Table("trades").Where("downstream_trace_id = ?", downstreamTraceID).Find(&es).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return nil, err
	}
	return es, nil
}

// CountTradeRequest
// Summary: This is function which get record count from trades by using downstream_operator_id.
// input: downstreamOperatorID(string) value of downstreamOperatorID
// output: (int) record count
// output: (error) error object
func (r *ouranosRepository) CountTradeRequest(downstreamOperatorID string) (int, error) {
	var count int64
	if err := r.db.Table("trades").Where("downstream_operator_id = ?", downstreamOperatorID).Count(&count).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return 0, err
	}
	return int(count), nil
}

// CountTradeRequest
// Summary: This is function which get record count from trades by using upstream_operator_id.
// input: upstreamOperatorID(string) value of upstreamOperatorID
// output: (int) record count
// output: (error) error object
func (r *ouranosRepository) CountTradeResponse(upstreamOperatorID string) (int, error) {
	var count int64
	if err := r.db.Table("trades").Where("upstream_operator_id = ?", upstreamOperatorID).Count(&count).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return 0, err
	}
	return int(count), nil
}

// PutTradeRequest
// Summary: This is function which update trades with TradeRequestEntityModel.
// input: tradeRequestEntityModel(TradeRequestEntityModel) TradeRequestEntityModel object
// output: (TradeRequestEntityModel) TradeRequestEntityModel object
// output: (error) error object
func (r *ouranosRepository) PutTradeRequest(tradeRequestEntityModel traceability.TradeRequestEntityModel) (traceability.TradeRequestEntityModel, error) {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		// upsert
		err := tx.Table("trades").Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
					{Name: "trade_id"},
				},
				DoUpdates: clause.AssignmentColumns(
					[]string{
						"downstream_operator_id",
						"upstream_operator_id",
						"downstream_trace_id",
						"upstream_trace_id",
						"trade_date",
						"deleted_at",
						"updated_at",
						"updated_user_id",
					}),
			}).Create(&tradeRequestEntityModel.TradeEntityModel).Error
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return err
		}

		err = tx.Table("request_status").Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
					{Name: "status_id"},
				},
				DoUpdates: clause.AssignmentColumns(
					[]string{
						"trade_id",
						"request_status",
						"message",
						"request_type",
						"response_due_date",
						"completed_count",
						"completed_count_modified_at",
						"trades_count",
						"trades_count_modified_at",
						"deleted_at",
						"updated_at",
						"updated_user_id",
					}),
			}).Create(&tradeRequestEntityModel.StatusEntityModel).Error
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return err
		}

		return err
	})
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeRequestEntityModel{}, err
	}

	return tradeRequestEntityModel, nil
}

// PutTradeResponse
// Summary: This is function which update trades with TradeRequestEntityModel.
// input: putTradeResponseInput(PutTradeResponseInput) PutTradeResponseInput object
// input: requestStatus(RequestStatus) RequestStatus object
// output: (TradeRequestEntityModel) TradeRequestEntityModel object
// output: (error) error object
func (r *ouranosRepository) PutTradeResponse(putTradeResponseInput traceability.PutTradeResponseInput, requestStatus traceability.RequestStatus) (traceability.TradeEntityModel, error) {
	now := time.Now()
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("trades").
			Where("trade_id = ?", putTradeResponseInput.TradeID).Updates(
			traceability.TradeEntityModel{
				UpstreamTraceID: &putTradeResponseInput.TraceID,
				UpdatedAt:       now,
			}).
			Error; err != nil {
			logger.Set(nil).Errorf(err.Error())
			return err
		}

		updates := traceability.StatusEntityModel{
			CfpResponseStatus: requestStatus.CfpResponseStatus.ToString(),
			TradeTreeStatus:   requestStatus.TradeTreeStatus.ToString(),
			UpdatedAt:         now,
		}

		if requestStatus.CompletedCount != nil {
			updates.CompletedCount = requestStatus.CompletedCount
			updates.CompletedCountModifiedAt = &now
		}

		if err := tx.Table("request_status").
			Where("trade_id = ?", putTradeResponseInput.TradeID).
			Updates(updates).
			Error; err != nil {
			logger.Set(nil).Errorf(err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.TradeEntityModel{}, err
	}

	var e traceability.TradeEntityModel
	if err := r.db.Table("trades").Where("trade_id = ?", putTradeResponseInput.TradeID).First(&e).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.TradeEntityModel{}, err
	}

	return e, nil
}

// ListTradesByOperatorID
// Summary: This is function which get TradeEntityModels from trades by using downstream_operator_id.
// input: operatorID(string) value of operatorID
// output: (TradeEntityModels) TradeEntityModels object
// output: (error) error object
func (r *ouranosRepository) ListTradesByOperatorID(operatorID string) (traceability.TradeEntityModels, error) {
	var es traceability.TradeEntityModels
	if err := r.db.Table("trades").Where("downstream_operator_id = ?", operatorID).Or("upstream_operator_id = ?", operatorID).Find(&es).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.TradeEntityModels{}, err
	}
	return es, nil

}

// DeleteTrade
// Summary: This is function which delete trades by using trade_id.
// input: tradeID(string) value of tradeID
// output: (error) error object
func (r *ouranosRepository) DeleteTrade(tradeID string) error {
	result := r.db.Unscoped().Table("trades").Where("trade_id = ?", tradeID).Delete(nil)
	if result.Error != nil {
		return fmt.Errorf(common.DeleteTableError("trades", result.Error))
	}
	return nil
}
