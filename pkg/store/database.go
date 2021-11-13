package store

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	pool *sql.DB
}

func NewDatabase(dsn string) (db *Database, err error) {
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	// Sanity check to confirm a connection can be made to the database.
	// There appears to be a bug in either the sql, or more likely mysql,
	// package that causes a race condition when Ping is called.
	err = pool.Ping()
	if err != nil {
		pool.Close()
		return
	}

	// Important! Required settings on the DB.
	// See: https://github.com/go-sql-driver/mysql#important-settings
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	db = &Database{pool: pool}

	return
}
