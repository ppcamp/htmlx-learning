package config

import (
	"github.com/spf13/viper"
)

func TemplatesPath() string {
	return viper.GetString("ui.templates.path")
}

func TemplateCaching() bool {
	return viper.GetBool("ui.templates.cache")
}
