package auth

import (
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/infrastructure/auth/client"
)

// authAPIRepository
// Summary: This is structure which defines AuthAPIRepository.
type authAPIRepository struct {
	cli *client.Client
}

// NewAuthAPIRepository
// Summary: This is function which creates new AuthAPIRepository.
// input: cli(*client.Client) client
// output: (repository.AuthAPIRepository) AuthAPIRepository object
func NewAuthAPIRepository(cli *client.Client) repository.AuthAPIRepository {
	return &authAPIRepository{cli: cli}
}
