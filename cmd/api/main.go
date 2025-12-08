package main

import (
	"go-transjakarta/internal/http/handlers"
	"go-transjakarta/internal/mqtt"
	"go-transjakarta/internal/rabbitmq"
	"log"
	"net/http"

	"go-transjakarta/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. KONEKSI DB
	database.ConnectPostgres()

	err := rabbitmq.Connect()
	if err != nil {
		log.Fatal("RabbitMQ Error:", err)
	}
	defer rabbitmq.Close()

	// Jalankan subscriber di background goroutine
	go mqtt.StartSubscriber()

	// HTTP server
	router := gin.Default()

	router.GET("/vehicles/:vehicle_id/location", handlers.GetLastLocation)
	router.GET("/vehicles/:vehicle_id/history", handlers.GetLocationHistory)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Jalankan server
	router.Run(":8088")
}
