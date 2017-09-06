package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

var (
	GolbalConfig *Config
)

type Config struct {
	RpcPort         int
	DBUrl           string
	RetrieveTimeout time.Duration
	MasterUrl       string
}

func init() {
	readConfig()
}

func readConfig() {
	project_path, _ := os.Getwd()
	viper.SetConfigName("slave-config")
	viper.AddConfigPath(project_path)
	viper.ReadInConfig()
	GolbalConfig = &Config{
		MasterUrl:       getVar("master.url").(string),
		DBUrl:           getVar("db.url").(string),
		RpcPort:         getVar("rpc.port").(int),
		RetrieveTimeout: time.Duration(getVar("jobs.sayhello.timeout").(int)) * time.Second,
	}
}

func getVar(key string) interface{} {
	val := os.Getenv(key)
	if val != "" {
		return parsString(val)
	} else {
		return viper.Get(key)
	}
}

func parsString(str string) interface{} {
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return int(i)
	}
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return f
	}
	b, err := strconv.ParseBool(str)
	if err == nil {
		return b
	}
	return str
}
