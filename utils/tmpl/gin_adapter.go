package tmpl

import (
	"net/http"
	"text/template"

	gin "github.com/gin-gonic/gin/render"
)

type ginRender struct{}

// Instance (HTMLProduction) returns an HTML instance which it realizes Render interface.
func (r *ginRender) Instance(name string, data any) gin.Render {
	return &htmlRender{name, data, nil}
}

func NewCustomRender() *ginRender {
	return &ginRender{}
}

type htmlRender struct {
	Name     string
	Data     any
	Template *template.Template
}

// Render (HTML) executes template and writes its result with custom ContentType for response.
func (r *htmlRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	return render(r.Name, r.Data, w)
}

// WriteContentType (HTML) writes HTML ContentType.
func (r *htmlRender) WriteContentType(w http.ResponseWriter) {
	var htmlContentType = []string{"text/html; charset=utf-8"}

	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = htmlContentType
	}
}
