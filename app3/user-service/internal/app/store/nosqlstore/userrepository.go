package nosqlstore

import (
	"context"
	"errors"
	"userservice/internal/app/model"
	"userservice/internal/app/utils/token"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user model.User) (string, error) {
	if err := user.BeforeCreate(); err != nil {
		return "", err
	}

	doc := bson.M{
		"_id":                primitive.NewObjectID(),
		"first_name":         user.FirstName,
		"last_name":          user.LastName,
		"email":              user.Email,
		"encrypted_password": user.EncryptedPassword,
	}

	res, err := r.store.db.Database("users").
		Collection("users").
		InsertOne(context.Background(), doc)
	if err != nil {
		return "", err
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("id is not an object id")
	}
	id := objID.Hex()
	return id, nil
}

func (r *UserRepository) Update(user model.User) error {
	objID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}
	updatedUser := bson.M{}
	if user.FirstName != "" {
		updatedUser["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		updatedUser["last_name"] = user.LastName
	}
	if user.Email != "" {
		updatedUser["email"] = user.Email
	}
	if user.Password != "" {
		if err := user.BeforeCreate(); err != nil {
			return err
		}
		updatedUser["encrypted_password"] = user.EncryptedPassword
	}

	res, err := r.store.db.Database("users").
		Collection("users").
		UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": updatedUser})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("no user found")
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res := r.store.db.Database("users").
		Collection("users").
		FindOneAndDelete(context.Background(), bson.M{"_id": objId})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (r *UserRepository) Find(id string) (*model.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user model.User
	err = r.store.db.Database("users").
		Collection("users").
		FindOne(context.Background(), bson.M{"_id": objId}).
		Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AuthCheck(email string, password string) (string, error) {
	var user model.User
	user.Email = email

	err := r.store.db.Database("users").
		Collection("users").
		FindOne(context.Background(), bson.M{"email": user.Email}).
		Decode(&user)
	if err != nil {
		return "", err
	}

	if user.ComparePassword(password) != nil {
		return "", errors.New("invalid password")
	}

	return token.GenerateToken(user.ID)
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	cursor, err := r.store.db.Database("users").
		Collection("users").
		Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
