package nosqlstore

import (
	"budgetservice/internal/app/model"
	"budgetservice/internal/app/response"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BudgetRepository struct {
	store *Store
}

func (r *BudgetRepository) Create(budget model.Budget, userID string) (string, error) {
	doc := bson.M{
		"_id":         primitive.NewObjectID(),
		"budget_name": budget.Name,
		"description": budget.Description,
		"amount":      budget.Amount,
	}
	res1, err := r.store.db.Database("budgets").
		Collection("budgets").
		InsertOne(context.Background(), doc)
	if err != nil {
		return "", err
	}
	objID, ok := res1.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", errors.New("id is not an object id")
	}
	res2, err := r.store.db.Database("budgets").
		Collection("user_budgets").
		InsertOne(context.Background(), bson.M{
			"budget_id": res1.InsertedID,
			"user_id":   userID,
			"ownership": true,
			"readonly":  false,
		})
	if err != nil {
		return "", err
	}
	_, ok = res2.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", errors.New("id is not an object id")
	}
	id := objID.Hex()
	return id, nil
}

func (r *BudgetRepository) FindAll() ([]response.BudgetResponse, error) {
	res1, err := r.store.db.Database("budgets").
		Collection("budgets").
		Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var budgets []response.BudgetResponse
	if err := res1.All(context.Background(), &budgets); err != nil {
		return nil, err
	}
	res2, err := r.store.db.Database("budgets").
		Collection("user_budgets").
		Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := res2.All(context.Background(), &budgets); err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *BudgetRepository) Find(id string) (*response.BudgetResponse, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res1, err := r.store.db.Database("budgets").
		Collection("budgets").
		Find(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}
	var budget []response.BudgetResponse
	if err := res1.All(context.Background(), &budget); err != nil {
		return nil, err
	}
	res2, err := r.store.db.Database("budgets").
		Collection("user_budgets").
		Find(context.Background(), bson.M{"budget_id": objId})
	if err != nil {
		return nil, err
	}
	if err := res2.All(context.Background(), &budget); err != nil {
		return nil, err
	}
	return &budget[0], nil
}

func (r *BudgetRepository) Edit(budget model.Budget) error {
	res, err := r.store.db.Database("budgets").
		Collection("budgets").
		Find(context.Background(), bson.M{"_id": budget.ID})
	if err != nil {
		return err
	}
	if err := res.All(context.Background(), &budget); err != nil {
		return err
	}
	return nil
}

func (r *BudgetRepository) Delete(id int) error {
	_, err := r.store.db.Database("budgets").
		Collection("budgets").
		DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *BudgetRepository) Share(budgetID int, userID int) error {
	panic("not implemented")
}
