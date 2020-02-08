package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// Channel The important bits of a channel
type Channel struct {
	Name    string
	BotName string
}

// GetChannel WAMP call helper for getting a channel by name
func GetChannel(service *service.Service, name string) (*Channel, error) {
	// Call get channel by name endpoint
	res, err := service.SimpleCall("private.channel.get", nil, wamp.Dict{"name": name})
	if err != nil {
		return nil, err
	}
	// Check and return response
	if len(res.Arguments) > 0 && res.Arguments[0] != nil {
		channel, err := conv.ToStringMap(res.Arguments[0])
		if err != nil {
			return nil, err
		}
		return &Channel{
			Name:    conv.ToString(channel["name"]),
			BotName: conv.ToString(channel["bot_name"]),
		}, nil
	}
	return nil, nil
}
