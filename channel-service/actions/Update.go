package actions

import (
	"errors"
	"log"

	"github.com/LiveSocket/bot/channel-service/helpers"
	"github.com/LiveSocket/bot/channel-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type updateInput struct {
	Name    string
	BotName string
}

// Update Updates a channel
//
// private.channel.update
// {name string, botName string}
//
// Returns [Channel]
func Update(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) socket.Result {
		// Get input args from call
		input, err := getUpdateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find channel
		old, err := models.GetChannel(service, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Update channel
		new := old
		new.BotName = input.BotName

		// Save update to channel
		err = models.UpdateChannel(service, new)
		if err != nil {
			return socket.Error(err)
		}

		// Emit channel updated
		if err := helpers.EmitChannelUpdated(service, old, new); err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return updated channel
		return socket.Success(new)
	}
}

func getUpdateInput(kwargs wamp.Dict) (*updateInput, error) {
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	if kwargs["botName"] == nil {
		return nil, errors.New("Missing botName")
	}

	return &updateInput{
		Name:    conv.ToString(kwargs["name"]),
		BotName: conv.ToString(kwargs["botName"]),
	}, nil
}
