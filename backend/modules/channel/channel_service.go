package channel

import "context"

type ChannelService interface {
	GetChannelByID(ctx context.Context, id string) (*ChannelEntity, error)
}
