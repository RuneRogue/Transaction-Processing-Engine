package repository

import (
	"github.com/RuneRogue/Transaction-Processing-Engine/model"
	"github.com/RuneRogue/Transaction-Processing-Engine/storage"
)

// TransactionRepository abstracts data access and makes the system independent of storage implementation.
type TransactionRepository struct {
	store *storage.MemoryStore
}

func NewTransactionRepository(store *storage.MemoryStore) *TransactionRepository {
	return &TransactionRepository{store: store}
}

func (r *TransactionRepository) GetTransactions(cardNumber string) []model.Transaction {
	return r.store.GetTransactions(cardNumber)
}

func (r *TransactionRepository) AddTransaction(tx model.Transaction) {
	r.store.AddTransaction(tx)
}
