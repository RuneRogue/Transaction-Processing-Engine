package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionId uuid.UUID `json:"transactionId"`
	CardNumber    string    `json:"cardNumber"`
	Type          string    `json:"type"`
	Amount        int64     `json:"amount"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}
