package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
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
}

func middlewares(m *gin.Engine) {
	log.Info("Configuring middlewares")
	m.Use(gin.Logger())
	m.Use(gin.Recovery())
	m.Use(gzip.Gzip(gzip.BestCompression))
}
