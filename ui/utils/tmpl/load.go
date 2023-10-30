package tmpl

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/movies-to-watch/common/files"
	"github.com/ppcamp/movies-to-watch/ui/config"
	log "github.com/sirupsen/logrus"
)

const (
	FileSuffix      = ".tmpl"
	LayoutSuffix    = "layout.tmpl"
	ComponentSuffix = "component.tmpl"
)

func load(tmpl string, layout ...string) (*template.Template, error) {
	log.Debug("Loading template: ", tmpl)

	withPath := config.TemplatesPath() + "/" + tmpl

	loadFiles := []string{withPath}
	loadFiles = append(loadFiles, layout...)

	log.Debug("Curr: ", withPath)
	log.Debug("Layout files: ", layout)

	t := template.New(tmpl).Funcs(functionMap)

	return t.ParseFiles(loadFiles...)
}

func LoadSingle(tmpl string) (*template.Template, error) {
	layout, err := files.GlobSuffix(config.TemplatesPath(), LayoutSuffix, ComponentSuffix)
	if err != nil {
		return nil, err
	}

	return load(tmpl, layout...)
}

func LoadMultiple(m *gin.Engine) (map[string]*template.Template, error) {
	tmpls := make(map[string]*template.Template)

	layout, err := files.GlobSuffix(config.TemplatesPath(), LayoutSuffix, ComponentSuffix)
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
