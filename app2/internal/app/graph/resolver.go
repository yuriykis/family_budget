package graph

import "app/internal/app/store"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store store.IStore
}

func NewResolver(store store.IStore) *Resolver {
	return &Resolver{store: store}
}
