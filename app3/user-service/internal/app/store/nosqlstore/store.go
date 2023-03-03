package nosqlstore

import (
	"userservice/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db             *mongo.Client
	UserRepository store.IUserRepository
}

func New(db *mongo.Client) *Store {
	return &Store{db: db}
}

func (s *Store) User() store.IUserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
