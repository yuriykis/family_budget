package store

import (
	"categoryservice/internal/app/model"
)

type ICategoryRepository interface {
	Create(model.Category) (int, error)
	FindAll() ([]model.Category, error)
	Find(int) (*model.Category, error)
	Edit(model.Category) error
	Delete(int) error
}
