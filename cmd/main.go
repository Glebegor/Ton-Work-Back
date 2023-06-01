package main

import (
	"os"

	TonWork "github.com/TonWork/back"
	handlers "github.com/TonWork/back/pkg/handler"
	repositoryes "github.com/TonWork/back/pkg/repository"
	services "github.com/TonWork/back/pkg/service"
	godotenv "github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initialization")
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error while loading env variables: %s", err.Error())
	}
	db, err := repositoryes.ConnectDB(repositoryes.ConfigDB{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
		Username: viper.GetString("db.Username"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Error while connected to database: %s", err.Error())
	}
	repository := repositoryes.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)
	server := new(TonWork.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {

	}
	return
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
