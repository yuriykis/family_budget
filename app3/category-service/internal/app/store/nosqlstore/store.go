package nosqlstore

import (
	"categoryservice/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db                 *mongo.Client
	categoryRepository store.ICategoryRepository
}

func New(db *mongo.Client) *Store {
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
