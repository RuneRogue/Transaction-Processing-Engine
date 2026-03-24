package model

type Card struct {
	CardNumber string `json:"cardNumber"`
	CardHolder string `json:"cardHolder"`
	PinHash    string `json:"pinHash"`
	Balance    int    `json:"balance"`
	Status     string `json:"status"`
}
