package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/htmlx-movies-to-watch/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	viper.AddConfigPath("./")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal(err)
	}

	if err := config.SetupLog(); err != nil {
		log.WithError(err).Fatal(err)
	}

	mux := gin.Default()
	if err := configureServer(mux); err != nil {
		log.WithError(err).Fatal("fail to register handlers")
	}

	s := &http.Server{
		Handler:     mux,
		Addr:        config.ServerPort(),
		BaseContext: func(_ net.Listener) context.Context { return ctx }, // close inner connections
	}

	go serve(ctx, s)
	gracefulStop(ctx, s)
}

// serve starts the server and logs any error returned by the server.
// It blocks until the server is closed, thus, you should call it in a goroutine.
func serve(ctx context.Context, s *http.Server) {
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithError(err).Fatal("fail to start server")
	}
}

func gracefulStop(ctx context.Context, s *http.Server) {
	// await for interrupt signal
	<-ctx.Done()
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), config.WaitTimeout())
	defer cancel()

	log.Infof("sending shutdown signal to server (timeout: %s)", config.WaitTimeout())
	if err := s.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("fail to stop server")
	}
}
