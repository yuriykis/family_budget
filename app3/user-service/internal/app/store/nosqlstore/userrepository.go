package nosqlstore

import (
	"context"
	"errors"
	"userservice/internal/app/model"
	"userservice/internal/app/utils/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user model.User) (string, error) {
	if err := user.BeforeCreate(); err != nil {
		return "", err
	}

	doc := model.User{
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		EncryptedPassword: user.EncryptedPassword,
	}

	res, err := r.store.db.Database("users").
		Collection("users").
		InsertOne(context.Background(), doc)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *UserRepository) Update(user model.User) error {
	res, err := r.store.db.Database("users").
		Collection("users").
		UpdateOne(context.Background(), model.User{ID: user.ID}, user)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("no user found")
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) AuthCheck(email string, password string) (string, error) {
	var user model.User
	user.Email = email
	user.Password = password

	if err := user.BeforeCreate(); err != nil {
		return "", err
	}

	err := r.store.db.Database("users").
		Collection("users").
		FindOne(context.Background(), model.User{Email: user.Email}).
		Decode(&user)
	if err != nil {
		return "", err
	}

	if user.EncryptedPassword != user.Password {
		return "", errors.New("invalid password")
	}

	return token.GenerateToken(user.ID)
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	return nil, nil
}
