package store_test

import (
	"os"
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

func TestNewDatabase(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping test when running in CI")
	}

	_, err := store.NewDatabase(testDSN)
	if err != nil {
		t.Errorf("Expected NewDatabase to not return an error, got: %s", err)
	}
}

func TestNewDatabaseBadDatabase(t *testing.T) {
	dsn := "root@tcp(nothing:3306)/invalid"
	_, err := store.NewDatabase(dsn)
	if err == nil {
		t.Error("Expected NewDatabase with invalid host and database to return error")
	}
}
