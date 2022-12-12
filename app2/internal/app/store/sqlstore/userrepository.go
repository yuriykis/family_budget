package sqlstore

import "app/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user model.User) error {

	return r.store.db.QueryRow(
		"INSERT INTO users (first_name, last_name, email, encrypted_password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Email,
		user.FirstName,
		user.LastName,
		user.EncryptedPassword,
	).Scan(&user.ID)

}
