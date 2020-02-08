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

type unignoreInput struct {
	Name string
}

// Unignore Marks the channel as unignored
//
// private.channel.unignore
// {name string}
//
// Returns nothing
func Unignore(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) socket.Result {
		// Get input args from call
		input, err := getIgnoreInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find channel
		channel, err := models.GetChannel(service, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Set unignored
		channel.Ignored = false
		err = models.UpdateChannel(service, channel)
		if err != nil {
			return socket.Error(err)
		}

		// Emit channel unignored
		err = helpers.EmitChannelUnignored(service, input.Name)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getUnignoreInput(kwargs wamp.Dict) (*ignoreInput, error) {
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &ignoreInput{
		Name: conv.ToString(kwargs["name"]),
	}, nil
}
