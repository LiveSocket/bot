package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitChannelDeleted Emits "event.channel.deleted"
func EmitChannelDeleted(service *service.Service, name string) error {
	return service.Publish("event.channel.deleted", nil, wamp.List{name}, nil)
}
