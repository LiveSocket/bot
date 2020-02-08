package actions

import (
	"errors"
	"log"

	"github.com/LiveSocket/bot/command-service/models"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type destroyAllInput struct {
	Channel string
}

// DestroyAll Destroys all commands for a channel
//
// public.command.destroyAll
// {channel string}
//
// Returns nothing
func DestroyAll(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDestroyAllInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find all commands for the channel
		err = models.DestroyAllCommands(service, input.Channel)
		if err != nil {
			return socket.Error(err)
		}

		// Emit destroyed event for each command
		if err := service.Publish("event.command.destroyedAll", nil, wamp.List{input.Channel}, nil); err != nil {
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getDestroyAllInput(kwargs wamp.Dict) (*destroyAllInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	return &destroyAllInput{Channel: conv.ToString(kwargs["channel"])}, nil
}
