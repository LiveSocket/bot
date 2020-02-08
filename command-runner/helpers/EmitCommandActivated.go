package helpers

import (
	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// EmitCommandActivated Emits "event.command.activated"
func EmitCommandActivated(service *service.Service, command *models.Custom) error {
	return service.Publish("event.command.activated", nil, wamp.List{command}, nil)
}
