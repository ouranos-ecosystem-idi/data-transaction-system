package usecase

import (
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/labstack/echo/v4"
)

// partsStructureUsecase
// Summary: This is structure which defines partsStructureUsecase.
type partsStructureUsecase struct {
	OuranosRepository repository.OuranosRepository
}

// NewPartsStructureDatastoreUsecase
// Summary: This is function to create new partsStructureUsecase.
// input: r(repository.OuranosRepository) repository interface
// output: (IPartsStructureUsecase) use case interface
func NewPartsStructureDatastoreUsecase(r repository.OuranosRepository) IPartsStructureUsecase {
	return &partsStructureUsecase{r}
}

// GetPartsStructure
// Summary: This is function which get request and response partsStructure.
// input: c(echo.Context) echo context
// input: getPartsStructureModel(traceability.GetPartsStructureModel) model for partsStructure retrieval
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (u *partsStructureUsecase) GetPartsStructure(c echo.Context, getPartsStructureModel traceability.GetPartsStructureModel) (traceability.PartsStructureModel, error) {
	partsStructures, err := u.OuranosRepository.GetPartsStructure(getPartsStructureModel)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.PartsStructureModel{}, err
	}

	m, err := partsStructures.ToModel()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.PartsStructureModel{}, err
	}

	return m, nil
}

// PutPartsStructure
// Summary: This is function which get request and response partsStructure.
// input: c(echo.Context) echo context
// input: partsStructureModel(traceability.PartsStructureModel) partsStructure model
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (u *partsStructureUsecase) PutPartsStructure(c echo.Context, partsStructureModel traceability.PartsStructureModel) (traceability.PartsStructureModel, error) {

	partsStructure, err := u.OuranosRepository.PutPartsStructure(partsStructureModel)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return traceability.PartsStructureModel{}, err
	}
	partsStructureModels, err := partsStructure.ToModel()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.PartsStructureModel{}, err
	}
	return partsStructureModels, nil
}
