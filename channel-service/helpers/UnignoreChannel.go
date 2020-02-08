package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// UnignoreChannel WAMP call helper for unignoring a channel
func UnignoreChannel(service *service.Service, name string) error {
	// Call get channel by name endpoint
	_, err := service.SimpleCall("private.channel.unignore", nil, wamp.Dict{"name": name})
	return err
}
