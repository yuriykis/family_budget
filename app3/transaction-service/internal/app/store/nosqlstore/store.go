package nosqlstore

import (
	"transactionservice/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db                    *mongo.Client
	TransactionRepository store.ITransactionRepository
}

func New(db *mongo.Client) *Store {
	return &Store{db: db}
}

func (s *Store) Transaction() store.ITransactionRepository {
	if s.TransactionRepository != nil {
		return s.TransactionRepository
	}

	s.TransactionRepository = &TransactionRepository{
		store: s,
	}

	return s.TransactionRepository
}
