package store

type IStore interface {
	Category() ICategoryRepository
}
