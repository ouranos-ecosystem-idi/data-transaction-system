package usecase

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// partsUsecase
// Summary: This struct defines traceability use cases for the parts.
type partsUsecase struct {
	OuranosRepository repository.OuranosRepository
}

// NewPartsUsecase
// Summary: This function creates a new partsUsecase.
// input: r(repository.OuranosRepository) ouranos api repository
// output: (IPartsUsecase) use case interface
func NewPartsUsecase(r repository.OuranosRepository) IPartsUsecase {
	return &partsUsecase{r}
}

// GetPartsList
// Summary: This function gets a partsList.
// input: c(echo.Context) echo context
// input: getPartsInput(traceability.GetPartsInput) getPartsInput object
// output: ([]traceability.PartsModel) list of partsModel
// output: (*string) next id
// output: (error) Error object
func (u *partsUsecase) GetPartsList(c echo.Context, getPartsInput traceability.GetPartsInput) ([]traceability.PartsModel, *string, error) {
	partsList, err := u.OuranosRepository.ListParts(getPartsInput)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}
	partsList = partsList.MaskAmountRequired()

	count, err := u.OuranosRepository.CountPartsList(getPartsInput)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return nil, nil, err
	}

	var dummyAfterPtr *string
	if count > getPartsInput.Limit {
		dummyAfter := uuid.New().String()
		dummyAfterPtr = &dummyAfter
	}
	partsListModels, err := partsList.ToModels()
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return nil, nil, err
	}
	return partsListModels, dummyAfterPtr, nil
}

// DeleteParts
// Summary: This function deletes a parts.
// input: c(echo.Context) echo context
// input: deletePartsInput(traceability.DeletePartsInput) deletePartsInput object
// output: (error) Error object
func (u *partsUsecase) DeleteParts(c echo.Context, deletePartsInput traceability.DeletePartsInput) (common.ResponseHeaders, error) {
	operatorID := c.Get("operatorID").(string)

	// If deleting parts is not exist, this API returns error.
	parts, err := u.OuranosRepository.GetPartByTraceID(deletePartsInput.TraceID)
	if err != nil {
		err = createTraceabilityError("MSGAECP0013", "指定された部品は存在しません。", []uuid.UUID{})
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	if *common.UUIDPtrToStringPtr(&parts.OperatorID) != operatorID {
		err = createTraceabilityError("MSGAECP0025", "認証情報と事業者識別子が一致しません。", []uuid.UUID{})
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	// If particular parts includes deleting parts, this API returns error.
	parents, err := u.OuranosRepository.ListParentPartsStructureByTraceId(deletePartsInput.TraceID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	} else if len(parents) > 0 {
		uuids := []uuid.UUID{parents[0].TraceID}
		err = createTraceabilityError("MSGAECP0014", "指定された部品は部品構成が存在するため削除できません。", uuids)
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	// If deleting parts consist of multiple parts, this API returns error.
	children, err := u.OuranosRepository.ListChildPartsStructureByTraceId(deletePartsInput.TraceID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	} else if len(children) > 0 {
		var uuids []uuid.UUID
		for _, child := range children {
			uuids = append(uuids, child.TraceID)
		}
		err = createTraceabilityError("MSGAECP0014", "指定された部品は部品構成が存在するため削除できません。", uuids)
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	// If deleting parts has already offered trade request, this API returns error.
	downstreams, err := u.OuranosRepository.ListTradeByDownstreamTraceID(deletePartsInput.TraceID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	} else if len(downstreams) > 0 {
		var uuids []uuid.UUID
		for _, downstream := range downstreams {
			uuids = append(uuids, *downstream.TradeID)
		}
		err = createTraceabilityError("MSGAECP0015", "指定された部品は依頼済みのため削除できません。", uuids)
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	// If deleting parts has already answered trade response, this API returns error.
	upstreams, err := u.OuranosRepository.ListTradeByUpstreamTraceID(deletePartsInput.TraceID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	} else if len(upstreams) > 0 {
		var uuids []uuid.UUID
		for _, upstream := range upstreams {
			uuids = append(uuids, *upstream.TradeID)
		}
		err = createTraceabilityError("MSGAECP0016", "指定された部品は受領済みの依頼に紐づいているため削除できません。", uuids)
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	err = u.OuranosRepository.DeletePartsWithCFP(deletePartsInput.TraceID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return common.ResponseHeaders{}, err
	}

	return common.ResponseHeaders{}, nil
}

// createTraceabilityError
// Summary: This function creates pseudo error of TraceabilityAPI
// input: errorCode(string) errorCode
// input: errorDescription(string) errorDescription
// input: relevantUUIDs([]uuid.UUID) array of UUID
// output: (error) Error object
func createTraceabilityError(errorCode string, errorDescription string, relevantUUIDs []uuid.UUID) error {
	err := common.TraceabilityAPIErrorDetailDelete{
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
	}

	if len(relevantUUIDs) > 0 {
		relevantData := common.UUIDsToStrings(relevantUUIDs)
		err.RelevantData = &relevantData
	}

	traceabilityError := common.TraceabilityAPIErrorDelete{
		Errors: []common.TraceabilityAPIErrorDetailDelete{
			err,
		},
	}
	return traceabilityError.ToCustomError(400)
}
