package sqlstore

import (
	"app/internal/app/store"
	"database/sql"
)

type Store struct {
	db                    *sql.DB
	userRepository        store.IUserRepository
	budgetRepository      store.IBudgetRepository
	categoryRepository    store.ICategoryRepository
	transactionRepository store.ITransactionRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) User() store.IUserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
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

func (s *Store) Category() store.ICategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &CategoryRepository{
		store: s,
	}

	return s.categoryRepository
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
