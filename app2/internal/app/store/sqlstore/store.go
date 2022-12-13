package sqlstore

import (
	"app/internal/app/store"
	"database/sql"
)

type Store struct {
	db             *sql.DB
	userRepository store.IUserRepository
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
