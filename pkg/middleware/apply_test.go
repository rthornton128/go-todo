package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rthornton128/go-todo/pkg/middleware"
)

func testMiddleware(count *int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*count++
	}
}

func nullMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func TestApply(t *testing.T) {
	middleware.Apply(nullMiddleware())
}

func TestApplyHandler(t *testing.T) {
	count := 0
	h := middleware.Apply(testMiddleware(&count), testMiddleware(&count))
	r := httptest.NewRequest(http.MethodGet, "http://fake.url", nil)
	w := httptest.NewRecorder()
	h(w, r)

	expect := 2
	if count != expect {
		t.Errorf("expected handlers to be called %d times, only called %d", expect, count)
	}

}
