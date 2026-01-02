package main

import (
	"os"

	todo "github.com/bearury/go-rest-api-postgres-todo"
	handler2 "github.com/bearury/go-rest-api-postgres-todo/pakage/handler"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("init config err: %v", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file err: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.HOST"),
		Port:     viper.GetString("db.PORT"),
		Username: viper.GetString("db.USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.DATABASE"),
		SSLMode:  viper.GetString("db.SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("Filed init db err: %v", err.Error())
	}

	repo := repository.NewRepository(db)
	newService := service.NewService(repo)
	handler := handler2.NewHandler(newService)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("PORT"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("failed to start server: %v", err.Error())
	}

	logrus.Printf("server started in port: %v", viper.GetString("PORT"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
