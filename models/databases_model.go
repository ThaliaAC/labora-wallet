package models

import "time"

type Log struct {
	ID          string    `json:"id"`
	NationalID  string    `json:"national_id"`
	Status      string    `json:"status_id"`
	Country     string    `json:"country"`
	RequestDate time.Time `json:"request_date"`
	RequestType string    `json:"request_type"`
}

type Wallet struct {
	ID          string    `json:"id"`
	NationalID  string    `json:"national_id"`
	Country     string    `json:"country"`
	RequestDate time.Time `json:"request_date"`
	Balance     int       `json:"balance"`
}
