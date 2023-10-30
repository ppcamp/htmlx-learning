package config

import (
	"github.com/spf13/viper"
)

func ServerPort() string {
	return viper.GetString("streamer.server.port")
}

func ServerAddr() string {
	return ":" + ServerPort()
}

func ServeFolder() string {
	return viper.GetString("streamer.server.serve_folder")
}

func VideosFolder() string {
	return viper.GetString("streamer.server.videos_folder")
}
