package store

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Database struct {
	pool *sql.DB
}

func NewDatabase(dsn string) (db *Database, err error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return
	}

	pool, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return
	}

	// Important! Required settings on the DB.
	// See: https://github.com/go-sql-driver/mysql#important-settings
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	db = &Database{pool: pool}
	err = db.Ping()

	return
}

func (db Database) Ping() error {
	return db.pool.Ping()
}

func (db Database) Close() error {
	return db.pool.Close()
}
