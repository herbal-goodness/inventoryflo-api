package db

import (
	"database/sql"
	"fmt"

	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB = nil

// CloseDb closes the db connection cleanly
func CloseDb() {
	db.Close()
}

func query(q string) (*sql.Rows, error) {
	var err error
	if db == nil {
		if err = newDb(); err != nil {
			return nil, err
		}
	}
	return db.Query(q)
}

func newDb() error {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get("dbHost"), config.Get("dbPort"), config.Get("dbUser"),
		config.Get("dbPass"), config.Get("dbName"))

	var err error
	db, err = sqlx.Connect("postgres", connection)
	if err != nil {
		return err
	}
	return db.Ping()
}
