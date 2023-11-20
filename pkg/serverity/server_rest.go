package serverity

import (
	"context"
	"net/http"
	"os"
	"time"
)

type HttpServer struct {
	httpServer *http.Server
}

func (s *HttpServer) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           os.Getenv("APP_IP") + ":" + os.Getenv("APP_PORT"),
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
