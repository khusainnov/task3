package service

import (
	"github.com/khusainnov/task3/pkg/repository"
)

type Upload interface {
	UploadFile(lines [][]string) error
}

type Service struct {
	Upload
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Upload: NewUploadService(repos),
	}
}
