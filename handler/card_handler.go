package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/service"
	"github.com/go-chi/chi/v5"
)

type CardHandler struct {
	cardService *service.CardService
}

type BalanceResponse struct {
	Status  string `json:"status"`
	Balance int64  `json:"balance,omitempty"`
	Message string `json:"message,omitempty"`
}

type TransactionHistoryResponse struct {
	Status             string              `json:"status"`
	Message            string              `json:"message,omitempty"`
	TransactionHistory []model.Transaction `json:"transactionHistory,omitempty"`
}

func NewCardHandler(cardService *service.CardService) *CardHandler {
	return &CardHandler{cardService: cardService}
}

func (h *CardHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	cardNumber := chi.URLParam(r, "cardNumber")

	balance, err := h.cardService.GetBalance(cardNumber)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BalanceResponse{
			Status:  model.TransactionStatusFailed,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(BalanceResponse{
		Status:  model.TransactionStatusSuccess,
		Balance: balance,
	})
}

func (h *CardHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	cardNumber := chi.URLParam(r, "cardNumber")
	transactions, err := h.cardService.GetTransactions(cardNumber)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(TransactionHistoryResponse{
			Status:  model.TransactionStatusFailed,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(TransactionHistoryResponse{
		Status:             model.TransactionStatusSuccess,
		TransactionHistory: transactions,
	})
}
