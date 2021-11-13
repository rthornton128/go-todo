package env_test

import (
	"os"
	"testing"

	"github.com/rthornton128/go-todo/pkg/env"
)

func TestNewEnvironment(t *testing.T) {
	e := env.NewEnvironment()
	if e == nil {
		t.Fail()
	}
}

func TestEnvironmentValue(t *testing.T) {
	tests := []struct {
		key, value string
	}{
		{"DSN", "user:pass@tcp(host:port)/database"},
		{"HOST", "host"},
		{"MODE", "production"},
		{"PORT", "80"},
	}

	e := env.NewEnvironment()
	for i, test := range tests {
		os.Setenv(test.key, test.value)
		result := e.Get(test.key)

		if test.value != result {
			t.Errorf("%d: for %s expected %s got %s", i, test.key, test.value, result)
		}
	}
}

func TestEnvironmentIsDevelopment(t *testing.T) {
	e := env.NewEnvironment()

	os.Setenv("MODE", "development")
	if !e.IsDevelopment() {
		t.Error("Expected IsDevelopment to be true")
	}

	os.Setenv("MODE", "other")
	if e.IsDevelopment() {
		t.Error("Expected IsDevelopment to be false")
	}
}

func TestEnvironmentIsProduction(t *testing.T) {
	e := env.NewEnvironment()

	os.Setenv("MODE", "production")
	if !e.IsProduction() {
		t.Error("Expected IsProduction to be true")
	}

	os.Setenv("MODE", "other")
	if e.IsProduction() {
		t.Error("Expected IsProduction to be false")
	}
}

func TestEnvironmentIsTest(t *testing.T) {
	e := env.NewEnvironment()

	os.Setenv("MODE", "test")
	if !e.IsTest() {
		t.Error("Expected IsTest to be true")
	}

	os.Setenv("MODE", "other")
	if e.IsTest() {
		t.Error("Expected IsTest to be false")
	}
}
