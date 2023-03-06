package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `json:"id"                 bson:"_id"`
	FirstName         string `json:"first_name"         bson:"first_name"`
	LastName          string `json:"last_name"          bson:"last_name"`
	Email             string `json:"email"              bson:"email"`
	Password          string `json:"password,omitempty" bson:"-"`
	EncryptedPassword string `json:"-"                  bson:"encrypted_password"`
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
