package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// serve starts the server and logs any error returned by the server.
// It blocks until the server is closed, thus, you should call it in a goroutine.
func Serve(ctx context.Context, s *http.Server, port string) {
	log.Infof("Server listening at http://localhost:%s", port)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithError(err).Fatal("Fail to start server")
	}
}

func GracefulStop(ctx context.Context, s *http.Server, timeout time.Duration) {
	// await for interrupt signal
	<-ctx.Done()
	log.Info("Stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Infof("Sending shutdown signal to server (timeout: %s)", timeout)
	if err := s.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("Fail to stop server")
	}
}
