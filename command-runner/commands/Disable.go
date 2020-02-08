package commands

import (
	"errors"
	"fmt"

	"github.com/LiveSocket/bot/command-runner/helpers"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

type disableInput struct {
	Channel string
	Name    string
}

// Disable Disables a command
//
// !disable !<name>
//
// command.disable
// [name string]{message twitch.PrivateMessage}
//
// Returns [string] as response for chat
func Disable(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *wamp.Invocation) client.InvokeResult {
		// Get input args from call
		input, err := parseDisable(service, invocation)
		if err != nil {
			return socket.Error(err)
		}

		// Find command
		command, err := helpers.GetCommandByID(service, input.Channel, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// If command doesn't exist
		if command == nil {
			return socket.Success(fmt.Sprintf("Command !%s doesn't exist", input.Name))
		}

		// Set command enabled to false
		command.Enabled = false

		// Update command
		_, err = helpers.UpdateCommand(service, command)
		if err != nil {
			return socket.Error(err)
		}

		// Return message to display in chat
		return socket.Success(fmt.Sprintf("!%s disabled", input.Name))
	}
}

func parseDisable(service *service.Service, invocation *wamp.Invocation) (*disableInput, error) {
	if invocation.ArgumentsKw["message"] == nil {
		return nil, errors.New("Missing message")
	}

	if len(invocation.Arguments) == 0 {
		return nil, errors.New("Missing command name")
	}

	name := conv.ToString(invocation.Arguments[0])
	message, err := conv.ToStringMap(invocation.ArgumentsKw["message"])
	if err != nil {
		return nil, err
	}
	input := &disableInput{
		Channel: conv.ToString(message["Channel"]),
		Name:    name[1:len(name)],
	}
	return validateDisable(service, input)
}

func validateDisable(service *service.Service, input *disableInput) (*disableInput, error) {
	// Find command
	result, err := service.SimpleCall("public.command.getById", nil, wamp.Dict{"channel": input.Channel, "name": input.Name})
	if err != nil {
		return nil, err
	}

	// Check if command exists
	if len(result.Arguments) == 0 {
		return nil, errors.New("Command does not exist")
	}

	// Validation passed
	return input, nil
}
