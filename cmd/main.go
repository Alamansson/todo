package main

import (
	"github.com/alamansson/todo"
	"log"
	"os"

	"github.com/alamansson/todo/pkg/handler"
	"github.com/alamansson/todo/pkg/repository"
	"github.com/alamansson/todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error initializeing config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initializer db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
