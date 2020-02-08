package helpers

import (
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gempir/go-twitch-irc/v2"
)

// IsMod WAMP call helper for checking if user is a mod
func IsMod(service *service.Service, message *twitch.PrivateMessage) (bool, error) {
	if _, ok := message.User.Badges["moderator"]; ok {
		return true, nil
	}
	if _, ok := message.User.Badges["broadcaster"]; ok {
		return true, nil
	}
	result, err := service.SimpleCall("private.mod.isMod", nil, wamp.Dict{"channel": message.Channel, "username": message.User.Name})
	if err != nil {
		return false, err
	}

	return conv.ToBool(result.Arguments[0])
}
