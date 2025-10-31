package channel

import "context"

type ChannelService interface {
	GetChannelByID(ctx context.Context, id string) (ChannelEntity, error)
	CreateChannel(ctx context.Context, channel *ChannelEntity) error
	GetAllChannels(ctx context.Context) (map[string]ChannelInfoDTO, error)
	GetAllChannelsWithImage(ctx context.Context) (map[string]ChannelInfoWithImageDTO, error)
	DeleteChannel(ctx context.Context, id string) error
	UpdateChannel(ctx context.Context, id string, channel UpdateChannelDTO) error
}

type channelService struct {
	repo ChannelRepository
}

func (c *channelService) UpdateChannel(ctx context.Context, id string, channel UpdateChannelDTO) error {
	existingChannel, err := c.repo.GetChannelByID(id)
	if err != nil {
		return err
	}

	if channel.Nome != nil {
		existingChannel.Nome = channel.Nome
	}
	if channel.Url != nil {
		existingChannel.Url = channel.Url
	}
	if channel.Youtube != nil {
		existingChannel.Youtube = channel.Youtube
	}
	if channel.Tag != nil {
		existingChannel.Tag = channel.Tag
	}
	if channel.Imagem != nil {
		existingChannel.Imagem = channel.Imagem
	}
	if channel.Site != nil {
		existingChannel.Site = channel.Site
	}

	return c.repo.UpdateChannel(&existingChannel)
}

func (c *channelService) DeleteChannel(ctx context.Context, id string) error {
	_, err := c.repo.GetChannelByID(id)
	if err != nil {
		return err
	}
	return c.repo.DeleteChannel(id)
}

func (c *channelService) GetAllChannelsWithImage(ctx context.Context) (map[string]ChannelInfoWithImageDTO, error) {
	channels, err := c.repo.GetAllChannelsWithImage()
	if err != nil {
		return nil, err
	}

	channelMap := make(map[string]ChannelInfoWithImageDTO)
	for _, channel := range channels {
		if channel.Nome != nil {
			info := ChannelInfoWithImageDTO{}
			if channel.Url != nil {
				info.Url = *channel.Url
			}
			if channel.Youtube != nil {
				info.Youtube = *channel.Youtube
			}
			if channel.Tag != nil {
				info.Tag = *channel.Tag
			}
			info.Imagem = channel.Imagem
			if channel.Site != nil {
				info.Site = *channel.Site
			}
			channelMap[*channel.Nome] = info
		}
	}

	return channelMap, nil
}

func (c *channelService) CreateChannel(ctx context.Context, channel *ChannelEntity) error {
	return c.repo.CreateChannel(*channel)
}

func (c *channelService) GetChannelByID(ctx context.Context, id string) (ChannelEntity, error) {
	return c.repo.GetChannelByID(id)
}

func (c *channelService) GetAllChannels(ctx context.Context) (map[string]ChannelInfoDTO, error) {
	channels, err := c.repo.GetAllChannels()
	if err != nil {
		return nil, err
	}

	channelMap := make(map[string]ChannelInfoDTO)
	for _, channel := range channels {
		if channel.Nome != nil {
			info := ChannelInfoDTO{}
			if channel.Url != nil {
				info.Url = *channel.Url
			}
			if channel.Youtube != nil {
				info.Youtube = *channel.Youtube
			}
			if channel.Tag != nil {
				info.Tag = *channel.Tag
			}
			if channel.Site != nil {
				info.Site = *channel.Site
			}
			channelMap[*channel.Nome] = info
		}
	}

	return channelMap, nil
}

func NewChannelService(repo ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}
