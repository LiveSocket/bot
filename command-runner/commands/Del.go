package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/LiveSocket/bot/command-runner/helpers"
	"github.com/LiveSocket/bot/conv"
	"github.com/LiveSocket/bot/service"
	"github.com/LiveSocket/bot/service/socket"
	"github.com/gammazero/nexus/v3/wamp"
)

type delInput struct {
	Channel string
	Name    string
}

// DelCommand Deletes a command
//
// !del !<name>
//
// command.del
// [name]{message twitch.PrivateMessage}
//
// Returns [string] as response for chat
func Del(service *service.Service) func(*socket.Invocation) socket.Result {
	return func(invocation *socket.Invocation) socket.Result {
		// Get input args from call
		input, err := getDelInput(service, invocation)
		if err != nil {
			return socket.Error(err)
		}

		// Destroy command
		err = helpers.DestroyCommand(service, input.Channel, input.Name)
		if err != nil {
			return socket.Error(err)
		}

		// Return message to display in chat
		return socket.Success(fmt.Sprintf("!%s deleted", input.Name))
	}
}

func getDelInput(service *service.Service, invocation *wamp.Invocation) (*delInput, error) {
	if invocation.ArgumentsKw["message"] == nil {
		return nil, errors.New("Missing message")
	}

	if len(invocation.Arguments) == 0 {
		return nil, errors.New("Must specify command to delete")
	}

	if invocation.Arguments[0] == nil {
		return nil, errors.New("Missing command name")
	}

	if !strings.HasPrefix(conv.ToString(invocation.Arguments[0]), "!") {
		return nil, errors.New("Command name must start with !")
	}

	name := conv.ToString(invocation.Arguments[0])
	message, err := conv.ToStringMap(invocation.ArgumentsKw["message"])
	if err != nil {
		return nil, err
	}
	input := &delInput{
		Channel: conv.ToString(message["Channel"]),
		Name:    name[1:len(name)],
	}
	return validateDel(service, input)
}

func validateDel(service *service.Service, input *delInput) (*delInput, error) {
	// Find command to delete
	result, err := service.SimpleCall("public.command.getById", nil, wamp.Dict{"channel": input.Channel, "name": input.Name[1:len(input.Name)]})
	if err != nil {
		return nil, err
	}

	// Check command exists
	if len(result.Arguments) == 0 {
		return nil, errors.New("Command does not exist")
	}

	return input, nil
}
