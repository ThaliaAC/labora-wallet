package models

import "time"

type Log struct {
	ID          string    `json:"id"`
	National_id string    `json:"national_id"`
	Country     string    `json:"country"`
	Status      string    `json:"status_id"`
	RequestDate time.Time `json:"request_date"`
	RequestType string    `json:"request_type"`
}

type Wallet struct {
	ID          string    `json:"id"`
	National_id string    `json:"national_id"`
	Country     string    `json:"country"`
	RequestDate time.Time `json:"request_date"`
	Balance     int       `json:"balance"`
}

type Transaction struct {
	ID               string    `json:"id"`
	Wallet_id        string    `json:"wallet_id"`
	Transaction_type string    `json:"transaction_type"`
	Amount           int       `json:"amount"`
	Transaction_date time.Time `json:"transaction_date"`
}
