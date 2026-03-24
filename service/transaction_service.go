package service

import (
	"time"

	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/repository"
	"github.com/RuneRogue/Transaction-Processing-Engine/utils"
	"github.com/google/uuid"
)

type TransactionService struct {
	cardRepo        *repository.CardRepository
	transactionRepo *repository.TransactionRepository
}

type TransactionRequest struct {
	CardNumber string
	Pin        string
	Type       string
	Amount     int64
}

type TransactionResponse struct {
	Status   string
	RespCode string
	Balance  int64
	Message  string
}

func NewTransactionService(cardRepo *repository.CardRepository, transactionRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{cardRepo: cardRepo, transactionRepo: transactionRepo}
}

func (s *TransactionService) ProcessTransaction(req TransactionRequest) (TransactionResponse, error) {
	//Step 1:- Get Card
	card, err := s.cardRepo.GetCard(req.CardNumber)
	if err != nil {
		s.LogTransaction(req.CardNumber, req.Type, model.TransactionStatusFailed, req.Amount)
		return TransactionResponse{
			Status:   model.TransactionStatusFailed,
			RespCode: "05",
			Message:  "Invalid card",
		}, nil
	}

	//Check if card is active or blocked
	if card.Status != model.CardStatusActive {
		s.LogTransaction(req.CardNumber, req.Type, model.TransactionStatusFailed, req.Amount)
		return TransactionResponse{
			Status:   model.TransactionStatusFailed,
			RespCode: "05",
			Message:  "Card is blocked",
		}, nil
	}

	//Step 2:- Validate PIN
	pin := utils.Hash(req.Pin)
	if pin != card.PinHash {
		s.LogTransaction(req.CardNumber, req.Type, model.TransactionStatusFailed, req.Amount)
		return TransactionResponse{
			Status:   model.TransactionStatusFailed,
			RespCode: "06",
			Message:  "Invalid PIN",
		}, nil
	}

	//Step 3:- Validate amount
	if req.Amount <= 0 {
		s.LogTransaction(req.CardNumber, req.Type, model.TransactionStatusFailed, req.Amount)
		return TransactionResponse{
			Status:   model.TransactionStatusFailed,
			RespCode: "07",
			Message:  "Invalid amount",
		}, nil
	}

	//Step 4:- Process Transaction
	switch req.Type {
	case model.TransactionTypeWithdraw:
		if card.Balance < req.Amount {
			s.LogTransaction(req.CardNumber, req.Type, model.TransactionStatusFailed, req.Amount)
			return TransactionResponse{
				Status:   model.TransactionStatusFailed,
				RespCode: "99",
				Message:  "Insufficient balance",
			}, nil
		}
		card.Balance -= req.Amount

	case model.TransactionTypeTopUp:
		card.Balance += req.Amount

	default:
		return TransactionResponse{
			Status:  model.TransactionStatusFailed,
			Message: "Invalid transaction type",
		}, nil
	}

	//Step 5:- Add Transaction
	s.LogTransaction(req.CardNumber, req.Type, model.TransactionStatusSuccess, req.Amount)

	return TransactionResponse{
		Status:   model.TransactionStatusSuccess,
		RespCode: "00",
		Balance:  card.Balance,
	}, nil
}

// LogTransaction function will log all the transaction to transaction history.
func (s *TransactionService) LogTransaction(cardNumber, txType, status string, amount int64) {
	tx := model.Transaction{
		TransactionId: uuid.New(),
		CardNumber:    cardNumber,
		Type:          txType,
		Amount:        amount,
		Status:        status,
		Timestamp:     time.Now(),
	}
	s.transactionRepo.AddTransaction(tx)
}
