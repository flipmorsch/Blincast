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

// DeleteChannel implements ChannelController.
func (c *channelController) DeleteChannel(ctx *gin.Context) {
	panic("unimplemented")
}

// GetAllChannels implements ChannelController.
func (c *channelController) GetAllChannels(ctx *gin.Context) {
	channels, err := c.channelService.GetAllChannels(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, channels)
}

// GetAllChannelsWithImage implements ChannelController.
func (c *channelController) GetAllChannelsWithImage(ctx *gin.Context) {
	panic("unimplemented")
}

// GetChannelByID implements ChannelController.
func (c *channelController) GetChannelByID(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateChannel implements ChannelController.
func (c *channelController) UpdateChannel(ctx *gin.Context) {
	panic("unimplemented")
}

func NewChannelController(channelService ChannelService) ChannelController {
	return &channelController{channelService: channelService}
}
