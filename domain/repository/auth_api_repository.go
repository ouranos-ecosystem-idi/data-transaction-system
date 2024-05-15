package repository

import "data-spaces-backend/domain/model/authentication"

//go:generate mockery --name AuthAPIRepository --output ../../test/mock --case underscore
type (
	AuthAPIRepository interface {
		VerifyAPIKey(request VerifyAPIKeyBody) (authentication.VeriryAPIKeyResponse, error)
		VerifyToken(request VerifyTokenBody) (authentication.VeriryTokenResponse, error)
	}
)

type VerifyAPIKeyBody struct {
	APIKey    string `json:"apiKey"`
	IPAddress string `json:"ipAddress"`
}

type VerifyTokenBody struct {
	Token string `json:"idToken"`
}
