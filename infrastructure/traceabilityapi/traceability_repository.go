package traceabilityapi

import (
	"data-spaces-backend/domain/repository"
	"data-spaces-backend/infrastructure/traceabilityapi/client"
)

// traceabilityRepository
// Summary: This is structure which defines traceabilityRepository.
type traceabilityRepository struct {
	cli *client.Client
}

// NewTraceabilityRepository
// Summary: This is function which creates new TraceabilityRepository.
// input: cli(*client.Client) client
// output: (repository.TraceabilityRepository) TraceabilityRepository object
func NewTraceabilityRepository(cli *client.Client) repository.TraceabilityRepository {
	return &traceabilityRepository{cli: cli}
}
