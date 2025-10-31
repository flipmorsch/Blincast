package config

import (
	"log"
	"time"

	"github.com/flipmorsch/Blincast/modules/channel"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeServer() {
	db, err := InitDB(&channel.ChannelEntity{})
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer CloseDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	channel.RegisterRoutes(router, db)

	router.Run()
}
