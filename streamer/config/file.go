package config

import "github.com/spf13/viper"

func VideoFileExt() string {
	return viper.GetString("streamer.video.ext")
}

func VideosFolder() string {
	return viper.GetString("streamer.video.path")
}

func PlaylistExt() string {
	return viper.GetString("streamer.video.playlist")
}

func PlaylistFolder() string {
	return viper.GetString("streamer.video.stream")
}
