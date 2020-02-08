package helpers

import (
	"github.com/LiveSocket/bot/mod-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitModCreated Emits "event.mod.created"
func EmitModCreated(service *service.Service, mod *models.Mod) error {
	return service.Publish("event.mod.created", nil, wamp.List{mod}, nil)
}
