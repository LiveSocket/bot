package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// CreateCommand WAMP call helper for creating a command
func CreateCommand(service *service.Service, channel string, name string, response string, username string) error {
	_, err := service.SimpleCall("public.command.create", nil, wamp.Dict{
		"channel":  channel,
		"name":     name,
		"response": response,
		"username": username,
	})
	return err
}
