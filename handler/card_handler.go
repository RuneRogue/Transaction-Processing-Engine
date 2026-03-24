package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/service"
	"github.com/go-chi/chi"
)

type CardHandler struct {
	cardService *service.CardService
}

type BalanceResponse struct {
	Status  string `json:"status"`
	Balance int64  `json:"balance,omitempty"`
	Message string `json:"message,omitempty"`
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
