package tmpl

import (
	"html/template"
	"io"

	"github.com/ppcamp/htmlx-movies-to-watch/config"
	log "github.com/sirupsen/logrus"
)

func render(tmpl string, data any, wr io.Writer) error {
	log.Debug("Rendering template: ", tmpl)

	var html *template.Template
	var err error

	if config.TemplateCaching() {
	} else {
		html, err = LoadSingle(tmpl)
		if err != nil {
			return err
		}
	}

	return html.Execute(wr, data)
}
