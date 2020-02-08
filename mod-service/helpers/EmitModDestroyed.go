package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitModDestroyed Emits "event.mod.destroyed"
func EmitModDestroyed(service *service.Service, channel, username string) error {
	return service.Publish("event.mod.destroyed", nil, wamp.List{channel, username}, nil)
}
