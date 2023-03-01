package sqlstore

import (
	"userservice/internal/app/model"
	"userservice/internal/app/utils/token"
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

func (r *UserRepository) Update(user model.User) error {
	_, err := r.store.db.Exec(
		"UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4",
		user.FirstName,
		user.LastName,
		user.Email,
		user.ID,
	)

	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.store.db.Exec(
		"DELETE FROM users WHERE id = $1",
		id,
	)

	return err
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

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, first_name, last_name, email FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	rows, err := r.store.db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
