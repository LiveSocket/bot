package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitChannelDestroyed Emits "event.channel.destroyed"
func EmitChannelDestroyed(service *service.Service, name string) error {
	return service.Publish("event.channel.destroyed", nil, wamp.List{name}, nil)
}
