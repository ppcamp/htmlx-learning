package tmpl

import "text/template"

var functionMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},

	"arr": func(args ...interface{}) []interface{} {
		return args
	},

	"dict": func(args ...interface{}) map[string]interface{} {
		dict := make(map[string]interface{})
		for i := 0; i < len(args); i += 2 {
			dict[args[i].(string)] = args[i+1]
		}
		return dict
	},
}
