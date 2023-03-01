package sqlstore

import (
	"database/sql"
	"transactionservice/internal/app/store"
)

type Store struct {
	db                    *sql.DB
	transactionRepository store.ITransactionRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Transaction() store.ITransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store: s,
	}

	return s.transactionRepository
}
