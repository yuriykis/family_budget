package sqlstore

import (
	"budgetservice/internal/app/store"
	"database/sql"
)

type Store struct {
	db               *sql.DB
	budgetRepository store.IBudgetRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Budget() store.IBudgetRepository {
	if s.budgetRepository != nil {
		return s.budgetRepository
	}

	s.budgetRepository = &BudgetRepository{
		store: s,
	}

	return s.budgetRepository
}
