package datastore

import (
	"data-spaces-backend/domain/repository"

	"gorm.io/gorm"
)

type ouranosRepository struct {
	db *gorm.DB
}

func NewOuranosRepository(
	db *gorm.DB,
) repository.OuranosRepository {
	return &ouranosRepository{db}
}
