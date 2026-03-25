package storage

import (
	"errors"
	"sync"

	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/utils"
)

// MemoryStore is in memory storage for card and transactions
type MemoryStore struct {
	Cards        map[string]*model.Card
	Transactions map[string][]model.Transaction
	mu           sync.RWMutex
}

// InitializeMemoryStore initializes the MemoryStore struct and adds an example card.
func InitializeMemoryStore() *MemoryStore {
	store := &MemoryStore{
		Cards:        make(map[string]*model.Card),
		Transactions: make(map[string][]model.Transaction),
	}
	store.Cards["4123456789012345"] = &model.Card{
		CardNumber: "4123456789012345",
		CardHolder: "John Doe",
		PinHash:    utils.Hash("1234"), //converting pin to hash
		Balance:    1000,
		Status:     model.CardStatusActive,
	}
	return store
}

// GetCard returns the card from cardNumber.
// We are using RLock here because we are only reading the card and not modifying it.
func (s *MemoryStore) GetCard(cardNumber string) (*model.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	card, ok := s.Cards[cardNumber]
	if !ok {
		return nil, errors.New("card not found")
	}
	return card, nil
}

// UpdateCard updates the card in the store.
// We are using Lock here because we are modifying the card.
func (s *MemoryStore) UpdateCard(card *model.Card) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Cards[card.CardNumber] = card
}

// AddCard adds a new card to the store.
// Returns an error if the card already exists.
func (s *MemoryStore) AddCard(card *model.Card) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.Cards[card.CardNumber]; exists {
		return errors.New("card already exists")
	}
	s.Cards[card.CardNumber] = card
	return nil
}

// AddTransaction adds a transaction to the store.
// We are using Lock here because we are modifying the Transaction.
func (s *MemoryStore) AddTransaction(tx model.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Transactions[tx.CardNumber] = append(s.Transactions[tx.CardNumber], tx)
}

// GetTransactions returns the transactions for a card.
// We are using RLock here because we are only reading the transactions and not modifying them.
func (s *MemoryStore) GetTransactions(cardNumber string) []model.Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()
	txs, ok := s.Transactions[cardNumber]
	if !ok {
		return []model.Transaction{}
	}
	// We are returning a copy to prevent callers from accidently mutating internal state.
	copyTxs := make([]model.Transaction, len(txs))
	copy(copyTxs, txs)
	return copyTxs
}
