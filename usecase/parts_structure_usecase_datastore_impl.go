package usecase

import (
	"data-spaces-backend/domain/common"
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
// input: getPartsStructureInput(traceability.getPartsStructureInput) model for partsStructure retrieval
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (u *partsStructureUsecase) GetPartsStructure(c echo.Context, getPartsStructureInput traceability.GetPartsStructureInput) (traceability.PartsStructureModel, error) {
	partsStructures, err := u.OuranosRepository.GetPartsStructure(getPartsStructureInput)
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
// input: putPartsStructureInput(traceability.PutPartsStructureInput) model for partsStructure retrieval
// output: (traceability.PartsStructureModel) partsStructure model
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *partsStructureUsecase) PutPartsStructure(c echo.Context, putPartsStructureInput traceability.PutPartsStructureInput) (traceability.PartsStructureModel, common.ResponseHeaders, error) {
	parentPartsModel, err := putPartsStructureInput.ParentPartsInput.ToModel()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return traceability.PartsStructureModel{}, common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}
	var childrenPartsModel []traceability.PartsModel
	if putPartsStructureInput.HasChild() {
		childrenPartsModel, err = putPartsStructureInput.ChildrenPartsInput.ToModels()
		if err != nil {
			logger.Set(c).Warnf(err.Error())
			errDetails := err.Error()

			return traceability.PartsStructureModel{}, common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
		}
	}
	partsStructureModel := traceability.PartsStructureModel{
		ParentPartsModel:   &parentPartsModel,
		ChildrenPartsModel: childrenPartsModel,
	}

	partsStructure, err := u.OuranosRepository.PutPartsStructure(partsStructureModel)
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return traceability.PartsStructureModel{}, common.ResponseHeaders{}, err
	}
	partsStructureModels, err := partsStructure.ToModel()
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return traceability.PartsStructureModel{}, common.ResponseHeaders{}, err
	}
	return partsStructureModels, common.ResponseHeaders{}, nil
}
