package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	pool *sql.DB
}

func NewDB(connStr string) *Database {
	connStr = "user=v dbname=sw-store sslmode=disable"
	pool, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = pool.Ping()
	if err != nil {
		panic(err)
	}

	return &Database{pool}
}

func (db *Database) Close() error {
	return db.pool.Close()
}
func (db *Database) Query(query string, params ...interface{}) (*sql.Rows, error) {
	if len(params) == 0 {
		return db.pool.Query(query)

	}
	return db.pool.Query(query, params)
}

var DB = NewDB("")
