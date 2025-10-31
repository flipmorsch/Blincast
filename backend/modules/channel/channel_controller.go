package channel

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func (c *channelController) CreateChannel(ctx *gin.Context) {
	var channel ChannelEntity
	if err := ctx.ShouldBindJSON(&channel); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateChannelEntity(&channel); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.channelService.CreateChannel(ctx, &channel); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(201)
}

func (c *channelController) DeleteChannel(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *channelController) GetAllChannels(ctx *gin.Context) {
	channels, err := c.channelService.GetAllChannels(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, channels)
}

func (c *channelController) GetAllChannelsWithImage(ctx *gin.Context) {
	channels, err := c.channelService.GetAllChannelsWithImage(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, channels)
}

func (c *channelController) GetChannelByID(ctx *gin.Context) {
	channel, err := c.channelService.GetChannelByID(ctx, ctx.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "channel not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, channel)
}

func (c *channelController) UpdateChannel(ctx *gin.Context) {
	panic("unimplemented")
}

func NewChannelController(channelService ChannelService) ChannelController {
	return &channelController{channelService: channelService}
}
