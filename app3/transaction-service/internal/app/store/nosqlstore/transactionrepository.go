package nosqlstore

import (
	"context"
	"transactionservice/internal/app/model"

	"go.mongodb.org/mongo-driver/bson"
)

type TransactionRepository struct {
	store *Store
}

func (r *TransactionRepository) Create(transaction model.Transaction, budgetID int) (int, error) {
	res, err := r.store.db.Database("transactions").
		Collection("transactions").
		InsertOne(context.Background(), bson.M{
			"budget_id":   budgetID,
			"amount":      transaction.Amount,
			"category_id": transaction.CategoryID,
		})

	if err != nil {
		return 0, err
	}

	objID, ok := res.InsertedID.(int)
	if !ok {
		return 0, err
	}

	return objID, nil
}

func (r *TransactionRepository) FindAll(budgetID int) ([]model.Transaction, error) {
	res, err := r.store.db.Database("transactions").
		Collection("transactions").
		Find(context.Background(), bson.M{"budget_id": budgetID})

	if err != nil {
		return nil, err
	}
	defer res.Close(context.Background())

	transactions := []model.Transaction{}
	for res.Next(context.Background()) {
		transaction := model.Transaction{}
		err := res.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) FindByCategory(categoryID int) ([]model.Transaction, error) {
	panic("not implemented")
}

func (r *TransactionRepository) Delete(id int) error {
	panic("not implemented")
}

func (r *TransactionRepository) DeleteByBudget(budgetID int) error {
	panic("not implemented")
}

func (r *TransactionRepository) Find(id int) (*model.Transaction, error) {
	panic("not implemented")
}
