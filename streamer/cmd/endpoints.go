package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/movies-to-watch/streamer/config"
	"github.com/ppcamp/movies-to-watch/streamer/handlers"
	log "github.com/sirupsen/logrus"
)

func configureServer(m *gin.Engine) error {
	log.Info("Configuring server")

	middlewares(m)
	routes(m)

	return nil
}

func routes(m *gin.Engine) {
	log.Info("Configuring routes")

	m.GET("/start/:name", handlers.Start)
	m.GET("/stream/:name", handlers.Stream)
	m.Static("/playlists", config.ServeFolder())
}

func middlewares(m *gin.Engine) {
	log.Info("Configuring middlewares")
	m.Use(gin.Logger())
	m.Use(gin.Recovery())
	m.Use(gzip.Gzip(gzip.BestCompression))
}
