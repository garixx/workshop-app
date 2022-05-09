package main

import (
	"fmt"
	tokenRepo "github.com/garixx/workshop-app/internal/authtoken/repository"
	tokenCase "github.com/garixx/workshop-app/internal/authtoken/usecase"
	"github.com/garixx/workshop-app/internal/config"
	"github.com/garixx/workshop-app/internal/database"
	"github.com/garixx/workshop-app/internal/delivery"
	"github.com/garixx/workshop-app/internal/inventory"
	userRepo "github.com/garixx/workshop-app/internal/user/repository"
	userCase "github.com/garixx/workshop-app/internal/user/usecase"
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

	pool, err := database.NewPostgresDB(config.GetPostgresDBConfig())
	if err != nil {
		logrus.Fatalf("connect to DB failed:%s", err.Error())
	}
	defer pool.Close()

	userRepository := userRepo.NewPostgresUserRepository(pool)
	userUseCase := userCase.NewUserUsecase(userRepository)

	authTokenRepository := tokenRepo.NewPostgresAuthTokenRepository(pool)
	tokenUseCase := tokenCase.NewAuthTokenUsecase(authTokenRepository)

	i := inventory.NewInventory(
		userUseCase,
		tokenUseCase,
	)

	err = delivery.RestFrontEnd{}.Start(*i)
	if err != nil {
		logrus.Fatalf("server start failed:%s", err.Error())
	}
}

func initConfig() error {
	// make qa environment default
	_ = viper.BindEnv("environment")
	viper.SetDefault("environment", "QA")

	path := fmt.Sprintf("./configs/%s/config.yml", strings.ToLower(viper.GetString("environment")))
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}
