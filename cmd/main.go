package main

import (
	"github.com/garixx/workshop-app/internal/config"
	"github.com/garixx/workshop-app/internal/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables:%s", err)
	}
	if err := initConfig(); err != nil {
		logrus.Fatalf("config parsing failed:%s", err.Error())
	}

	db, err := repository.NewPostgresDB(config.GetPostgresDBConfig())
	if err != nil {
		logrus.Fatalf("connect to DB failed:%s", err.Error())
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	
}

func initConfig() error {
	// make qa environment default
	_ = viper.BindEnv("environment")
	viper.SetDefault("environment", "QA")

	viper.SetConfigFile(`./configs/` + strings.ToLower(viper.GetString("environment")) + `/config.yml`)
	return viper.ReadInConfig()
}
