package store

import (
	"categoryservice/internal/app/model"
)

type ICategoryRepository interface {
	Create(model.Category) (string, error)
	FindAll() ([]model.Category, error)
	Find(string) (*model.Category, error)
	Edit(model.Category) error
	Delete(int) error
}
