package sqlstore

import "app/internal/app/model"

type CategoryRepository struct {
	store *Store
}

func (r *CategoryRepository) Create(category model.Category) (int, error) {
	err := r.store.db.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id",
		category.Name,
		category.Description,
	).Scan(&category.ID)

	return category.ID, err
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryRepository) Find(int) (*model.Category, error) {
	panic("not implemented") // TODO: Implement
}
