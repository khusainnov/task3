package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/task3/entity"
)

type Upload interface {
	UploadFile(dscv []entity.CSVData) (string, error)
}

type Repository struct {
	Upload
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Upload: NewUploadPostgres(db),
	}
}
