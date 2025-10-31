package config

import (
	"log"

	"github.com/flipmorsch/Blincast/modules/channel"
	"github.com/gin-gonic/gin"
)

func InitializeServer() {
	db, err := InitDB(&channel.ChannelEntity{})
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer CloseDB()

	router := gin.Default()

	channel.RegisterRoutes(router, db)

	router.Run()
}
