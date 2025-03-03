package middlewares

import (
	"context"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type contextKey string

const dbKey contextKey = "db"

// Middleware to inject DB into request context
func WithDB(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeoutDuration := 3 * time.Second // Default timeout

			// Allow dynamic timeout via header (e.g., for long-running reports)
			if r.Header.Get("X-Request-Timeout") != "" {
				if customTimeout, err := time.ParseDuration(r.Header.Get("X-Request-Timeout")); err == nil {
					timeoutDuration = customTimeout
				}
			}

			ctx, cancel := context.WithTimeout(r.Context(), timeoutDuration)
			defer cancel()

			ctx = context.WithValue(ctx, dbKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Retrieve DB from context
func GetDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return db
}
