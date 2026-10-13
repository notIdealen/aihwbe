package middleware_task

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/notIdealen/aihwbe.git/internal/middleware"
)

func ContentType() middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
				fmt.Println(errors.New("unexpected content type"))
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
