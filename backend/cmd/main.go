package main

import (
	"github.com/Monqwa/smartshop/database"
	"github.com/Monqwa/smartshop/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
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

	RunHttpServer()

	log.Println("Приложение запущено, база данных подключена")
}

func RunHttpServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	handlers.NewItemRouter(r)
	handlers.NewListRouter(r)

	r.Run()
}
