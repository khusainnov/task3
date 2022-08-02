package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/task3"
	"github.com/khusainnov/task3/pkg/handler"
	"github.com/khusainnov/task3/pkg/repository"
	"github.com/khusainnov/task3/pkg/service"
	"github.com/spf13/viper"
)

//func csvFileReader(fl *os.File) {
//	dcsv := make([]CSVData, 0, 1000)
//
//	csvLines, err := csv.NewReader(fl).ReadAll()
//	if err != nil {
//		log.Fatalf("Error due to read file: %s", err.Error())
//	}
//
//	for _, line := range csvLines {
//		dcsv = append(dcsv, *&CSVData{
//			State:                 line[0],
//			ZipCode:               line[1],
//			TaxRegionName:         line[2],
//			StateRate:             line[3],
//			EstimatedCombinedRate: line[4],
//			EstimatedCountyRate:   line[5],
//			EstimatedCityRate:     line[6],
//			EstimatedSpecialRate:  line[7],
//			RiskLevel:             line[8],
//		})
//	}
//
//	for i, v := range dcsv {
//		fmt.Printf("%d: zip_code: %s, rate: %s\n", i, v.ZipCode, v.StateRate)
//	}
//}

func main() {

	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatalf("Cannot read .env config, due to error: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("Cannot read .yml config, due to error: %s", err.Error())
	}

	log.Println("Initializing DB")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	defer func() {
		if r := recover(); r != nil {
			db.Close()
			log.Println("DB closed, due to panic")
		}
	}()

	if err != nil {
		log.Fatalf("Cannot run db, due to error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(task3.Server)
	if err = s.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error due starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
