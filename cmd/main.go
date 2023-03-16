package main

import (
	"log"

	"github.com/gmalka/rest_api"
	"github.com/gmalka/rest_api/pkg/handler"
	"github.com/gmalka/rest_api/pkg/repository"
	"github.com/gmalka/rest_api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error config read: %s", err.Error())
	}
	repos := repositoty.NewRepository()
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