package config

import (
	"github.com/spf13/viper"
)

func ServerPort() string {
	return viper.GetString("ui.server.port")
}

func ServerAddr() string {
	return ":" + ServerPort()
}

func StaticFolder() string {
	return viper.GetString("ui.server.static_folder")
}
