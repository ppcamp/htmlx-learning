package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/htmlx-movies-to-watch/config"
	"github.com/ppcamp/htmlx-movies-to-watch/handlers"
	"github.com/ppcamp/htmlx-movies-to-watch/utils/templ"
	log "github.com/sirupsen/logrus"
)

func configureServer(mux *gin.Engine) error {
	tmplPath := config.TemplatesPath() + "/*.tmpl"
	log.Info(tmplPath)

	mux.FuncMap = templ.FunctionMap()
	mux.LoadHTMLGlob(tmplPath)
	mux.Use(gzip.Gzip(gzip.BestCompression))

	mux.Static("/assets", "./"+config.StaticFolder())
	mux.GET("/", handlers.Index)

	return nil
}
