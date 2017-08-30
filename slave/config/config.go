package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Message      string
	ResponseTime int
	ExitTime     int
}

func init() {
	readConfig()
}

func readConfig() {
	project_path, _ := os.Getwd()
	viper.SetConfigName("master-config")
	viper.AddConfigPath(project_path)
	viper.ReadInConfig()
}


func GetConfig() *Config {
	return &Config{
		Message:      viper.GetString("message"),
		ResponseTime: viper.GetInt("response-time"),
		ExitTime:     viper.GetInt("exit-time"),
	}
}