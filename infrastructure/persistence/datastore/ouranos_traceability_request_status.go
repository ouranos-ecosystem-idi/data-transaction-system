package datastore

import (
	"fmt"
	"time"

	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/extension/logger"

	"gorm.io/gorm"
)

// GetStatus
// Summary: This function gets the status of a request and response.
// input: operatorID(string) ID of the operator
// input: limit(int) upper threshold
// input: statusID(*string) ID of the status
// input: traceID(*string) ID of the trace
// input: statusTarget(string) target of the status
// output: (traceability.StatusEntityModels) StatusEntityModels object
// output: (error) error object
func (r *ouranosRepository) GetStatus(operatorID string, limit int, statusID *string, traceID *string, statusTarget string) (traceability.StatusEntityModels, error) {
	var statuses traceability.StatusEntityModels

	db := r.db.Table("request_status").
		Joins("INNER JOIN trades ON trades.trade_id = request_status.trade_id")

	switch statusTarget {
	case traceability.Request.ToString():
		if traceID != nil {
			db = db.Where("trades.downstream_trace_id = ?", traceID)
		}
		db = db.Where("trades.downstream_operator_id = ?", operatorID)
	case traceability.Response.ToString():
		if statusID != nil {
			db = db.Where("request_status.status_id = ?", statusID)
		}
		db = db.Where("trades.upstream_operator_id = ?", operatorID)
	default:
		if statusID != nil {
			db = db.Where("request_status.status_id = ?", statusID)
		}
		db = db.Where(`(trades.upstream_operator_id = ?
			OR trades.downstream_operator_id = ?)`, operatorID, operatorID)
	}

	err := db.Order("request_status.created_at DESC").
		Limit(limit).
		Find(&statuses).Error
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusEntityModels{}, err
	}
	return statuses, nil
}

// CountStatus
// Summary: This function counts the status of a request and response.
// input: operatorID(string) ID of the operator
// input: statusID(*string) ID of the status
// input: traceID(*string) ID of the trace
// input: statusTarget(string) target of the status
// output: (int) count of status
// output: (error) error object
func (r *ouranosRepository) CountStatus(operatorID string, statusID *string, traceID *string, statusTarget string) (int, error) {
	var count int64
	db := r.db.Table("request_status").
		Joins("INNER JOIN trades ON trades.trade_id = request_status.trade_id")

	switch statusTarget {
	case traceability.Request.ToString():
		if traceID != nil {
			db = db.Where("trades.downstream_trace_id = ?", traceID)
		}
		db = db.Where("trades.downstream_operator_id = ?", operatorID)
	case traceability.Response.ToString():
		if statusID != nil {
			db = db.Where("request_status.status_id = ?", statusID)
		}
		db = db.Where("trades.upstream_operator_id = ?", operatorID)
	default:
		if statusID != nil {
			db = db.Where("request_status.status_id = ?", statusID)
		}
		db = db.Where(`(trades.upstream_operator_id = ?
			OR trades.downstream_operator_id = ?)`, operatorID, operatorID)
	}

	err := db.Count(&count).Error
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return 0, err
	}
	return int(count), nil
}

// GetStatusByTradeID
// Summary: This function gets the status by trade ID.
// input: tradeID(string) ID of the trade
// output: (traceability.StatusEntityModel) StatusEntityModel object
// output: (error) error object
func (r *ouranosRepository) GetStatusByTradeID(tradeID string) (traceability.StatusEntityModel, error) {
	var status traceability.StatusEntityModel
	if err := r.db.Table("request_status").Where(`trade_id = ?`, tradeID).First(&status).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return status, err
	}
	return status, nil
}

// PutStatusCancel
// Summary: This function updates the status to "cancel".
// input: statusID(string) ID of the status
// input: operatorID(string) ID of the operator
// output: (error) error object
func (r *ouranosRepository) PutStatusCancel(statusID string, operatorID string) error {
	var status traceability.StatusEntityModel
	if err := r.db.Table("request_status").Where("status_id = ?", statusID).First(&status).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return err
	}

	var trade traceability.TradeEntityModel
	if err := r.db.Table("trades").Where("trade_id = ?", status.TradeID).Where("downstream_operator_id = ?", operatorID).First(&trade).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return err
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("request_status").Where("status_id = ?", statusID).Delete(nil).Error; err != nil {
			logger.Set(nil).Errorf(err.Error())
			return err
		}

		if err := tx.Table("trades").Where("trade_id = ?", status.TradeID).Where("downstream_operator_id = ?", operatorID).Delete(nil).Error; err != nil {
			logger.Set(nil).Errorf(err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return err
	}

	return nil
}

// PutStatusReject
// Summary: This function updates the status to "reject".
// input: statusID(string) ID of the status
// input: replyMessage(*string) reply message
// input: operatorID(string) ID of the operator
// output: (traceability.StatusEntityModel) StatusEntityModel object
// output: (error) error object
func (r *ouranosRepository) PutStatusReject(statusID string, replyMessage *string, operatorID string) (traceability.StatusEntityModel, error) {
	var status traceability.StatusEntityModel
	if err := r.db.Table("request_status").Where("status_id = ?", statusID).First(&status).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.StatusEntityModel{}, err
	}

	var trade traceability.TradeEntityModel
	if err := r.db.Table("trades").Where("trade_id = ?", status.TradeID).Where("upstream_operator_id = ?", operatorID).First(&trade).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())

		return traceability.StatusEntityModel{}, err
	}

	now := time.Now()
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("request_status").Where("status_id = ?", statusID).Updates(
			traceability.StatusEntityModel{
				CfpResponseStatus: traceability.CfpResponseStatusReject.ToString(),
				TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
				ReplyMessage:      replyMessage,
				UpdatedAt:         now,
			}).
			Error; err != nil {
			return err
		}

		if err := tx.Table("trades").Where("trade_id = ?", status.TradeID).Where("upstream_operator_id = ?", operatorID).Updates(
			traceability.TradeEntityModel{
				UpstreamOperatorID: nil,
				UpstreamTraceID:    nil,
			}).
			Error; err != nil {
			logger.Set(nil).Errorf(err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.StatusEntityModel{}, err
	}

	var statusResult traceability.StatusEntityModel
	if err := r.db.Table("request_status").Where("status_id = ?", statusID).First(&statusResult).Error; err != nil {
		logger.Set(nil).Errorf(err.Error())
		return traceability.StatusEntityModel{}, err
	}

	return statusResult, nil
}

// DeleteRequestStatusByTradeID
// Summary: This function deletes the status by trade ID.
// input: tradeID(string) ID of the trade
// output: (error) error object
func (r *ouranosRepository) DeleteRequestStatusByTradeID(tradeID string) error {
	result := r.db.Unscoped().Table("request_status").Where("trade_id = ?", tradeID).Delete(nil)
	if result.Error != nil {
		return fmt.Errorf("failed to physically delete record from table request_status: %v", result.Error)
	}
	return nil
}
