package middleware

import (
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func Access(h http.Handler, w io.Writer) http.HandlerFunc {
	logger := zerolog.New(w).With().Timestamp().Logger()

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info().
			Str("method", r.Method).
			Str("host", r.RemoteAddr).
			Str("url", r.URL.String()).
			Msg("test")
		h.ServeHTTP(w, r)
	}
}
