package sqlstore

import (
	"categoryservice/internal/app/store"
	"database/sql"
)

type Store struct {
	db                 *sql.DB
	categoryRepository store.ICategoryRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
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
