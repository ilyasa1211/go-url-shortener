package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "data/data.db")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sqlite connected")

	return db
}
