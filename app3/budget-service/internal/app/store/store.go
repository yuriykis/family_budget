package store

type IStore interface {
	Budget() IBudgetRepository
}
