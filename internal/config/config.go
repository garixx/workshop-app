package config

import (
	"github.com/spf13/viper"
	"os"
)

type PostgresDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func GetPostgresDBConfig() PostgresDBConfig {
	return PostgresDBConfig{
		Host:     viper.GetString("database.postgres.host"),
		Port:     viper.GetString("database.postgres.port"),
		Username: viper.GetString("database.postgres.user"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("database.postgres.dbname"),
		SSLMode:  viper.GetString("database.postgres.sslmode"),
	}
}
