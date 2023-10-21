package tmpl

import (
	"text/template"

	log "github.com/sirupsen/logrus"
)

var functionMap = template.FuncMap{
	"slice": func(args ...any) []any {
		return args
	},
	"until": func(args ...any) []int {
		if len(args) != 1 {
			log.WithField("args", args).Panic("until() requires one argument")
		}
		n := args[0].(int)
		f := make([]int, 0, n)
		for i := 0; i < n; i++ {
			f = append(f, i)
		}
		return f
	},

	"map": func(args ...any) map[string]any {
		dict := make(map[string]any)
		for i := 0; i < len(args); i += 2 {
			dict[args[i].(string)] = args[i+1]
		}
		return dict
	},
}
