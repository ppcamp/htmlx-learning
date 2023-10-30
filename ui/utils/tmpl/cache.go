package tmpl

import "text/template"

// TODO: implement cache for templates basing on time
type Cache interface {
	Get(tmpl string) (*template.Template, error)

	Set(tmpl string, t *template.Template) error
}
