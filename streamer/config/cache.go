package config

import (
	"time"

	"github.com/spf13/viper"
)

func PlaylistExpire() time.Duration {
	return viper.GetDuration("streamer.video.expire")
}

func CleanInterval() time.Duration {
	return viper.GetDuration("streamer.video.clean_interval")
}
