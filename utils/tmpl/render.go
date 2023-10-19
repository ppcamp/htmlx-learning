package tmpl

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/ppcamp/htmlx-movies-to-watch/config"
	log "github.com/sirupsen/logrus"
)

func render(tmpl string, data any, wr io.Writer) error {
	stripped := tmpl[strings.LastIndex(tmpl, "/")+1:]

	log.Info("Rendering template: ", stripped)

	var html *template.Template
	var err error

	if config.TemplateCaching() {
	} else {
		html, err = LoadSingle(tmpl)
		if err != nil {
			return err
		}
	}

	for _, v := range html.Templates() {
		fmt.Println(v.Name())
	}

	return html.ExecuteTemplate(wr, stripped, data)
}
