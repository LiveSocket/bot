package helpers

import (
	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitCommandDeactivated Emits "event.command.deactivated"
func EmitCommandDeactivated(service *service.Service, command *models.Custom) error {
	return service.Publish("event.command.deactivated", nil, wamp.List{command}, nil)
}
