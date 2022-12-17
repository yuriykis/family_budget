package sqlstore

import (
	"app/internal/app/model"
	"app/internal/app/utils/token"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user model.User) (int, error) {

	if err := user.BeforeCreate(); err != nil {
		return 0, err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO users (first_name, last_name, email, encrypted_password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName,
		user.LastName,
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.ID)

	return user.ID, err
}

func (r *UserRepository) AuthCheck(email string, password string) (string, error) {

	var user model.User
	user.Email = email
	user.Password = password

	if err := user.BeforeCreate(); err != nil {
		return "", err
	}
	err := r.store.db.QueryRow(
		"SELECT id, encrypted_password FROM users WHERE email = $1",
		user.Email,
	).Scan(&user.ID, &user.EncryptedPassword)

	if err != nil {
		return "", err
	}

	err = user.ComparePassword(password)

	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(uint(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
