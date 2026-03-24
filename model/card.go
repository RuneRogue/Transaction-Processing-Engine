package model

type Card struct {
	CardNumber string `json:"cardNumber"`
	CardHolder string `json:"cardHolder"`
	PinHash    string `json:"pinHash"`
	Balance    int64  `json:"balance"`
	Status     string `json:"status"`
}
