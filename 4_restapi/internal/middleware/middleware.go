package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func WrapRouter(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; 0 <= i; i-- {
		h = mws[i](h)
	}
	return h
}
