package sqlstore

import (
	"categoryservice/internal/app/model"
	"fmt"
)

type CategoryRepository struct {
	store *Store
}

func (r *CategoryRepository) Create(category model.Category) (string, error) {
	err := r.store.db.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id",
		category.Name,
		category.Description,
	).Scan(&category.ID)
	return fmt.Sprint(category.ID), err
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	rows, err := r.store.db.Query("SELECT * FROM categories")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []model.Category{}
	for rows.Next() {
		category := model.Category{}
		var tmp string
		var tmp2 string
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&tmp,
			&tmp2,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Find(categoryID string) (*model.Category, error) {
	category := model.Category{}
	category.ID = fmt.Sprint(categoryID)
	err := r.store.db.QueryRow(
		"SELECT * FROM categories WHERE id = $1",
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)

	return &category, err
}

func (r *CategoryRepository) Edit(category model.Category) error {
	_, err := r.store.db.Exec(
		"UPDATE categories SET name = $1, description = $2 WHERE id = $3",
		category.Name,
		category.Description,
		category.ID,
	)
	return err
}

func (r *CategoryRepository) Delete(categoryID int) error {

	_, err := r.store.db.Exec(
		"DELETE FROM categories WHERE id = $1",
		categoryID,
	)
	return err
}
