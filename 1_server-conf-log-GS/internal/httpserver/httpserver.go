package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/notIdealen/simple_server-config-logger/internal/logger"
)

type Server struct {
	// mux    *http.ServeMux
	*http.Server
	Logger *logger.Logger
}

func NewHTTPServer(router http.Handler, l *logger.Logger) *Server {
	cfg, err := NewServerConfig()
	if err != nil {
		return nil
	}

	return &Server{
		// mux: router,
		Server: &http.Server{
			Addr:    cfg.Addr,
			Handler: router,
		},
		Logger: l,
	}
}

func (s *Server) Run(ctx context.Context) {
	go s.ListenAndServe()

	<-ctx.Done()
	timectx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := s.Shutdown(timectx); err != nil {
		s.Logger.Error(fmt.Sprintf("Shutdown limit overflow, %v", err.Error()))
	}
	s.Logger.Info("Server stopped",
		slog.String("error", ctx.Err().Error()),
	)
	s.Logger.CloseFile()
}
