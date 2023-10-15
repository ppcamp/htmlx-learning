package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/htmlx-movies-to-watch/config"
	"github.com/ppcamp/htmlx-movies-to-watch/handlers"
	"github.com/ppcamp/htmlx-movies-to-watch/utils/tmpl"
	log "github.com/sirupsen/logrus"
)

func configureServer(m *gin.Engine) error {
	log.Info("Configuring server")

	if err := templates(m); err != nil {
		return err
	}

	middlewares(m)
	routes(m)

	return nil
}

func routes(m *gin.Engine) {
	log.Info("Configuring routes")

	m.Static("/assets", "./"+config.StaticFolder())

	m.GET("/", handlers.Index)
	m.GET("/bookmarks", handlers.Bookmarks)
	m.GET("/watch", handlers.Watch)
	m.GET("/search", handlers.Search)
	m.GET("/todos", handlers.Todos)

	m.GET("/videos", handlers.Videos)
}

func middlewares(m *gin.Engine) {
	log.Info("Configuring middlewares")
	m.Use(gin.Logger())
	m.Use(gin.Recovery())
	m.Use(gzip.Gzip(gzip.BestCompression))
}

func templates(m *gin.Engine) error {
	m.HTMLRender = tmpl.NewCustomRender()
	return nil
}
