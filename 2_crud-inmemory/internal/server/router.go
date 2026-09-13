package server

import (
	"fmt"
	"net/http"
)

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (r *Router) SetHandlers(routes []Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.HandleFunc(pattern, route.Handler)
	}
}
