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

type deactivateInput struct {
	Channel string
	Name    string
}

// Deactivate Deactivates a registered command
//
// private.command.deactivate
// {channel string, name string}
//
// Returns nothing
func Deactivate(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDeactivateInput(invocation.ArgumentsKw)
		if err != nil {
			return socket.Error(err)
		}

		// Find custom command
		command, err := models.GetCustom(service, input.Name, input.Channel)
		if err != nil {
			return socket.Error(err)
		}

		// Find and deactivate command
		command.Enabled = false
		err = models.UpdateCustom(service, command)
		if err != nil {
			return socket.Error(err)
		}

		// Emit command deactivated
		err = helpers.EmitCommandDeactivated(service, command)
		if err != nil {
			// Don't fail on error
			log.Print(err)
		}

		// return success but nothing
		return socket.Success()
	}
}

func getDeactivateInput(kwargs wamp.Dict) (*deactivateInput, error) {
	if kwargs["channel"] == nil {
		return nil, errors.New("Missing channel")
	}
	if kwargs["name"] == nil {
		return nil, errors.New("Missing name")
	}

	return &deactivateInput{
		Channel: conv.ToString(kwargs["channel"]),
		Name:    conv.ToString(kwargs["name"]),
	}, nil
}
