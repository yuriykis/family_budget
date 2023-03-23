package apiserver

import (
	"net/http"
	"transactionservice/internal/app/store/nosqlstore"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func Start(config *Config) error {
// 	db, err := newDB(config.DatabaseURL)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	store := sqlstore.New(db)
// 	server := newServer(store)
// 	return http.ListenAndServe(config.BindAddr, server)
// }

// func newDB(databaseURL string) (*sql.DB, error) {
// 	db, err := sql.Open("postgres", databaseURL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Disconnect(nil)

	store := nosqlstore.New(db)
	server := newServer(store)
	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseURL string) (*mongo.Client, error) {
	db, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		return nil, err
	}

	if err := db.Connect(nil); err != nil {
		return nil, err
	}
	return db, nil
}
