package usecase

import (
	"data-spaces-backend/usecase/input"
	"data-spaces-backend/usecase/output"
)

// IVerifyUsecase
// Summary: This is interface which defines VerifyUsecase.
//
//go:generate mockery --name IVerifyUsecase --output ../test/mock --case underscore
type IVerifyUsecase interface {
	VerifyAPIKey(input input.VerifyAPIKey) (output.VerifyAPIKey, error)
	VerifyToken(input input.VerifyToken) (output.VerifyToken, error)
}
