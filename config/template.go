package config

import (
	"github.com/spf13/viper"
)

func TemplatesPath() string {
	return viper.GetString("templates.path")
}

func TemplateCaching() bool {
	return viper.GetBool("templates.cache")
}
