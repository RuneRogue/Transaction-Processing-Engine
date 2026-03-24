package service

import (
	"errors"

	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/repository"
)

type CardService struct {
	cardRepo *repository.CardRepository
}

func NewCardService(cardRepo *repository.CardRepository) *CardService {
	return &CardService{cardRepo: cardRepo}
}

// GetBalance returns the balance of a card.
// It first validates the card and then returns the balance.
func (s *CardService) GetBalance(cardNumber string) (int64, error) {
	card, err := s.ValidateCard(cardNumber)
	if err != nil {
		return 0, err
	}
	return card.Balance, nil
}

// ValidateCard validates the card and returns the card if it is active.
func (s *CardService) ValidateCard(cardNumber string) (*model.Card, error) {
	card, err := s.cardRepo.GetCard(cardNumber)
	if err != nil {
		return nil, err
	}
	if card.Status != model.CardStatusActive {
		return nil, errors.New("card is not active")
	}
	return card, nil
}
