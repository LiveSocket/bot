package helpers

import (
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// DestroyChannel WAMP call to destroy a channel
func DestroyChannel(service *service.Service, name string) error {
	// Call destroy channel
	_, err := service.SimpleCall("private.channel.destroy", nil, wamp.Dict{"name": name})
	return err
}
