package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/task3/entity"
)

var wg sync.WaitGroup
var mu sync.Mutex

type UploadPostgres struct {
	db *sqlx.DB
}

//var Queue chan entity.CSVData
//func init() {
//	Queue = make(chan entity.CSVData, 999)
//}

func NewUploadPostgres(db *sqlx.DB) *UploadPostgres {
	return &UploadPostgres{db: db}
}

func (u *UploadPostgres) UploadFile(dcsv []entity.CSVData) (string, error) {
	/*queryInsert := fmt.Sprintf("INSERT INTO %s VALUES($1, $2);", table)
	queryUpdate := fmt.Sprintf("UPDATE %s SET rate=$1 WHERE zip_code=$2;", table)*/
	/*	go func() {
		for i := 1; i < len(dcsv); i++ {
			_, err := u.db.Query(queryInsert, dcsv[i].ZipCode, dcsv[i].StateRate)
			if err != nil {
				_, err = u.db.Query(queryUpdate, dcsv[i].StateRate, dcsv[i].ZipCode)
			}
		}
	}()*/
	start := time.Now()
	Queue := make(chan entity.CSVData, 999)
	fmt.Println(len(dcsv))
	/*	for i := 1; i < len(dcsv); i++ {
		Queue <- dcsv[i]
		u.startProcess(queryInsert, queryUpdate)
	}*/

	for i := 1; i < len(dcsv); i++ {
		go u.worker(Queue)
	}

	for i := 1; i < len(dcsv); i++ {
		Queue <- dcsv[i]
	}
	close(Queue)

	if r := recover(); r != nil {
		u.db.Close()
		return "Recovered", errors.New("recovered. DB Closed")
	}

	wg.Wait()

	fmt.Printf("after wg.Wait: %v\n", time.Now().Sub(start).Milliseconds())

	if r := recover(); r != nil {
		u.db.Close()
		return "Recovered", errors.New("recovered. DB Closed")
	}

	//for i := 1; i < len(dcsv); i++ {
	//	_, err := u.db.Query(queryInsert, dcsv[i].ZipCode, dcsv[i].StateRate)
	//	if err != nil {
	//		_, err = u.db.Query(queryUpdate, dcsv[i].StateRate, dcsv[i].ZipCode)
	//		if err != nil {
	//			return "", err
	//		}
	//	}
	//}

	//for i := 1; i < len(dcsv); i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		_, err := u.db.Query(queryInsert, dcsv[i].ZipCode, dcsv[i].StateRate)
	//		if err != nil {
	//			_, _ = u.db.Query(queryUpdate, dcsv[i].StateRate, dcsv[i].ZipCode)
	//			//if err != nil {
	//			//	log.Fatalf("U_P error: %s", err.Error())
	//			//}
	//		}
	//		wg.Done()
	//	}(i)
	//}

	//wg.Wait()

	fmt.Printf("in the end: %v\n", time.Now().Sub(start).Milliseconds())
	return "uploaded", nil
}

var counter = 0

func (u *UploadPostgres) worker(queue <-chan entity.CSVData) {
	queryInsert := fmt.Sprintf("INSERT INTO %s VALUES($1, $2);", table)
	queryUpdate := fmt.Sprintf("UPDATE %s SET rate=$1 WHERE zip_code=$2;", table)
	for job := range queue {
		wg.Add(1)
		mu.Lock()
		param, err := u.db.Query(queryInsert, job.ZipCode, job.StateRate)
		if err != nil {
			counter++
			//log.Printf("%d, %s", counter, err.Error())
			_, err = u.db.Query(queryUpdate, job.StateRate, job.ZipCode)
			if err != nil {
				log.Printf("%s", err.Error())
			}
		}
		mu.Unlock()
		param.Close()
	}

	wg.Done()
}

func (u *UploadPostgres) startProcess(queryInsert, queryUpdate string) {
	Queue := make(chan entity.CSVData, 999)

	select {
	case job := <-Queue:
		wg.Add(1)
		mu.Lock()
		_, err := u.db.Query(queryInsert, job.ZipCode, job.StateRate)
		if err != nil {
			_, _ = u.db.Query(queryUpdate, job.StateRate, job.ZipCode)
		}
		mu.Unlock()
		wg.Done()
	}

}
