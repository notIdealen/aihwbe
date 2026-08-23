package httpserver

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// func NewRoute(m string, p string, h http.HandlerFunc) Route {
// 	return Route{
// 		Method:  m,
// 		Path:    p,
// 		Handler: h,
// 	}
// }
