package handler

import (
	"net/http"

	"github.com/notIdealen/simple_server-config-logger/internal/httpserver"
	"github.com/notIdealen/simple_server-config-logger/internal/middleware"
)

func UserHealth(w http.ResponseWriter, r *http.Request) {
	rw, ok := w.(*middleware.MyResponseWriter)
	if ok {
		rw.SetStatusCode(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Send data!!!"))
}

func GetRoutes() []httpserver.Route {
	return []httpserver.Route{
		{Method: "GET", Path: "/user/health", Handler: UserHealth},
	}
}
