package usecase

import (
	"errors"
	"fmt"
	"time"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// tradeUsecase
// Summary: This is structure which defines tradeUsecase.
type tradeUsecase struct {
	OuranosRepository repository.OuranosRepository
}

// NewTradeUsecase
// Summary: This is function to create new TradeUsecase.
// input: r(repository.OuranosRepository) repository interface
// output: (ITradeUsecase) usecase interface
func NewTradeUsecase(r repository.OuranosRepository) ITradeUsecase {
	return &tradeUsecase{r}
}

// GetTradeRequest
// Summary: This is function which get a list of trade and response status.
// input: c(echo.Context) echo context
// input: getTradeRequestInput(traceability.GetTradeRequestInput) GetTradeRequestInput object
// output: ([]traceability.TradeModel) list of TradeModel
// output: (*string) next id
// output: (error) error object
func (u *tradeUsecase) GetTradeRequest(c echo.Context, getTradeRequestInput traceability.GetTradeRequestInput) ([]traceability.TradeModel, *string, error) {
	es, err := u.OuranosRepository.GetTradeRequest(
		getTradeRequestInput.OperatorID.String(),
		getTradeRequestInput.Limit,
		common.UUIDsToStrings(getTradeRequestInput.TraceIDs),
	)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return []traceability.TradeModel{}, nil, err
	}
	count, err := u.OuranosRepository.CountTradeRequest(getTradeRequestInput.OperatorID.String())
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return []traceability.TradeModel{}, nil, err
	}

	var dummyAfterPtr *string
	if count > getTradeRequestInput.Limit {
		// If the total number of cases is greater than the limit, set a value in After.
		dummyAfter := uuid.New().String()
		dummyAfterPtr = &dummyAfter
	}

	return es.ToModels(), dummyAfterPtr, nil
}

// GetTradeResponse
// Summary: This is function which get a list of trade and response status.
// input: c(echo.Context) echo context
// input: getTradeResponseInput(traceability.GetTradeResponseInput) GetTradeResponseInput object
// output: ([]traceability.TradeResponseModel) list of TradeResponseModel
// output: (*string) next id
// output: (error) error object
func (u *tradeUsecase) GetTradeResponse(c echo.Context, getTradeResponseInput traceability.GetTradeResponseInput) ([]traceability.TradeResponseModel, *string, error) {
	trades, err := u.OuranosRepository.GetTradeResponse(
		getTradeResponseInput.OperatorID.String(),
		getTradeResponseInput.Limit,
	)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}
	count, err := u.OuranosRepository.CountTradeResponse(getTradeResponseInput.OperatorID.String())
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}

	res := make([]traceability.TradeResponseModel, len(trades))
	for i, trade := range trades {
		status, err := u.OuranosRepository.GetStatusByTradeID(trade.TradeID.String())
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, nil, err
		}

		downstreamParts, err := u.OuranosRepository.GetPartByTraceID(trade.DownstreamTraceID.String())
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, nil, err
		}
		statusModel, err := status.ToModel()
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, nil, err
		}
		partsModel, err := downstreamParts.ToModelWithMasking()
		if err != nil {
			logger.Set(c).Errorf(err.Error())

			return nil, nil, err
		}

		tr := traceability.TradeResponseModel{
			TradeModel:  trade.ToModel(),
			StatusModel: statusModel,
			PartsModel:  partsModel,
		}
		res[i] = tr
	}
	var dummyAfterPtr *string
	if count > getTradeResponseInput.Limit {
		// If the total number of cases is greater than the limit, set a value in After.
		dummyAfter := uuid.New().String()
		dummyAfterPtr = &dummyAfter
	}

	return res, dummyAfterPtr, nil
}

// PutTradeRequest
// Summary: This is function which put trade request with TradeRequestModel.
// input: c(echo.Context) echo context
// input: tradeRequestModel(traceability.TradeRequestModel) TradeRequestModel object
// output: (traceability.TradeRequestModel) TradeRequestModel object
// output: (error) error object
func (u *tradeUsecase) PutTradeRequest(c echo.Context, tradeRequestModel traceability.TradeRequestModel) (traceability.TradeRequestModel, error) {
	// If TradeID is Null, generate a new ID
	if tradeRequestModel.TradeModel.TradeID == nil || *tradeRequestModel.TradeModel.TradeID == uuid.Nil {
		tradeID, _ := uuid.NewRandom()
		tradeRequestModel.TradeModel.TradeID = &tradeID

		statusID, _ := uuid.NewRandom()

		tradeRequestModel.StatusModel.StatusID = statusID
		tradeRequestModel.StatusModel.TradeID = tradeID
	}

	fmt.Printf("Put Usecase TradeBody: %+v\n", tradeRequestModel)

	now := time.Now()
	// Convert to timestamp compliant with "ISO8601" compliant string and "UTC" for TradeDate
	utcNow := now.UTC()
	iso8601Format := "2006-01-02T15:04:05Z" // ISO8601
	isoUtcTime := utcNow.Format(iso8601Format)

	tradeEntityModel := traceability.TradeEntityModel{
		TradeID:              tradeRequestModel.TradeModel.TradeID,
		DownstreamOperatorID: tradeRequestModel.TradeModel.DownstreamOperatorID,
		UpstreamOperatorID:   &tradeRequestModel.TradeModel.UpstreamOperatorID,
		DownstreamTraceID:    tradeRequestModel.TradeModel.DownstreamTraceID,
		UpstreamTraceID:      tradeRequestModel.TradeModel.UpstreamTraceID,
		TradeDate:            &isoUtcTime,
		CreatedAt:            now,
		CreatedUserID:        "sample",
		UpdatedAt:            now,
		UpdatedUserID:        "sample",
	}

	tradeRequestModel.StatusModel.RequestType = traceability.RequestTypeCFP.ToString()
	statusEntityModel := traceability.StatusEntityModel{
		StatusID:          tradeRequestModel.StatusModel.StatusID,
		TradeID:           *tradeRequestModel.TradeModel.TradeID,
		CfpResponseStatus: traceability.CfpResponseStatusPending.ToString(),
		TradeTreeStatus:   traceability.TradeTreeStatusUnterminated.ToString(),
		Message:           tradeRequestModel.StatusModel.Message,
		RequestType:       tradeRequestModel.StatusModel.RequestType,
		CreatedUserId:     "sample",
		UpdatedAt:         now,
		UpdatedUserId:     "sample",
	}

	tradeRequestEntityModel := traceability.TradeRequestEntityModel{
		TradeEntityModel:  tradeEntityModel,
		StatusEntityModel: statusEntityModel,
	}

	fmt.Printf("Put Usecase tradeRequestEntityModel: %+v\n", tradeRequestEntityModel)

	res, err := u.OuranosRepository.PutTradeRequest(tradeRequestEntityModel)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return tradeRequestModel, err
	}

	statusModel, err := res.ToModel()
	if err != nil {
		return tradeRequestModel, err
	}
	return statusModel, nil
}

// PutTradeResponse
// Summary: This is function which put trade response with PutTradeResponse.
// input: c(echo.Context) echo context
// input: putTradeResponseInput(traceability.PutTradeResponseInput) PutTradeResponseInput object
// output: (traceability.TradeModel) TradeModel object
// output: (error) error object
func (u *tradeUsecase) PutTradeResponse(c echo.Context, putTradeResponseInput traceability.PutTradeResponseInput) (traceability.TradeModel, error) {
	CfpResponseStatus := traceability.CfpResponseStatusComplete
	TradeTreeStatus := traceability.TradeTreeStatusTerminated
	requestStatusValue := traceability.RequestStatus{
		CfpResponseStatus: CfpResponseStatus,
		TradeTreeStatus:   TradeTreeStatus,
	}

	_, err := u.OuranosRepository.GetCFPInformation(putTradeResponseInput.TraceID.String())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			trade, err := u.OuranosRepository.GetTrade(putTradeResponseInput.TradeID.String())
			if err != nil {
				logger.Set(nil).Error(err.Error())
				return traceability.TradeModel{}, err
			}
			_, err = u.OuranosRepository.GetCFPInformation(trade.DownstreamTraceID.String())
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					requestStatusValue.CfpResponseStatus = traceability.CfpResponseStatusPending
				} else {
					logger.Set(nil).Error(err.Error())
					return traceability.TradeModel{}, err
				}
			}
		} else {
			logger.Set(nil).Error(err.Error())

			return traceability.TradeModel{}, err
		}
	}

	tradePart, err := u.OuranosRepository.GetPartByTraceID(putTradeResponseInput.TraceID.String())
	if err != nil {
		logger.Set(nil).Error(err.Error())

		return traceability.TradeModel{}, err
	}
	if tradePart.TerminatedFlag {
		requestStatusValue.TradeTreeStatus = traceability.TradeTreeStatusUnterminated
	}

	fmt.Printf("requestStatusValue: %+v\n", requestStatusValue)

	trade, err := u.OuranosRepository.PutTradeResponse(putTradeResponseInput, requestStatusValue)
	if err != nil {
		logger.Set(nil).Error(err.Error())

		return traceability.TradeModel{}, err
	}

	return trade.ToModel(), nil
}
