package tmpl

import (
	"html/template"
	"io"
	"strings"

	"github.com/ppcamp/movies-to-watch/ui/config"
	log "github.com/sirupsen/logrus"
)

func render(tmpl string, data any, wr io.Writer) error {
	stripped := tmpl[strings.LastIndex(tmpl, "/")+1:]

	log.WithField("data", data).Info("Rendering template: ", stripped)

	var html *template.Template
	var err error

	if config.TemplateCaching() {
	} else {
		html, err = LoadSingle(tmpl)
		if err != nil {
			return err
		}
	}

	return html.ExecuteTemplate(wr, stripped, data)
}
