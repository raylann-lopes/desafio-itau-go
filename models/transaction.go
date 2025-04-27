package models

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	TransactionType string    `json:"transactionType"`
	Value           float64   `json:"value" binding:"required"`
	Date            time.Time `json:"date"`
}
