package channel

import "github.com/gin-gonic/gin"

type (
	ChannelController interface {
		GetAllChannels(ctx *gin.Context)
		GetAllChannelsWithImage(ctx *gin.Context)
		GetChannelByID(ctx *gin.Context)
		CreateChannel(ctx *gin.Context)
		UpdateChannel(ctx *gin.Context)
		DeleteChannel(ctx *gin.Context)
	}

	channelController struct {
		channelService ChannelService
	}
)
