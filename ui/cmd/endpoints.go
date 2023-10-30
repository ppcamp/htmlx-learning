package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/movies-to-watch/ui/config"
	"github.com/ppcamp/movies-to-watch/ui/handlers"
	"github.com/ppcamp/movies-to-watch/ui/utils/tmpl"
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
	m.GET("/watch/:video", handlers.Watch)

	htmlx := m.Group("/htmlx")
	htmlx.GET("/videos", handlers.HtmlxVideos)
	htmlx.GET("/todos", handlers.HtmlxTodos)
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
