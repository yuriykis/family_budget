package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                string `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) BeforeCreate() error {
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
