package helpers

import (
	"github.com/LiveSocket/bot/channel-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/gammazero/nexus/v3/wamp"
)

// CreateChannel WAMP call to create a channel
func CreateChannel(service *service.Service, name, botName string) (*models.Channel, error) {
	// Call create channel
	res, err := service.SimpleCall("private.channel.create", nil, wamp.Dict{"name": name, "botName": botName})
	if err != nil {
		return nil, err
	}

	//Check and return result
	if len(res.Arguments) > 0 {
		channel, err := conv.ToStringMap(res.Arguments[0])
		if err != nil {
			return nil, err
		}
		ignored, err := conv.ToBool(channel["ignored"])
		if err != nil {
			return nil, err
		}
		return &models.Channel{
			Name:    conv.ToString(channel["name"]),
			BotName: conv.ToString(channel["bot_name"]),
			Notes:   conv.ToString(channel["notes"]),
			Ignored: ignored,
		}, nil
	}
	return nil, nil
}
