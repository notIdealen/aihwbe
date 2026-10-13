package domain

import "net/http"

type MyResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Log        string
}
