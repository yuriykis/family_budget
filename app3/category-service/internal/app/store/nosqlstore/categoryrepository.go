package nosqlstore

import (
	"categoryservice/internal/app/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepository struct {
	store *Store
}

func (r *CategoryRepository) Create(category model.Category) (string, error) {
	res, err := r.store.db.Database("categories").
		Collection("categories").
		InsertOne(context.Background(), category)
	if err != nil {
		return "", err
	}
	objId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("id is not an object id")
	}
	id := objId.Hex()
	return id, nil
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	res, err := r.store.db.Database("categories").
		Collection("categories").
		Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var categories []model.Category
	err = res.All(context.Background(), &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) Find(categoryID string) (*model.Category, error) {
	objId, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, err
	}
	var category model.Category
	err = r.store.db.Database("categories").
		Collection("categories").
		FindOne(context.Background(), bson.M{"_id": objId}).
		Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Edit(category model.Category) error {
	panic("implement me")
}

func (r *CategoryRepository) Delete(categoryID int) error {
	panic("implement me")
}
