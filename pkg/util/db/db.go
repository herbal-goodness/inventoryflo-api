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

func Query(q string, args ...interface{}) ([]map[string]interface{}, error) {
	if db == nil {
		if err := newDb(); err != nil {
			return nil, err
		}
	}
	rows, err := db.Queryx(q, args...)
	if err != nil {
		return nil, err
	}
	return rowsToMap(rows)
}

func rowsToMap(rows *sqlx.Rows) ([]map[string]interface{}, error) {
	cols, _ := rows.Columns()
	var result []map[string]interface{}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		r := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			r[colName] = *val
		}
		result = append(result, r)
	}
	return result, nil
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
