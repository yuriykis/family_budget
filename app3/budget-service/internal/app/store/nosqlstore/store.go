package nosqlstore

import (
	"budgetservice/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db               *mongo.Client
	BudgetRepository store.IBudgetRepository
}

func New(db *mongo.Client) *Store {
	return &Store{db: db}
}

func (s *Store) Budget() store.IBudgetRepository {
	if s.BudgetRepository != nil {
		return s.BudgetRepository
	}

	s.BudgetRepository = &BudgetRepository{
		store: s,
	}

	return s.BudgetRepository
}
