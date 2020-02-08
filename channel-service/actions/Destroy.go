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

type destroyInput struct {
	Name string
}

// Destroy Destroys a channel
//
// private.channel.destroy
// {name string}
//
// Returns nothing
func Destroy(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDestroyInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Destroy channel
		err = models.DestroyChannel(service, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Emit channel destroyed
		if err := helpers.EmitChannelDestroyed(service, input.Name); err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getDestroyInput(kwargs wamp.Dict) (*destroyInput, error) {
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &destroyInput{Name: conv.ToString(kwargs["name"])}, nil
}
