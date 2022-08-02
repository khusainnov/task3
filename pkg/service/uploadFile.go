package service

import (
	"log"

	"github.com/khusainnov/task3/entity"
	"github.com/khusainnov/task3/pkg/repository"
)

type UploadService struct {
	repo repository.Upload
}

func NewUploadService(repo repository.Upload) *UploadService {
	return &UploadService{repo: repo}
}

func (u *UploadService) UploadFile(csvLines [][]string) error {
	dcsv := make([]entity.CSVData, 0, 1000)

	for _, line := range csvLines {
		dcsv = append(dcsv, *&entity.CSVData{
			State:                 line[0],
			ZipCode:               line[1],
			TaxRegionName:         line[2],
			StateRate:             line[3],
			EstimatedCombinedRate: line[4],
			EstimatedCountyRate:   line[5],
			EstimatedCityRate:     line[6],
			EstimatedSpecialRate:  line[7],
			RiskLevel:             line[8],
		})
	}

	_, err := u.repo.UploadFile(dcsv)

	log.Println("some error " + err.Error())

	return err
}
