package helpers

import (
	"github.com/LiveSocket/bot/channel-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitChannelUpdated Emits "event.channel.updated"
func EmitChannelUpdated(service *service.Service, old *models.Channel, new *models.Channel) error {
	return service.Publish("event.channel.updated", nil, wamp.List{old, new}, nil)
}
