package store

type IStore interface {
	User() IUserRepository
}
