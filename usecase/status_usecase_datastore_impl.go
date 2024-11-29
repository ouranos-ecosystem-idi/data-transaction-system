package usecase

import (
	"data-spaces-backend/domain/common"
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// statusUsecase
// Summary: This is structure which defines statusUsecase.
type statusUsecase struct {
	OuranosRepository repository.OuranosRepository
}

// NewStatusUsecase
// Summary: This is function to create new statusUsecase.
// input: r(repository.OuranosRepository) repository interface
// output: (IStatusUsecase) use case interface
func NewStatusUsecase(r repository.OuranosRepository) IStatusUsecase {
	return &statusUsecase{r}
}

// GetStatus
// Summary: This is function which get a list of request and response status.
// input: c(echo.Context) echo context
// input: getStatusInput(traceability.GetStatusInput) GetStatusInput object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func (u *statusUsecase) GetStatus(c echo.Context, getStatusInput traceability.GetStatusInput) ([]traceability.StatusModel, *string, error) {
	statusID := common.UUIDPtrToStringPtr(getStatusInput.StatusID)
	traceID := common.UUIDPtrToStringPtr(getStatusInput.TraceID)
	statuses, err := u.OuranosRepository.GetStatus(getStatusInput.OperatorID.String(), getStatusInput.Limit, statusID, traceID, getStatusInput.StatusTarget.ToString())
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return []traceability.StatusModel{}, nil, err
	}

	count, err := u.OuranosRepository.CountStatus(getStatusInput.OperatorID.String(), statusID, traceID, getStatusInput.StatusTarget.ToString())
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return []traceability.StatusModel{}, nil, err
	}
	// NOTE: Return dummy since it is a fake implementation.
	var dummyAfterPtr *string
	if count > getStatusInput.Limit {
		dummyAfter := uuid.New().String()
		dummyAfterPtr = &dummyAfter
	}
	statusModels, err := statuses.ToModels(traceability.PathStatus)
	if err != nil {
		return nil, nil, err
	}
	return statusModels, dummyAfterPtr, nil
}

// PutStatusCancel
// Summary: This is function which cancels the status of a request.
// input: c(echo.Context) echo context
// input: putStatusInput(traceability.PutStatusInput) PutStatusInput object
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *statusUsecase) PutStatusCancel(c echo.Context, putStatusInput traceability.PutStatusInput) (common.ResponseHeaders, error) {
	statusModel, err := putStatusInput.ToModel()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}

	operatorID := c.Get("operatorID").(string)
	err = u.OuranosRepository.PutStatusCancel(statusModel.StatusID.String(), operatorID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return common.ResponseHeaders{}, err
	}

	return common.ResponseHeaders{}, nil
}

// PutStatusReject
// Summary: This is function which rejects the status of a request.
// input: c(echo.Context) echo context
// input: putStatusInput(traceability.PutStatusInput) PutStatusInput object
// output: (common.ResponseHeaders) response headers
// output: (error) error object
func (u *statusUsecase) PutStatusReject(c echo.Context, putStatusInput traceability.PutStatusInput) (common.ResponseHeaders, error) {
	statusModel, err := putStatusInput.ToModel()
	if err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := err.Error()

		return common.ResponseHeaders{}, common.NewCustomError(common.CustomErrorCode400, common.Err400Validation, &errDetails, common.HTTPErrorSourceDataspace)
	}

	operatorID := c.Get("operatorID").(string)
	_, err = u.OuranosRepository.PutStatusReject(statusModel.StatusID.String(), statusModel.ReplyMessage, operatorID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return common.ResponseHeaders{}, err
	}

	return common.ResponseHeaders{}, nil
}
