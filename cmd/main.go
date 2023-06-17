package main

import (
	"os"
	"os/signal"
	"syscall"

	TonWork "github.com/TonWork/back"
	handlers "github.com/TonWork/back/pkg/handler"
	repositoryes "github.com/TonWork/back/pkg/repository"
	services "github.com/TonWork/back/pkg/service"
	godotenv "github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"

	_ "github.com/lib/pq"
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
	hub := TonWork.NewHub()
	handler := handlers.NewHandler(service, hub)
	server := new(TonWork.Server)
	go func() {
		err := server.Run(viper.GetString("Port"), handler.InitRoutes())
		if err != nil {
			logrus.Fatalf("Error while running server: %s", err.Error())
		}
	}()
	logrus.Printf("Server is loading on port %s.", viper.GetString("Port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
