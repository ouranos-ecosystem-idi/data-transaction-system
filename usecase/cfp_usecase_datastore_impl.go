package usecase

import (
	"errors"
	"fmt"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// cfpUsecase
// Summary: This is structure which defines cfpUsecase.
type cfpUsecase struct {
	r repository.OuranosRepository
}

// NewCfpUsecase
// Summary: This is function to create new cfpUsecase.
// input: r(repository.OuranosRepository) repository interface
// output: (ICfpUsecase) use case interface
func NewCfpUsecase(r repository.OuranosRepository) ICfpUsecase {
	return &cfpUsecase{r}
}

// GetCfp
// Summary: This is function which get a list of cfp.
// input: c(echo.Context) echo context
// input: getCfpInput(traceability.GetCfpInput) GetCfpInput object
// output: ([]traceability.CfpModel) list of CfpModel
// output: (error) error object
func (u *cfpUsecase) GetCfp(c echo.Context, getCfpInput traceability.GetCfpInput) ([]traceability.CfpModel, error) {
	var res []traceability.CfpModel = []traceability.CfpModel{}

	for _, traceID := range getCfpInput.TraceIDs {
		parts, err := u.r.GetPartByTraceID(traceID.String())
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}

		partsStructureModel, err := u.r.GetPartsStructureByTraceId(traceID.String())
		if err != nil {
			logger.Set(nil).Errorf(err.Error())

			return nil, err
		}

		if parts.TerminatedFlag {
			// Pattern A. For terminated parts
			// A-1. If terminated, get only its own CFP value
			logger.Set(c).Debugf("TraceID: %#v is a terminated part. Only its own CFP information is retrieved.", traceID.String())
			cfps, err := u.r.ListCFPsByTraceID(traceID.String())
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					logger.Set(c).Debugf("Do not process because CFP information for TraceID: %#v of the end parts is not registered.", traceID.String())
					continue
				}

				logger.Set(c).Errorf(err.Error())

				return nil, err
			}

			if len(cfps) == 0 {
				logger.Set(c).Debugf("Do not process because CFP information for TraceID: %#v of the end parts is not registered.", traceID.String())
				continue
			}

			// For parent, add total CFP
			if partsStructureModel.IsParent() {
				parentCfpSet, err := traceability.NewCfpEntityModelSetFromCfpEntityModels(cfps, true)
				if err != nil {
					logger.Set(c).Errorf(err.Error())

					return nil, err
				}
				ms, err := parentCfpSet.ToModels()
				if err != nil {
					logger.Set(c).Errorf(err.Error())

					return nil, err
				}
				res = append(res, ms...)

			} else {
				ms, err := cfps.ToModels()
				if err != nil {
					logger.Set(c).Errorf(err.Error())

					return nil, err
				}
				res = append(res, ms...)
			}

			continue
		}

		if !partsStructureModel.IsParent() {
			logger.Set(c).Debugf("TraceID: %#v does not have any child parts, so we get the upstream CFP value associated with the trade information", traceID.String())
			// Pattern B. For child parts
			// B-1. get trade information of a part registered as a child part
			trade, err := u.r.GetTradeByDownstreamTraceID(traceID.String())
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					logger.Set(c).Debugf("TraceID of the part: %#v is not processed because the trade information is not registered", traceID.String())
					cfpResponse := traceability.NewEmptyCfpResponse(traceID)
					cfps, err := cfpResponse.ToModels()
					if err != nil {
						logger.Set(c).Errorf(err.Error())

						return nil, err
					}
					res = append(res, cfps...)
					continue
				}
				logger.Set(c).Errorf(err.Error())

				return nil, err
			}
			// B-2. get CFP from the upstream traceID of trade
			if trade.UpstreamTraceID == nil {
				logger.Set(c).Debugf("Not processed because the upstream TraceID of the part is not registered.")
				continue
			}
			cfps, err := u.r.ListCFPsByTraceID(trade.UpstreamTraceID.String())
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, err
			}

			if len(cfps) == 0 {
				logger.Set(c).Debugf("Not processed because the CFP information for TraceID: %#v in the upstream of the part is not registered.", trade.UpstreamTraceID.String())
				continue
			}

			// B-3. convert the CFP value of the trading partner to xxxResponse format (cfpID is a zero value and traceID is a downstream component to recreate the cfp model)
			cfpResponse, err := cfps.MakeCfpResponse(traceID)
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, err
			}
			ms, err := cfpResponse.ToModels()
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, err
			}

			res = append(res, ms...)
			continue
		}

		// Pattern C. For parent component
		logger.Set(c).Debugf("TraceID: %#v has child parts, so it gets its own CFP value and the child parts' CFP values", traceID.String())

		// Get child parts
		getPartsStructureInput := traceability.GetPartsStructureInput{OperatorID: getCfpInput.OperatorID.String(), TraceID: traceID}
		partsStructure, err := u.r.GetPartsStructure(getPartsStructureInput)
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, err
		}

		// C-1. get CFP of parent parts
		parentCfps, err := u.r.ListCFPsByTraceID(partsStructure.ParentPartsEntity.TraceID.String())
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, err
		}
		if len(parentCfps) == 0 {
			logger.Set(c).Debugf("Not processed because CFP information for TraceID: %#v in parent parts is not registered", partsStructure.ParentPartsEntity.TraceID.String())
			continue
		}
		parentCfpSet, err := traceability.NewCfpEntityModelSetFromCfpEntityModels(parentCfps, true)
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, err
		}

		for _, childParts := range partsStructure.ChildrenPartsEntity {
			var childCfps traceability.CfpEntityModels
			// C-2. get CFP from transaction information of child parts
			if childParts.TerminatedFlag {
				// C-2-A. If terminated, get CFP directly
				childCfps, err = u.r.ListCFPsByTraceID(childParts.TraceID.String())
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						logger.Set(c).Debugf("TraceID of child parts: %#v is not processed because the CFP is not registered", childParts.TraceID.String())
						continue
					}
					logger.Set(c).Errorf(err.Error())

					return nil, err
				}
			} else {
				// C-2-B. If not terminated, get upstream traceID from trade information
				trade, err := u.r.GetTradeByDownstreamTraceID(childParts.TraceID.String())
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						logger.Set(c).Debugf("TraceID of child parts: %#v is not processed because the trade information is not registered", childParts.TraceID.String())
						continue
					}
					logger.Set(c).Errorf(err.Error())

					return nil, err
				}
				if trade.UpstreamTraceID == nil {
					logger.Set(c).Debugf("TraceID of child parts: %#v is not processed because the trade information is not registered", childParts.TraceID.String())
					continue
				}
				childCfps, err = u.r.ListCFPsByTraceID(trade.UpstreamTraceID.String())
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						logger.Set(c).Debugf("Not processed because the CFP information for TraceID: %#v on the upstream of the child parts is not registered", trade.UpstreamTraceID.String())
						continue
					}
					logger.Set(c).Errorf(err.Error())

					return nil, err
				}
			}
			// C-3. Add the CFP value of a group of child parts to the CFP value of the parent part
			if childCfps.GetPreProductionCfp() != nil && childCfps.GetPreProductionCfp().GhgEmission != nil {
				ghgEmission := *parentCfpSet.GetPreComponentTotalCfp().GhgEmission + *childCfps.GetPreProductionCfp().GhgEmission
				parentCfpSet.GetPreComponentTotalCfp().GhgEmission = &ghgEmission
			}
			if childCfps.GetMainProductionCfp() != nil && childCfps.GetMainProductionCfp().GhgEmission != nil {
				ghgEmission := *parentCfpSet.GetMainComponentTotalCfp().GhgEmission + *childCfps.GetMainProductionCfp().GhgEmission
				parentCfpSet.GetMainComponentTotalCfp().GhgEmission = &ghgEmission
			}
		}
		ms, err := parentCfpSet.ToModels()
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, err
		}

		res = append(res, ms...)
	}
	return res, nil
}

// PutCfp
// Summary: This is function which put a list of cfp.
// input: c(echo.Context) echo context
// input: putCfpInputs(traceability.PutCfpInputs) PutCfpInputs object
// input: operatorID(string) ID of the operator
// output: ([]traceability.CfpModel) list of CfpModel
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *cfpUsecase) PutCfp(c echo.Context, putCfpInputs traceability.PutCfpInputs, operatorID string) ([]traceability.CfpModel, common.ResponseHeaders, error) {
	cfpModels, err := putCfpInputs.ToModels()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return nil, common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}

	cfpID, err := cfpModels.GetCommonCfpID()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, common.ResponseHeaders{}, err
	}
	traceID, err := cfpModels.GetCommonTraceID()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, common.ResponseHeaders{}, err
	}

	if cfpID == nil {
		cfps, err := u.r.ListCFPsByTraceID(traceID.String())
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, common.ResponseHeaders{}, err
		}
		if len(cfps) != 0 {
			logger.Set(c).Errorf(common.TraceIDAlreadyHasCfpsError(traceID.String()))

			return nil, common.ResponseHeaders{}, fmt.Errorf(common.TraceIDAlreadyHasCfpsError(traceID.String()))
		}

		es := traceability.GenerateCfpEntitisFromModels(cfpModels)
		resCfpEs, err := u.r.BatchCreateCFP(es)
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, common.ResponseHeaders{}, err
		}

		trades, err := u.r.ListTradeByUpstreamTraceID(traceID.String())
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, common.ResponseHeaders{}, err
		}
		for _, trade := range trades {
			tradePart, err := u.r.GetPartByTraceID((*trade.UpstreamTraceID).String())
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, common.ResponseHeaders{}, err
			}

			putTradeResponseInput := traceability.PutTradeResponseInput{
				TradeID: *trade.TradeID,
				TraceID: *trade.UpstreamTraceID,
			}

			var tradeTreeStatus traceability.TradeTreeStatus
			if tradePart.TerminatedFlag {
				tradeTreeStatus = traceability.TradeTreeStatusTerminated
			} else {
				tradeTreeStatus = traceability.TradeTreeStatusUnterminated
			}
			cfpResponseStatusComplete := traceability.CfpResponseStatusComplete
			requestStatusValue := traceability.RequestStatus{
				CfpResponseStatus: &cfpResponseStatusComplete,
				TradeTreeStatus:   &tradeTreeStatus,
				// set fixed value due to complexity of calculateion process
				CompletedCount: common.IntPtr(1),
			}

			_, err = u.r.PutTradeResponse(putTradeResponseInput, requestStatusValue)
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, common.ResponseHeaders{}, err
			}
		}
		models, err := resCfpEs.ToModels()
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, common.ResponseHeaders{}, err
		}
		return models, common.ResponseHeaders{}, nil
	} else {
		res := make(traceability.CfpModels, len(cfpModels))
		for i, m := range cfpModels {
			e, err := u.r.GetCFP(m.CfpID.String(), m.CfpType)
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, common.ResponseHeaders{}, err
			}
			e.Update(
				m.GhgEmission,
				m.GhgDeclaredUnit.ToString(),
				m.DqrType,
				m.DqrValue.TeR,
				m.DqrValue.GeR,
				m.DqrValue.TiR,
			)
			e, err = u.r.PutCFP(e)
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, common.ResponseHeaders{}, err
			}

			r, err := e.ToModel()
			if err != nil {
				logger.Set(c).Errorf(err.Error())

				return nil, common.ResponseHeaders{}, err
			}
			res[i] = r

		}
		return res, common.ResponseHeaders{}, nil
	}
}
