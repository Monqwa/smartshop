package main

import (
	"github.com/Monqwa/smartshop/database"
	"github.com/spf13/viper"
	"log"
)

func initConfig() {
	viper.SetConfigName("config") // имя без .yml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // смотрит в корне проекта
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения config.yml: %v", err)
	}
}

func main() {

	initConfig()

	database.Connect()

	log.Println("Приложение запущено, база данных подключена")
}
