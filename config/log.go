package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupLog() error {
	lvl, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)
	return nil
}
