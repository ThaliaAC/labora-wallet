package services

import (
	"fmt"

	"github.com/ThaliaAC/labora-wallet/models"
)

func WalletStatusFromLog(id int, log models.Log) (string, error) {
	var err error
	if log.Status == "APPROVED" {
		return "APPROVED", err
	}
	return "REJECTED", err
}

func CreateLog(log models.Log, db PostgresDBHandler) error {
	_, err := db.Exec("INSERT INTO logsTable (national_id, status_id, country, request_date, request_type) VALUES ($1, $2, $3, $4, $5)", log.National_id, log.Status, log.Country, log.RequestDate, log.RequestType)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
