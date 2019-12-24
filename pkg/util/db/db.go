package db

import (
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

func Query(q string) (*sqlx.Rows, error) {
	if db == nil {
		if err := newDb(); err != nil {
			return nil, err
		}
	}
	return db.Queryx(q)
}

func Select(dest interface{}, q string, args ...interface{}) error {
	if db == nil {
		if err := newDb(); err != nil {
			return err
		}
	}
	return db.Select(dest, q, args)
}

func Get(dest interface{}, q string, args ...interface{}) error {
	if db == nil {
		if err := newDb(); err != nil {
			return err
		}
	}
	return db.Get(dest, q, args)
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
