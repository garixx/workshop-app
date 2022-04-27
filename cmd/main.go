package main

import (
	"fmt"
	"github.com/garixx/workshop-app/internal/config"
	repository "github.com/garixx/workshop-app/internal/user/repository/postgres"
	"github.com/garixx/workshop-app/internal/user/usecase"
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

	pool, err := repository.NewPostgresDB(config.GetPostgresDBConfig())
	if err != nil {
		logrus.Fatalf("connect to DB failed:%s", err.Error())
	}
	defer pool.Close()

	repo := repository.NewPostgresUserRepository(pool)
	useCase := usecase.NewUserUsecase(repo)

}

func initConfig() error {
	// make qa environment default
	_ = viper.BindEnv("environment")
	viper.SetDefault("environment", "QA")

	path := fmt.Sprintf("./configs/%s/config.yml", strings.ToLower(viper.GetString("environment")))
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}
