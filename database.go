package main

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func (app *App) openDB() error {
	cfg, err := mysql.ParseDSN(app.env.DataSourceName)
	if err != nil {
		return err
	}
	app.logger.Info().Msg(cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	// Important, required settings on the DB
	// See: https://github.com/go-sql-driver/mysql#important-settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	app.db = db

	return app.db.Ping()
}

func (app App) closeDB() {
	app.db.Close()
}
