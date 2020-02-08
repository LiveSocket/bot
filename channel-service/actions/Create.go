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

type createInput struct {
	Name    string
	BotName string
}

// Create Creates a new channel
//
// private.channel.create
// {name string, botName string}
//
// Returns [Channel]
func Create(service *service.Service) func(invocation *socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getCreateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Create channel
		channel, err := models.CreateChannel(service, input.Name, input.BotName)
		if err != nil {
			return socket.Error(err)
		}

		// Emit channel created
		err = helpers.EmitChannelCreated(service, channel)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return result
		return socket.Success(channel)
	}
}

func getCreateInput(kwargs wamp.Dict) (*createInput, error) {
	if kwargs == nil {
		return nil, errors.New("Missing Kwargs")
	}

	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	if kwargs["botName"] == nil {
		return nil, errors.New("Missing botName")
	}

	return &createInput{
		Name:    conv.ToString(kwargs["name"]),
		BotName: conv.ToString(kwargs["botName"]),
	}, nil
}
