package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// DestroyCommand WAMP call helper for destroying a command
func DestroyCommand(service *service.Service, channel string, name string) error {
	_, err := service.SimpleCall("public.command.destroy", nil, wamp.Dict{"channel": channel, "name": name})
	return err
}
