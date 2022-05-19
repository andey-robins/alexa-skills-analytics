package db

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

var instance *Database
var once sync.Once

func Connect() *Database {
	once.Do(func() {
		// create and connect to db
		database, err := sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		instance = &Database{Db: database}

		// create db if it doesn't exist
		_, err = os.Stat("./data.db")
		if errors.Is(err, os.ErrNotExist) {
			// load schema from file
			content, err := ioutil.ReadFile("./db/schema.sql")
			if err != nil {
				log.Fatal(err)
			}
			query := string(content)

			// create DB
			_, err = instance.Db.Exec(query)
			if err != nil {
				log.Printf("Failed to initialize database: %v", err)
			}
		} else if err != nil {
			log.Printf("[Err]: Error is checking for sqlite db - %v\n", err)
		} else {
			log.Println("[+]: Successfully connected to db")
		}

	})
	return instance
}
