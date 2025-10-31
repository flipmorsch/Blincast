package channel

import "context"

type ChannelService interface {
	GetChannelByID(ctx context.Context, id string) (*ChannelEntity, error)
	CreateChannel(ctx context.Context, channel *ChannelEntity) error
}

type channelService struct {
	repo ChannelRepository
}

func (c *channelService) CreateChannel(ctx context.Context, channel *ChannelEntity) error {
	return c.repo.CreateChannel(*channel)
}

func (c *channelService) GetChannelByID(ctx context.Context, id string) (*ChannelEntity, error) {
	panic("unimplemented")
}

func NewChannelService(repo ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}
