package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	Message         string
	RetrieveTimeout time.Duration
}

func init() {
	readConfig()
}

func readConfig() {
	project_path, _ := os.Getwd()
	viper.SetConfigName("slave-config")
	viper.AddConfigPath(project_path)
	viper.ReadInConfig()
}

func GetConfig() *Config {
	return &Config{
		Message:         viper.GetString("message"),
		RetrieveTimeout: viper.GetDuration("retrieve-timeout"),
	}
}
