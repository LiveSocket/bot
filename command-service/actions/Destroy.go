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

type destroyInput struct {
	Channel string
	Name    string
}

// Destroy Destroys a channel
//
// public.command.destroy
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

		// Destroy command
		err = models.DestroyCommand(service, input.Channel, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Emit event that command was destroyed
		if err := service.Publish("event.command.destroyed", nil, wamp.List{input.Channel, input.Name}, nil); err != nil {
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getDestroyInput(kwargs wamp.Dict) (*destroyInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}

	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &destroyInput{
		Channel: conv.ToString(kwargs["channel"]),
		Name:    conv.ToString(kwargs["name"]),
	}, nil
}
