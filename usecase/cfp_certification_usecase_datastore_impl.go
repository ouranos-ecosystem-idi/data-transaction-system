package usecase

import (
	"data-spaces-backend/domain/model/traceability"
	"data-spaces-backend/domain/repository"

	"github.com/labstack/echo/v4"
)

// cfpCertificationUsecase
// Summary: This is structure which defines cfpCertificationUsecase.
type cfpCertificationUsecase struct {
	r repository.OuranosRepository
}

// NewCfpCertificationUsecase
// Summary: This is function to create new cfpCertificationUsecase.
// input: r(repository.OuranosRepository) repository interface
// output: (ICfpCertificationUsecase) use case interface
func NewCfpCertificationUsecase(r repository.OuranosRepository) ICfpCertificationUsecase {
	return &cfpCertificationUsecase{r}
}

// GetCfpCertification
// Summary: This is function which get cfp certification.
// input: c(echo.Context) echo context
// input: getCfpCertificationModel(traceability.GetCfpCertificationModel) GetCfpCertificationModel object
// output: (traceability.CfpCertificationModels) CfpCertificationModels object
// output: (error) error object
func (u *cfpCertificationUsecase) GetCfpCertification(c echo.Context, getCfpCertificationModel traceability.GetCfpCertificationModel) (traceability.CfpCertificationModels, error) {
	cfpCertificationModels, _ := u.r.GetCFPCertifications(getCfpCertificationModel.OperatorID.String(), getCfpCertificationModel.TraceID.String())

	return cfpCertificationModels, nil
}
