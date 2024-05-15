package usecase

import (
	"errors"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/model/traceability/traceabilityentity"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/labstack/echo/v4"
)

// partsTraceabilityUsecase
// Summary: This struct defines traceability use cases for the parts.
type partsTraceabilityUsecase struct {
	TraceabilityRepository repository.TraceabilityRepository
}

// NewPartsTraceabilityUsecase
// Summary: This function creates a new partsTraceabilityUsecase.
// input: r(repository.TraceabilityRepository) traceability api repository
// output: (IPartsUsecase) use case interface
func NewPartsTraceabilityUsecase(r repository.TraceabilityRepository) IPartsUsecase {
	return &partsTraceabilityUsecase{r}
}

// GetPartsList
// Summary: This function gets a list of parts.
// input: c(echo.Context) echo context
// input: getPartsModel(traceability.GetPartsModel) get parts model
// output: partsModels([]traceability.PartsModel) list of partsModel
// output: after(*string) next id
// output: err(error) Error object
func (u *partsTraceabilityUsecase) GetPartsList(c echo.Context, getPartsModel traceability.GetPartsModel) (partsModels []traceability.PartsModel, after *string, err error) {
	request := traceabilityentity.GetPartsRequest{
		OperatorID:       getPartsModel.OperatorID,
		TraceID:          getPartsModel.TraceID,
		PartsItem:        getPartsModel.PartsName,
		SupportPartsItem: nil,
		PlantID:          getPartsModel.PlantID,
		ParentFlag:       getPartsModel.ParentFlag,
		After:            common.UUIDPtrToStringPtr(getPartsModel.After),
	}

	res, err := u.TraceabilityRepository.GetParts(c, request, getPartsModel.Limit)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}

		return nil, nil, err
	}

	partsModels, err = res.ToModel()
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return nil, nil, err
	}
	after = res.NextPrt()
	return partsModels, after, nil
}
