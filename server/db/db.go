package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Db *pgxpool.Pool
}

var instance *Database
var once sync.Once

func Connect() *Database {
	once.Do(func() {
		database, err := pgxpool.Connect(context.Background(), os.Getenv("DB_CONN"))
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		instance = &Database{Db: database}
	})

	return instance
}
