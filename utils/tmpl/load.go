package tmpl

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/htmlx-movies-to-watch/config"
	"github.com/ppcamp/htmlx-movies-to-watch/utils/files"
	log "github.com/sirupsen/logrus"
)

const (
	FileSuffix   = ".tmpl"
	LayoutSuffix = "layout.tmpl"
)

func load(tmpl string, layout ...string) (*template.Template, error) {
	log.Debug("Loading template: ", tmpl)

	basefile := config.TemplatesPath() + "/" + tmpl

	loadFiles := []string{basefile}
	loadFiles = append(loadFiles, layout...)

	html, err := template.ParseFiles(loadFiles...)
	if err != nil {
		return nil, err
	}

	return html.Funcs(functionMap), nil
}

func LoadSingle(tmpl string) (*template.Template, error) {
	layout, err := files.GlobSuffix(config.TemplatesPath(), LayoutSuffix)
	if err != nil {
		return nil, err
	}

	return load(tmpl, layout...)
}

func LoadMultiple(m *gin.Engine) (map[string]*template.Template, error) {
	tmpls := make(map[string]*template.Template)

	layout, err := files.GlobSuffix(config.TemplatesPath(), LayoutSuffix)
	if err != nil {
		return nil, err
	}

	all, err := files.GlobSuffix(config.TemplatesPath(), FileSuffix)
	if err != nil {
		return nil, err
	}

	for _, s := range all {
		if strings.HasSuffix(s, LayoutSuffix) {
			continue
		}

		html, err := load(s, layout...)
		if err != nil {
			return nil, err
		}
		tmpls[s] = html
	}
	return tmpls, err
}
