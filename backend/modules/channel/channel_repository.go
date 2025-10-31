package channel

import "gorm.io/gorm"

type ChannelRepository interface {
	GetChannelByID(id string) (ChannelEntity, error)
	CreateChannel(channel ChannelEntity) error
	UpdateChannel(channel ChannelEntity) error
	DeleteChannel(id string) error
	GetAllChannels() ([]ChannelEntity, error)
	GetAllChannelsWithImage() ([]ChannelEntity, error)
}

type channelRepository struct {
	db *gorm.DB
}

func (c *channelRepository) CreateChannel(channel ChannelEntity) error {
	return c.db.Create(&channel).Error
}

func (c *channelRepository) DeleteChannel(id string) error {
	panic("unimplemented")
}

func (c *channelRepository) GetAllChannels() ([]ChannelEntity, error) {
	panic("unimplemented")
}

func (c *channelRepository) GetAllChannelsWithImage() ([]ChannelEntity, error) {
	panic("unimplemented")
}

func (c *channelRepository) GetChannelByID(id string) (ChannelEntity, error) {
	panic("unimplemented")
}

func (c *channelRepository) UpdateChannel(channel ChannelEntity) error {
	panic("unimplemented")
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	return &channelRepository{db: db}
}
