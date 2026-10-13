package server_mock

import (
	"fmt"
	"net/http"
)

func ServerForTestHandler(path string, h http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(path, h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
