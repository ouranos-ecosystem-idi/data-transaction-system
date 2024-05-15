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

// partsStructureTraceabilityUsecase
// Summary: This struct defines traceability use cases for the partsStructure.
type partsStructureTraceabilityUsecase struct {
	TraceabilityRepository repository.TraceabilityRepository
}

// NewPartsStructureTraceabilityUsecase
// Summary: This function creates a new partsStructureTraceabilityUsecase.
// input: r(repository.TraceabilityRepository) traceability repository
// output: (IPartsStructureUsecase) partsStructure use case interface
func NewPartsStructureTraceabilityUsecase(r repository.TraceabilityRepository) IPartsStructureUsecase {
	return &partsStructureTraceabilityUsecase{r}
}

// GetPartsStructure
// Summary: This function get request and response partsStructure.
// input: c(echo.Context) echo context
// input: getInput(traceability.GetPartsStructureModel) model for partsStructure retrieval
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (u *partsStructureTraceabilityUsecase) GetPartsStructure(c echo.Context, getPartsStructureModel traceability.GetPartsStructureModel) (traceability.PartsStructureModel, error) {
	request := traceabilityentity.GetPartsStructuresRequest{
		OperatorID:    getPartsStructureModel.OperatorID,
		ParentTraceID: getPartsStructureModel.TraceID.String(),
	}

	res, err := u.TraceabilityRepository.GetPartsStructures(c, request)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}
		return traceability.PartsStructureModel{}, err
	}

	return res.ToModel()
}

// PutPartsStructure
// Summary: This function get request and response partsStructure.
// input: c(echo.Context) echo context
// input: partsStructureModel(traceability.PartsStructureModel) partsStructure model
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (u *partsStructureTraceabilityUsecase) PutPartsStructure(c echo.Context, partsStructureModel traceability.PartsStructureModel) (traceability.PartsStructureModel, error) {
	req := traceabilityentity.NewPostPartsStructureRequestFromModel(partsStructureModel)

	rt, err := u.TraceabilityRepository.PostPartsStructures(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}
		return traceability.PartsStructureModel{}, err
	}

	partsStructureModel.ParentPartsModel.TraceID, err = rt.ParentTraceID()
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return traceability.PartsStructureModel{}, err
	}

	for i, child := range partsStructureModel.ChildrenPartsModel {
		childTraceID, err := rt.ChildTraceID(child.PartsName, child.SupportPartsName)
		if err != nil {
			logger.Set(c).Errorf(err.Error())
			return traceability.PartsStructureModel{}, err
		}
		partsStructureModel.ChildrenPartsModel[i].TraceID = childTraceID
	}

	return partsStructureModel, nil
}
