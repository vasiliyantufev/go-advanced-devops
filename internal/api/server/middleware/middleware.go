// Package middleware - gzip module
package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer will be responsible for gzip compression
	return w.Writer.Write(b)
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Checks if the client supports gzip compression
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// If gzip is not supported, transfer control
			next.ServeHTTP(w, r)
			return
		}

		// Create a gzip.Writer on top of the current w
		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		// Passes a gzipWriter type variable to the page handler to output data
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
