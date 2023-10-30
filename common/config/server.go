package config

import (
	"time"

	"github.com/spf13/viper"
)

func WaitTimeout() time.Duration {
	v := viper.GetDuration("common.server.graceful_stop_timeout")
	return v * time.Second
}
