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

type ignoreInput struct {
	Name string
}

// Ignore Marks the channel as ignored
//
// private.channel.ignore
// {name string}
//
// Returns nothing
func Ignore(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
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

		// Set ignored
		channel.Ignored = true
		err = models.UpdateChannel(service, channel)
		if err != nil {
			return socket.Error(err)
		}

		// Emit channel ignored
		err = helpers.EmitChannelIgnored(service, input.Name)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getIgnoreInput(kwargs wamp.Dict) (*ignoreInput, error) {
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &ignoreInput{
		Name: conv.ToString(kwargs["name"]),
	}, nil
}
