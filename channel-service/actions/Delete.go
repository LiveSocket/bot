package actions

import (
	"errors"
	"log"

	"github.com/LiveSocket/bot/channel-service/helpers"
	"github.com/LiveSocket/bot/channel-service/models"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type deleteInput struct {
	Name string
}

// Delete Deletes a channel
//
// private.channel.delete
// {name string}
//
// Returns nothing
func Delete(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDeleteInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find channel by name
		err = models.DeleteChannel(service, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Emit channel deleted event
		if err := helpers.EmitChannelDeleted(service, input.Name); err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getDeleteInput(kwargs wamp.Dict) (*deleteInput, error) {
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &deleteInput{Name: kwargs["name"].(string)}, nil
}
