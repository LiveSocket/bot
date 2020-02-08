package actions

import (
	"errors"
	"log"

	"github.com/LiveSocket/bot/command-runner/helpers"
	"github.com/LiveSocket/bot/command-runner/models"
	"github.com/LiveSocket/bot/conv"

	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type activateInput struct {
	Channel string
	Name    string
}

// Activate Activates a custom command
//
// private.command.activate
// {channel string, name string}
//
// Returns nothing
func Activate(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getActivateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find custom command
		command, err := models.GetCustom(service, input.Name, input.Channel)
		if err != nil {
			return socket.Error(err)
		}

		if command == nil {
			return socket.Error("Custom command not found")
		}

		// Activate command
		command.Enabled = true
		err = models.UpdateCustom(service, command)
		if err != nil {
			return socket.Error(err)
		}

		// Emit command activated
		err = helpers.EmitCommandActivated(service, command)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// Return success but nothing
		return socket.Success()
	}
}

func getActivateInput(kwargs wamp.Dict) (*activateInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &activateInput{
		Channel: conv.ToString(kwargs["channel"]),
		Name:    conv.ToString(kwargs["name"]),
	}, nil
}
