package usecase

import (
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
// input: getPartsModel(traceability.GetPartsModel) get partsModel
// output: ([]traceability.PartsModel) list of partsModel
// output: (*string) next id
// output: (error) Error object
func (u *partsUsecase) GetPartsList(c echo.Context, getPartsModel traceability.GetPartsModel) ([]traceability.PartsModel, *string, error) {
	partsList, err := u.OuranosRepository.ListParts(getPartsModel)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return nil, nil, err
	}
	partsList = partsList.MaskAmountRequired()

	count, err := u.OuranosRepository.CountPartsList(getPartsModel)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return nil, nil, err
	}

	var dummyAfterPtr *string
	if count > getPartsModel.Limit {
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
