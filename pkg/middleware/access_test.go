package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/rthornton128/go-todo/pkg/middleware"
)

func countHandler(count *int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*count++
	}
}

func TestLogAccess(t *testing.T) {
	buffer := &bytes.Buffer{}
	count := 0
	h := middleware.Access(countHandler(&count), buffer)
	r := httptest.NewRequest(http.MethodGet, "http://test.url", nil)
	w := httptest.NewRecorder()

	h(w, r)

	t.Log(buffer.String())
	matched, err := regexp.Match("test", buffer.Bytes())
	if err != nil {
		t.Errorf("error reading body from writer: %s", err)
	}
	if !matched {
		t.Errorf("expected test to be present in log output")
	}
}
