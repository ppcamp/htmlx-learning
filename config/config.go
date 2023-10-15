package config

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("env")

	err := viper.ReadInConfig()
	return err
}
