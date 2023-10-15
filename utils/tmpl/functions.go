package tmpl

import "text/template"

var functionMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
}
