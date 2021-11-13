package store_test

import (
	"testing"

	"github.com/rthornton128/go-todo/pkg/store"
)

var testDSN string = "root@tcp(go-todo.railgun:3306)/test"

func TestNewDatabaseBadDSN(t *testing.T) {
	dsn := "*"
	_, err := store.NewDatabase(dsn)
	if err == nil {
		t.Error("Expected NewDatabase with bad data source name to return error")
	}
}

func TestNewDatabaseConnects(t *testing.T) {
	db, err := store.NewDatabase(testDSN)
	if err != nil {
		t.Errorf("Expected NewDatabase to not return an error, got: %s", err)
	}

	if err = db.Ping(); err != nil {
		t.Errorf("Expected Ping to be nil, got: %s", err)
	}
}

func TestNewDatabaseBadDatabase(t *testing.T) {
	dsn := "root@tcp(nothing:3306)/invalid"
	_, err := store.NewDatabase(dsn)
	if err == nil {
		t.Error("Expected NewDatabase with invalid host and database to return error")
	}
}

func TestNewDatabaseClose(t *testing.T) {
	db, _ := store.NewDatabase(testDSN)
	err := db.Close()
	if err != nil {
		t.Errorf("Expected Close to be nil, got: %s", err)
	}
}
