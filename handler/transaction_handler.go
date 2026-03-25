package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/service"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

// HandleTransaction handles the POST /api/transaction endpoint
func (h *TransactionHandler) HandleTransaction(w http.ResponseWriter, r *http.Request) {
	req := service.TransactionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  model.TransactionStatusFailed,
			"message": "Invalid request body",
		})
		return
	}

	resp, _ := h.transactionService.ProcessTransaction(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
