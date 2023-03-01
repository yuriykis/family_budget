package store

type IStore interface {
	Transaction() ITransactionRepository
}
