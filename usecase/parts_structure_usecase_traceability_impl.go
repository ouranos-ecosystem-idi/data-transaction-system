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
// input: getPartsStructureInput(traceability.GetPartsStructureInput) model for partsStructure retrieval
// output: (traceability.PartsStructureModel) partsStructure model
// output: (error) error object
func (u *partsStructureTraceabilityUsecase) GetPartsStructure(c echo.Context, getPartsStructureInput traceability.GetPartsStructureInput) (traceability.PartsStructureModel, error) {
	request := traceabilityentity.GetPartsStructuresRequest{
		OperatorID:    getPartsStructureInput.OperatorID,
		ParentTraceID: getPartsStructureInput.TraceID.String(),
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
// input: putPartsStructureInput(traceability.PutPartsStructureInput) model for partsStructure retrieval
// output: (traceability.PartsStructureModel) partsStructure model
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *partsStructureTraceabilityUsecase) PutPartsStructure(c echo.Context, putPartsStructureInput traceability.PutPartsStructureInput) (traceability.PartsStructureModel, common.ResponseHeaders, error) {
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

	req := traceabilityentity.NewPostPartsStructureRequestFromModel(partsStructureModel)

	rt, headers, err := u.TraceabilityRepository.PostPartsStructures(c, req)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) && customErr.IsWarn() {
			logger.Set(c).Warnf(err.Error())
		} else {
			logger.Set(c).Errorf(err.Error())
		}
		return traceability.PartsStructureModel{}, common.ResponseHeaders{}, err
	}

	partsStructureModel.ParentPartsModel.TraceID, err = rt.ParentTraceID()
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return traceability.PartsStructureModel{}, common.ResponseHeaders{}, err
	}

	for i, child := range partsStructureModel.ChildrenPartsModel {
		childTraceID, err := rt.ChildTraceID(child.PartsName, child.SupportPartsName, *common.UUIDPtrToStringPtr(child.PlantID))
		if err != nil {
			logger.Set(c).Errorf(err.Error())
			return traceability.PartsStructureModel{}, common.ResponseHeaders{}, err
		}
		partsStructureModel.ChildrenPartsModel[i].TraceID = childTraceID
	}

	return partsStructureModel, headers, nil
}
