package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gmalka/rest_api"
	"github.com/gmalka/rest_api/pkg/handler"
	"github.com/gmalka/rest_api/pkg/repository"
	"github.com/gmalka/rest_api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("HELLO")
	if err := initConfig(); err != nil {
		log.Fatalf("Error config read: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Fail load env: %s", err.Error())
	}

	db, err := repositoty.NewPostgresDB(repositoty.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}) 

	if err != nil {
		log.Fatalf("Error while db connecting: %s", err.Error())	
	}

	repos := repositoty.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)


	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error accured while server rungging :%s", err)
	}
}
 
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}