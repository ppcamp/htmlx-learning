package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupLog() error {
	lvl, err := logrus.ParseLevel(viper.GetString("common.log.level"))
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)
	gin.SetMode(viper.GetString("common.gin.mode"))

	return nil
}
