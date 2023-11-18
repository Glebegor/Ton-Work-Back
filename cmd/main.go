package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	handlers "github.com/Glebegor/Ton-Work-Back/pkg/handler"
	repositoryes "github.com/Glebegor/Ton-Work-Back/pkg/repository"
	services "github.com/Glebegor/Ton-Work-Back/pkg/service"
	TonWork "github.com/Glebegor/Ton-Work-Back/structint"
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
		Server:   viper.GetString("db.Server"),
		Port:     viper.GetString("db.Port"),
		Database: viper.GetString("db.DBName"),
		User:     viper.GetString("db.Username"),
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
		for {
			if err := repository.Subscribes.UpdateTimeOfSub(); err != nil {
				logrus.Fatalf("Error while was updating sub time: %s", err.Error())
			}
			fmt.Printf("--Updated time of sub: %s--", time.Now())
			time.Sleep(time.Hour * 24)
		}
	}()

	go hub.Run()

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
