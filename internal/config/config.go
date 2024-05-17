package config

import (
	"github.com/leetcode-golang-classroom/gin-auth/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	MYSQL_URI       string `mapstructure:"MYSQL_URI"`
	SERVER_PORT     uint64 `mapstructure:"SERVER_PORT"`
	JWT_SIGN_SECRET string `mapstructure:"JWT_SIGN_SECRET"`
}

var C *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		util.FailOnError(err, "Failed to load config")
	}
	v.AutomaticEnv()

	err = v.Unmarshal(&C)
	if err != nil {
		util.FailOnError(err, "Failed to read environment")
	}
}
