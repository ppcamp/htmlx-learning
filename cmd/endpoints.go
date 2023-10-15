package main

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/htmlx-movies-to-watch/config"
	"github.com/ppcamp/htmlx-movies-to-watch/handlers"
	"github.com/ppcamp/htmlx-movies-to-watch/utils/templ"
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
	m.Static("/assets", "./"+config.StaticFolder())
	m.GET("/", handlers.Index)
}

func middlewares(m *gin.Engine) {
	m.Use(gin.Logger())
	m.Use(gin.Recovery())
	m.Use(gzip.Gzip(gzip.BestCompression))
}

func templates(m *gin.Engine) error {
	tmpl, err := glob(config.TemplatesPath())
	if err != nil {
		return err
	}
	html := template.Must(template.ParseFiles(tmpl...))
	m.SetHTMLTemplate(html)

	m.FuncMap = templ.FunctionMap()
	return nil
}

// glob returns all the templates files in the templates folder.
// We need this because the filepath subdirectory glob (**/*.tmpl) only works if we don't pass the
// prefix folder (templates/**/*.tmpl), otherwise, it returns only the files under subdirectories.
//
// This function is a workaround to match all files.
func glob(p string) ([]string, error) {
	tmpl := make([]string, 0, 10)
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			log.Debug("Loading template: ", path)
			tmpl = append(tmpl, path)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
