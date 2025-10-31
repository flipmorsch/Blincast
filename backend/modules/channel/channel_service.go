package channel

import "context"

type ChannelService interface {
	GetChannelByID(ctx context.Context, id string) (ChannelEntity, error)
	CreateChannel(ctx context.Context, channel *ChannelEntity) error
	GetAllChannels(ctx context.Context) ([]ChannelEntity, error)
	GetAllChannelsWithImage(ctx context.Context) ([]ChannelEntity, error)
	DeleteChannel(ctx context.Context, id string) error
}

type channelService struct {
	repo ChannelRepository
}

func (c *channelService) DeleteChannel(ctx context.Context, id string) error {
	_, err := c.repo.GetChannelByID(id)
	if err != nil {
		return err
	}
	return c.repo.DeleteChannel(id)
}

func (c *channelService) GetAllChannelsWithImage(ctx context.Context) ([]ChannelEntity, error) {
	return c.repo.GetAllChannelsWithImage()
}

func (c *channelService) CreateChannel(ctx context.Context, channel *ChannelEntity) error {
	return c.repo.CreateChannel(*channel)
}

func (c *channelService) GetChannelByID(ctx context.Context, id string) (ChannelEntity, error) {
	return c.repo.GetChannelByID(id)
}

func (c *channelService) GetAllChannels(ctx context.Context) ([]ChannelEntity, error) {
	return c.repo.GetAllChannels()
}

func NewChannelService(repo ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}
