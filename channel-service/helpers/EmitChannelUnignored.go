package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitChannelUnignored Emits "event.channel.unignored"
func EmitChannelUnignored(service *service.Service, channel string) error {
	return service.Publish("event.channel.unignored", nil, wamp.List{channel}, nil)
}
