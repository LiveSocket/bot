package helpers

import (
	"github.com/LiveSocket/bot/channel-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitChannelCreated Emits "event.channel.created"
func EmitChannelCreated(service *service.Service, channel *models.Channel) error {
	return service.Publish("event.channel.created", nil, wamp.List{channel}, nil)
}
