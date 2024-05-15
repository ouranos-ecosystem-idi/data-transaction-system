package auth

import (
	"bytes"
	"encoding/json"

	"data-spaces-backend/domain/model/authentication"
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/extension/logger"
	"data-spaces-backend/infrastructure/auth/client"
)

// VerifyAPIKey
// Summary: This is function which verifies API key.
// input: request(repository.VerifyApiKeyBody) request
// output: (authentication.VeriryApiKeyResponse) response
// output: (error) error object
func (r *authAPIRepository) VerifyAPIKey(request repository.VerifyAPIKeyBody) (authentication.VeriryAPIKeyResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return authentication.VeriryAPIKeyResponse{}, err
	}
	body := bytes.NewBuffer(jsonData)

	resString, err := r.cli.Post(client.PathSystemAuthAPIKey, body)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return authentication.VeriryAPIKeyResponse{}, err
	}

	var response authentication.VeriryAPIKeyResponse
	if err := json.Unmarshal([]byte(resString), &response); err != nil {
		logger.Set(nil).Errorf(err.Error())

		return authentication.VeriryAPIKeyResponse{}, err
	}

	return response, nil
}

// VerifyToken
// Summary: This is function which verifies token.
// input: request(repository.VerifyTokenBody) request
// output: (authentication.VeriryTokenResponse) response
// output: (error) error object
func (r *authAPIRepository) VerifyToken(request repository.VerifyTokenBody) (authentication.VeriryTokenResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return authentication.VeriryTokenResponse{}, err
	}
	body := bytes.NewBuffer(jsonData)

	resString, err := r.cli.Post(client.PathSystemToken, body)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return authentication.VeriryTokenResponse{}, err
	}

	var response authentication.VeriryTokenResponse
	if err := json.Unmarshal([]byte(resString), &response); err != nil {
		logger.Set(nil).Errorf(err.Error())

		return authentication.VeriryTokenResponse{}, err
	}

	return response, nil
}
