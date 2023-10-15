package config

import (
	"time"

	"github.com/spf13/viper"
)

func ServerPort() string {
	return viper.GetString("server.port")
}

func ServerAddr() string {
	return ":" + ServerPort()
}

func WaitTimeout() time.Duration {
	v := viper.GetDuration("server.graceful_stop_timeout")
	return v * time.Second
}

func StaticFolder() string {
	return viper.GetString("server.static_folder")
}
