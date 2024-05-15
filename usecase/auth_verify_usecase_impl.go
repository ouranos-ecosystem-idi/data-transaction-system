package usecase

import (
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/usecase/input"
	"data-spaces-backend/usecase/output"
)

// verifyUsecase
// Summary: This is structure which defines verifyUsecase.
type verifyUsecase struct {
	authAPIRepository repository.AuthAPIRepository
}

// NewVerifyUsecase
// Summary: This is function which creates new VerifyUsecase.
// input: r(repository.AuthAPIRepository) AuthAPIRepository
// output: (IVerifyUsecase) VerifyUsecase object
func NewVerifyUsecase(r repository.AuthAPIRepository) IVerifyUsecase {
	return &verifyUsecase{r}
}

// VerifyAPIKey
// Summary: This is function which verifies the API key.
// input: input(input.VerifyAPIKey) VerifyAPIKey
// output: (output.VerifyAPIKey) VerifyAPIKey
// output: (error) error object
func (u verifyUsecase) VerifyAPIKey(input input.VerifyAPIKey) (output.VerifyAPIKey, error) {
	param := repository.VerifyAPIKeyBody{
		APIKey:    input.APIKey,
		IPAddress: input.IPAddress,
	}
	res, err := u.authAPIRepository.VerifyAPIKey(param)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return output.VerifyAPIKey{}, err
	}
	output := output.VerifyAPIKey{
		IsAPIKeyValid:    res.IsAPIKeyValid,
		IsIPAddressValid: res.IsIPAddressValid,
	}
	return output, nil
}

// VerifyToken
// Summary: This is function which verifies the token.
// input: input(input.VerifyToken) VerifyToken
// output: (output.VerifyToken) VerifyToken
// output: (error) error object
func (u verifyUsecase) VerifyToken(input input.VerifyToken) (output.VerifyToken, error) {
	req := repository.VerifyTokenBody{
		Token: input.Token,
	}
	res, err := u.authAPIRepository.VerifyToken(req)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return output.VerifyToken{}, err
	}
	output := output.VerifyToken{
		OperatorID: res.OperatorID,
	}

	return output, nil
}
