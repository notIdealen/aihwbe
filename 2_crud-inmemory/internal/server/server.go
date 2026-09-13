package server

import (
	"context"
	"net/http"
	"time"

	"github.com/notIdealen/aihwbe.git/internal/logger"
)

type CServer struct {
	http.Server
}

func NewCServer(config *Config) *CServer {
	s := &CServer{}
	if config == nil {
		s.Addr = "8080"
	}
	s.Addr = config.Addr
	return s
}

func (s *CServer) AttachHandler(h http.Handler) {
	s.Handler = h
}

func (s *CServer) Run(stopCtx context.Context, logger *logger.Logger) {
	logger.Info("Server run...")

	go s.ListenAndServe()

	<-stopCtx.Done()

	waitCtx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()

	if err := s.Shutdown(waitCtx); err != nil {
		logger.Error("Invalid server shutdown")
	}
	logger.Info("...Server stopped")
	logger.CloseFile()
}
