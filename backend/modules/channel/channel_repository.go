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
	var channels []ChannelEntity
	err := c.db.Find(&channels).Error
	return channels, err
}

func (c *channelRepository) GetAllChannelsWithImage() ([]ChannelEntity, error) {
	var channels []ChannelEntity
	err := c.db.Where("imagem IS NOT NULL").Find(&channels).Error
	return channels, err
}

func (c *channelRepository) GetChannelByID(id string) (ChannelEntity, error) {
	var channel ChannelEntity
	err := c.db.First(&channel, "id = ?", id).Error
	return channel, err
}

func (c *channelRepository) UpdateChannel(channel ChannelEntity) error {
	panic("unimplemented")
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	return &channelRepository{db: db}
}
