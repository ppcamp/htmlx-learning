package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	commoncfg "github.com/ppcamp/movies-to-watch/common/config"
	common "github.com/ppcamp/movies-to-watch/common/server"
	"github.com/ppcamp/movies-to-watch/ui/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := commoncfg.LoadConfig(); err != nil {
		log.WithError(err).Fatal(err)
	}

	if err := commoncfg.SetupLog(); err != nil {
		log.WithError(err).Fatal(err)
	}
	log.Info("Configs loaded")

	mux := gin.Default()
	if err := configureServer(mux); err != nil {
		log.WithError(err).Fatal("Fail to configure server")
	}

	s := &http.Server{
		Handler:     mux,
		Addr:        config.ServerAddr(),
		BaseContext: func(_ net.Listener) context.Context { return ctx }, // close inner connections
	}

	go common.Serve(ctx, s, config.ServerPort())
	common.GracefulStop(ctx, s, commoncfg.WaitTimeout())
}
