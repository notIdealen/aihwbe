package httpserver

import (
	"fmt"
	"net/http"
)

type Router struct {
	*http.ServeMux
	// version int
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (r *Router) RegistRoutes(routes []Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.ServeMux.HandleFunc(pattern, route.Handler)
	}
}
