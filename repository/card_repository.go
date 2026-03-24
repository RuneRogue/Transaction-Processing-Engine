package repository

import (
	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/storage"
)

// CardRepository abstracts data access and makes the system independent of storage implementation.
type CardRepository struct {
	store *storage.MemoryStore
}

func NewCardRepository(store *storage.MemoryStore) *CardRepository {
	return &CardRepository{store: store}
}

func (r *CardRepository) GetCard(cardNumber string) (*model.Card, error) {
	return r.store.GetCard(cardNumber)
}

func (r *CardRepository) UpdateCard(card *model.Card) {
	r.store.UpdateCard(card)
}
