package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitChannelIgnored Emits "event.channel.ignored"
func EmitChannelIgnored(service *service.Service, channel string) error {
	return service.Publish("event.channel.ignored", nil, wamp.List{channel}, nil)
}
