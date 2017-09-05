package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Node Nodes `yaml:"nodes"`
	DB   DB    `yaml:"db"`
}

type DB struct {
	Migrations Migrations
	Connection Connection
}

type Migrations struct {
	Folder           string
	ConnectionString string
	Enable           bool
}

type Connection struct {
	URL string
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

func GetNodesConfig() *Nodes {
	return &Nodes{
		RetrieveTimeout: viper.GetDuration("nodes.retrieve-timeout") * time.Second,
		ExitTimeout:     viper.GetDuration("nodes.exit-timeout") * time.Second,
	}
}

func GetMigrationsConfig() *Migrations {
	return &Migrations{
		Folder:           viper.GetString("db.migrations.folder"),
		ConnectionString: viper.GetString("db.migrations.connectionString"),
		Enable:           viper.GetBool("db.migrations.enable"),
	}
}

func GetConnectionConfig() *Connection {
	return &Connection{
		URL: viper.GetString("db.connection.url"),
	}
}
