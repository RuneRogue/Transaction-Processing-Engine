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
	Balance int64  `json:"balance"`
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

// GetBalance handles the GET /api/card/balance/{cardNumber} endpoint
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

// GetTransactions handles the GET /api/card/transactions/{cardNumber} endpoint
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

type CreateCardResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// CreateCard handles the POST /api/card endpoint
func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var req service.CreateCardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	
	w.Header().Set("Content-Type", "application/json")
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateCardResponse{
			Status:  model.TransactionStatusFailed,
			Message: "Invalid request payload",
		})
		return
	}

	if req.CardNumber == "" || req.Pin == "" || req.CardHolder == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateCardResponse{
			Status:  model.TransactionStatusFailed,
			Message: "Missing required fields",
		})
		return
	}

	err = h.cardService.CreateCard(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateCardResponse{
			Status:  model.TransactionStatusFailed,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateCardResponse{
		Status: model.TransactionStatusSuccess,
	})
}
