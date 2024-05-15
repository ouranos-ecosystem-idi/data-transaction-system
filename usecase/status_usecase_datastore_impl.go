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
// input: getStatusModel(traceability.GetStatusModel) GetStatusModel object
// output: ([]traceability.StatusModel) list of StatusModel
// output: (*string) next id
// output: (error) error object
func (u *statusUsecase) GetStatus(c echo.Context, getStatusModel traceability.GetStatusModel) ([]traceability.StatusModel, *string, error) {
	statusID := common.UUIDPtrToStringPtr(getStatusModel.StatusID)
	traceID := common.UUIDPtrToStringPtr(getStatusModel.TraceID)
	statuses, err := u.OuranosRepository.GetStatus(getStatusModel.OperatorID.String(), getStatusModel.Limit, statusID, traceID, getStatusModel.StatusTarget.ToString())
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return []traceability.StatusModel{}, nil, err
	}

	count, err := u.OuranosRepository.CountStatus(getStatusModel.OperatorID.String(), statusID, traceID, getStatusModel.StatusTarget.ToString())
	if err != nil {
		logger.Set(c).Errorf(err.Error())
		return []traceability.StatusModel{}, nil, err
	}
	// NOTE: Return dummy since it is a fake implementation.
	var dummyAfterPtr *string
	if count > getStatusModel.Limit {
		dummyAfter := uuid.New().String()
		dummyAfterPtr = &dummyAfter
	}
	statusModels, err := statuses.ToModels()
	if err != nil {
		return nil, nil, err
	}
	return statusModels, dummyAfterPtr, nil
}

// PutStatusCancel
// Summary: This is function which cancels the status of a request.
// input: c(echo.Context) echo context
// input: statusModel(traceability.StatusModel) StatusModel object
// output: (error) error object
func (u *statusUsecase) PutStatusCancel(c echo.Context, statusModel traceability.StatusModel) error {
	operatorID := c.Get("operatorID").(string)
	_, err := u.OuranosRepository.PutStatusCancel(statusModel.StatusID.String(), operatorID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return err
	}

	return nil
}

// PutStatusReject
// Summary: This is function which rejects the status of a request.
// input: c(echo.Context) echo context
// input: statusModel(traceability.StatusModel) StatusModel object
// output: (error) error object
func (u *statusUsecase) PutStatusReject(c echo.Context, statusModel traceability.StatusModel) error {
	operatorID := c.Get("operatorID").(string)
	_, err := u.OuranosRepository.PutStatusReject(statusModel.StatusID.String(), statusModel.ReplyMessage, operatorID)
	if err != nil {
		logger.Set(c).Errorf(err.Error())

		return err
	}

	return nil
}
