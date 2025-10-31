package main

import (
	"log"

	"github.com/flipmorsch/Blincast/config"
	"github.com/flipmorsch/Blincast/modules/channel"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.InitDB(&channel.ChannelEntity{})
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer config.CloseDB()

	g := gin.Default()

	g.Run()
	println("Server starting")
}
