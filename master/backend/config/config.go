package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Node Nodes `yaml:"nodes"`
}

type Nodes struct {
	RetrieveTimeout time.Duration
	ExitTimeout     time.Duration
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
		Node: Nodes{
			RetrieveTimeout: viper.GetDuration("nodes.retrieve-timeout"),
			ExitTimeout:     viper.GetDuration("nodes.exit-timeout"),
		},
	}
}
