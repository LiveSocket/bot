package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// IgnoreChannel WAMP call helper for ignoring a channel
func IgnoreChannel(service *service.Service, name string) error {
	// Call get channel by name endpoint
	_, err := service.SimpleCall("private.channel.ignore", nil, wamp.Dict{"name": name})
	return err
}
