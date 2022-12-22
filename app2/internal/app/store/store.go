package store

type IStore interface {
	User() IUserRepository
	Budget() IBudgetRepository
	Category() ICategoryRepository
	Transaction() ITransactionRepository
}
