package channel

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	channelRepo := NewChannelRepository(db)
	channelService := NewChannelService(channelRepo)
	channelController := NewChannelController(channelService)

	channelRoutes := router.Group("/channel")
	channelRoutes.GET("", channelController.GetChannelByID)
	channelRoutes.POST("/", channelController.CreateChannel)
}
